package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"strings"
	"testing"
)

func TestParameterModel(t *testing.T) {
	Convey("Testing json vs parameter conversions", t, func() {
		Convey("Given a parameter", func() {
			parameter := Parameter{
				Local:      "en_EN",
				TimeZone:   "UTC+2",
				SleepStart: 280,
				SleepEnd:   12,
			}
			Convey("Converting parameter to json then json to parameter should provide same parameter information", func() {
				json := parameter.toJson()
				test_parameter := parameterFromJson(strings.NewReader(json))
				So(parameter.Local, ShouldEqual, test_parameter.Local)
				So(parameter.TimeZone, ShouldEqual, test_parameter.TimeZone)
				So(parameter.SleepStart, ShouldEqual, test_parameter.SleepStart)
				So(parameter.SleepEnd, ShouldEqual, test_parameter.SleepEnd)
			})
		})
	})

	Convey("Testing isValid function", t, func() {
		Convey("Given a correct parameter. Parameter should be validate", func() {
			parameter := Parameter{
				Local:      "en_EN",
				TimeZone:   "UTC+2",
				SleepStart: 280,
				SleepEnd:   12,
			}
			So(parameter.isValid(), ShouldBeNil)
			So(parameter.isValid(), ShouldNotResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_name.app_error", nil, "id="+strconv.FormatUint(parameter.IdParameter, 10)))
			So(parameter.isValid(), ShouldNotResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.not_alphanum_parameter_name.app_error", nil, "id="+strconv.FormatUint(parameter.IdParameter, 10)))
			So(parameter.isValid(), ShouldNotResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.description.app_error", nil, "id="+strconv.FormatUint(parameter.IdParameter, 10)))
		})
		Convey("Given an incorrect parameter. Parameter should be refused", func() {
			empty := Parameter{}
			parameter := Parameter{
				Local:      "en_EN",
				TimeZone:   "UTC+2",
				SleepStart: 280,
				SleepEnd:   12,
			}
			Convey("Empty parameter should return first error from is valid error", func() {
				So(empty.isValid(), ShouldResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
			})
			parameter.Local = "en_ENG"
			Convey("Given empty local or too long local should return Local error", func() {
				So(parameter.isValid(), ShouldResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
				parameter.Local = ""
				So(parameter.isValid(), ShouldResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_local.app_error", nil, ""))
			})
			parameter.Local = "en_EN"
			parameter.TimeZone = "UTF+134"
			Convey("Given empty time zone or too long time zone should return Time Zone error", func() {
				So(parameter.isValid(), ShouldResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_timezone.app_error", nil, ""))
				parameter.TimeZone = ""
				So(parameter.isValid(), ShouldResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_timezone.app_error", nil, ""))
			})
			parameter.TimeZone = "UTF+12"
			parameter.SleepEnd = -1
			Convey("Given negative or too big Sleep timers should return sleep error", func() {
				So(parameter.isValid(), ShouldResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_sleep_end.app_error", nil, ""))
				parameter.SleepEnd = 1441
				So(parameter.isValid(), ShouldResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_sleep_end.app_error", nil, ""))
				parameter.SleepEnd = 10
				parameter.SleepStart = -10
				So(parameter.isValid(), ShouldResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_sleep_start.app_error", nil, ""))
				parameter.SleepStart = 2000
				So(parameter.isValid(), ShouldResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_sleep_start.app_error", nil, ""))
			})
		})
	})

	Convey("Testing PreSave function", t, func() {
		parameter := Parameter{}
		Convey("If parameter is empty, should fill some fields - webId, ParameterName, UpdatedAt, Avatar and type, and parameter should be valid", func() {
			parameter.preSave()
			So(parameter.isValid(), ShouldBeNil)
			So(parameter.isValid(), ShouldNotResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_name.app_error", nil, "id="+strconv.FormatUint(parameter.IdParameter, 10)))
			So(parameter.isValid(), ShouldNotResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.not_alphanum_parameter_name.app_error", nil, "id="+strconv.FormatUint(parameter.IdParameter, 10)))
			So(parameter.isValid(), ShouldNotResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.description.app_error", nil, "id="+strconv.FormatUint(parameter.IdParameter, 10)))
			So(parameter.Local, ShouldEqual, "fr_FR")
			So(parameter.TimeZone, ShouldEqual, "UTC-0")
		})
		Convey("If parameter is filled, nothing should happen", func() {
			parameter = Parameter{
				Local:      "en_EN",
				TimeZone:   "UTC+2",
				SleepStart: 280,
				SleepEnd:   12,
			}
			So(parameter.isValid(), ShouldBeNil)
			So(parameter.isValid(), ShouldNotResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_name.app_error", nil, "id="+strconv.FormatUint(parameter.IdParameter, 10)))
			So(parameter.isValid(), ShouldNotResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.not_alphanum_parameter_name.app_error", nil, "id="+strconv.FormatUint(parameter.IdParameter, 10)))
			So(parameter.isValid(), ShouldNotResemble, NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.description.app_error", nil, "id="+strconv.FormatUint(parameter.IdParameter, 10)))
			So(parameter.Local, ShouldEqual, "en_EN")
			So(parameter.TimeZone, ShouldEqual, "UTC+2")
			So(parameter.SleepStart, ShouldEqual, 280)
			So(parameter.SleepEnd, ShouldEqual, 12)
		})
	})
}
