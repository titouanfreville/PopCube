package models

import (
	"encoding/json"
	"io"
	u "utils"
)

type Emoji struct {
	IdEmoji  uint64 `gorm:"primary_key;column:idEmoji;AUTO_INCREMENT" json:"-"`
	Name     string `gorm:"column:name;not null;unique" json:"name"`
	Shortcut string `gorm:"column:shortcut;not null;unique" json:"shortcut"`
	Link     string `gorm:"column:link;not null;unique" json:"link"`
}

func (emoji *Emoji) IsValid() *u.AppError {
	if len(emoji.Name) == 0 || len(emoji.Name) > 64 {
		return u.NewLocAppError("Emoji.IsValid", "model.emoji.name.app_error", nil, "")
	}

	if len(emoji.Shortcut) == 0 || len(emoji.Shortcut) > 20 {
		return u.NewLocAppError("Emoji.IsValid", "model.emoji.shortcut.app_error", nil, "")
	}

	if len(emoji.Link) == 0 {
		return u.NewLocAppError("Emoji.IsValid", "model.emoji.link.app_error", nil, "")
	}

	return nil
}

func (emoji *Emoji) ToJson() string {
	b, err := json.Marshal(emoji)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func EmojiFromJson(data io.Reader) *Emoji {
	decoder := json.NewDecoder(data)
	var emoji Emoji
	err := decoder.Decode(&emoji)
	if err == nil {
		return &emoji
	} else {
		return nil
	}
}

func EmojiListToJson(emojiList []*Emoji) string {
	b, err := json.Marshal(emojiList)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func EmojiListFromJson(data io.Reader) []*Emoji {
	decoder := json.NewDecoder(data)
	var emojiList []*Emoji
	err := decoder.Decode(&emojiList)
	if err == nil {
		return emojiList
	} else {
		return nil
	}
}
