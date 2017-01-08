package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
	u "utils"
)

func TestMessageModel(t *testing.T) {
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

	Convey("Testing json vs message conversions", t, func() {
		Convey("Given a message", func() {
			message := Message{
				Date:    GetMillis(),
				Content: "Message test",
				Creator: user_test,
				Channel: channel_test,
			}

			Convey("Converting message to json then json to message should provide same message information (empty fields if ignore in JSON).", func() {
				json := message.ToJson()
				test_message := MessageFromJson(strings.NewReader(json))
				So(test_message.Date, ShouldEqual, message.Date)
				So(test_message.Content, ShouldEqual, message.Content)
				So(test_message.Creator, ShouldResemble, User{})
				So(test_message.Channel, ShouldResemble, Channel{})
			})
		})

		Convey("Given an message list", func() {
			message1 := Message{Date: GetMillis(),
				Content: "Message test",
			}
			message2 := Message{Date: GetMillis(),
				Content: "Message test",
			}
			message3 := Message{
				Date:    GetMillis(),
				Content: "Message test",
			}
			message_list := []*Message{&message1, &message2, &message3}

			Convey("Transfoming it in JSON then back to EMOJI LIST shoud give ressembling objects", func() {
				json := MessageListToJson(message_list)
				new_message_list := MessageListFromJson(strings.NewReader(json))
				So(new_message_list, ShouldResemble, message_list)
			})
		})
	})

	Convey("Testing pre Save function", t, func() {
		Convey("Given any message, it should update date with current date", func() {
			m1 := Message{}
			m2 := Message{Date: 10}
			m3 := Message{Date: 20, Content: "Test presave"}
			m4 := Message{Content: "Test presave"}
			m5 := Message{Date: 20, Content: "Test presave", Creator: user_test, Channel: channel_test}
			d1 := GetMillis()
			m1.PreSave()
			d2 := GetMillis()
			m2.PreSave()
			d3 := GetMillis()
			m3.PreSave()
			d4 := GetMillis()
			m4.PreSave()
			d5 := GetMillis()
			m5.PreSave()
			So(m1.Date, ShouldEqual, d1)
			So(m2.Date, ShouldEqual, d2)
			So(m3.Date, ShouldEqual, d3)
			So(m4.Date, ShouldEqual, d4)
			So(m5.Date, ShouldEqual, d5)
		})
	})

	Convey("Testing IsValid function", t, func() {
		Convey("Given a correct message. Message should be validate", func() {
			message := Message{
				Date:    GetMillis(),
				Content: "Message test",
				Creator: user_test,
				Channel: channel_test,
			}
			So(message.IsValid(), ShouldBeNil)
			So(message.IsValid(), ShouldNotResemble, u.NewLocAppError("Message.IsValid", "model.message.date.app_error", nil, ""))
			So(message.IsValid(), ShouldNotResemble, u.NewLocAppError("Message.IsValid", "model.message.creator.app_error", nil, ""))
			So(message.IsValid(), ShouldNotResemble, u.NewLocAppError("Message.IsValid", "model.message.channel.app_error", nil, ""))
		})
		Convey("Given an incorrect message. Message should be refused", func() {
			empty := Message{}
			message := Message{
				Content: "Message test",
				Creator: user_test,
				Channel: channel_test,
			}

			Convey("Empty message or no date message should return No Date error", func() {
				So(empty.IsValid(), ShouldNotBeNil)
				So(empty.IsValid(), ShouldResemble, u.NewLocAppError("Message.IsValid", "model.message.date.app_error", nil, ""))
				So(empty.IsValid(), ShouldNotResemble, u.NewLocAppError("Message.IsValid", "model.message.creator.app_error", nil, ""))
				So(empty.IsValid(), ShouldNotResemble, u.NewLocAppError("Message.IsValid", "model.message.channel.app_error", nil, ""))
				So(message.IsValid(), ShouldNotBeNil)
				So(message.IsValid(), ShouldResemble, u.NewLocAppError("Message.IsValid", "model.message.date.app_error", nil, ""))
				So(message.IsValid(), ShouldNotResemble, u.NewLocAppError("Message.IsValid", "model.message.creator.app_error", nil, ""))
				So(message.IsValid(), ShouldNotResemble, u.NewLocAppError("Message.IsValid", "model.message.channel.app_error", nil, ""))
			})

			message.Date = GetMillis()
			message.Creator = User{}

			Convey("Empty creator messages must return creator error", func() {
				So(message.IsValid(), ShouldNotBeNil)
				So(message.IsValid(), ShouldResemble, u.NewLocAppError("Message.IsValid", "model.message.creator.app_error", nil, ""))
				So(message.IsValid(), ShouldNotResemble, u.NewLocAppError("Message.IsValid", "model.message.date.app_error", nil, ""))
				So(message.IsValid(), ShouldNotResemble, u.NewLocAppError("Message.IsValid", "model.message.channel.app_error", nil, ""))
			})

			message.Creator = user_test
			message.Channel = Channel{}

			Convey("Empty channel message must return channel error", func() {
				So(message.IsValid(), ShouldNotBeNil)
				So(message.IsValid(), ShouldResemble, u.NewLocAppError("Message.IsValid", "model.message.channel.app_error", nil, ""))
				So(message.IsValid(), ShouldNotResemble, u.NewLocAppError("Message.IsValid", "model.message.date.app_error", nil, ""))
				So(message.IsValid(), ShouldNotResemble, u.NewLocAppError("Message.IsValid", "model.message.creator.app_error", nil, ""))
			})
		})
	})
}
