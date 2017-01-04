package models

import (
	"encoding/json"
	"io"
	"strings"
	"unicode/utf8"
)

const (
	DEFAULT_CHANNEL                = "general"
	CHANNEL_DISPLAY_NAME_MAX_RUNES = 64
	CHANNEL_NAME_MAX_LENGTH        = 64
	CHANNEL_DESCRIPTION_MAX_RUNES  = 1024
	CHANNEL_SUBJECT_MAX_RUNES      = 250
)

var (
	CHANNNEL_AVAILABLE_TYPES = []string{"direct", "text", "audio", "video"}
)

type Channel struct {
	ChannelId   uint64 `gorm:"primary_key;column:idChannel;AUTO_INCREMENT" json:"-"`
	WebId       string `gorm:"column:webId;not null;unique" json:"web_id"`
	ChannelName string `gorm:"column:channelName;not null;unique" json:"display_name"`
	Type        string `gorm:"column:type;not null" json:"type"`
	UpdatedAt   int64  `gorm:"column:updatedAt;not null" json:"updated_at"`
	Private     bool   `gorm:"column:private;not null" json:"private"`
	Description string `gorm:"column:desciption" json:"description,omitempty"`
	Subject     string `gorm:"column:subject" json:"subject,omitempty"`
	Avatar      string `gorm:"column:avatar" json:"avatar,omitempty"`
}

func (channel *Channel) toJson() string {
	b, err := json.Marshal(channel)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func channelFromJson(data io.Reader) *Channel {
	decoder := json.NewDecoder(data)
	var channel Channel
	err := decoder.Decode(&channel)
	if err == nil {
		return &channel
	} else {
		return nil
	}
}

func (channel *Channel) etag() string {
	return Etag(channel.WebId, channel.UpdatedAt)
}

func (channel *Channel) isValid() *AppError {

	if len(channel.WebId) != 26 {
		return NewLocAppError("Channel.IsValid", "model.channel.is_valid.id.app_error", nil, "")
	}

	if channel.UpdatedAt == 0 {
		return NewLocAppError("Channel.IsValid", "model.channel.is_valid.update_at.app_error", nil, "id="+channel.WebId)
	}

	if utf8.RuneCountInString(channel.ChannelName) > CHANNEL_DISPLAY_NAME_MAX_RUNES || utf8.RuneCountInString(channel.ChannelName) == 0 {
		return NewLocAppError("Channel.IsValid", "model.channel.is_valid.channel_name.app_error", nil, "id="+channel.WebId)
	}

	if !IsValidChannelIdentifier(channel.ChannelName) {
		return NewLocAppError("Channel.IsValid", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebId)
	}

	if utf8.RuneCountInString(channel.Description) > CHANNEL_DESCRIPTION_MAX_RUNES {
		return NewLocAppError("Channel.IsValid", "model.channel.is_valid.description.app_error", nil, "id="+channel.WebId)
	}

	if utf8.RuneCountInString(channel.Subject) > CHANNEL_SUBJECT_MAX_RUNES {
		return NewLocAppError("Channel.IsValid", "model.channel.is_valid.subject.app_error", nil, "id="+channel.WebId)
	}

	if !StringInArray(channel.Type, CHANNNEL_AVAILABLE_TYPES) {
		return NewLocAppError("Channel.IsValid", "model.channel.is_valid.type.app_error", nil, "id="+channel.WebId)
	}

	return nil
}

func (channel *Channel) preSave() {
	if channel.WebId == "" {
		channel.WebId = NewId()
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

func (channel *Channel) preUpdate() {
	channel.UpdatedAt = GetMillis()
}

func getDMNameFromIds(userId1, userId2 string) string {
	if userId1 > userId2 {
		return userId2 + "__" + userId1
	} else {
		return userId1 + "__" + userId2
	}
}
