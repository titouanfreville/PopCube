// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.
// Testing base tools for DB models.-

package model

import (
	"strings"
	"testing"
	"strconv"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUtilities(t *testing.T) {
	Convey("Testing Id generation", t, func() {
		number_of_generation := 1000
		assertion := "Checking validity of " + strconv.Itoa(number_of_generation) + "random ids"

		Convey(assertion, func() {
				for i := 0; i < number_of_generation; i++ {
					id := NewId()
					So(len(id), ShouldBeLessThan, 27)
				}
			})
	})

	Convey("Test that random String function correctly generate string (1000 generation to test)", t, func() {
		number_of_generation := 1000
		assertion := "Checking generation of " + strconv.Itoa(number_of_generation) +" random string"

		Convey(assertion, func() {
			for i := 0; i < number_of_generation; i++ {
				r := NewRandomString(32)
				So(len(r), ShouldEqual, 32)
			}
		})
	})

	Convey("Testing message error formating", t, func() {

		Convey("From an error formated error generating a json formated from the error and but it back as error formated error should give the same object", func() {
			err := NewLocAppError("TestAppError", "message", nil, "")
			json := err.ToJson()
			rerr := AppErrorFromJson(strings.NewReader(json))
			err.Error()
			So(err.Message, ShouldEqual, rerr.Message);
		})

		Convey("Generating json error error message from html information should work", func() {
			rerr := AppErrorFromJson(strings.NewReader("<html><body>This is a broken test</body></html>"))
			So("body: <html><body>This is a broken test</body></html>", ShouldEqual, rerr.DetailedError)
		})
	})

	Convey("Testing Map from/to Json conversions", t, func() {

		Convey("Convert a map to json then back to map should provide same map", func() {
			m := make(map[string]string)
			m["id"] = "test_id"
			json := MapToJson(m)
			correct := MapFromJson(strings.NewReader(json))
			So(correct["id"], ShouldEqual, "test_id")
		})

		Convey("Using an empty json to generate map should provide empty map", func() {
			invalid := MapFromJson(strings.NewReader(""))
			So(len(invalid), ShouldEqual, 0)
		})
	})

	Convey("Testing email validation", t, func() {
		correct_mail := "test.test+xala@something.co"
		invalid_mail := "@test.test+xala@something.co"

		Convey("Validating a correctly formated email should be accepted", func() {
			So(IsValidEmail(correct_mail), ShouldBeTrue)
		})

		Convey("Validating a non correctly formated email should correctly be refused", func() {
			So(IsValidEmail(invalid_mail), ShouldBeFalse)
			So(IsValidEmail("Corey+test@hulen.com"), ShouldBeFalse)
		})
	})

	Convey("Testing Lower case string checker", t, func() {

		Convey("Providing a lower case test to lowercase checker should return true", func () {
			So(IsLower("corey+test@hulen.com"), ShouldBeTrue)
		})

		Convey("Providing a non lower case test to lowercase checker should return false", func () {
			So(IsLower("Corey+test@hulen.com"), ShouldBeFalse)
		})
	})

	Convey("Testing Hastags parsing ", t, func() {

		var hashtags = map[string]string{
			"#test":           "#test",
			"test":            "",
			"#test123":        "#test123",
			"#123test123":     "",
			"#test-test":      "#test-test",
			"#test?":          "#test",
			"hi #there":       "#there",
			"#bug #idea":      "#bug #idea",
			"#bug or #gif!":   "#bug #gif",
			"#hüllo":          "#hüllo",
			"#?test":          "",
			"#-test":          "",
			"#yo_yo":          "#yo_yo",
			"(#brakets)":      "#brakets",
			")#stekarb(":      "#stekarb",
			"<#less_than<":    "#less_than",
			">#greater_than>": "#greater_than",
			"-#minus-":        "#minus",
			"_#under_":        "#under",
			"+#plus+":         "#plus",
			"=#equals=":       "#equals",
			"%#pct%":          "#pct",
			"&#and&":          "#and",
			"^#hat^":          "#hat",
			"##brown#":        "#brown",
			"*#star*":         "#star",
			"|#pipe|":         "#pipe",
			":#colon:":        "#colon",
			";#semi;":         "#semi",
			"#Mötley;":        "#Mötley",
			".#period.":       "#period",
			"¿#upside¿":       "#upside",
			"\"#quote\"":      "#quote",
			"/#slash/":        "#slash",
			"\\#backslash\\":  "#backslash",
			"#a":              "",
			"#1":              "",
			"foo#bar":         "",
		}

		Convey("A text containing or not containing # should be correctly parse", func() {
			for input, output := range hashtags {
				o, _ := ParseHashtags(input)
				So(o, ShouldEqual, output)
			}
		})
	})
}

// func TestEtag(t *testing.T) {
// 	etag := Etag("hello", 24)
// 	if len(etag) <= 0 {
// 		t.Fatal()
// 	}
// }