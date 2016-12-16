// This file is used to test if user model is working correctly.
// A user is always linked to an organisation
// He has basic channel to join

package model

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Test correction test for user ;)

// Test Password functionalities from User Model
func TestUserModel(t *testing.T) {
	Convey("Testing password management ...", t, func() {
		Convey("Given a password", func() {
			hash := HashPassword("Test")

			Convey("Compare it with correct entry shoud be true", func() {
				So(ComparePassword(hash, "Test"), ShouldBeTrue)
			})

			Convey("Compare it with correct entry shoud be false", func() {
				So(ComparePassword(hash, "Test1"), ShouldBeFalse)
			})

			Convey("Compare it with empty entry shoud be false", func() {
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
				So(IsLower(user.Username), ShouldBeTrue)
				So(IsLower(user.Email), ShouldBeTrue)
			})

			Convey("Etag should be correctly generated", func() {
				user.PreSave()
				etag := user.Etag(true, true)
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
				So(IsLower(user.Username), ShouldBeTrue)
				So(IsLower(user.Email), ShouldBeTrue)
			})

			Convey("Etag should be correctly generated", func() {
				user.PreSave()
				etag := user.Etag(true, true)
				expected := CURRENT_VERSION + "." + user.Id + "." + strconv.FormatInt(user.UpdatedAt, 10) + "." + "true" + "." + "true"
				So(etag, ShouldEqual, expected)
			})
		})

		Convey("Given a full user entry", func() {
			user := User{
				Id:                 "testId",
				UpdatedAt:          10,
				Deleted:            true,
				Username:           "TesT",
				Password:           "test",
				Email:              "Test@poPcube.fr",
				EmailVerified:      true,
				Nickname:           "NickName",
				FirstName:          "Test",
				LastName:           "L",
				Roles:              "Owner",
				LastPasswordUpdate: 20,
				FailedAttempts:     1,
				Locale:             "vi",
				Channels:           []string{"chef", "is", "back"},
				PrivateChannels:    []string{"Newbie"},
				LastActivityAt:     5,
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
				etag := user.Etag(true, true)
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
				Username:  "TesT",
				Password:  "test",
				Email:     "test@popcube.fr",
				Nickname:  "NickName",
				FirstName: "Test",
				LastName:  "L",
				Roles:     "Owner",
			}
			correct_user.PreSave()
			So(correct_user.IsValid(), ShouldBeNil)
			So(correct_user.IsValid(), ShouldNotResemble, NewLocAppError("User.IsValid", "model.user.is_valid.id.app_error", nil, ""))
			So(correct_user.IsValid(), ShouldNotResemble, NewLocAppError("User.IsValid", "model.user.is_valid.email.app_error", nil, "user_id="+correct_user.Id))
			So(correct_user.IsValid(), ShouldNotResemble, NewLocAppError("User.IsValid", "model.user.is_valid.nickname.app_error", nil, "user_id="+correct_user.Id))
			So(correct_user.IsValid(), ShouldNotResemble, NewLocAppError("User.IsValid", "model.user.is_valid.first_name.app_error", nil, "user_id="+correct_user.Id))
			So(correct_user.IsValid(), ShouldNotResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+correct_user.Id))
			So(correct_user.IsValid(), ShouldNotResemble, NewLocAppError("User.IsValid", "model.user.is_valid.last_name.app_error", nil, "user_id="+correct_user.Id))
			So(correct_user.IsValid(), ShouldNotResemble, NewLocAppError("User.IsValid", "model.user.auth_data_pwd.username.app_error", nil, "user_id="+correct_user.Id))
		})
		Convey("Given an incorrect user, validation should return error message", func() {
			Convey("Incorrect ID user should return a message invalid id", func() {
				user := User{
					Id:        "Nimp",
					Username:  "TesT",
					Password:  "test",
					Email:     "test@popcube.fr",
					Nickname:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Roles:     "Owner",
				}
				So(user.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.id.app_error", nil, ""))
			})
			Convey("Incorrect Username user should return error Invalid username", func() {
				user1 := User{
					Username:  "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone",
					Password:  "test",
					Email:     "test@popcube.fr",
					Nickname:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Roles:     "Owner",
				}
				user1.PreSave()
				So(user1.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user1.Id))
				user2 := User{
					Id:        NewId(),
					Password:  "test",
					Email:     "test@popcube.fr",
					Nickname:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Roles:     "Owner",
				}
				So(user2.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user2.Id))
				user3 := User{
					Id:        NewId(),
					Username:  "xD/",
					Password:  "test",
					Email:     "test@popcube.fr",
					Nickname:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Roles:     "Owner",
				}
				So(user3.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user3.Id))
				user3 = User{
					Id:        NewId(),
					Username:  "xD\\",
					Password:  "test",
					Email:     "test@popcube.fr",
					Nickname:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Roles:     "Owner",
				}
				So(user3.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user3.Id))
				user3 = User{
					Id:        NewId(),
					Username:  "xD*",
					Password:  "test",
					Email:     "test@popcube.fr",
					Nickname:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Roles:     "Owner",
				}
				So(user3.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user3.Id))
				user3 = User{
					Id:        NewId(),
					Username:  "xD{",
					Password:  "test",
					Email:     "test@popcube.fr",
					Nickname:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Roles:     "Owner",
				}
				So(user3.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user3.Id))
				user3 = User{
					Id:        NewId(),
					Username:  "xD}",
					Password:  "test",
					Email:     "test@popcube.fr",
					Nickname:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Roles:     "Owner",
				}
				So(user3.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user3.Id))
				user3 = User{
					Id:        NewId(),
					Username:  "xD#",
					Password:  "test",
					Email:     "test@popcube.fr",
					Nickname:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Roles:     "Owner",
				}
				So(user3.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user3.Id))
				user3 = User{
					Id:        NewId(),
					Username:  "xD_",
					Password:  "test",
					Email:     "test@popcube.fr",
					Nickname:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Roles:     "Owner",
				}
				So(user3.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.username.app_error", nil, "user_id="+user3.Id))
			})
		})

		Convey("Incorrect Email user should return error Invalid email", func() {
			user := User{
				Password:  "test",
				Email:     "testpopcube.fr",
				Nickname:  "NickName",
				FirstName: "Test",
				LastName:  "L",
				Roles:     "Owner",
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.email.app_error", nil, "user_id="+user.Id))
			user = User{
				Password:  "test",
				Email:     "test/popcube.fr",
				Nickname:  "NickName",
				FirstName: "Test",
				LastName:  "L",
				Roles:     "Owner",
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.email.app_error", nil, "user_id="+user.Id))
			user = User{
				Password:  "test",
				Email:     "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone@popcube.fr",
				Nickname:  "NickName",
				FirstName: "Test",
				LastName:  "L",
				Roles:     "Owner",
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.email.app_error", nil, "user_id="+user.Id))
		})

		Convey("Nickname, Firstname and Lastname should be less than 64 characters long", func() {
			user := User{
				Password:  "test",
				Email:     "test@popcube.fr",
				Nickname:  "NickNameéèéééééééééééétroplongazdazdzadazdazdzadz_>_<azdazdzadazdazz",
				FirstName: "Test",
				LastName:  "L",
				Roles:     "Owner",
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.nickname.app_error", nil, "user_id="+user.Id))
			user = User{
				Password:  "test",
				Email:     "test@popcube.fr",
				Nickname:  "Nickname",
				FirstName: "TestéèéèéèéèèéèéèéèéèéèèéèéèéèèéèéèNJnefiznfidsdfnpdsjfazddrfazdzadzadzadzadazd",
				LastName:  "L",
				Roles:     "Owner",
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.first_name.app_error", nil, "user_id="+user.Id))
			user = User{
				Password:  "test",
				Email:     "test@popcube.fr",
				Nickname:  "Nickname",
				FirstName: "Test",
				LastName:  "TestéèéèéèéèèéèéèéèéèéèèéèéèéèèéèéèNJnefiznfidsdfdazdzadzadzadzadzadzadazdazdazdzadazdzanpdsjf",
				Roles:     "Owner",
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.last_name.app_error", nil, "user_id="+user.Id))
		})

		Convey("Password can]t be empty", func() {
			user := User{
				Email:     "test@popcube.fr",
				Nickname:  "Nickname",
				FirstName: "Test",
				LastName:  "L",
				Roles:     "Owner",
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, NewLocAppError("User.IsValid", "model.user.is_valid.auth_data_pwd.app_error", nil, "user_id="+user.Id))
		})
	})

	Convey("Testing Full Name getter", t, func() {
		Convey("Providing an user without full name should return an empty string", func() {
			user := User{}
			So(user.GetFullName(), ShouldBeBlank)
			user.Password = "test"
			user.PreSave()
			So(user.GetFullName(), ShouldBeBlank)
		})
		Convey("Providing only first name should return a string containing only first name", func() {
			user := User{
				FirstName: "Test",
			}
			So(user.GetFullName(), ShouldEqual, "Test")
			user.Password = "test"
			user.PreSave()
			So(user.GetFullName(), ShouldEqual, "Test")
		})
		Convey("Providing only last name should return a string containing only first name", func() {
			user := User{
				LastName: "Test",
			}
			So(user.GetFullName(), ShouldEqual, "Test")
			user.Password = "test"
			user.PreSave()
			So(user.GetFullName(), ShouldEqual, "Test")
		})
		Convey("Providing both first and last name should return a string containing first then last name", func() {
			user := User{
				LastName:  "Last",
				FirstName: "First",
			}
			So(user.GetFullName(), ShouldEqual, "First Last")
			user.Password = "test"
			user.PreSave()
			So(user.GetFullName(), ShouldEqual, "First Last")
		})

	})

	Convey("Testing GetDisplayName function", t, func() {
		Convey("Given a correct user", func() {
			u := User{Password: "test", Username: "test"}
			u.PreSave()
			Convey("User without First/Last/Nick name should have Username as display name", func() {
				So(u.GetDisplayName(), ShouldEqual, "test")
			})
			Convey("User with First/Last name but no nickname should have full name as displayname", func() {
				u.LastName = "Troll"
				So(u.GetDisplayName(), ShouldEqual, "Troll")
				u.FirstName = "Min"
				So(u.GetDisplayName(), ShouldEqual, "Min Troll")
				u.LastName = ""
				So(u.GetDisplayName(), ShouldEqual, "Min")
			})
			Convey("User having a nickname should have their nickname diplayed", func() {
				u.Nickname = "nOOb"
				So(u.GetDisplayName(), ShouldEqual, "nOOb")
			})
		})
	})

	Convey("Testing IsValidUsername function", t, func() {
		Convey("Given an user name :", func() {
			Convey("Containing caps -> refused", func() {
				So(IsValidUsername("IamContaingCaps"), ShouldBeFalse)
				So(IsValidUsername("amContaingCaps"), ShouldBeFalse)
				So(IsValidUsername("FULLCAPS"), ShouldBeFalse)
				So(IsValidUsername("capsattheenD"), ShouldBeFalse)
			})
			Convey("Reserved -> refused", func() {
				for _, uname := range restrictedUsernames {
					So(IsValidUsername(uname), ShouldBeFalse)
				}
			})
			Convey("Containing illegal characters ( * ] \\ space ( ) { } [ ] .... -> refused)", func() {
				So(IsValidUsername("i contain spaces"), ShouldBeFalse)
				So(IsValidUsername("one space"), ShouldBeFalse)
				So(IsValidUsername(" "), ShouldBeFalse)
				So(IsValidUsername("iama*"), ShouldBeFalse)
				So(IsValidUsername("*"), ShouldBeFalse)
				So(IsValidUsername("some*things"), ShouldBeFalse)
				So(IsValidUsername("]"), ShouldBeFalse)
				So(IsValidUsername("]citation"), ShouldBeFalse)
				So(IsValidUsername("ci]tation"), ShouldBeFalse)
				So(IsValidUsername("citation]"), ShouldBeFalse)
				So(IsValidUsername("{"), ShouldBeFalse)
				So(IsValidUsername("{citation"), ShouldBeFalse)
				So(IsValidUsername("ci{tation"), ShouldBeFalse)
				So(IsValidUsername("citation{"), ShouldBeFalse)
				So(IsValidUsername("}"), ShouldBeFalse)
				So(IsValidUsername("}citation"), ShouldBeFalse)
				So(IsValidUsername("ci}tation"), ShouldBeFalse)
				So(IsValidUsername("citation}"), ShouldBeFalse)
				So(IsValidUsername("("), ShouldBeFalse)
				So(IsValidUsername("(citation"), ShouldBeFalse)
				So(IsValidUsername("ci(tation"), ShouldBeFalse)
				So(IsValidUsername("citation("), ShouldBeFalse)
				So(IsValidUsername(")"), ShouldBeFalse)
				So(IsValidUsername(")citation"), ShouldBeFalse)
				So(IsValidUsername("ci)tation"), ShouldBeFalse)
				So(IsValidUsername("citation)"), ShouldBeFalse)
				So(IsValidUsername("["), ShouldBeFalse)
				So(IsValidUsername("[citation"), ShouldBeFalse)
				So(IsValidUsername("ci[tation"), ShouldBeFalse)
				So(IsValidUsername("citation["), ShouldBeFalse)
				So(IsValidUsername("]"), ShouldBeFalse)
				So(IsValidUsername("]citation"), ShouldBeFalse)
				So(IsValidUsername("ci]tation"), ShouldBeFalse)
				So(IsValidUsername("citation]"), ShouldBeFalse)
				So(IsValidUsername("\\"), ShouldBeFalse)
				So(IsValidUsername("\\citation"), ShouldBeFalse)
				So(IsValidUsername("ci\\tation"), ShouldBeFalse)
				So(IsValidUsername("citation\\"), ShouldBeFalse)
			})
			Convey("Correct -> accepted", func() {
				So(IsValidUsername("je-suis"), ShouldBeTrue)
				So(IsValidUsername("je_suis"), ShouldBeTrue)
				So(IsValidUsername("je-suis_"), ShouldBeTrue)
				So(IsValidUsername("je-suis-"), ShouldBeTrue)
				So(IsValidUsername("je_suis-"), ShouldBeTrue)
				So(IsValidUsername("je_suis_"), ShouldBeTrue)
				So(IsValidUsername("_jesuis"), ShouldBeTrue)
				So(IsValidUsername("_je-suis"), ShouldBeTrue)
				So(IsValidUsername("-jesuis"), ShouldBeTrue)
				So(IsValidUsername("-je_suis"), ShouldBeTrue)
				So(IsValidUsername("je.suis"), ShouldBeTrue)
				So(IsValidUsername("je.suis."), ShouldBeTrue)
				So(IsValidUsername("jesuis."), ShouldBeTrue)
				So(IsValidUsername("unnomcommeca"), ShouldBeTrue)
			})
		})
	})

	Convey("Testing Clean Username function function", t, func() {
		Convey("Given an user name :", func() {
			Convey("Containing caps -> should lower them", func() {
				So(CleanUsername("IamContaingCaps"), ShouldEqual, "iamcontaingcaps")
				So(CleanUsername("amContaingCaps"), ShouldEqual, "amcontaingcaps")
				So(CleanUsername("FULLCAPS"), ShouldEqual, "fullcaps")
				So(CleanUsername("capsattheenD"), ShouldEqual, "capsattheend")
			})
			Convey("Reserved -> should return a random name starting with a", func() {
				for _, uname := range restrictedUsernames {
					So(len(CleanUsername(uname)), ShouldEqual, 27)
				}
			})
			Convey("Containing illegal characters ( * ] \\ space ( ) { } [ ] .... -> should transform them in -)", func() {
				So(CleanUsername("i contain spaces"), ShouldEqual, "i-contain-spaces")
				So(CleanUsername("one space"), ShouldEqual, "one-space")
				So(CleanUsername(" "), ShouldEqual, "-")
				So(CleanUsername("iama*"), ShouldEqual, "iama-")
				So(CleanUsername("*"), ShouldEqual, "-")
				So(CleanUsername("some*things"), ShouldEqual, "some-things")
				So(CleanUsername("]"), ShouldEqual, "-")
				So(CleanUsername("]citation"), ShouldEqual, "-citation")
				So(CleanUsername("ci]tation"), ShouldEqual, "ci-tation")
				So(CleanUsername("citation]"), ShouldEqual, "citation-")
				So(CleanUsername("{"), ShouldEqual, "-")
				So(CleanUsername("{citation"), ShouldEqual, "-citation")
				So(CleanUsername("ci{tation"), ShouldEqual, "ci-tation")
				So(CleanUsername("citation{"), ShouldEqual, "citation-")
				So(CleanUsername("}"), ShouldEqual, "-")
				So(CleanUsername("}citation"), ShouldEqual, "-citation")
				So(CleanUsername("ci}tation"), ShouldEqual, "ci-tation")
				So(CleanUsername("citation}"), ShouldEqual, "citation-")
				So(CleanUsername("("), ShouldEqual, "-")
				So(CleanUsername("(citation"), ShouldEqual, "-citation")
				So(CleanUsername("ci(tation"), ShouldEqual, "ci-tation")
				So(CleanUsername("citation("), ShouldEqual, "citation-")
				So(CleanUsername(")"), ShouldEqual, "-")
				So(CleanUsername(")citation"), ShouldEqual, "-citation")
				So(CleanUsername("ci)tation"), ShouldEqual, "ci-tation")
				So(CleanUsername("citation)"), ShouldEqual, "citation-")
				So(CleanUsername("["), ShouldEqual, "-")
				So(CleanUsername("[citation"), ShouldEqual, "-citation")
				So(CleanUsername("ci[tation"), ShouldEqual, "ci-tation")
				So(CleanUsername("citation["), ShouldEqual, "citation-")
				So(CleanUsername("]"), ShouldEqual, "-")
				So(CleanUsername("]citation"), ShouldEqual, "-citation")
				So(CleanUsername("ci]tation"), ShouldEqual, "ci-tation")
				So(CleanUsername("citation]"), ShouldEqual, "citation-")
				So(CleanUsername("\\"), ShouldEqual, "-")
				So(CleanUsername("\\citation"), ShouldEqual, "-citation")
				So(CleanUsername("ci\\tation"), ShouldEqual, "ci-tation")
				So(CleanUsername("citation\\"), ShouldEqual, "citation-")
			})
			Convey("Correct -> should stay the same", func() {
				So(CleanUsername("je-suis"), ShouldEqual, "je-suis")
				So(CleanUsername("je_suis"), ShouldEqual, "je_suis")
				So(CleanUsername("je-suis_"), ShouldEqual, "je-suis_")
				So(CleanUsername("je-suis-"), ShouldEqual, "je-suis-")
				So(CleanUsername("je_suis-"), ShouldEqual, "je_suis-")
				So(CleanUsername("je_suis_"), ShouldEqual, "je_suis_")
				So(CleanUsername("_jesuis"), ShouldEqual, "_jesuis")
				So(CleanUsername("_je-suis"), ShouldEqual, "_je-suis")
				So(CleanUsername("-jesuis"), ShouldEqual, "-jesuis")
				So(CleanUsername("-je_suis"), ShouldEqual, "-je_suis")
				So(CleanUsername("je.suis"), ShouldEqual, "je.suis")
				So(CleanUsername("je.suis."), ShouldEqual, "je.suis.")
				So(CleanUsername("jesuis."), ShouldEqual, "jesuis.")
				So(CleanUsername("unnomcommeca"), ShouldEqual, "unnomcommeca")
			})
		})
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
