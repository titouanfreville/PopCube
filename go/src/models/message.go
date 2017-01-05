package models

import (
	"encoding/json"
	"io"
)

type Message struct {
	MessageId uint64 `gorm:"primary_key;column:idMessage;AUTO_INCREMENT" json:"-"`
	Name      string `gorm:"column:name;not null;unique" json:"name"`
	Shortcut  string `gorm:"column:shortcut;not null;unique" json:"shortcut"`
	Link      string `gorm:"column:link;not null;unique" json:"link"`
}

func (message *Message) isValid() *AppError {
	if len(message.Name) == 0 || len(message.Name) > 64 {
		return NewLocAppError("Message.IsValid", "model.message.name.app_error", nil, "")
	}

	if len(message.Shortcut) == 0 || len(message.Shortcut) > 20 {
		return NewLocAppError("Message.IsValid", "model.message.shortcut.app_error", nil, "")
	}

	if len(message.Link) == 0 {
		return NewLocAppError("Message.IsValid", "model.message.link.app_error", nil, "")
	}

	return nil
}

func (message *Message) toJson() string {
	b, err := json.Marshal(message)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func messageFromJson(data io.Reader) *Message {
	decoder := json.NewDecoder(data)
	var message Message
	err := decoder.Decode(&message)
	if err == nil {
		return &message
	} else {
		return nil
	}
}

func messageListToJson(messageList []*Message) string {
	b, err := json.Marshal(messageList)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func messageListFromJson(data io.Reader) []*Message {
	decoder := json.NewDecoder(data)
	var messageList []*Message
	err := decoder.Decode(&messageList)
	if err == nil {
		return messageList
	} else {
		return nil
	}
}
