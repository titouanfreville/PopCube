// This file is used to test if user model is working correctly.
// A user is always linked to an parameter
// He has bpsic channel to join
package datastores

import (
	. "github.com/smartystreets/goconvey/convey"
	. "models"
	// "strings"
	"testing"
	u "utils"
)

func TestParameterStore(t *testing.T) {
	ds := dbStore{}
	ds.InitConnection("root", "popcube_test", "popcube_dev")
	db := *ds.Db
	osi := ParameterStoreImpl{}
	Convey("Testing save function", t, func() {
		dbError := u.NewLocAppError("parameterStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyexistError := u.NewLocAppError("parameterStoreImpl.Save", "save.transaction.create.already_exist", nil, "")
		parameter := Parameter{
			Local:      "en_EN",
			TimeZone:   "UTC+2",
			SleepStart: 280,
			SleepEnd:   12,
		}
		Convey("Given a correct parameter.", func() {
			appError := osi.Save(&parameter, ds)
			Convey("Trying to add it for the first time, should be accepted", func() {
				So(appError, ShouldBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyexistError)
			})
			Convey("Trying to add it a second time should return duplicate error", func() {
				appError2 := osi.Save(&parameter, ds)
				So(appError2, ShouldNotBeNil)
				So(appError2, ShouldResemble, alreadyexistError)
				So(appError2, ShouldNotResemble, dbError)
			})
		})
		Convey("Given an incorrect parameter.", func() {
			empty := Parameter{}
			Convey("Empty parameter should return first error from is valid error", func() {
				appError := osi.Save(&empty, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
			})
			parameter.Local = "en_ENG"
			Convey("Given empty local or too long local should return Local error", func() {
				appError := osi.Save(&parameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
				parameter.Local = ""
				appError = osi.Save(&parameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
			})
			parameter.Local = "en_EN"
			parameter.TimeZone = "UTF+134"
			Convey("Given empty time zone or too long time zone should return Time Zone error", func() {
				appError := osi.Save(&parameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_timezone.app_error", nil, ""))
				parameter.TimeZone = ""
				appError = osi.Save(&parameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_timezone.app_error", nil, ""))
			})
			parameter.TimeZone = "UTF+12"
			parameter.SleepEnd = -1
			Convey("Given negative or too big Sleep timers should return sleep error", func() {
				appError := osi.Save(&parameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_sleep_end.app_error", nil, ""))
				parameter.SleepEnd = 1441
				appError = osi.Save(&parameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_sleep_end.app_error", nil, ""))
				parameter.SleepEnd = 10
				parameter.SleepStart = -10
				appError = osi.Save(&parameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_sleep_start.app_error", nil, ""))
				parameter.SleepStart = 2000
				appError = osi.Save(&parameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Save.parameter.PreSave", "model.parameter.is_valid.parameter_sleep_start.app_error", nil, ""))
			})
		})
		db.Delete(&parameter)
	})

	Convey("Testing update function", t, func() {
		parameter := Parameter{
			Local:      "en_EN",
			TimeZone:   "UTC+2",
			SleepStart: 280,
			SleepEnd:   12,
		}
		newParameter := Parameter{
			Local:      "vi_VI",
			TimeZone:   "UTC+6",
			SleepStart: 260,
			SleepEnd:   24,
		}
		appError := osi.Save(&parameter, ds)
		dbError := u.NewLocAppError("parameterStoreImpl.Update", "update.transaction.updates.encounterError", nil, "")
		So(appError, ShouldBeNil)
		So(appError, ShouldNotResemble, dbError)
		Convey("Providing a correct user to update", func() {
			appError := osi.Update(&parameter, &newParameter, ds)
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, dbError)
		})
		Convey("Providing an incorrect user as new should result in errors", func() {
			empty := Parameter{}
			Convey("Empty parameter should return first error from is valid error", func() {
				appError := osi.Update(&parameter, &empty, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
			})
			newParameter.Local = "en_ENG"
			Convey("Given empty local or too long local should return Local error", func() {
				appError := osi.Update(&parameter, &newParameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
				newParameter.Local = ""
				appError = osi.Update(&parameter, &newParameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
			})
			newParameter.Local = "en_EN"
			newParameter.TimeZone = "UTF+134"
			Convey("Given empty time zone or too long time zone should return Time Zone error", func() {
				appError := osi.Update(&parameter, &newParameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_timezone.app_error", nil, ""))
				newParameter.TimeZone = ""
				appError = osi.Update(&parameter, &newParameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_timezone.app_error", nil, ""))
			})
			newParameter.TimeZone = "UTF+12"
			newParameter.SleepEnd = -1
			Convey("Given negative or too big Sleep timers should return sleep error", func() {
				appError := osi.Update(&parameter, &newParameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_sleep_end.app_error", nil, ""))
				newParameter.SleepEnd = 1441
				appError = osi.Update(&parameter, &newParameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_sleep_end.app_error", nil, ""))
				newParameter.SleepEnd = 10
				newParameter.SleepStart = -10
				appError = osi.Update(&parameter, &newParameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_sleep_start.app_error", nil, ""))
				newParameter.SleepStart = 2000
				appError = osi.Update(&parameter, &newParameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterNew.PreSave", "model.parameter.is_valid.parameter_sleep_start.app_error", nil, ""))
			})
		})

		Convey("Providing an incorrect user as old should result in errors", func() {
			empty := Parameter{}
			Convey("Empty parameter should return first error from is valid error", func() {
				appError := osi.Update(&empty, &newParameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
			})
			parameter.Local = "en_ENG"
			Convey("Given empty local or too long local should return Local error", func() {
				appError := osi.Update(&parameter, &newParameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
				parameter.Local = ""
				appError = osi.Update(&parameter, &newParameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
			})
			parameter.Local = "en_EN"
			parameter.TimeZone = "UTF+134"
			Convey("Given empty time zone or too long time zone should return Time Zone error", func() {
				appError := osi.Update(&parameter, &newParameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_timezone.app_error", nil, ""))
				parameter.TimeZone = ""
				appError = osi.Update(&parameter, &newParameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_timezone.app_error", nil, ""))
			})
			parameter.TimeZone = "UTF+12"
			parameter.SleepEnd = -1
			Convey("Given negative or too big Sleep timers should return sleep error", func() {
				appError := osi.Update(&parameter, &newParameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_sleep_end.app_error", nil, ""))
				parameter.SleepEnd = 1441
				appError = osi.Update(&parameter, &newParameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_sleep_end.app_error", nil, ""))
				parameter.SleepEnd = 10
				parameter.SleepStart = -10
				appError = osi.Update(&parameter, &newParameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_sleep_start.app_error", nil, ""))
				parameter.SleepStart = 2000
				appError = osi.Update(&parameter, &newParameter, ds)
				So(appError, ShouldResemble, u.NewLocAppError("parameterStoreImpl.Update.parameterOld.PreSave", "model.parameter.is_valid.parameter_sleep_start.app_error", nil, ""))
			})
		})
		db.Delete(&parameter)
		db.Delete(&newParameter)
	})

	Convey("Testing Get function", t, func() {
		parameter := Parameter{
			Local:      "vi_VI",
			TimeZone:   "UTC+6",
			SleepStart: 260,
			SleepEnd:   24,
		}
		Convey("Trying to get parameter from empty DB should return empty", func() {
			So(&Parameter{}, ShouldResemble, osi.Get(ds))
		})
		appError := osi.Save(&parameter, ds)
		So(appError, ShouldBeNil)
		Convey("Trying to get parameter from non empty DB should return a correct parameter object", func() {
			got := osi.Get(ds)
			So(&parameter, ShouldResemble, got)
			So(got.IsValid(), ShouldBeNil)
		})
		db.Delete(&parameter)
	})
}
