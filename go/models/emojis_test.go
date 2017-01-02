package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestEmojisModel(t *testing.T) {
	Convey("Testing isValid function", t, func() {
		Convey("Given a correct emojis. Should be validated", func() {
			emoji := Emoji{
				Name:     "Troll Face",
				ShortCut: ":troll-face:",
				Link:     "emojis/trollface.svg",
			}
			So(emoji.isValid(), ShouldBeNil)
			So(emoji.isValid(), ShouldNotResemble, NewLocAppError("Emoji.IsValid", "model.emoji.name.app_error", nil, ""))
			So(emoji.isValid(), ShouldNotResemble, NewLocAppError("Emoji.IsValid", "model.emoji.shortcut.app_error", nil, ""))
			So(emoji.isValid(), ShouldNotResemble, NewLocAppError("Emoji.IsValid", "model.emoji.link.app_error", nil, ""))
		})

		Convey("Given incorrect emojis. Should be refused", func() {
			emoji := Emoji{
				Name:     "Troll Face",
				ShortCut: ":this-is-a-tool-long-shortcut:",
				Link:     "emojis/trollface.svg",
			}

			Convey("Too long shortcut or empty shorctcut should return Shortcut error", func() {
				So(emoji.isValid(), ShouldResemble, NewLocAppError("Emoji.IsValid", "model.emoji.shortcut.app_error", nil, ""))
				emoji.ShortCut = ""
				So(emoji.isValid(), ShouldResemble, NewLocAppError("Emoji.IsValid", "model.emoji.shortcut.app_error", nil, ""))
			})
			emoji.ShortCut = ":goodone:"
			emoji.Name = ""
			Convey("Too long or empty Name should return name error", func() {
				So(emoji.isValid(), ShouldResemble, NewLocAppError("Emoji.IsValid", "model.emoji.name.app_error", nil, ""))
				emoji.Name = "thishastobeatoolongname.For this, it need to be more than 64 char lenght .............. So long. Plus it should be alpha numeric. I'll add the test later on."
				So(emoji.isValid(), ShouldResemble, NewLocAppError("Emoji.IsValid", "model.emoji.name.app_error", nil, ""))
			})
			emoji.Name = "Correct Name"
			emoji.Link = ""
			Convey("Empty link should result in link error", func() {
				So(emoji.isValid(), ShouldResemble, NewLocAppError("Emoji.IsValid", "model.emoji.link.app_error", nil, ""))
			})
		})
	})

	Convey("Testing json VS emoji transformations", t, func() {
		Convey("Given an emoji", func() {
			emoji := Emoji{
				Name:     "Troll Face",
				ShortCut: ":troll-face:",
				Link:     "emojis/trollface.svg",
			}
			Convey("Transforming it in JSON then back to EMOJI should provide similar objects", func() {
				json := emoji.toJson()
				new_emoji := emojiFromJson(strings.NewReader(json))
				So(new_emoji, ShouldResemble, &emoji)
			})
		})

		Convey("Given an emoji list", func() {
			emoji1 := Emoji{
				Name:     "Troll Face",
				ShortCut: ":troll:",
				Link:     "emojis/trollface.svg",
			}
			emoji2 := Emoji{
				Name:     "Joy Face",
				ShortCut: ":)",
				Link:     "emojis/joyface.svg",
			}
			emoji3 := Emoji{
				Name:     "Face Palm",
				ShortCut: ":facepalm:",
				Link:     "emojis/facepalm.svg",
			}
			emoji_list := []*Emoji{&emoji1, &emoji2, &emoji3}

			Convey("Transfoming it in JSON then back to EMOJI LIST shoud give ressembling objects", func() {
				json := emojiListToJson(emoji_list)
				new_emoji_list := emojiListFromJson(strings.NewReader(json))
				So(new_emoji_list, ShouldResemble, emoji_list)
			})

		})
	})

}
