// This file is used to test if user model is working correctly.
// A user is always linked to an organisation
// He has basic channel to join
package data_stores

import (
	// "github.com/jinzhu/gorm"
	. "github.com/smartystreets/goconvey/convey"
	. "models"
	"strconv"
	"strings"
	"testing"
	. "utils"
)

func TestOrganisationStore(t *testing.T) {
	ds := DataStore{}
	ds.initConnection("root", "popcube_test", "popcube_dev")
	db := *ds.Db
	osi := OrganisationStoreImpl{}
	Convey("Testing save function", t, func() {
		db_error := NewLocAppError("organisation_store_impl.Save", "save.transaction.create.encounter_error", nil, "")
		alreadyexist_error := NewLocAppError("organisation_store_impl.Save", "save.transaction.create.already_exist", nil, "Organisation Name: zeus")
		organisation := Organisation{
			IdOrganisation:   0,
			DockerStack:      1,
			OrganisationName: "zeus",
			Description:      "Testing organisation description :O",
			Avatar:           "zeus.svg",
			Domain:           "zeus.popcube",
		}
		Convey("Given a correct organisation.", func() {
			appError := osi.Save(&organisation, ds)
			Convey("Trying to add it for the first time, should be accepted", func() {
				So(appError, ShouldBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldNotResemble, alreadyexist_error)
			})
			Convey("Trying to add it a second time should return duplicate error", func() {
				appError2 := osi.Save(&organisation, ds)
				So(appError2, ShouldNotBeNil)
				So(appError2, ShouldResemble, alreadyexist_error)
				So(appError2, ShouldNotResemble, db_error)
			})
		})
		Convey("Given an incorrect organisation.", func() {
			empty := Organisation{}
			organisation.OrganisationName = ""
			Convey("Empty organisation or no Organisation Name organisation should return No name error", func() {
				appError := osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldNotResemble, alreadyexist_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Save.organisation.PreSave", "model.organisation.is_valid.organisation_name.app_error", nil,
					"id="+strconv.FormatUint(organisation.IdOrganisation, 10)))
				appError = osi.Save(&empty, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldNotResemble, alreadyexist_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Save.organisation.PreSave", "model.organisation.is_valid.organisation_name.app_error", nil,
					"id="+strconv.FormatUint(organisation.IdOrganisation, 10)))
			})
			organisation.OrganisationName = strings.ToLower("ThisShouldBeAFreakingLongEnougthStringToRefuse.BahNon, pas tout seul. C'est long 64 caractères en vrai  ~#~")
			Convey("Too long organisation name should return Too Long organisation name error", func() {
				appError := osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldNotResemble, alreadyexist_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Save.organisation.PreSave", "model.organisation.is_valid.organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IdOrganisation, 10)))
			})
			Convey("Incorect Alpha Num organisation name should be refused ", func() {
				organisation.OrganisationName = "?/+*"
				appError := osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldNotResemble, alreadyexist_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Save.organisation.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IdOrganisation, 10)))
				organisation.OrganisationName = "("
				appError = osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldNotResemble, alreadyexist_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Save.organisation.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IdOrganisation, 10)))
				organisation.OrganisationName = "{"
				appError = osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldNotResemble, alreadyexist_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Save.organisation.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IdOrganisation, 10)))
				organisation.OrganisationName = "}"
				appError = osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldNotResemble, alreadyexist_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Save.organisation.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IdOrganisation, 10)))
				organisation.OrganisationName = ")"
				appError = osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldNotResemble, alreadyexist_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Save.organisation.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IdOrganisation, 10)))
				organisation.OrganisationName = "["
				appError = osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldNotResemble, alreadyexist_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Save.organisation.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IdOrganisation, 10)))
				organisation.OrganisationName = "]"
				appError = osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldNotResemble, alreadyexist_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Save.organisation.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IdOrganisation, 10)))
				organisation.OrganisationName = " "
				appError = osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldNotResemble, alreadyexist_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Save.organisation.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(organisation.IdOrganisation, 10)))
			})
			organisation.OrganisationName = "electra"

			organisation.Description = "Il Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face:"
			Convey("Given a too long description, should return too long description error :p", func() {

				appError := osi.Save(&organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldNotResemble, alreadyexist_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Save.organisation.PreSave", "model.organisation.is_valid.description.app_error",
					nil, "id="+strconv.FormatUint(organisation.IdOrganisation, 10)))
			})
			organisation.Description = "Stoppppppp"
		})
		db.Delete(&organisation)
	})
	Convey("Testing update function", t, func() {
		organisation := Organisation{
			IdOrganisation:   0,
			DockerStack:      1,
			OrganisationName: "zeus",
			Description:      "Testing organisation description :O",
			Avatar:           "zeus.svg",
			Domain:           "zeus.popcube",
		}
		new_organisation := Organisation{
			DockerStack:      4,
			OrganisationName: "NewZeus",
		}
		appError := osi.Save(&organisation, ds)
		db_error := NewLocAppError("organisation_store_impl.Update", "update.transaction.updates.encounter_error", nil, "")
		So(appError, ShouldBeNil)
		So(appError, ShouldNotResemble, db_error)
		Convey("Providing a correct user to update", func() {
			appError = osi.Update(&organisation, &new_organisation, ds)
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, db_error)
		})
		Convey("Providing an incorrect user should result in errors", func() {
			empty := Organisation{}
			new_organisation.OrganisationName = ""
			Convey("Empty organisation or no Organisation Name organisation should return No name error", func() {
				appError := osi.Update(&organisation, &new_organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Update.organisation_new.PreSave", "model.organisation.is_valid.organisation_name.app_error", nil,
					"id="+strconv.FormatUint(new_organisation.IdOrganisation, 10)))
				appError = osi.Update(&organisation, &empty, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Update.organisation_new.PreSave", "model.organisation.is_valid.organisation_name.app_error", nil,
					"id="+strconv.FormatUint(empty.IdOrganisation, 10)))
			})
			new_organisation.OrganisationName = strings.ToLower("ThisShouldBeAFreakingLongEnougthStringToRefuse.BahNon, pas tout seul. C'est long 64 caractères en vrai  ~#~")
			Convey("Too long organisation name should return Too Long organisation name error", func() {
				appError = osi.Update(&organisation, &new_organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Update.organisation_new.PreSave", "model.organisation.is_valid.organisation_name.app_error",
					nil, "id="+strconv.FormatUint(new_organisation.IdOrganisation, 10)))
			})
			Convey("Incorect Alpha Num organisation name should be refused ", func() {
				new_organisation.OrganisationName = "?/+*"
				appError = osi.Update(&organisation, &new_organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Update.organisation_new.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(new_organisation.IdOrganisation, 10)))
				new_organisation.OrganisationName = "("
				appError = osi.Update(&organisation, &new_organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Update.organisation_new.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(new_organisation.IdOrganisation, 10)))
				new_organisation.OrganisationName = "{"
				appError = osi.Update(&organisation, &new_organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Update.organisation_new.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(new_organisation.IdOrganisation, 10)))
				new_organisation.OrganisationName = "}"
				appError = osi.Update(&organisation, &new_organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Update.organisation_new.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(new_organisation.IdOrganisation, 10)))
				new_organisation.OrganisationName = ")"
				appError = osi.Update(&organisation, &new_organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Update.organisation_new.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(new_organisation.IdOrganisation, 10)))
				new_organisation.OrganisationName = "["
				appError = osi.Update(&organisation, &new_organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Update.organisation_new.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(new_organisation.IdOrganisation, 10)))
				new_organisation.OrganisationName = "]"
				appError = osi.Update(&organisation, &new_organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Update.organisation_new.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(new_organisation.IdOrganisation, 10)))
				new_organisation.OrganisationName = " "
				appError = osi.Update(&organisation, &new_organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Update.organisation_new.PreSave", "model.organisation.is_valid.not_alphanum_organisation_name.app_error",
					nil, "id="+strconv.FormatUint(new_organisation.IdOrganisation, 10)))
			})
			new_organisation.OrganisationName = "electra"

			new_organisation.Description = "Il Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face:"
			Convey("Given a too long description, should return too long description error :p", func() {
				appError = osi.Update(&organisation, &new_organisation, ds)
				So(appError, ShouldNotBeNil)
				So(appError, ShouldNotResemble, db_error)
				So(appError, ShouldResemble, NewLocAppError("organisation_store_impl.Update.organisation_new.PreSave", "model.organisation.is_valid.description.app_error",
					nil, "id="+strconv.FormatUint(new_organisation.IdOrganisation, 10)))
			})
			new_organisation.Description = "Stoppppppp"

		})
		db.Delete(&organisation)
	})
	Convey("Testing Get function", t, func() {
		organisation := Organisation{
			IdOrganisation:   0,
			DockerStack:      1,
			OrganisationName: "zeus",
			Description:      "Testing organisation description :O",
			Avatar:           "zeus.svg",
			Domain:           "zeus.popcube",
		}
		Convey("Trying to get organisation from empty DB should return empty", func() {
			So(&Organisation{}, ShouldResemble, osi.Get(ds))
		})
		appError := osi.Save(&organisation, ds)
		So(appError, ShouldBeNil)
		Convey("Trying to get organisation from non empty DB should return a correct organisation object", func() {
			got := osi.Get(ds)
			So(&organisation, ShouldResemble, got)
			So(got.IsValid(), ShouldBeNil)
		})
		db.Delete(&organisation)
	})
}
