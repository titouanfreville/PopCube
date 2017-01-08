// This file is used to test if user model is working correctly.
// A user is always linked to an organisation
// He has basic channel to join
package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"strings"
	"testing"
	u "utils"
)

// Test correction test for user ;)

// Test Password functionalities from user Model
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
			user := User{WebId: NewId(), Username: NewId()}
			Convey("Converting user to json then json to user should provide same user information", func() {
				json := user.ToJson()
				test_user := UserFromJson(strings.NewReader(json))
				So(user.WebId, ShouldEqual, test_user.WebId)
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
				So(user.WebId, ShouldNotBeBlank)
				So(user.Username, ShouldNotBeBlank)
				So(user.EmailVerified, ShouldBeFalse)
				So(user.Deleted, ShouldBeFalse)
				So(user.UpdatedAt, ShouldNotBeNil)
				So(user.UpdatedAt, ShouldBeGreaterThan, 0)
				So(user.UpdatedAt, ShouldEqual, user.LastPasswordUpdate)
				So(user.Locale, ShouldNotBeBlank)
				So(ComparePassword(user.Password, "test"), ShouldBeTrue)
			})

			Convey("Data should be correctly formated", func() {
				user.PreSave()
				So(IsLower(user.Username), ShouldBeTrue)
				So(IsLower(user.Email), ShouldBeTrue)
			})

			Convey("Etag should be correctly generated", func() {
				user.PreSave()
				Etag := user.Etag(true, true)
				expected := CURRENT_VERSION + "." + user.WebId + "." + strconv.FormatInt(user.UpdatedAt, 10) + "." + "true" + "." + "true"
				So(Etag, ShouldEqual, expected)
			})
		})

		Convey("Given an user with email and username", func() {
			user := User{Password: "test", Username: "TesT", Email: "Test@poPcube.fr"}
			Convey("Applying PreSave should fill blank required fields and concerve overs", func() {
				user.PreSave()
				So(user.WebId, ShouldNotBeBlank)
				So(user.Username, ShouldEqual, "test")
				So(user.Email, ShouldEqual, "test@popcube.fr")
				So(user.EmailVerified, ShouldBeFalse)
				So(user.Deleted, ShouldBeFalse)
				So(user.UpdatedAt, ShouldNotBeNil)
				So(user.UpdatedAt, ShouldBeGreaterThan, 0)
				So(user.UpdatedAt, ShouldEqual, user.LastPasswordUpdate)
				So(user.Locale, ShouldNotBeBlank)
				So(ComparePassword(user.Password, "test"), ShouldBeTrue)
			})

			Convey("Data should be correctly formated", func() {
				user.PreSave()
				So(IsLower(user.Username), ShouldBeTrue)
				So(IsLower(user.Email), ShouldBeTrue)
			})

			Convey("Etag should be correctly generated", func() {
				user.PreSave()
				Etag := user.Etag(true, true)
				expected := CURRENT_VERSION + "." + user.WebId + "." + strconv.FormatInt(user.UpdatedAt, 10) + "." + "true" + "." + "true"
				So(Etag, ShouldEqual, expected)
			})
		})

		Convey("Given a full user entry", func() {
			user := User{
				WebId:              "testId",
				UpdatedAt:          10,
				Deleted:            true,
				Username:           "TesT",
				Password:           "test",
				Email:              "Test@poPcube.fr",
				EmailVerified:      true,
				NickName:           "NickName",
				FirstName:          "Test",
				LastName:           "L",
				Role:               OWNER,
				LastPasswordUpdate: 20,
				FailedAttempts:     1,
				Locale:             "vi",
				LastActivityAt:     5,
			}

			Convey("Applying PreSave should only correctly format field and use good time for last Updates", func() {
				user.PreSave()
				So(user.WebId, ShouldEqual, "testId")
				So(user.UpdatedAt, ShouldNotEqual, 10)
				So(user.Deleted, ShouldBeTrue)
				So(user.Username, ShouldEqual, "test")
				So(ComparePassword(user.Password, "test"), ShouldBeTrue)
				So(user.Email, ShouldEqual, "test@popcube.fr")
				So(user.EmailVerified, ShouldBeTrue)
				So(user.NickName, ShouldEqual, "NickName")
				So(user.FirstName, ShouldEqual, "Test")
				So(user.LastName, ShouldEqual, "L")
				So(user.Role, ShouldResemble, OWNER)
				So(user.LastPasswordUpdate, ShouldNotEqual, 20)
				So(user.LastPasswordUpdate, ShouldEqual, user.UpdatedAt)
				So(user.FailedAttempts, ShouldEqual, 1)
				So(user.Locale, ShouldEqual, "vi")
				So(user.LastActivityAt, ShouldEqual, 5)
			})

			Convey("Etag should be correctly generated", func() {
				user.PreSave()
				Etag := user.Etag(true, true)
				expected := CURRENT_VERSION + "." + user.WebId + "." + strconv.FormatInt(user.UpdatedAt, 10) + "." + "true" + "." + "true"
				So(Etag, ShouldEqual, expected)
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
				NickName:  "NickName",
				FirstName: "Test",
				LastName:  "L",
				Role:      OWNER,
			}
			correct_user.PreSave()
			So(correct_user.IsValid(), ShouldBeNil)
			So(correct_user.IsValid(), ShouldNotResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.WebId.app_error", nil, ""))
			So(correct_user.IsValid(), ShouldNotResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Email.app_error", nil, "user_webId="+correct_user.WebId))
			So(correct_user.IsValid(), ShouldNotResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.NickName.app_error", nil, "user_webId="+correct_user.WebId))
			So(correct_user.IsValid(), ShouldNotResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.first_name.app_error", nil, "user_webId="+correct_user.WebId))
			So(correct_user.IsValid(), ShouldNotResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webId="+correct_user.WebId))
			So(correct_user.IsValid(), ShouldNotResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.last_name.app_error", nil, "user_webId="+correct_user.WebId))
			So(correct_user.IsValid(), ShouldNotResemble, u.NewLocAppError("user.IsValid", "model.user.auth_data_pwd.Username.app_error", nil, "user_webId="+correct_user.WebId))
		})
		Convey("Given an incorrect user, validation should return error message", func() {
			Convey("Incorrect ID user should return a message invalid id", func() {
				user := User{
					WebId:     "Nimp",
					Username:  "TesT",
					Password:  "test",
					Email:     "test@popcube.fr",
					NickName:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Role:      OWNER,
				}
				So(user.IsValid(), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.WebId.app_error", nil, ""))
			})
			Convey("Incorrect username user should return error Invalid username", func() {
				user1 := User{
					Username:  "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone",
					Password:  "test",
					Email:     "test@popcube.fr",
					NickName:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Role:      OWNER,
				}
				user1.PreSave()
				So(user1.IsValid(), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webId="+user1.WebId))
				user2 := User{
					WebId:     NewId(),
					Password:  "test",
					Email:     "test@popcube.fr",
					NickName:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Role:      OWNER,
				}
				So(user2.IsValid(), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webId="+user2.WebId))
				user3 := User{
					WebId:     NewId(),
					Username:  "xD/",
					Password:  "test",
					Email:     "test@popcube.fr",
					NickName:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Role:      OWNER,
				}
				So(user3.IsValid(), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webId="+user3.WebId))
				user3 = User{
					WebId:     NewId(),
					Username:  "xD\\",
					Password:  "test",
					Email:     "test@popcube.fr",
					NickName:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Role:      OWNER,
				}
				So(user3.IsValid(), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webId="+user3.WebId))
				user3 = User{
					WebId:     NewId(),
					Username:  "xD*",
					Password:  "test",
					Email:     "test@popcube.fr",
					NickName:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Role:      OWNER,
				}
				So(user3.IsValid(), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webId="+user3.WebId))
				user3 = User{
					WebId:     NewId(),
					Username:  "xD{",
					Password:  "test",
					Email:     "test@popcube.fr",
					NickName:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Role:      OWNER,
				}
				So(user3.IsValid(), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webId="+user3.WebId))
				user3 = User{
					WebId:     NewId(),
					Username:  "xD}",
					Password:  "test",
					Email:     "test@popcube.fr",
					NickName:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Role:      OWNER,
				}
				So(user3.IsValid(), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webId="+user3.WebId))
				user3 = User{
					WebId:     NewId(),
					Username:  "xD#",
					Password:  "test",
					Email:     "test@popcube.fr",
					NickName:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Role:      OWNER,
				}
				So(user3.IsValid(), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webId="+user3.WebId))
				user3 = User{
					WebId:     NewId(),
					Username:  "xD_",
					Password:  "test",
					Email:     "test@popcube.fr",
					NickName:  "NickName",
					FirstName: "Test",
					LastName:  "L",
					Role:      OWNER,
				}
				So(user3.IsValid(), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webId="+user3.WebId))
			})
		})

		Convey("Incorrect Email user should return error Invalid email", func() {
			user := User{
				Password:  "test",
				Email:     "testpopcube.fr",
				NickName:  "NickName",
				FirstName: "Test",
				LastName:  "L",
				Role:      OWNER,
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Email.app_error", nil, "user_webId="+user.WebId))
			user = User{
				Password:  "test",
				Email:     "test/popcube.fr",
				NickName:  "NickName",
				FirstName: "Test",
				LastName:  "L",
				Role:      OWNER,
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Email.app_error", nil, "user_webId="+user.WebId))
			user = User{
				Password:  "test",
				Email:     "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone@popcube.fr",
				NickName:  "NickName",
				FirstName: "Test",
				LastName:  "L",
				Role:      OWNER,
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Email.app_error", nil, "user_webId="+user.WebId))
		})

		Convey("NickName, FirstName: and Lastname should be less than 64 characters long", func() {
			user := User{
				Password:  "test",
				Email:     "test@popcube.fr",
				NickName:  "NickNameéèéééééééééééétroplongazdazdzadazdazdzadz_>_<azdazdzadazdazz",
				FirstName: "Test",
				LastName:  "L",
				Role:      OWNER,
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.NickName.app_error", nil, "user_webId="+user.WebId))
			user = User{
				Password:  "test",
				Email:     "test@popcube.fr",
				NickName:  "NickName",
				FirstName: "TestéèéèéèéèèéèéèéèéèéèèéèéèéèèéèéèNJnefiznfidsdfnpdsjfazddrfazdzadzadzadzadazd",
				LastName:  "L",
				Role:      OWNER,
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.first_name.app_error", nil, "user_webId="+user.WebId))
			user = User{
				Password:  "test",
				Email:     "test@popcube.fr",
				NickName:  "NickName",
				FirstName: "Test",
				LastName:  "TestéèéèéèéèèéèéèéèéèéèèéèéèéèèéèéèNJnefiznfidsdfdazdzadzadzadzadzadzadazdazdazdzadazdzanpdsjf",
				Role:      OWNER,
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.last_name.app_error", nil, "user_webId="+user.WebId))
		})

		Convey("Password can]t be empty", func() {
			user := User{
				Email:     "test@popcube.fr",
				NickName:  "NickName",
				FirstName: "Test",
				LastName:  "L",
				Role:      OWNER,
			}
			user.PreSave()
			So(user.IsValid(), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.auth_data_pwd.app_error", nil, "user_webId="+user.WebId))
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
			Convey("user without First/Last/Nick name should have username as display name", func() {
				So(u.GetDisplayName(), ShouldEqual, "test")
			})
			Convey("user with First/Last name but no nickname should have full name as displayname", func() {
				u.LastName = "Troll"
				So(u.GetDisplayName(), ShouldEqual, "Troll")
				u.FirstName = "Min"
				So(u.GetDisplayName(), ShouldEqual, "Min Troll")
				u.LastName = ""
				So(u.GetDisplayName(), ShouldEqual, "Min")
			})
			Convey("user having a nickname should have their nickname diplayed", func() {
				u.NickName = "nOOb"
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

	Convey("Testing Clean username function function", t, func() {
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
