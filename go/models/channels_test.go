package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestChannelModel(t *testing.T) {
	Convey("Testing json vs channel conversions", t, func() {
		Convey("Given a channel", func() {
			channel := Channel{WebId: NewId(), ChannelName: NewId()}
			Convey("Converting channel to json then json to channel should provide same channel information", func() {
				json := channel.toJson()
				test_channel := channelFromJson(strings.NewReader(json))
				So(channel.WebId, ShouldEqual, test_channel.WebId)
				So(channel.ChannelName, ShouldEqual, test_channel.ChannelName)
			})
		})
	})

	Convey("Testing isValid function", t, func() {
		Convey("Given a correct channel. Channel should be validate", func() {
			channel := Channel{
				WebId:       NewId(),
				ChannelName: "electra",
				UpdatedAt:   GetMillis(),
				Type:        "audio",
				Private:     false,
				Description: "Testing channel description :O",
				Subject:     "Sujet",
				Avatar:      "jesuiscool.svg",
			}
			So(channel.isValid(), ShouldBeNil)
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.id.app_error", nil, ""))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.channel_name.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.update_at.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.description.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.subject.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.type.app_error", nil, "id="+channel.WebId))
		})
		Convey("Given an incorrect channel. Channel should be refused", func() {
			empty := Channel{}
			channel := Channel{
				ChannelName: "Electra",
				UpdatedAt:   GetMillis(),
				Type:        "audio",
				Private:     false,
				Description: "Testing channel description :O",
				Subject:     "Sujet",
				Avatar:      "jesuiscool.svg",
			}
			Convey("Empty channel or no WebId channel should return No Id error", func() {
				So(empty.isValid(), ShouldResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.id.app_error", nil, ""))
				So(channel.isValid(), ShouldResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.id.app_error", nil, ""))
			})
			channel.WebId = NewId()
			channel.ChannelName = strings.ToLower("ThisShouldBeAFreakingLongEnougthStringToRefuse.BahNon, pas tout seul. C'est long 64 caractères en vrai  ~#~")
			Convey("Too long channel name should return Too Long channel name error", func() {
				So(channel.isValid(), ShouldResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.channel_name.app_error", nil, "id="+channel.WebId))
			})
			Convey("Incorect Alpha Num channel name should be refused (no CAPS)", func() {
				channel.ChannelName = "JeSuisCaps"
				So(channel.isValid(), ShouldResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebId))
				channel.ChannelName = "?/+*"
				So(channel.isValid(), ShouldResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebId))
				channel.ChannelName = "("
				So(channel.isValid(), ShouldResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebId))
				channel.ChannelName = "{"
				So(channel.isValid(), ShouldResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebId))
				channel.ChannelName = "}"
				So(channel.isValid(), ShouldResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebId))
				channel.ChannelName = ")"
				So(channel.isValid(), ShouldResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebId))
				channel.ChannelName = "["
				So(channel.isValid(), ShouldResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebId))
				channel.ChannelName = "]"
				So(channel.isValid(), ShouldResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebId))
				channel.ChannelName = " "
				So(channel.isValid(), ShouldResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebId))
			})
			channel.ChannelName = "electra"
			channel.UpdatedAt = 0
			Convey("Given an incorrect update date should be refuse", func() {
				So(channel.isValid(), ShouldResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.update_at.app_error", nil, "id="+channel.WebId))
			})
			channel.UpdatedAt = GetMillis()
			channel.Description = "Il Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face:"
			Convey("Given a too long description, should return too long description error :p", func() {
				So(channel.isValid(), ShouldResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.description.app_error", nil, "id="+channel.WebId))
			})
			channel.Description = "Stoppppppp"
			channel.Subject = "Encore beaucoup de caractere pour rien .... mais un peu moins cette fois. Il n'en faut que 250 ........... Fait dodo, cola mon p'tit frere. Fais dodo, j'ai pêté un cable. Swing du null, Swing du null, c'est le swing du null ..... :guitare: :singer: :music: Je suis un main troll :O"
			Convey("Given a too long subject, should return too long description error :p", func() {
				So(channel.isValid(), ShouldResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.subject.app_error", nil, "id="+channel.WebId))
			})
			channel.Subject = "Safe"
			channel.Type = "Outside of Range"
			Convey("Providing a wrong type should not work", func() {
				So(channel.isValid(), ShouldResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.type.app_error", nil, "id="+channel.WebId))
			})
		})
	})

	Convey("Testing PreSave function", t, func() {
		channel := Channel{}
		Convey("If channel is empty, should fill some fields - webId, ChannelName, UpdatedAt, Avatar and type, and user should be valid", func() {
			channel.preSave()
			So(channel.isValid(), ShouldBeNil)
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.id.app_error", nil, ""))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.channel_name.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.update_at.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.description.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.subject.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.type.app_error", nil, "id="+channel.WebId))
			So(channel.Avatar, ShouldEqual, "default_channel_avatar.svg")
			So(channel.Type, ShouldEqual, "text")
		})
		Convey("If provided ChannelName contain caps, they should be lowered", func() {
			channel.ChannelName = "JeSuisCaps"
			channel.preSave()
			So(channel.isValid(), ShouldBeNil)
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.id.app_error", nil, ""))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.channel_name.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.update_at.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.description.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.subject.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.type.app_error", nil, "id="+channel.WebId))
			So(channel.ChannelName, ShouldEqual, "jesuiscaps")
			channel.ChannelName = "nocapsshouldnotbemodified"
			channel.preSave()
			So(channel.isValid(), ShouldBeNil)
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.id.app_error", nil, ""))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.channel_name.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.not_alphanum_channel_name.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.update_at.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.description.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.subject.app_error", nil, "id="+channel.WebId))
			So(channel.isValid(), ShouldNotResemble, NewLocAppError("Channel.IsValid", "model.channel.is_valid.type.app_error", nil, "id="+channel.WebId))
			So(channel.ChannelName, ShouldEqual, "nocapsshouldnotbemodified")
		})
	})

	Convey("Testing PreUpdate function", t, func() {
		Convey("PreUpdating a channel should not modify channel, only update time.", func() {
			channel := Channel{
				WebId:       "TestWebId",
				ChannelName: "TestChannelName",
				UpdatedAt:   GetMillis() - 20,
				Type:        "audio",
				Private:     true,
				Description: "Testing channel description",
				Subject:     "Sujet",
				Avatar:      "jesuiscool.svg",
			}
			old_updatedat := channel.UpdatedAt
			channel.preUpdate()
			So(channel.UpdatedAt, ShouldBeGreaterThan, old_updatedat)
			So(channel.WebId, ShouldEqual, "TestWebId")
			So(channel.ChannelName, ShouldEqual, "TestChannelName")
			So(channel.Type, ShouldEqual, "audio")
			So(channel.Private, ShouldEqual, true)
			So(channel.Description, ShouldEqual, "Testing channel description")
			So(channel.Subject, ShouldEqual, "Sujet")
			So(channel.Avatar, ShouldEqual, "jesuiscool.svg")
		})
	})
}
