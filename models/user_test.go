// This file is used to test if user model is working correctly.
// A user is always linked to an organisation
// He has basic channel to join

package model

import (
	"strings"
	"testing"
	"strconv"
	. "github.com/smartystreets/goconvey/convey"
)

// Test correction test for user ;)


// Test Password functionalities from User Model
func TestUserModel(t *testing.T) {
	Convey("Testing password management ...", t, func() {
		Convey("Given a password", func() {
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
	})

	Convey("Testing user format", t, func() {
		Convey("Given an user", func() {
			user := User{Id: NewId(), Username: NewId()}
			Convey("Converting user to json then json to user should provide same user information", func() {
				json := user.ToJson()
				test_user := UserFromJson(strings.NewReader(json))
				So(user.Id, ShouldEqual, test_user.Id)
				So(user.Username, ShouldEqual, test_user.Username)
			})
		})
	})

	Convey("Testing Pre Save and Pre Update function", t, func() {
		user1 := User{Password: "test"}
		Convey("Given an incomplete user", func() {
			user := User{Password: "test"}
			Convey("Applying PreSave should fill required fields", func() {
				user.PreSave()
				So(user.Id, ShouldNotBeBlank)
				So(user.Username, ShouldNotBeBlank)
				So(user.EmailVerified, ShouldBeFalse)
				So(user.Deleted, ShouldBeFalse)
				So(user.UpdatedAt, ShouldNotBeNil)
				So(user.UpdatedAt, ShouldBeGreaterThan, 0)
				So(user.UpdatedAt, ShouldEqual, user.LastPasswordUpdate)
				So(user.Locale, ShouldNotBeBlank)
				So(user.Channels, ShouldNotBeNil)
				So(len(user.Channels), ShouldBeGreaterThan, 0)
				So(ComparePassword(user.Password, "test"), ShouldBeTrue)
			})

			Convey("Data should be correctly formated", func() {
				user.PreSave()
				So(IsLower(user.Username),ShouldBeTrue)
				So(IsLower(user.Email),ShouldBeTrue)
			})

			Convey("Etag should be correctly generated", func() {
				user.PreSave()
				etag := user.Etag(true,true)
				expected := CURRENT_VERSION + "." + user.Id + "." + strconv.FormatInt(user.UpdatedAt, 10) + "." + "true" + "." + "true"
				So(etag, ShouldEqual, expected)
			})
		})

		Convey("Given an user with email and username", func() {
			user := User{Password: "test", Username: "TesT", Email: "Test@poPcube.fr"}
			Convey("Applying PreSave should fill blank required fields and concerve overs", func() {
				user.PreSave()
				So(user.Id, ShouldNotBeBlank)
				So(user.Username, ShouldEqual, "test")
				So(user.Email, ShouldEqual, "test@popcube.fr")
				So(user.EmailVerified, ShouldBeFalse)
				So(user.Deleted, ShouldBeFalse)
				So(user.UpdatedAt, ShouldNotBeNil)
				So(user.UpdatedAt, ShouldBeGreaterThan, 0)
				So(user.UpdatedAt, ShouldEqual, user.LastPasswordUpdate)
				So(user.Locale, ShouldNotBeBlank)
				So(user.Channels, ShouldNotBeNil)
				So(len(user.Channels), ShouldBeGreaterThan, 0)
				So(ComparePassword(user.Password, "test"), ShouldBeTrue)
			})

			Convey("Data should be correctly formated", func() {
				user.PreSave()
				So(IsLower(user.Username),ShouldBeTrue)
				So(IsLower(user.Email),ShouldBeTrue)
			})

			Convey("Etag should be correctly generated", func() {
				user.PreSave()
				etag := user.Etag(true,true)
				expected := CURRENT_VERSION + "." + user.Id + "." + strconv.FormatInt(user.UpdatedAt, 10) + "." + "true" + "." + "true"
				So(etag, ShouldEqual, expected)
			})
		})

		Convey("Given a full user entry", func() {
			user := User{
				Id: "testId",
				UpdatedAt: 10,
				Deleted: true,
				Username: "TesT",
				Password: "test",
				Email: "Test@poPcube.fr",
				EmailVerified: true,
				Nickname: "NickName",
				FirstName: "Test",
				LastName: "L",
				Roles: "Owner",
				LastPasswordUpdate: 20,
				FailedAttempts: 1,
				Locale: "vi",
				Channels: []string{"chef", "is", "back"},
				PrivateChannels: []string{"Newbie"},
				LastActivityAt: 5,
			}

			Convey("Applying PreSave should only correctly format field and use good time for last Updates", func() {
				user.PreSave()
				So(user.Id, ShouldEqual, "testId")
				So(user.UpdatedAt, ShouldNotEqual, 10)
				So(user.Deleted, ShouldBeTrue)
				So(user.Username, ShouldEqual, "test")
				So(ComparePassword(user.Password, "test"), ShouldBeTrue)
				So(user.Email, ShouldEqual, "test@popcube.fr")
				So(user.EmailVerified, ShouldBeTrue)
				So(user.Nickname, ShouldEqual, "NickName")
				So(user.FirstName, ShouldEqual, "Test")
				So(user.LastName, ShouldEqual, "L")
				So(user.Roles, ShouldEqual, "Owner")
				So(user.LastPasswordUpdate, ShouldNotEqual, 20)
				So(user.LastPasswordUpdate, ShouldEqual, user.UpdatedAt)
				So(user.FailedAttempts, ShouldEqual, 1)
				So(user.Locale, ShouldEqual, "vi")
				So(user.Channels, ShouldResemble, []string{"chef", "is", "back"})
				So(user.PrivateChannels, ShouldResemble, []string{"Newbie"})
				So(user.LastActivityAt, ShouldEqual, 5)
			})

			Convey("Etag should be correctly generated", func() {
				user.PreSave()
				etag := user.Etag(true,true)
				expected := CURRENT_VERSION + "." + user.Id + "." + strconv.FormatInt(user.UpdatedAt, 10) + "." + "true" + "." + "true"
				So(etag, ShouldEqual, expected)
			})
		})

		Convey("Given an user.", func() {
			oldUpdated := user1.UpdatedAt
			user1.Password = "NewPassword"
			user1.PreSave()

			Convey("Applying PreSave should correctly update values", func() {
				So(ComparePassword(user1.Password, "NewPassword"), ShouldBeTrue)
				So(user1.UpdatedAt, ShouldBeGreaterThan, oldUpdated)
			})

			Convey("Applying PreSave should correctly format values", func() {
				So(IsLower(user1.Username), ShouldBeTrue)
				So(IsLower(user1.Email), ShouldBeTrue)
			})
		})
	})

	Convey("Testing fonction IsValid", t, func() {
		Convey("Given a correct user, validation should work", func() {
			correct_user := User{
				Username: "TesT",
				Password: "test",
				Email: "test@popcube.fr",
				Nickname: "NickName",
				FirstName: "Test",
				LastName: "L",
				Roles: "Owner",
			}
			correct_user.PreSave()
			So(correct_user.IsValid(),ShouldNotBeNil);
		})
		Convey("Given an incorrect user, validation should return error message", func() {
			Convey("Incorrect ID user should return a message invalid id", func() {
				user := User{
					Id: "Nimp",
					Username: "TesT",
					Password: "test",
					Email: "test@popcube.fr",
					Nickname: "NickName",
					FirstName: "Test",
					LastName: "L",
					Roles: "Owner",
				}
				So(user.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.id.app_error", nil, ""))
			})
			Convey("Incorrect Username user should return error Invalid username", func() {
				user1 := User{
					Username: "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone",
					Password: "test",
					Email: "test@popcube.fr",
					Nickname: "NickName",
					FirstName: "Test",
					LastName: "L",
					Roles: "Owner",
				}
				user1.PreSave()
				So(user1.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user1.Id))
				user2 := User{
					Id: NewId(),
					Password: "test",
					Email: "test@popcube.fr",
					Nickname: "NickName",
					FirstName: "Test",
					LastName: "L",
					Roles: "Owner",
				}
				So(user2.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user2.Id))
				user3 := User{
					Id: NewId(),
					Username: "xD/",
					Password: "test",
					Email: "test@popcube.fr",
					Nickname: "NickName",
					FirstName: "Test",
					LastName: "L",
					Roles: "Owner",
				}
				So(user3.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user3.Id))
				user3 = User{
					Id: NewId(),
					Username: "xD\\",
					Password: "test",
					Email: "test@popcube.fr",
					Nickname: "NickName",
					FirstName: "Test",
					LastName: "L",
					Roles: "Owner",
				}
				So(user3.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user3.Id))
				user3 = User{
					Id: NewId(),
					Username: "xD*",
					Password: "test",
					Email: "test@popcube.fr",
					Nickname: "NickName",
					FirstName: "Test",
					LastName: "L",
					Roles: "Owner",
				}
				So(user3.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user3.Id))
				user3 = User{
					Id: NewId(),
					Username: "xD{",
					Password: "test",
					Email: "test@popcube.fr",
					Nickname: "NickName",
					FirstName: "Test",
					LastName: "L",
					Roles: "Owner",
				}
				So(user3.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user3.Id))
				user3 = User{
					Id: NewId(),
					Username: "xD}",
					Password: "test",
					Email: "test@popcube.fr",
					Nickname: "NickName",
					FirstName: "Test",
					LastName: "L",
					Roles: "Owner",
				}
				So(user3.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user3.Id))
				user3 = User{
					Id: NewId(),
					Username: "xD#",
					Password: "test",
					Email: "test@popcube.fr",
					Nickname: "NickName",
					FirstName: "Test",
					LastName: "L",
					Roles: "Owner",
				}
				So(user3.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user3.Id))
				user3 = User{
					Id: NewId(),
					Username: "xD_",
					Password: "test",
					Email: "test@popcube.fr",
					Nickname: "NickName",
					FirstName: "Test",
					LastName: "L",
					Roles: "Owner",
				}
				So(user3.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user3.Id))
			})
		})

		Convey("Incorrect Email user should return error Invalid email", func() {
			user := User{
				Password: "test",
				Email: "testpopcube.fr",
				Nickname: "NickName",
				FirstName: "Test",
				LastName: "L",
				Roles: "Owner",
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.email.app_error", nil, "user_id="+user.Id))
			user = User{
				Password: "test",
				Email: "test/popcube.fr",
				Nickname: "NickName",
				FirstName: "Test",
				LastName: "L",
				Roles: "Owner",
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.email.app_error", nil, "user_id="+user.Id))
			user = User{
				Password: "test",
				Email: "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone@popcube.fr",
				Nickname: "NickName",
				FirstName: "Test",
				LastName: "L",
				Roles: "Owner",
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.email.app_error", nil, "user_id="+user.Id))
		})

		Convey("Nickname, Firstanem and Lastname should be less than 64 characters long", func() {
			user := User{
				Password: "test",
				Email: "testpopcube.fr",
				Nickname: "NickName",
				FirstName: "Test",
				LastName: "L",
				Roles: "Owner",
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.email.app_error", nil, "user_id="+user.Id))
	})
}


// func TestUserUpdateMentionKeysFromUsername(t *testing.T) {
// 	user := User{Username: "user"}
// 	user.SetDefaultNotifications()

// 	if user.NotifyProps["mention_keys"] != "user,@user" {
// 		t.Fatal("default mention keys are invalid: %v", user.NotifyProps["mention_keys"])
// 	}

// 	user.Username = "person"
// 	user.UpdateMentionKeysFromUsername("user")
// 	if user.NotifyProps["mention_keys"] != "person,@person" {
// 		t.Fatal("mention keys are invalid after changing username: %v", user.NotifyProps["mention_keys"])
// 	}

// 	user.NotifyProps["mention_keys"] += ",mention"
// 	user.UpdateMentionKeysFromUsername("person")
// 	if user.NotifyProps["mention_keys"] != "person,@person,mention" {
// 		t.Fatal("mention keys are invalid after adding extra mention keyword: %v", user.NotifyProps["mention_keys"])
// 	}

// 	user.Username = "user"
// 	user.UpdateMentionKeysFromUsername("person")
// 	if user.NotifyProps["mention_keys"] != "user,@user,mention" {
// 		t.Fatal("mention keys are invalid after changing username with extra mention keyword: %v", user.NotifyProps["mention_keys"])
// 	}
// }

// func TestUserIsValid(t *testing.T) {
// 	user := User{}

// 	if err := user.IsValid(); err == nil {
// 		t.Fatal()
// 	}

// 	user.Id = NewId()
// 	if err := user.IsValid(); err == nil {
// 		t.Fatal()
// 	}

// 	user.CreateAt = GetMillis()
// 	if err := user.IsValid(); err == nil {
// 		t.Fatal()
// 	}

// 	user.UpdateAt = GetMillis()
// 	if err := user.IsValid(); err == nil {
// 		t.Fatal()
// 	}

// 	user.Username = NewId() + "^hello#"
// 	if err := user.IsValid(); err == nil {
// 		t.Fatal()
// 	}

// 	user.Username = NewId()
// 	user.Email = strings.Repeat("01234567890", 20)
// 	if err := user.IsValid(); err == nil {
// 		t.Fatal()
// 	}

// 	user.Email = "test@nowhere.com"
// 	user.Nickname = strings.Repeat("01234567890", 20)
// 	if err := user.IsValid(); err == nil {
// 		t.Fatal()
// 	}

// 	user.Nickname = ""
// 	if err := user.IsValid(); err != nil {
// 		t.Fatal(err)
// 	}

// 	user.FirstName = ""
// 	user.LastName = ""
// 	if err := user.IsValid(); err != nil {
// 		t.Fatal(err)
// 	}

// 	user.FirstName = strings.Repeat("01234567890", 20)
// 	if err := user.IsValid(); err == nil {
// 		t.Fatal(err)
// 	}

// 	user.FirstName = ""
// 	user.LastName = strings.Repeat("01234567890", 20)
// 	if err := user.IsValid(); err == nil {
// 		t.Fatal(err)
// 	}
// }

// func TestUserGetFullName(t *testing.T) {
// 	user := User{}

// 	if fullName := user.GetFullName(); fullName != "" {
// 		t.Fatal("Full name should be blank")
// 	}

// 	user.FirstName = "first"
// 	if fullName := user.GetFullName(); fullName != "first" {
// 		t.Fatal("Full name should be first name")
// 	}

// 	user.FirstName = ""
// 	user.LastName = "last"
// 	if fullName := user.GetFullName(); fullName != "last" {
// 		t.Fatal("Full name should be last name")
// 	}

// 	user.FirstName = "first"
// 	if fullName := user.GetFullName(); fullName != "first last" {
// 		t.Fatal("Full name should be first name and last name")
// 	}
// }

// func TestUserGetDisplayName(t *testing.T) {
// 	user := User{Username: "user"}

// 	if displayName := user.GetDisplayName(); displayName != "user" {
// 		t.Fatal("Display name should be username")
// 	}

// 	user.FirstName = "first"
// 	user.LastName = "last"
// 	if displayName := user.GetDisplayName(); displayName != "first last" {
// 		t.Fatal("Display name should be full name")
// 	}

// 	user.Nickname = "nickname"
// 	if displayName := user.GetDisplayName(); displayName != "nickname" {
// 		t.Fatal("Display name should be nickname")
// 	}
// }

// var usernames = []struct {
// 	value    string
// 	expected bool
// }{
// 	{"spin-punch", true},
// 	{"Spin-punch", false},
// 	{"spin punch-", false},
// 	{"spin_punch", true},
// 	{"spin", true},
// 	{"PUNCH", false},
// 	{"spin.punch", true},
// 	{"spin'punch", false},
// 	{"spin*punch", false},
// 	{"all", false},
// }

// func TestValidUsername(t *testing.T) {
// 	for _, v := range usernames {
// 		if IsValidUsername(v.value) != v.expected {
// 			t.Errorf("expect %v as %v", v.value, v.expected)
// 		}
// 	}
// }

// func TestCleanUsername(t *testing.T) {
// 	if CleanUsername("Spin-punch") != "spin-punch" {
// 		t.Fatal("didn't clean name properly")
// 	}
// 	if CleanUsername("PUNCH") != "punch" {
// 		t.Fatal("didn't clean name properly")
// 	}
// 	if CleanUsername("spin'punch") != "spin-punch" {
// 		t.Fatal("didn't clean name properly")
// 	}
// 	if CleanUsername("spin") != "spin" {
// 		t.Fatal("didn't clean name properly")
// 	}
// 	if len(CleanUsername("all")) != 27 {
// 		t.Fatal("didn't clean name properly")
// 	}
// }

// func TestRoles(t *testing.T) {

// 	if IsValidUserRoles("admin") {
// 		t.Fatal()
// 	}

// 	if IsValidUserRoles("junk") {
// 		t.Fatal()
// 	}

// 	if !IsValidUserRoles("system_user system_admin") {
// 		t.Fatal()
// 	}

// 	if IsInRole("system_admin junk", "admin") {
// 		t.Fatal()
// 	}

// 	if !IsInRole("system_admin junk", "system_admin") {
// 		t.Fatal()
// 	}

// 	if IsInRole("admin", "system_admin") {
// 		t.Fatal()
// 	}
// }