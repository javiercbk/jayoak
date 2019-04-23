package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// User is the logged user
type User struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	OrganizationID int64  `json:"organizationId"`
}

const userKey = "user"

// RetrieveUser retrieves the user currently logged or nil if no session.
func RetrieveUser(c *gin.Context) User {
	session := sessions.Default(c)
	return session.Get(userKey).(User)
}

// SaveUser saves the user to in session
func SaveUser(c *gin.Context, user User) error {
	session := sessions.Default(c)
	session.Set(userKey, user)
	return session.Save()
}
