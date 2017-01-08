package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
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

	Convey("Testing isValid function", t, func() {
		Convey("Given a correct member. Should be validated", func() {
			member := Member{
				User:    user_test,
				Channel: channel_test,
			}
			So(member.isValid(), ShouldBeNil)
			So(member.isValid(), ShouldNotResemble, NewLocAppError("Member.IsValid", "model.member.user.app_error", nil, ""))
			So(member.isValid(), ShouldNotResemble, NewLocAppError("Member.IsValid", "model.member.channel.app_error", nil, ""))
		})

		Convey("Given incorrect member. Should be refused", func() {
			empty := Member{}
			member := Member{
				User:    user_test,
				Channel: channel_test,
			}
			member.User = User{}
			Convey("Empty member or member without User should return User error", func() {
				So(member.isValid(), ShouldResemble, NewLocAppError("Member.IsValid", "model.member.user.app_error", nil, ""))
				So(member.isValid(), ShouldNotResemble, NewLocAppError("Member.IsValid", "model.member.channel.app_error", nil, ""))
				So(empty.isValid(), ShouldResemble, NewLocAppError("Member.IsValid", "model.member.user.app_error", nil, ""))
				So(empty.isValid(), ShouldNotResemble, NewLocAppError("Member.IsValid", "model.member.channel.app_error", nil, ""))
			})

			member.User = user_test
			member.Channel = Channel{}
			Convey("Empty link should result in link error", func() {
				So(member.isValid(), ShouldNotResemble, NewLocAppError("Member.IsValid", "model.member.user.app_error", nil, ""))
				So(member.isValid(), ShouldResemble, NewLocAppError("Member.IsValid", "model.member.channel.app_error", nil, ""))
			})
		})
	})
}
