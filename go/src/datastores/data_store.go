/*Package datastores implements the basics databases communication functions used by PopCube chat api.

Stores

The following is a list of stores described:
	Avatar: Contain all informations for avatar management
	Channel: Contain all informations for channel management
	Emojis: Contain all informations for emojis management
	Organisation: Contain all informations for organisation management
	Parameter: Contain all informations for parmeters management
	Role: Contain all informations for roles management
	User: Contain all informations for users management
*/
// Created by Titouan FREVILLE <titouanfreville@gmail.com>
//
// Inspired by mattermost project
package datastores

import (
	// Importing sql driver. They are used by gorm package and used by default from blank.
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"models"
	u "utils"
)

// DataStore Struct to manage Db knowledge
type DataStore struct {
	Db  *gorm.DB
	Err error
}

// type StoreResult struct {
// 	Data interface{}
// 	Err  *models.AppError
// }

// type StoreChannel chan StoreResult

// func Must(sc StoreChannel) interface{} {
// 	r := <-sc
// 	if r.Err != nil {
// 		l4g.Close()
// 		time.Sleep(time.Second)
// 		panic(r.Err)
// 	}

// 	return r.Data
// }

// InitConnection init Database connection && database models
func (ds *DataStore) InitConnection(user string, dbname string, password string) {
	connectionChain := user + ":" + password + "@(database:3306)/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connectionChain)
	// db.AutoMigrate( &models.Channel{}, &models.Emoji{}, &models.Folder{},
	// 	&models.Member{}, &models.Message{}, &models.Organisation{}, &models.Parameter{},
	// 	&models.Role{}, &models.User{})
	db.AutoMigrate(&models.Avatar{}, &models.Emoji{}, &models.Organisation{})
	ds.Db = db
	ds.Err = err
}

// CloseConnection close database connection
func (ds *DataStore) CloseConnection() {
	db := *ds.Db
	defer db.Close()
	ds.Db = &gorm.DB{}
}

// type Store interface {
// 	Team() TeamStore
// 	Channel() ChannelStore
// 	Post() PostStore
// 	User() UserStore
// 	Audit() AuditStore
// 	Compliance() ComplianceStore
// 	Session() SessionStore
// 	OAuth() OAuthStore
// 	System() SystemStore
// 	Webhook() WebhookStore
// 	Command() CommandStore
// 	Preference() PreferenceStore
// 	License() LicenseStore
// 	PasswordRecovery() PasswordRecoveryStore
// 	Emoji() EmojiStore
// 	Status() StatusStore
// 	FileInfo() FileInfoStore
// 	Reaction() ReactionStore
// 	MarkSystemRanUnitTests()
// 	Close()
// 	DropAllTables()
// 	TotalMasterDbConnections() int
// 	TotalReadDbConnections() int
// }

/*OrganisationStore interface the organisation communication
Organisation is unique in the database. So they are no use of providing an user to get.
Delete is useless as we will down the docker stack in case an organisation leace.
*/
type OrganisationStore interface {
	Save(organisation *models.Organisation, ds DataStore) *u.AppError
	Update(organisation *models.Organisation, newOrganisation *models.Organisation, ds DataStore) *u.AppError
	Get(ds DataStore) *models.Organisation
}

/*AvatarStore interface the avatar communication */
type AvatarStore interface {
	Save(avatar *models.Avatar, ds DataStore) *u.AppError
	Update(avatar *models.Avatar, newAvatar *models.Avatar, ds DataStore) *u.AppError
	GetByName(avatarName string, ds DataStore) *models.Avatar
	GetByLink(avatarLink string, ds DataStore) *models.Avatar
	GetAll(ds DataStore) *models.Avatar
	Delete(avatar *models.Avatar, ds DataStore) *u.AppError
}

/*EmojiStore interface the emoji communication*/
type EmojiStore interface {
	Save(emoji *models.Avatar, ds DataStore) *u.AppError
	Update(emoji *models.Avatar, newEmoji *models.Avatar, ds DataStore) *u.AppError
	GetByName(emojiName string, ds DataStore) *models.Avatar
	GetByShortcut(emojiShortcut string, ds DataStore) *models.Avatar
	GetByLink(emojiLink string, ds DataStore) *models.Avatar
	GetAll(ds DataStore) *models.Avatar
	Delete(emoji *models.Avatar, ds DataStore) *u.AppError
}

// type UserStore interface {
// 	Save(user *models.User) StoreChannel
// 	Update(user *models.User, allowRoleUpdate bool) StoreChannel
// 	UpdateLastPictureUpdate(userID string) StoreChannel
// 	UpdateUpdateAt(userID string) StoreChannel
// 	UpdatePassword(userID, newPassword string) StoreChannel
// 	Get(id string) StoreChannel
// 	GetAll() StoreChannel
// 	InvalidateProfilesInChannelCacheByUser(userID string)
// 	InvalidateProfilesInChannelCache(channelID string)
// 	GetProfilesInChannel(channelID string, offset int, limit int, allowFromCache bool) StoreChannel
// 	GetProfilesNotInChannel(teamID string, channelID string, offset int, limit int) StoreChannel
// 	GetProfilesByUsernames(usernames []string, teamID string) StoreChannel
// 	GetAllProfiles(offset int, limit int) StoreChannel
// 	GetProfiles(teamID string, offset int, limit int) StoreChannel
// 	GetProfileByIDs(userID []string, allowFromCache bool) StoreChannel
// 	InvalidatProfileCacheForUser(userID string)
// 	GetByEmail(email string) StoreChannel
// 	GetByUsername(username string) StoreChannel
// 	GetForLogin(loginID string, allowSignInWithUsername, allowSignInWithEmail, ldapEnabled bool) StoreChannel
// 	VerifyEmail(userID string) StoreChannel
// 	GetEtagForAllProfiles() StoreChannel
// 	GetEtagForProfiles(teamID string) StoreChannel
// 	UpdateFailedPasswordAttempts(userID string, attempts int) StoreChannel
// 	GetTotalUsersCount() StoreChannel
// 	GetSystemAdminProfiles() StoreChannel
// 	PermanentDelete(userID string) StoreChannel
// 	AnalyticsUniqueUserCount(teamID string) StoreChannel
// 	GetUnreadCount(userID string) StoreChannel
// 	GetUnreadCountForChannel(userID string, channelID string) StoreChannel
// 	GetRecentlyActiveUsersForTeam(teamID string) StoreChannel
// 	Search(teamID string, term string, options map[string]bool) StoreChannel
// 	SearchInChannel(channelID string, term string, options map[string]bool) StoreChannel
// 	SearchNotInChannel(teamID string, channelID string, term string, options map[string]bool) StoreChannel
// }
