// This file is used to test if user model is working correctly.
// A user is always linked to an organisation
// He has basic channel to join

package model

import (
	// "strings"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

// Test correction test for user ;)


// Test Password functionalities from User Model
func TestPasswordHash(t *testing.T) {
	Convey("Given a password", t, func() {
		hash := HashPassword("Test")

		Convey("Compare it with correct entry shoud be true", func () {
			So(ComparePassword(hash, "Test"), ShouldBeTrue)
		})

		Convey("Compare it with correct entry shoud be false", func () {
			So(ComparePassword(hash, "Test1"), ShouldBeFalse)
		})

		Convey("Compare it with empty entry shoud be false", func () {
			So(ComparePassword(hash, ""), ShouldBeFalse)
		})
	})
}