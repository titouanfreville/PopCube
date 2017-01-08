package models

import (
	"encoding/json"
	"io"
	u "utils"
)

type Message struct {
	IdMessage uint64  `gorm:"primary_key;column:idMessage;AUTO_INCREMENT" json:"-"`
	Date      int64   `gorm:"column:date;not null" json:"date"`
	Content   string  `gorm:"column:content;type:longtext" json:"content"`
	Creator   User    `gorm:"column:creator; not null;ForeignKey:IdUser;" json:"-"`
	Channel   Channel `gorm:"column:channel; not null;ForeignKey:IdChannel;" json:"-"`
}

// IsValid function is used to check that the provided message correspond to the message model. It has to be use before tring to store it in the db.
func (message *Message) IsValid() *u.AppError {
	if message.Date == 0 {
		return u.NewLocAppError("Message.IsValid", "model.message.date.app_error", nil, "")
	}
	if message.Creator == (User{}) {
		return u.NewLocAppError("Message.IsValid", "model.message.creator.app_error", nil, "")
	}
	if message.Channel == (Channel{}) {
		return u.NewLocAppError("Message.IsValid", "model.message.channel.app_error", nil, "")
	}

	return nil
}

// PreSave need to be called before saving a new or an updated mesage in the DB so it will have good time store.
func (message *Message) PreSave() {
	message.Date = GetMillis()
}

// ToJson take the message object and transfor it into a json object for api usage.
func (message *Message) ToJson() string {
	b, err := json.Marshal(message)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func MessageFromJson(data io.Reader) *Message {
	decoder := json.NewDecoder(data)
	var message Message
	err := decoder.Decode(&message)
	if err == nil {
		return &message
	} else {
		return nil
	}
}

func MessageListToJson(messageList []*Message) string {
	b, err := json.Marshal(messageList)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func MessageListFromJson(data io.Reader) []*Message {
	decoder := json.NewDecoder(data)
	var messageList []*Message
	err := decoder.Decode(&messageList)
	if err == nil {
		return messageList
	} else {
		return nil
	}
}
