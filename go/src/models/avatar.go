package models

import (
	"encoding/json"
	"io"
	u "utils"
)

type Avatar struct {
	IdAvatar uint64 `gorm:"primary_key;column:idAvatar;AUTO_INCREMENT" json:"-"`
	Name     string `gorm:"column:name;not null;unique" json:"name"`
	Link     string `gorm:"column:link;not null;unique" json:"link"`
}

func (avatar *Avatar) IsValid() *u.AppError {
	if len(avatar.Name) == 0 || len(avatar.Name) > 64 {
		return u.NewLocAppError("Avatar.IsValid", "model.avatar.name.app_error", nil, "")
	}

	if len(avatar.Link) == 0 {
		return u.NewLocAppError("Avatar.IsValid", "model.avatar.link.app_error", nil, "")
	}

	return nil
}

func (avatar *Avatar) ToJson() string {
	b, err := json.Marshal(avatar)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func AvatarFromJson(data io.Reader) *Avatar {
	decoder := json.NewDecoder(data)
	var avatar Avatar
	err := decoder.Decode(&avatar)
	if err == nil {
		return &avatar
	} else {
		return nil
	}
}

func AvatarListToJson(avatarList []*Avatar) string {
	b, err := json.Marshal(avatarList)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func AvatarListFromJson(data io.Reader) []*Avatar {
	decoder := json.NewDecoder(data)
	var avatarList []*Avatar
	err := decoder.Decode(&avatarList)
	if err == nil {
		return avatarList
	} else {
		return nil
	}
}
