package models

import (
	u "utils"
)

type Member struct {
	User    User    `gorm:"column:user; not null;ForeignKey:IdUser;" json:"-`
	Channel Channel `gorm:"column:channel; not null;ForeignKey:IdChannel;" json:"-"`
	Role    Role    `gorm:"column:role; ForeignKey:IdRole;" json:"-"`
}

func (member *Member) IsValid() *u.AppError {
	if member.User == (User{}) {
		return u.NewLocAppError("Member.IsValid", "model.member.user.app_error", nil, "")
	}
	if member.Channel == (Channel{}) {
		return u.NewLocAppError("Member.IsValid", "model.member.channel.app_error", nil, "")
	}
	return nil
}