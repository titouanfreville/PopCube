// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.
// Testing base tools for DB models.-

package model

import (
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
					So(len(id),ShouldBeLessThan,27)
				}
			})
	})

	Convey("Test that random String function correctly generate string (1000 generation to test)", t, func() {
		number_of_generation := 1000
		assertion := "Checking generation of" + strconv.Itoa(number_of_generation) +" random string"
		Convey(assertion, func() {
			for i := 0; i < number_of_generation; i++ {
				r := NewRandomString(32)
				So(len(r),ShouldEqual,32)
			}
		})
	})

	Convey("Test that random String function correctly generate string (1000 generation to test)", t, func() {

	})
}

// func TestAppError(t *testing.T) {
// 	err := NewLocAppError("TestAppError", "message", nil, "")
// 	json := err.ToJson()
// 	rerr := AppErrorFromJson(strings.NewReader(json))
// 	if err.Message != rerr.Message {
// 		t.Fatal()
// 	}

// 	err.Error()
// }

// func TestAppErrorJunk(t *testing.T) {
// 	rerr := AppErrorFromJson(strings.NewReader("<html><body>This is a broken test</body></html>"))
// 	if "body: <html><body>This is a broken test</body></html>" != rerr.DetailedError {
// 		t.Fatal()
// 	}
// }

// func TestMapJson(t *testing.T) {

// 	m := make(map[string]string)
// 	m["id"] = "test_id"
// 	json := MapToJson(m)

// 	rm := MapFromJson(strings.NewReader(json))

// 	if rm["id"] != "test_id" {
// 		t.Fatal("map should be valid")
// 	}

// 	rm2 := MapFromJson(strings.NewReader(""))
// 	if len(rm2) > 0 {
// 		t.Fatal("make should be ivalid")
// 	}
// }

// func TestValidEmail(t *testing.T) {
// 	if !IsValidEmail("corey+test@hulen.com") {
// 		t.Error("email should be valid")
// 	}

// 	if IsValidEmail("@corey+test@hulen.com") {
// 		t.Error("should be invalid")
// 	}
// }

// func TestValidLower(t *testing.T) {
// 	if !IsLower("corey+test@hulen.com") {
// 		t.Error("should be valid")
// 	}

// 	if IsLower("Corey+test@hulen.com") {
// 		t.Error("should be invalid")
// 	}
// }

// func TestEtag(t *testing.T) {
// 	etag := Etag("hello", 24)
// 	if len(etag) <= 0 {
// 		t.Fatal()
// 	}
// }

// var hashtags = map[string]string{
// 	"#test":           "#test",
// 	"test":            "",
// 	"#test123":        "#test123",
// 	"#123test123":     "",
// 	"#test-test":      "#test-test",
// 	"#test?":          "#test",
// 	"hi #there":       "#there",
// 	"#bug #idea":      "#bug #idea",
// 	"#bug or #gif!":   "#bug #gif",
// 	"#hüllo":          "#hüllo",
// 	"#?test":          "",
// 	"#-test":          "",
// 	"#yo_yo":          "#yo_yo",
// 	"(#brakets)":      "#brakets",
// 	")#stekarb(":      "#stekarb",
// 	"<#less_than<":    "#less_than",
// 	">#greater_than>": "#greater_than",
// 	"-#minus-":        "#minus",
// 	"_#under_":        "#under",
// 	"+#plus+":         "#plus",
// 	"=#equals=":       "#equals",
// 	"%#pct%":          "#pct",
// 	"&#and&":          "#and",
// 	"^#hat^":          "#hat",
// 	"##brown#":        "#brown",
// 	"*#star*":         "#star",
// 	"|#pipe|":         "#pipe",
// 	":#colon:":        "#colon",
// 	";#semi;":         "#semi",
// 	"#Mötley;":        "#Mötley",
// 	".#period.":       "#period",
// 	"¿#upside¿":       "#upside",
// 	"\"#quote\"":      "#quote",
// 	"/#slash/":        "#slash",
// 	"\\#backslash\\":  "#backslash",
// 	"#a":              "",
// 	"#1":              "",
// 	"foo#bar":         "",
// }

// func TestParseHashtags(t *testing.T) {
// 	for input, output := range hashtags {
// 		if o, _ := ParseHashtags(input); o != output {
// 			t.Fatal("failed to parse hashtags from input=" + input + " expected=" + output + " actual=" + o)
// 		}
// 	}
// }