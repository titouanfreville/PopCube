package model

import (
	"encoding/json"
	"io"
)

type Emoji struct {
	Name     string `json:"name"`
	ShortCut string `json:"shortcut"`
	Link     string `json:"link"`
}

func (emoji *Emoji) isValid() *AppError {
	if len(emoji.Name) == 0 || len(emoji.Name) > 64 {
		return NewLocAppError("Emoji.IsValid", "model.emoji.name.app_error", nil, "")
	}

	if len(emoji.ShortCut) == 0 || len(emoji.ShortCut) > 20 {
		return NewLocAppError("Emoji.IsValid", "model.emoji.shortcut.app_error", nil, "")
	}

	if len(emoji.Link) == 0 {
		return NewLocAppError("Emoji.IsValid", "model.emoji.link.app_error", nil, "")
	}

	return nil
}

func (emoji *Emoji) toJson() string {
	b, err := json.Marshal(emoji)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func emojiFromJson(data io.Reader) *Emoji {
	decoder := json.NewDecoder(data)
	var emoji Emoji
	err := decoder.Decode(&emoji)
	if err == nil {
		return &emoji
	} else {
		return nil
	}
}

func emojiListToJson(emojiList []*Emoji) string {
	b, err := json.Marshal(emojiList)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func emojiListFromJson(data io.Reader) []*Emoji {
	decoder := json.NewDecoder(data)
	var emojiList []*Emoji
	err := decoder.Decode(&emojiList)
	if err == nil {
		return emojiList
	} else {
		return nil
	}
}
