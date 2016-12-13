// Here is the file which describe the user model.
// It provides bascis function to manipulate the model.
package model

import (
	"encoding/json"
	// "fmt"
	"io"
	// "regexp"
	"strings"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

const (
	USER_NOTIFY_ALL            = "all"
	USER_NOTIFY_MENTION        = "mention"
	USER_NOTIFY_NONE           = "none"
	DEFAULT_LOCALE             = "en"
	USER_AUTH_SERVICE_EMAIL    = "email"
	USER_AUTH_SERVICE_USERNAME = "username"
)

var (
	USER_CHANNEL               = []string{"general", "random"}
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
// AllowMarketing     bool      `json:"allow_marketing,omitempty"`

// User object
//
// - ID: String unique and non null to identify the user.
//
// - UpdatedAt: Time of the last update. Used to create tag for browser cache.
//
// - Deleted: True if user is deleted.
//
// - Username: Store the user username used to log into the service.
//
// - Password: Hashed password.
//
// - Email: User mail ;).
//
// - EmailVerified: true if email was verified by user.
//
// - Nickname: Name to use in communication channel (by default : username).
//
// - First name: User true first name.
//
// - Last name: User true last name.
//
// - Roles: User role in the organisation (Owner, Admin, User, Invité).
//
// - LastPasswordUpdate: Date of the last password modification.
//
// - FailedAttemps: Number of fail try to connect to account.
//
// - Locale: User favorite langage.
//
// - Channels: List of the open channel the user is in.
//
// - PrivateChannels: List of the private channel the user is in.
//
// - LastActivityAt: Date && Time of the last activity of the user.
type User struct {
	Id                 string   `json:"id"`
	UpdatedAt          int64    `json:"update_at,omitempty"`
	Deleted            bool     `json:"deleted"`
	Username           string   `json:"username"`
	Password           string   `json:"password,omitempty"`
	Email              string   `json:"email,omitempty"`
	EmailVerified      bool     `json:"email_verified,omitempty"`
	Nickname           string   `json:"nickname"`
	FirstName          string   `json:"first_name"`
	LastName           string   `json:"last_name"`
	Roles              string   `json:"roles,omitempty"`
	LastPasswordUpdate int64    `json:"last_password_update,omitempty"`
	FailedAttempts     int      `json:"failed_attempts,omitempty"`
	Locale             string   `json:"locale"`
	Channels           []string `json:"channels,omitempty"`
	PrivateChannels    []string `json:"private_channels"`
	LastActivityAt     int64    `db:"-" json:"last_activity_at,omitempty"`
}

// IsValid validates the user and returns an error if it isn't configured
// correctly.
func (u *User) IsValid() *AppError {

	if len(u.Id) != 26 {
		return NewLocAppError("User.IsValid", "model.user.is_valid.id.app_error", nil, "")
	}

	if len(u.Username) > 128 || !IsValidUsername(u.Username) {
		return NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+u.Id)
	}

	if len(u.Email) == 0 || len(u.Email) > 128 || !IsValidEmail(u.Email) {
		return NewLocAppError("User.IsValid", "model.user.is_valid.email.app_error", nil, "user_id="+u.Id)
	}

	if utf8.RuneCountInString(u.Nickname) > 64 {
		return NewLocAppError("User.IsValid", "model.user.is_valid.nickname.app_error", nil, "user_id="+u.Id)
	}

	if utf8.RuneCountInString(u.FirstName) > 64 {
		return NewLocAppError("User.IsValid", "model.user.is_valid.first_name.app_error", nil, "user_id="+u.Id)
	}

	if utf8.RuneCountInString(u.LastName) > 64 {
		return NewLocAppError("User.IsValid", "model.user.is_valid.last_name.app_error", nil, "user_id="+u.Id)
	}

	// if u.AuthData != nil && len(*u.AuthData) > 128 {
	// 	return NewLocAppError("User.IsValid", "model.user.is_valid.auth_data.app_error", nil, "user_id="+u.Id)
	// }

	// if u.AuthData != nil && len(*u.AuthData) > 0 && len(u.AuthService) == 0 {
	// 	return NewLocAppError("User.IsValid", "model.user.is_valid.auth_data_type.app_error", nil, "user_id="+u.Id)
	// }

	// && u.AuthData != nil && len(*u.AuthData) > 0
	if len(u.Password) > 0 {
		return NewLocAppError("User.IsValid", "model.user.is_valid.auth_data_pwd.app_error", nil, "user_id="+u.Id)
	}

	return nil
}

// PreSave have to be run before saving user in DB. It will fill necessary information (id, username, etc. ) and hash password
func (u *User) PreSave() {
	if u.Id == "" {
		u.Id = NewId()
	}

	if u.Username == "" {
		u.Username = NewId()
	}

	// if u.AuthData != nil && *u.AuthData == "" {
	// 	u.AuthData = nil
	// }

	u.Username = strings.ToLower(u.Username)
	u.Email = strings.ToLower(u.Email)

	u.UpdatedAt = GetMillis()
	u.LastPasswordUpdate = u.UpdatedAt

	if u.Locale == "" {
		u.Locale = DEFAULT_LOCALE
	}

	if u.Channels == nil || len(u.Channels) == 0 {
		u.Channels = USER_CHANNEL
	}

	if len(u.Password) > 0 {
		u.Password = HashPassword(u.Password)
	}
}

// PreSave will set the Id and Username if missing.  It will also fill
// in the CreateAt, UpdateAt times.  It will also hash the password.  It should
// be run before saving the user to the db.
// PreUpdate should be run before updating the user in the db.
func (u *User) PreUpdate() {
	u.Username = strings.ToLower(u.Username)
	u.Email = strings.ToLower(u.Email)
	u.UpdatedAt = GetMillis()
}

// ToJson convert a User to a json string
func (u *User) ToJson() string {
	b, err := json.Marshal(u)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

// UserFromJson will decode the input and return a User
func UserFromJson(data io.Reader) *User {
	decoder := json.NewDecoder(data)
	var user User
	err := decoder.Decode(&user)
	if err == nil {
		return &user
	} else {
		return nil
	}
}

// Generate a valid strong etag so the browser can cache the results
func (u *User) Etag(showFullName, showEmail bool) string {
	return Etag(u.Id, u.UpdatedAt, showFullName, showEmail)
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
