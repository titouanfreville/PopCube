
package utils

import (
  "testing"

  . "github.com/smartystreets/goconvey/convey"
)

func TestUtilsPackage(t *testing.T) {
  Convey("Testing Array intersections", t, func() {
    Convey("Given an array", func() {
      a := []string{
        "abc",
        "def",
        "ghi",
      }

      empty := []string{}

      Convey("Trying intersection with an empty array or an array without common parts, it should be empty", func() {

        b := []string{
          "jkl",
          "mnp",
        }

        c := []string {
          "jkl",
        }

        So(stringArrayIntersection(a, empty), ShouldBeEmpty)
        So(stringArrayIntersection(a, b), ShouldBeEmpty)
        So(stringArrayIntersection(a, c), ShouldBeEmpty)

      })

      Convey("Trying intersection with common point should return a table containing the common elements", func() {
        b := []string {
          "jkl",
          "abc",
        }  

        c := []string {
          "def",
          "mno",
        }

        d := []string {
          "abc",
          "ghi",
          "ameno",
        }

        So(stringArrayIntersection(a, a), ShouldResemble, a)
        So(stringArrayIntersection(a, b), ShouldContain, "abc")
        So(stringArrayIntersection(a, c), ShouldContain, "def")
        So(stringArrayIntersection(a, d), ShouldContain, "abc")
        So(stringArrayIntersection(a, d), ShouldContain, "ghi")
      })  
    })
  })
}

// func TestRemoveDuplicatesFromStringArray(t *testing.T) {
//   a := []string{
//     "a",
//     "b",
//     "a",
//     "a",
//     "b",
//     "c",
//     "a",
//   }

//   if len(removeDuplicatesFromStringArray(a)) != 3 {
//     t.Fatal("should be 3")
//   }
// }