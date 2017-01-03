package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestAvatarsModel(t *testing.T) {
	Convey("Testing isValid function", t, func() {
		Convey("Given a correct avatars. Should be validated", func() {
			avatar := Avatar{
				Name: "Troll Face",
				Link: "avatars/trollface.svg",
			}
			So(avatar.isValid(), ShouldBeNil)
			So(avatar.isValid(), ShouldNotResemble, NewLocAppError("Avatar.IsValid", "model.avatar.name.app_error", nil, ""))
			So(avatar.isValid(), ShouldNotResemble, NewLocAppError("Avatar.IsValid", "model.avatar.link.app_error", nil, ""))
		})

		Convey("Given incorrect avatars. Should be refused", func() {
			avatar := Avatar{
				Name: "Troll Face",
				Link: "avatars/trollface.svg",
			}

			avatar.Name = ""

			Convey("Too long or empty Name should return name error", func() {
				So(avatar.isValid(), ShouldResemble, NewLocAppError("Avatar.IsValid", "model.avatar.name.app_error", nil, ""))
				avatar.Name = "thishastobeatoolongname.For this, it need to be more than 64 char lenght .............. So long. Plus it should be alpha numeric. I'll add the test later on."
				So(avatar.isValid(), ShouldResemble, NewLocAppError("Avatar.IsValid", "model.avatar.name.app_error", nil, ""))
			})

			avatar.Name = "Correct Name"
			avatar.Link = ""

			Convey("Empty link should result in link error", func() {
				So(avatar.isValid(), ShouldResemble, NewLocAppError("Avatar.IsValid", "model.avatar.link.app_error", nil, ""))
			})
		})
	})

	Convey("Testing json VS avatar transformations", t, func() {
		Convey("Given an avatar", func() {
			avatar := Avatar{
				Name: "Troll Face",
				Link: "avatars/trollface.svg",
			}
			Convey("Transforming it in JSON then back to EMOJI should provide similar objects", func() {
				json := avatar.toJson()
				new_avatar := avatarFromJson(strings.NewReader(json))
				So(new_avatar, ShouldResemble, &avatar)
			})
		})

		Convey("Given an avatar list", func() {
			avatar1 := Avatar{
				Name: "Troll Face",
				Link: "avatars/trollface.svg",
			}
			avatar2 := Avatar{
				Name: "Joy Face",
				Link: "avatars/joyface.svg",
			}
			avatar3 := Avatar{
				Name: "Face Palm",
				Link: "avatars/facepalm.svg",
			}
			avatar_list := []*Avatar{&avatar1, &avatar2, &avatar3}

			Convey("Transfoming it in JSON then back to EMOJI LIST shoud give ressembling objects", func() {
				json := avatarListToJson(avatar_list)
				new_avatar_list := avatarListFromJson(strings.NewReader(json))
				So(new_avatar_list, ShouldResemble, avatar_list)
			})

		})
	})

}
