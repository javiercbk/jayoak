package user

import (
	"context"
	"testing"

	"github.com/javiercbk/jayoak/models"
	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/crypto/bcrypt"

	"github.com/javiercbk/jayoak/http/session"
	testHelper "github.com/javiercbk/jayoak/testing"
)

func TestMain(m *testing.M) {
	testHelper.InitializeDB(m)
}

const sessionUserEmail = "test@test.com"

func setUp(ctx context.Context) (*Handlers, session.User, error) {
	nullLogger := testHelper.NullLogger()
	var sessUser session.User
	db, err := testHelper.DB()
	if err != nil {
		return nil, sessUser, err
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, sessUser, err
	}
	org := &models.Organization{
		Name: "Test",
	}
	err = org.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return nil, sessUser, err
	}
	user := &models.User{
		Name:           "Test",
		Email:          sessionUserEmail,
		Password:       "unhashedPassword",
		OrganizationID: org.ID,
	}
	err = user.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return nil, sessUser, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, sessUser, err
	}
	sessUser.ID = user.ID
	sessUser.Name = user.Name
	sessUser.OrganizationID = user.OrganizationID
	h := NewHandlers(nullLogger, db)
	return h, sessUser, nil
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	h, user, err := setUp(ctx)
	if err != nil {
		t.Fatalf("error setting up test %v\n", err)
	}
	testTable := []struct {
		description string
		pAccount    prospectAccount
		err         error
	}{
		{
			description: "create new user in new organization",
			pAccount: prospectAccount{
				Email:    "test@test.com",
				Name:     "Test",
				Password: "nicePassword",
				Organization: prospectOrganization{
					Name: "niceOrganization",
				},
			},
			err: nil,
		}, {
			description: "create duplicated user with in new organization",
			pAccount: prospectAccount{
				Email:    "test@test.com",
				Name:     "Test",
				Password: "nicePassword",
				Organization: prospectOrganization{
					Name: "niceOrganization",
				},
			},
			err: nil,
		}, {
			description: "create new user with in existing organization",
			pAccount: prospectAccount{
				Email:    "test1@test.com",
				Name:     "Test",
				Password: "nicePassword",
				Organization: prospectOrganization{
					ID: user.OrganizationID,
				},
			},
			err: nil,
		}, {
			description: "error on creating duplicated user in existing organization",
			pAccount: prospectAccount{
				Email:    sessionUserEmail,
				Name:     "Test2",
				Password: "nicePassword",
				Organization: prospectOrganization{
					ID: user.OrganizationID,
				},
			},
			err: ErrUserAlreadyExists,
		},
	}
	for _, test := range testTable {
		newUser, err := h.CreateNewAccount(ctx, test.pAccount)
		if err != test.err {
			t.Fatalf("'%s': error, expected %v to be %v", test.description, err, test.err)
		} else if err == nil {
			if newUser.ID <= 0 {
				t.Fatalf("'%s': error, expected user id %d to greater than zero", test.description, newUser.ID)
			}
			if newUser.OrganizationID <= 0 {
				t.Fatalf("'%s': error, expected organization id %d to greater than zero", test.description, newUser.OrganizationID)
			}
			if newUser.Password != "" && bcrypt.CompareHashAndPassword([]byte(test.pAccount.Password), []byte(newUser.Password)) != nil {
				t.Fatalf("'%s': error, bcrypt hash do not match password %s", test.description, newUser.Password)
			}
		}
	}
}
