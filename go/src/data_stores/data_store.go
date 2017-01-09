// Created by Titouan FREVILLE <titouanfreville@gmail.com>
//
// Inspired by mattermost project
/*
	Package Data Stores.
	This package implements the basics databases communication functions used by PopCube chat api.

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
package data_stores

import (
	// l4g "github.com/alecthomas/log4go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"models"
	. "utils"
	// "time"
)

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

func (ds *DataStore) initConnection(user string, dbname string, password string) {
	connection_chain := user + ":" + password + "@(database:3306)/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connection_chain)
	// db.AutoMigrate( &models.Channel{}, &models.Emoji{}, &models.Folder{},
	// 	&models.Member{}, &models.Message{}, &models.Organisation{}, &models.Parameter{},
	// 	&models.Role{}, &models.User{})
	db.AutoMigrate(&models.Avatar{}, &models.Emoji{}, &models.Organisation{})
	ds.Db = db
	ds.Err = err
}

func (ds *DataStore) closeConnection() {
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

// Organisation is unique in the database. So they are no use of providing an user to get.
// Delete is useless as we will down the docker stack in case an organisation leace.
type OrganisationStore interface {
	Save(organisation *models.Organisation, ds DataStore) *AppError
	Update(organisation *models.Organisation, new_organisation *models.Organisation, ds DataStore) *AppError
	Get(ds DataStore) *models.Organisation
}

type AvatarStore interface {
	Save(avatar *models.Avatar, ds DataStore) *AppError
	Update(avatar *models.Avatar, new_avatar *models.Avatar, ds DataStore) *AppError
	GetByName(avatarName string, ds DataStore) *models.Avatar
	GetByLink(avatarLink string, ds DataStore) *models.Avatar
	GetAll(ds DataStore) *models.Avatar
	Delete(avatar *models.Avatar, ds DataStore) *AppError
}

type EmojiStore interface {
	Save(emoji *models.Avatar, ds DataStore) *AppError
	Update(emoji *models.Avatar, new_emoji *models.Avatar, ds DataStore) *AppError
	GetByName(emojiName string, ds DataStore) *models.Avatar
	GetByShortcut(emojiShortcut string, ds DataStore) *models.Avatar
	GetByLink(emojiLink string, ds DataStore) *models.Avatar
	GetAll(ds DataStore) *models.Avatar
	Delete(emoji *models.Avatar, ds DataStore) *AppError
}

// type UserStore interface {
// 	Save(user *models.User) StoreChannel
// 	Update(user *models.User, allowRoleUpdate bool) StoreChannel
// 	UpdateLastPictureUpdate(userId string) StoreChannel
// 	UpdateUpdateAt(userId string) StoreChannel
// 	UpdatePassword(userId, newPassword string) StoreChannel
// 	Get(id string) StoreChannel
// 	GetAll() StoreChannel
// 	InvalidateProfilesInChannelCacheByUser(userId string)
// 	InvalidateProfilesInChannelCache(channelId string)
// 	GetProfilesInChannel(channelId string, offset int, limit int, allowFromCache bool) StoreChannel
// 	GetProfilesNotInChannel(teamId string, channelId string, offset int, limit int) StoreChannel
// 	GetProfilesByUsernames(usernames []string, teamId string) StoreChannel
// 	GetAllProfiles(offset int, limit int) StoreChannel
// 	GetProfiles(teamId string, offset int, limit int) StoreChannel
// 	GetProfileByIds(userId []string, allowFromCache bool) StoreChannel
// 	InvalidatProfileCacheForUser(userId string)
// 	GetByEmail(email string) StoreChannel
// 	GetByUsername(username string) StoreChannel
// 	GetForLogin(loginId string, allowSignInWithUsername, allowSignInWithEmail, ldapEnabled bool) StoreChannel
// 	VerifyEmail(userId string) StoreChannel
// 	GetEtagForAllProfiles() StoreChannel
// 	GetEtagForProfiles(teamId string) StoreChannel
// 	UpdateFailedPasswordAttempts(userId string, attempts int) StoreChannel
// 	GetTotalUsersCount() StoreChannel
// 	GetSystemAdminProfiles() StoreChannel
// 	PermanentDelete(userId string) StoreChannel
// 	AnalyticsUniqueUserCount(teamId string) StoreChannel
// 	GetUnreadCount(userId string) StoreChannel
// 	GetUnreadCountForChannel(userId string, channelId string) StoreChannel
// 	GetRecentlyActiveUsersForTeam(teamId string) StoreChannel
// 	Search(teamId string, term string, options map[string]bool) StoreChannel
// 	SearchInChannel(channelId string, term string, options map[string]bool) StoreChannel
// 	SearchNotInChannel(teamId string, channelId string, term string, options map[string]bool) StoreChannel
// }
