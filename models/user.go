// Here is the file which describe the user model. It
// It provides bascis function to manipulate the model.
package model

import (
	// "encoding/json"
	// "fmt"
	// "io"
	// "regexp"
	// "strings"
	// "unicode/utf8"
	"golang.org/x/crypto/bcrypt"

)

const (
	USER_NOTIFY_ALL            = "all"
	USER_NOTIFY_MENTION        = "mention"
	USER_NOTIFY_NONE           = "none"
	DEFAULT_LOCALE             = "en"
	USER_AUTH_SERVICE_EMAIL    = "email"
	USER_AUTH_SERVICE_USERNAME = "username"
	USER_CHANNEL 							 = "general"
)

// Used in mattermost project ... Don't think they are relevant for us.
// AuthData           *string   `json:"auth_data,omitempty"`
// AuthService        string    `json:"auth_service"`
//	LastPictureUpdate  int64     `json:"last_picture_update,omitempty"`
//	 Props              StringMap `json:"props,omitempty"`
// NotifyProps        StringMap `json:"notify_props,omitempty"`
//	MfaActive          bool      `json:"mfa_active,omitempty"`
//	MfaSecret          string    `json:"mfa_secret,omitempty"`
//	CreateAt           int64     `json:"create_at,omitempty"`
//	UpdateAt           int64     `json:"update_at,omitempty"`

// User object.
// - ID: String unique and non null to identify the user
// - Deleted: True if user is deleted
// - Username: Store the user username used to log into the service
// - Password: Hashed password
// - Email: User mail ;)
// - EmailVerified: true if email was verified by user
// - Nickname: Name to use in communication channel (by default : username)
// - First name:
type User struct {
	Id                 string    `json:"id"`
	Deleted            bool     `json:"deleted"`
	Username           string    `json:"username"`
	Password           string    `json:"password,omitempty"`
	Email              string    `json:"email"`
	EmailVerified      bool      `json:"email_verified,omitempty"`
	Nickname           string    `json:"nickname"`
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	Roles              string    `json:"roles"`
	AllowMarketing     bool      `json:"allow_marketing,omitempty"`
	LastPasswordUpdate int64     `json:"last_password_update,omitempty"`
	FailedAttempts     int       `json:"failed_attempts,omitempty"`
	Locale             string    `json:"locale"`
	LastActivityAt     int64     `db:"-" json:"last_activity_at,omitempty"`
}


// HashPassword generates a hash using the bcrypt.GenerateFromPassword
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

// ComparePassword compares the hash
func ComparePassword(hash string, password string) bool {

	if len(password) == 0 || len(hash) == 0 {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}