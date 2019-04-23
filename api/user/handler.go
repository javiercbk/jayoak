package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/javiercbk/jayoak/http/response"
	"github.com/javiercbk/jayoak/http/session"
	"github.com/javiercbk/jayoak/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// ErrUserAlreadyExists is returned when attempting to create a user that already exists
var ErrUserAlreadyExists = errors.New("user already exists")

// BCryptCost is the ammount of iterations applied to bcrypt
const BCryptCost = 12

// HashPassword hashes a password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), BCryptCost)
	return string(hash[0:]), err
}

// Handlers defines a handler for user
type Handlers struct {
	logger *log.Logger
	db     *sql.DB
}

// NewHandlers creates a new handler for sound
func NewHandlers(logger *log.Logger, db *sql.DB) *Handlers {
	return &Handlers{
		logger: logger,
		db:     db,
	}
}

// Routes initializes the routes for the audio handlers
func (h *Handlers) Routes(rg *gin.RouterGroup) {
	rg.POST("/user", h.CreateAccount)
	rg.GET("/user/:userID", h.Retrieve)
}

type userSearch struct {
	ID string `uri:"userID" binding:"required"`
}

type prospectOrganization struct {
	ID   int64  `json:"id" binding:"numeric"`
	Name string `json:"name" binding:"length(1,256),optional"`
}

type prospectAccount struct {
	Name         string               `json:"name" binding:"required,length(1,256)"`
	Email        string               `json:"email" binding:"required,length(1,256)"`
	Password     string               `json:"password" binding:"required,length(10,256)"`
	Organization prospectOrganization `json:"organization"`
}

// Retrieve a user from the database in the same organization
func (h *Handlers) Retrieve(c *gin.Context) {
	var err error
	var userID int64
	var search userSearch
	user := session.RetrieveUser(c)
	err = c.ShouldBind(&search)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "error binding user search params")
		return
	}
	if search.ID == "current" {
		userID = user.ID
	} else {
		userID, err = strconv.ParseInt(search.ID, 10, 64)
		if err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest, "error user ID must be current or a valid id")
			return
		}
	}
	dbUser, err := models.Users(qm.Where("id = ? AND organization_id = ?", userID, user.OrganizationID)).One(c, h.db)
	if err != nil {
		h.logger.Printf("error retriving user %s: %v\n", search.ID, err)
		response.NewErrorResponse(c, http.StatusInternalServerError, "error retrieving user from the database")
		return
	} else if dbUser == nil {
		response.NewNotFoundResponse(c)
		return
	}
	response.NewSuccessResponse(c, *dbUser)
}

// CreateAccount creates a user account
func (h *Handlers) CreateAccount(c *gin.Context) {
	var pAccount prospectAccount
	var err error
	var orgIdentifier string
	user := session.RetrieveUser(c)
	err = c.ShouldBind(&pAccount)
	if err != nil {
		h.logger.Printf("error binding user search: %v\n", err)
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}
	if user.ID == 0 {
		// if no user in session, an organization name must have been provided
		orgName := strings.Trim(pAccount.Organization.Name, " ")
		if orgName == "" {
			response.NewErrorResponse(c, http.StatusBadRequest, "non empty organization name is required")
			return
		}
		pAccount.Organization.Name = orgName
		orgIdentifier = orgName
	} else {
		pAccount.Organization.ID = user.OrganizationID
		orgIdentifier = strconv.FormatInt(pAccount.Organization.ID, 10)
	}
	h.logger.Printf("creating new account for Organization %s with email %s\n", orgIdentifier, pAccount.Email)
	newUser, err := h.CreateNewAccount(c, pAccount)
	if err != nil {
		if err == ErrUserAlreadyExists {
			response.NewErrorResponse(c, http.StatusConflict, fmt.Sprintf("There is an already existing user %s in organization %s", pAccount.Email, orgIdentifier))
			return
		}
		h.logger.Printf("error saving user: %v\n", err)
		response.NewErrorResponse(c, http.StatusInternalServerError, "error creating user")
		return
	}
	response.NewSuccessResponse(c, *newUser)
}

// CreateNewAccount creates a new account
func (h *Handlers) CreateNewAccount(ctx context.Context, pAccount prospectAccount) (*models.User, error) {
	var user *models.User
	var err error
	if pAccount.Organization.ID != 0 {
		user, err = h.createUserInOrganization(ctx, pAccount)
	} else {
		user, err = h.createNewOrganization(ctx, pAccount)
	}
	if err == nil {
		// do not allow password to go out of this scope
		user.Password = ""
	}
	return user, err
}

func (h *Handlers) createUserInOrganization(ctx context.Context, pAccount prospectAccount) (*models.User, error) {
	userExists, err := models.Users(qm.Select("id"),
		qm.Where("email = ? AND organization_id = ?", pAccount.Email, pAccount.Organization.ID)).Exists(ctx, h.db)
	if err != nil {
		return nil, err
	}
	if userExists {
		return nil, ErrUserAlreadyExists
	}
	return createUser(ctx, h.db, pAccount)
}

func (h *Handlers) createNewOrganization(ctx context.Context, pAccount prospectAccount) (*models.User, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	newOrganization := models.Organization{
		Name: pAccount.Organization.Name,
	}
	err = newOrganization.Insert(ctx, tx, boil.Infer())
	if err != nil {
		rbErr := tx.Rollback()
		h.logger.Printf("error inserting new Organization. Error %v. Rollback error (might be nil): %v", err, rbErr)
		return nil, err
	}
	pAccount.Organization.ID = newOrganization.ID
	newUser, err := createUser(ctx, tx, pAccount)
	if err != nil {
		rbErr := tx.Rollback()
		h.logger.Printf("error inserting new User. Error %v. Rollback error (might be nil): %v", err, rbErr)
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		h.logger.Printf("error commiting user create transaction. Error %v", err)
		return nil, err
	}
	return newUser, err
}

func createUser(ctx context.Context, exec boil.ContextExecutor, pAccount prospectAccount) (*models.User, error) {
	hashPassword, err := HashPassword(pAccount.Password)
	if err != nil {
		return nil, err
	}
	newUser := &models.User{
		Email:          pAccount.Email,
		Name:           pAccount.Name,
		Password:       hashPassword,
		OrganizationID: pAccount.Organization.ID,
	}
	err = newUser.Insert(ctx, exec, boil.Infer())
	return newUser, err
}
