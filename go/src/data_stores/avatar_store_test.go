// This file is used to test if user model is working correctly.
// A user is always linked to an avatar
// He has basic channel to join
package data_stores

import (
	// "github.com/jinzhu/gorm"
	. "github.com/smartystreets/goconvey/convey"
	. "models"
	// "strconv"
	// "strings"
	"testing"
	. "utils"
)

func TestAvatarStore(t *testing.T) {
	ds := DataStore{}
	ds.initConnection("root", "popcube_test", "popcube_dev")
	db := *ds.Db
	asi := AvatarStoreImpl{}
	Convey("Testing save function", t, func() {
		db_error := NewLocAppError("avatar_store_impl.Save", "save.transaction.create.encounter_error", nil, "")
		alreadyexist_error := NewLocAppError("avatar_store_impl.Save", "save.transaction.create.already_exist", nil, "Avatar Name: Troll Face")
		avatar := Avatar{
			Name: "Troll Face",
			Link: "avatars/trollface.svg",
		}
		Convey("Given a correct avatar.", func() {
			appError := asi.Save(&avatar, ds)
			Convey("Trying to add it for the first time, should be accepted", func() {
				So(appError, ShouldBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldNotResemble, alreadyexist_error)
			})
			Convey("Trying to add it a second time should return duplicate error", func() {
				appError2 := asi.Save(&avatar, ds)
				So(appError2, ShouldNotBeNil)
				So(appError2, ShouldResemble, alreadyexist_error)
				So(appError2, ShouldNotResemble, db_error)
			})
		})
		db.Delete(&avatar)
	})

	Convey("Testing update function", t, func() {
		db_error := NewLocAppError("avatar_store_impl.Save", "save.transaction.create.encounter_error", nil, "")
		alreadyexist_error := NewLocAppError("avatar_store_impl.Save", "save.transaction.create.already_exist", nil, "Avatar Name: Troll Face")
		avatar := Avatar{
			Name: "Troll Face",
			Link: "avatars/trollface.svg",
		}
		avatar_new := Avatar{
			Name: "TrollFace2",
			Link: "avatars/trollface2.svg",
		}

		appError := asi.Save(&avatar, ds)
		So(appError, ShouldBeNil)
		So(appError, ShouldNotResemble, db_error)
		So(appError, ShouldNotResemble, alreadyexist_error)

		Convey("Provided correct Avatar to modify should not return errors", func() {
			appError := asi.Update(&avatar, &avatar_new, ds)
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, db_error)
			So(appError, ShouldNotResemble, alreadyexist_error)
		})

		Convey("Provided wrong Avatar to modify should result in old_avatar error", func() {
			avatar.Name = ""
			Convey("Too long or empty Name should return name error", func() {
				appError := asi.Update(&avatar, &avatar_new, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("avatar_store_impl.Update.avatar_old.PreSave", "model.avatar.name.app_error", nil, ""))
				avatar.Name = "thishastobeatoolongname.For this, it need to be more than 64 char lenght .............. So long. Plus it should be alpha numeric. I'll add the test later on."
				appError = asi.Update(&avatar, &avatar_new, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("avatar_store_impl.Update.avatar_old.PreSave", "model.avatar.name.app_error", nil, ""))
			})

			avatar.Name = "Correct Name"
			avatar.Link = ""

			Convey("Empty link should result in link error", func() {
				appError = asi.Update(&avatar, &avatar_new, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("avatar_store_impl.Update.avatar_old.PreSave", "model.avatar.link.app_error", nil, ""))
			})
		})

		Convey("Provided wrong Avatar to modify should result in new_avatar error", func() {
			avatar_new.Name = ""
			Convey("Too long or empty Name should return name error", func() {
				appError := asi.Update(&avatar, &avatar_new, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("avatar_store_impl.Update.avatar_new.PreSave", "model.avatar.name.app_error", nil, ""))
				avatar_new.Name = "thishastobeatoolongname.For this, it need to be more than 64 char lenght .............. So long. Plus it should be alpha numeric. I'll add the test later on."
				appError = asi.Update(&avatar, &avatar_new, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("avatar_store_impl.Update.avatar_new.PreSave", "model.avatar.name.app_error", nil, ""))
			})

			avatar_new.Name = "Correct Name"
			avatar_new.Link = ""

			Convey("Empty link should result in link error", func() {
				appError = asi.Update(&avatar, &avatar_new, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("avatar_store_impl.Update.avatar_new.PreSave", "model.avatar.link.app_error", nil, ""))
			})
		})
		db.Delete(&avatar)
		db.Delete(&avatar_new)
	})

	Convey("Testing Getters", t, func() {
		avatar0 := Avatar{
			Name: "Troll Face",
			Link: "avatars/trollface.svg",
		}
		avatar1 := Avatar{
			Name: "Face Palm Old",
			Link: "avatars/facepalmold.svg",
		}
		avatar1_new := Avatar{
			Name: "Face Palm",
			Link: "avatars/facepalm.svg",
		}
		avatar2 := Avatar{
			Name: "God",
			Link: "avatars/docker.svg",
		}
		avatar3 := Avatar{
			Name: "nOOb",
			Link: "avatars/sparadra.svg",
		}
		asi.Save(&avatar0, ds)
		asi.Save(&avatar1, ds)
		asi.Update(&avatar1, &avatar1_new, ds)
		asi.Save(&avatar2, ds)
		asi.Save(&avatar3, ds)
		// Have to be after save so ID are up to date :O
		avatar_list := []Avatar{
			avatar0,
			avatar1,
			avatar2,
			avatar3,
		}

		Convey("We have to be able to find all avatars in the db", func() {
			avatars := asi.GetAll(ds)
			So(avatars, ShouldNotResemble, &[]Avatar{})
			So(avatars, ShouldResemble, &avatar_list)
		})

		Convey("We have to be able to find an avatar from is name", func() {
			avatar := asi.GetByName(avatar0.Name, ds)
			So(avatar, ShouldNotResemble, &Avatar{})
			So(avatar, ShouldResemble, &avatar0)
			avatar = asi.GetByName(avatar2.Name, ds)
			So(avatar, ShouldNotResemble, &Avatar{})
			So(avatar, ShouldResemble, &avatar2)
			avatar = asi.GetByName(avatar3.Name, ds)
			So(avatar, ShouldNotResemble, &Avatar{})
			So(avatar, ShouldResemble, &avatar3)
			Convey("Should also work from updated value", func() {
				avatar = asi.GetByName(avatar1.Name, ds)
				So(avatar, ShouldNotResemble, &Avatar{})
				So(avatar, ShouldResemble, &avatar1)
			})
		})

		Convey("We have to be able to find an avatar from is link", func() {
			avatar := asi.GetByLink(avatar0.Link, ds)
			So(avatar, ShouldNotResemble, &Avatar{})
			So(avatar, ShouldResemble, &avatar0)
			avatar = asi.GetByLink(avatar2.Link, ds)
			So(avatar, ShouldNotResemble, &Avatar{})
			So(avatar, ShouldResemble, &avatar2)
			avatar = asi.GetByLink(avatar3.Link, ds)
			So(avatar, ShouldNotResemble, &Avatar{})
			So(avatar, ShouldResemble, &avatar3)
			Convey("Should also work from updated value", func() {
				avatar = asi.GetByLink(avatar1.Link, ds)
				So(avatar, ShouldNotResemble, &Avatar{})
				So(avatar, ShouldResemble, &avatar1)
			})
		})

		Convey("Searching for non existent avatar should return empty", func() {
			avatar := asi.GetByLink("The Mask", ds)
			So(avatar, ShouldResemble, &Avatar{})
			avatar = asi.GetByName("Fantôme", ds)
			So(avatar, ShouldResemble, &Avatar{})
		})

		db.Delete(&avatar0)
		db.Delete(&avatar1)
		db.Delete(&avatar2)
		db.Delete(&avatar3)

		Convey("Searching all in empty table should return empty", func() {
			avatars := asi.GetAll(ds)
			So(avatars, ShouldResemble, &[]Avatar{})
		})
	})

	Convey("Testing delete avatar", t, func() {
		dberror := NewLocAppError("avatar_store_impl.Delete", "update.transaction.delete.encounter_error", nil, "")
		avatar0 := Avatar{
			Name: "Troll Face",
			Link: "avatars/trollface.svg",
		}
		avatar1 := Avatar{
			Name: "Face Palm",
			Link: "avatars/facepal.svg",
		}
		avatar2 := Avatar{
			Name: "God",
			Link: "avatars/docker.svg",
		}
		avatar3 := Avatar{
			Name: "nOOb",
			Link: "avatars/sparadra.svg",
		}
		asi.Save(&avatar0, ds)
		asi.Save(&avatar1, ds)
		asi.Save(&avatar2, ds)
		asi.Save(&avatar3, ds)
		avatar3_old := avatar3
		avatar_list1 := []Avatar{
			avatar0,
			avatar1,
			avatar2,
			avatar3_old,
		}

		Convey("Deleting a known avatar should work", func() {
			appError := asi.Delete(&avatar2, ds)
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, dberror)
			So(asi.GetByName("God", ds), ShouldResemble, &Avatar{})
		})

		Convey("Trying to delete from non conform avatar should return specific avatar error and should not delete avatars.", func() {
			avatar3.Name = ""
			Convey("Too long or empty Name should return name error", func() {
				appError := asi.Delete(&avatar3, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dberror)
				So(appError, ShouldResemble, NewLocAppError("avatar_store_impl.Delete.avatar.PreSave", "model.avatar.name.app_error", nil, ""))
				So(asi.GetAll(ds), ShouldResemble, &avatar_list1)
				avatar3.Name = "thishastobeatoolongname.For this, it need to be more than 64 char lenght .............. So long. Plus it should be alpha numeric. I'll add the test later on."
				appError = asi.Delete(&avatar3, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dberror)
				So(appError, ShouldResemble, NewLocAppError("avatar_store_impl.Delete.avatar.PreSave", "model.avatar.name.app_error", nil, ""))
				So(asi.GetAll(ds), ShouldResemble, &avatar_list1)
			})

			avatar3.Name = "nOOb"
			avatar3.Link = ""

			Convey("Empty link should result in link error", func() {
				appError := asi.Delete(&avatar3, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, dberror)
				So(appError, ShouldResemble, NewLocAppError("avatar_store_impl.Delete.avatar.PreSave", "model.avatar.link.app_error", nil, ""))
				So(asi.GetAll(ds), ShouldResemble, &avatar_list1)
			})
		})

		db.Delete(&avatar0)
		db.Delete(&avatar1)
		db.Delete(&avatar2)
		db.Delete(&avatar3)
	})
}
