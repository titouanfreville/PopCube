package models

import (
	u "utils"
)

// Member describe the associtive table member between USER, CHANNEL, and ROLE
type Member struct {
	User    User    `gorm:"column:user; not null;ForeignKey:IDUser;" json:"-`
	Channel Channel `gorm:"column:channel; not null;ForeignKey:IDChannel;" json:"-"`
	Role    Role    `gorm:"column:role; ForeignKey:IDRole;" json:"-"`
}

// IsValid check validity of member object
func (member *Member) IsValid() *u.AppError {
	if member.User == (User{}) {
		return u.NewLocAppError("Member.IsValid", "model.member.user.app_error", nil, "")
	}
	if member.Channel == (Channel{}) {
		return u.NewLocAppError("Member.IsValid", "model.member.channel.app_error", nil, "")
	}
	return nil
}
