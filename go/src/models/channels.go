package models

import (
	"encoding/json"
	"io"
	"strings"
	"unicode/utf8"
	u "utils"
)

const (
	defaultChannel             = "general"
	channelDislayNameMaxRunes  = 64
	channelNameMaxLength       = 64
	channelDescriptionMaxRunes = 1024
	channelSubjectMaxRunes     = 250
)

var (
	// ChannelAvailableTypes Used to have knowsledge on type a channel can take
	ChannelAvailableTypes = []string{"direct", "text", "audio", "video"}
)

// Channel type is a model for DB Channel table
type Channel struct {
	ChannelID   uint64 `gorm:"primary_key;column:idChannel;AUTO_INCREMENT" json:"-"`
	WebID       string `gorm:"column:webID;not null;unique" json:"web_id"`
	ChannelName string `gorm:"column:channelName;not null;unique" json:"display_name"`
	Type        string `gorm:"column:type;not null" json:"type"`
	UpdatedAt   int64  `gorm:"column:updatedAt;not null" json:"updated_at"`
	Private     bool   `gorm:"column:private;not null" json:"private"`
	Description string `gorm:"column:desciption" json:"description,omitempty"`
	Subject     string `gorm:"column:subject" json:"subject,omitempty"`
	Avatar      string `gorm:"column:avatar" json:"avatar,omitempty"`
}

// ToJSON Take a channel and convert it into json
func (channel *Channel) ToJSON() string {
	b, err := json.Marshal(channel)
	if err != nil {
		return ""
	}
	return string(b)
}

// ChannelFromJSON try to parse a json object as channel object
func ChannelFromJSON(data io.Reader) *Channel {
	decoder := json.NewDecoder(data)
	var channel Channel
	err := decoder.Decode(&channel)
	if err == nil {
		return &channel
	}
	return nil
}

// Etag is a small function used to create cache ID
func (channel *Channel) Etag() string {
	return Etag(channel.WebID, channel.UpdatedAt)
}

// IsValid check the correctness of a channel object
func (channel *Channel) IsValid() *u.AppError {

	if len(channel.WebID) != 26 {
		return u.NewLocAppError("Channel.IsValid", "model.channel.is_valid.id.app_error", nil, "")
	}

	if channel.UpdatedAt == 0 {
		return u.NewLocAppError("Channel.IsValid", "model.channel.is_valid.update_at.app_error", nil, "id="+channel.WebID)
	}

	if utf8.RuneCountInString(channel.ChannelName) > channelDislayNameMaxRunes || utf8.RuneCountInString(channel.ChannelName) == 0 {
		return u.NewLocAppError("Channel.IsValid", "model.channel.is_valid.channel_name.app_error", nil, "id="+channel.WebID)
	}

	if !IsValidChannelIDentifier(channel.ChannelName) {
		return u.NewLocAppError("Channel.IsValid", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebID)
	}

	if utf8.RuneCountInString(channel.Description) > channelDescriptionMaxRunes {
		return u.NewLocAppError("Channel.IsValid", "model.channel.is_valid.description.app_error", nil, "id="+channel.WebID)
	}

	if utf8.RuneCountInString(channel.Subject) > channelSubjectMaxRunes {
		return u.NewLocAppError("Channel.IsValid", "model.channel.is_valid.subject.app_error", nil, "id="+channel.WebID)
	}

	if !StringInArray(channel.Type, ChannelAvailableTypes) {
		return u.NewLocAppError("Channel.IsValid", "model.channel.is_valid.type.app_error", nil, "id="+channel.WebID)
	}

	return nil
}

// PreSave Is used to add default values to channel before saving it in DB
func (channel *Channel) PreSave() {
	if channel.WebID == "" {
		channel.WebID = NewID()
	}

	channel.ChannelName = strings.ToLower(channel.ChannelName)

	channel.UpdatedAt = GetMillis()

	if channel.Avatar == "" {
		channel.Avatar = "default_channel_avatar.svg"
	}

	if channel.Type == "" {
		channel.Type = "text"
	}

	if channel.Type == "direct" {
		channel.Private = true
	}
}

// PreUpdate Is used to add default values to channel before updating it in DB
func (channel *Channel) PreUpdate() {
	channel.UpdatedAt = GetMillis()
}

// GetDMNameFromIDs Create Direct message name from 2 userIDs
func GetDMNameFromIDs(userID1, userID2 string) string {
	if userID1 > userID2 {
		return userID2 + "__" + userID1
	}
	return userID1 + "__" + userID2

}
