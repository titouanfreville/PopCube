package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	u "utils"
)

func TestMemberModel(t *testing.T) {
	user_test := User{
		WebId:              NewId(),
		UpdatedAt:          10,
		Deleted:            true,
		Username:           "l",
		Password:           "test",
		Email:              "test@popcube.fr",
		EmailVerified:      true,
		NickName:           "NickName",
		FirstName:          "Test",
		LastName:           "L",
		Role:               OWNER,
		LastPasswordUpdate: 20,
		FailedAttempts:     1,
		Locale:             "vi",
	}

	channel_test := Channel{
		WebId:       NewId(),
		ChannelName: "electra",
		UpdatedAt:   GetMillis(),
		Type:        "audio",
		Private:     false,
		Description: "Testing channel description :O",
		Subject:     "Sujet",
		Avatar:      "jesuiscool.svg",
	}

	Convey("Testing IsValid function", t, func() {
		Convey("Given a correct member. Should be validated", func() {
			member := Member{
				User:    user_test,
				Channel: channel_test,
			}
			So(member.IsValid(), ShouldBeNil)
			So(member.IsValid(), ShouldNotResemble, u.NewLocAppError("Member.IsValid", "model.member.user.app_error", nil, ""))
			So(member.IsValid(), ShouldNotResemble, u.NewLocAppError("Member.IsValid", "model.member.channel.app_error", nil, ""))
		})

		Convey("Given incorrect member. Should be refused", func() {
			empty := Member{}
			member := Member{
				User:    user_test,
				Channel: channel_test,
			}
			member.User = User{}
			Convey("Empty member or member without User should return User error", func() {
				So(member.IsValid(), ShouldResemble, u.NewLocAppError("Member.IsValid", "model.member.user.app_error", nil, ""))
				So(member.IsValid(), ShouldNotResemble, u.NewLocAppError("Member.IsValid", "model.member.channel.app_error", nil, ""))
				So(empty.IsValid(), ShouldResemble, u.NewLocAppError("Member.IsValid", "model.member.user.app_error", nil, ""))
				So(empty.IsValid(), ShouldNotResemble, u.NewLocAppError("Member.IsValid", "model.member.channel.app_error", nil, ""))
			})

			member.User = user_test
			member.Channel = Channel{}
			Convey("Empty link should result in link error", func() {
				So(member.IsValid(), ShouldNotResemble, u.NewLocAppError("Member.IsValid", "model.member.user.app_error", nil, ""))
				So(member.IsValid(), ShouldResemble, u.NewLocAppError("Member.IsValid", "model.member.channel.app_error", nil, ""))
			})
		})
	})
}
