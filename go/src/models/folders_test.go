package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestFolderModel(t *testing.T) {
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

	message_test := Message{
		Date:    GetMillis(),
		Content: "Message test",
		Creator: user_test,
		Channel: channel_test,
	}

	Convey("Testing isValid function", t, func() {
		name_error := NewLocAppError("Folder.IsValid", "model.folder.name.app_error", nil, "")
		link_error := NewLocAppError("Folder.IsValid", "model.folder.link.app_error", nil, "")
		type_error := NewLocAppError("Folder.IsValid", "model.folder.type.app_error", nil, "")
		message_error := NewLocAppError("Folder.IsValid", "model.folder.message.app_error", nil, "")
		Convey("Given a correct folders. Should be validated", func() {
			folder := Folder{
				Name:    "Je suis .... Cailloux",
				Link:    "folders/youtube.com/watch?v=1JQE4YZS1Cg&t=1966s",
				Type:    "Video",
				Message: message_test,
			}
			So(folder.isValid(), ShouldBeNil)
			So(folder.isValid(), ShouldNotResemble, name_error)
			So(folder.isValid(), ShouldNotResemble, link_error)
			So(folder.isValid(), ShouldNotResemble, type_error)
			So(folder.isValid(), ShouldNotResemble, message_error)
		})

		Convey("Given incorrect folders. Should be refused", func() {
			folder := Folder{
				Name:    "Je suis .... Cailloux",
				Link:    "folders/youtube.com/watch?v=1JQE4YZS1Cg&t=1966s",
				Type:    "Video",
				Message: message_test,
			}
			empty := Folder{}
			folder.Name = ""

			Convey("empty Name or folder should return name error", func() {
				So(folder.isValid(), ShouldNotBeNil)
				So(folder.isValid(), ShouldResemble, name_error)
				So(folder.isValid(), ShouldNotResemble, link_error)
				So(folder.isValid(), ShouldNotResemble, type_error)
				So(folder.isValid(), ShouldNotResemble, message_error)
				So(empty.isValid(), ShouldNotBeNil)
				So(empty.isValid(), ShouldResemble, name_error)
				So(empty.isValid(), ShouldNotResemble, link_error)
				So(empty.isValid(), ShouldNotResemble, type_error)
				So(empty.isValid(), ShouldNotResemble, message_error)
			})

			folder.Name = "Correct Name"
			folder.Link = ""

			Convey("Empty link should result in link error", func() {
				So(folder.isValid(), ShouldNotBeNil)
				So(folder.isValid(), ShouldNotResemble, name_error)
				So(folder.isValid(), ShouldResemble, link_error)
				So(folder.isValid(), ShouldNotResemble, type_error)
				So(folder.isValid(), ShouldNotResemble, message_error)
			})

			folder.Link = "folder/corretc.xml"
			folder.Type = ""

			Convey("Empty type should result in type error", func() {
				So(folder.isValid(), ShouldNotBeNil)
				So(folder.isValid(), ShouldNotResemble, name_error)
				So(folder.isValid(), ShouldNotResemble, link_error)
				So(folder.isValid(), ShouldResemble, type_error)
				So(folder.isValid(), ShouldNotResemble, message_error)
			})

			folder.Type = "xml"
			folder.Message = Message{}

			Convey("Empty message should result in message", func() {
				So(folder.isValid(), ShouldNotBeNil)
				So(folder.isValid(), ShouldNotResemble, name_error)
				So(folder.isValid(), ShouldNotResemble, link_error)
				So(folder.isValid(), ShouldNotResemble, type_error)
				So(folder.isValid(), ShouldResemble, message_error)
			})
		})
	})

	Convey("Testing json VS folder transformations", t, func() {
		Convey("Given an folder", func() {
			folder := Folder{
				Name: "Je suis .... Cailloux",
				Link: "folders/youtube.com/watch?v=1JQE4YZS1Cg&t=1966s",
				Type: "Video",
			}
			Convey("Transforming it in JSON then back to FOLDER should provide similar objects", func() {
				json := folder.toJson()
				new_folder := folderFromJson(strings.NewReader(json))
				So(new_folder, ShouldResemble, &folder)
			})
		})

		Convey("Given an folder list", func() {
			folder1 := Folder{
				Name: "Je suis .... Cailloux",
				Link: "folders/youtube.com/watch?v=1JQE4YZS1Cg&t=1966s",
				Type: "Video",
			}
			folder2 := Folder{
				Name: "Somethi,g",
				Link: "folders/something.sql",
				Type: "sql",
			}
			folder3 := Folder{
				Name: "Should automatize type recon",
				Link: "folders/facepalm.svg",
				Type: "facepalm?",
			}
			folder_list := []*Folder{&folder1, &folder2, &folder3}

			Convey("Transfoming it in JSON then back to FOLDER LIST shoud give ressembling objects", func() {
				json := folderListToJson(folder_list)
				new_folder_list := folderListFromJson(strings.NewReader(json))
				So(new_folder_list, ShouldResemble, folder_list)
			})

		})
	})

}
