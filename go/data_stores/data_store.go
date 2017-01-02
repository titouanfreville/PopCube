package data_store

import (
	"time"

	"../models/"
	//l4g "github.com/alecthomas/log4go"
)

type StoreResult struct {
	Data interface{}
	Err  *model.AppError
}

type StoreChannel chan StoreResult

func Must(sc StoreChannel) interface{} {
	r := <-sc
	if r.Err != nil {
		//l4g.Close()
		time.Sleep(time.Second)
		panic(r.Err)
	}

	return r.Data
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

// type UserStore interface {
// 	Save(user *model.User) StoreChannel
// 	Update(user *model.User, allowRoleUpdate bool) StoreChannel
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
