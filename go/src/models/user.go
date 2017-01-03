// Here is the file which describe the user model.
// It provwebIdes bascis function to manipulate the model.
package models

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

const (
	user_NOTIFY_ALL            = "all"
	user_NOTIFY_MENTION        = "mention"
	user_NOTIFY_NONE           = "none"
	DEFAULT_locale             = "en"
	user_AUTH_SERVICE_email    = "email"
	user_AUTH_SERVICE_username = "username"
)

var (
	user_CHANNEL        = []string{"general", "random"}
	restrictedUsernames = []string{
		"all",
		"channel",
		"popcubebot",
		"here",
	}
	validUsernameChars = regexp.MustCompile(`^[a-z0-9\.\-_]+$`)
)

// Used in mattermost project ... Don't think they are relevant for us.
//	MfaActive          bool      `json:"mfa_active,omitempty"`
//	MfaSecret          string    `json:"mfa_secret,omitempty"`

// user object
//
// - webwebId: String unique and non null to webIdentify the user on application services. - REQUIRED
//
// - username: Store the user username used to log into the service. - REQUIRED
//
// - email: user mail ;). - REQUIRED
//
// - emailVerified: true if email was verified by user. - REQUIRED
//
// - updatedAt: Time of the last update. Used to create tag for browser cache. - REQUIRED
//
// - deleted: True if user is deleted. - REQUIRED
//
// - password: Hashed password. - REQUIRED
//
// - lastpasswordUpdate: Date of the last password modification. - REQUIRED
//
// - failedAttemps: Number of fail try to connect to account. - REQUIRED
//
// - locale: user favorite langage. - REQUIRED
//
// - role : int referencing a user role existing in the database. - REQUIRED
//
// - nickname: Name to use in communication channel (by default : username).
//
// - first name: user true first name.
//
// - last name: user true last name.
//
// - lastActivityAt: Date && Time of the last activity of the user.
type User struct {
	WebId              string `json:"webId"`
	UpdatedAt          int64  `json:"update_at,omitempty"`
	Deleted            bool   `json:"deleted"`
	Username           string `json:"username"`
	Password           string `json:"password,omitempty"`
	Email              string `json:"email,omitempty"`
	EmailVerified      bool   `json:"email_verified,omitempty"`
	Nickname           string `json:"nickname"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Avatar             string `json:"avatar"`
	Role               int64  `json:"roles,omitempty"`
	LastPasswordUpdate int64  `json:"last_password_update,omitempty"`
	FailedAttempts     int    `json:"failed_attempts,omitempty"`
	Locale             string `json:"locale"`
	LastActivityAt     int64  `db:"-" json:"last_activity_at,omitempty"`
}

// isValid valwebIdates the user and returns an error if it isn't configured
// correctly.
func (u *User) isValid() *AppError {

	if len(u.WebId) != 26 {
		return NewLocAppError("user.isValid", "model.user.is_valid.WebId.app_error", nil, "")
	}

	if !isValidUsername(u.Username) {
		return NewLocAppError("user.isValid", "model.user.is_valid.Username.app_error", nil, "user_webId="+u.WebId)
	}

	if len(u.Email) == 0 || len(u.Email) > 128 || !IsValidEmail(u.Email) {
		return NewLocAppError("user.isValid", "model.user.is_valid.Email.app_error", nil, "user_webId="+u.WebId)
	}

	if utf8.RuneCountInString(u.Nickname) > 64 {
		return NewLocAppError("user.isValid", "model.user.is_valid.Nickname.app_error", nil, "user_webId="+u.WebId)
	}

	if utf8.RuneCountInString(u.FirstName) > 64 {
		return NewLocAppError("user.isValid", "model.user.is_valid.first_name.app_error", nil, "user_webId="+u.WebId)
	}

	if utf8.RuneCountInString(u.LastName) > 64 {
		return NewLocAppError("user.isValid", "model.user.is_valid.last_name.app_error", nil, "user_webId="+u.WebId)
	}

	if len(u.Password) == 0 {
		return NewLocAppError("user.isValid", "model.user.is_valid.auth_data_pwd.app_error", nil, "user_webId="+u.WebId)
	}

	return nil
}

// preSave have to be run before saving user in DB. It will fill necessary information (webId, username, etc. ) and hash password
func (u *User) preSave() {
	if u.WebId == "" {
		u.WebId = NewId()
	}

	if u.Username == "" {
		u.Username = NewId()
	}

	u.Username = strings.ToLower(u.Username)
	u.Email = strings.ToLower(u.Email)

	u.UpdatedAt = GetMillis()
	u.LastPasswordUpdate = u.UpdatedAt

	if u.Locale == "" {
		u.Locale = DEFAULT_locale
	}

	if len(u.Password) > 0 {
		u.Password = hashPassword(u.Password)
	}
}

// preSave will set the webId and username if missing.  It will also fill
// in the CreateAt, UpdateAt times.  It will also hash the password.  It should
// be run before saving the user to the db.
// PreUpdate should be run before updating the user in the db.
func (u *User) preUpdate() {
	u.Username = strings.ToLower(u.Username)
	u.Email = strings.ToLower(u.Email)
	u.UpdatedAt = GetMillis()
}

// ToJson convert a user to a json string
func (u *User) toJson() string {
	b, err := json.Marshal(u)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

// userFromJson will decode the input and return a user
func userFromJson(data io.Reader) *User {
	decoder := json.NewDecoder(data)
	var user User
	err := decoder.Decode(&user)
	if err == nil {
		return &user
	} else {
		return nil
	}
}

func isValidUsername(u string) bool {
	if len(u) == 0 || len(u) > 64 {
		return false
	}

	if !validUsernameChars.MatchString(u) {
		return false
	}

	for _, restrictedUsername := range restrictedUsernames {
		if u == restrictedUsername {
			return false
		}
	}

	return true
}

// Generate a valwebId strong etag so the browser can cache the results
func (u *User) etag(showFullName, showemail bool) string {
	return Etag(u.WebId, u.UpdatedAt, showFullName, showemail)
}

// Get full name of the user
func (u *User) getFullName() string {
	if u.LastName == "" {
		return u.FirstName
	}
	if u.FirstName == "" {
		return u.LastName
	}
	return u.FirstName + " " + u.LastName
}

// Get full name of the user
func (u *User) getDisplayName() string {
	if u.Nickname != "" {
		return u.Nickname
	}
	if u.getFullName() != "" {
		return u.getFullName()
	}
	return u.Username
}

// hashpassword generates a hash using the bcrypt.GenerateFrompassword
func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

// Comparepassword compares the hash
func comparePassword(hash string, password string) bool {

	if len(password) == 0 || len(hash) == 0 {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Transform user name to meet requirement
func cleanUsername(s string) string {
	s = strings.ToLower(strings.Replace(s, " ", "-", -1))

	for _, value := range reservedName {
		if s == value {
			s = strings.Replace(s, value, "", -1)
		}
	}

	s = strings.TrimSpace(s)

	for _, c := range s {
		char := fmt.Sprintf("%c", c)
		if !validUsernameChars.MatchString(char) {
			s = strings.Replace(s, char, "-", -1)
		}
	}

	if !isValidUsername(s) {
		s = "a" + NewId()
	}

	return s
}
