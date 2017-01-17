// This file is used to test if user model is working correctly.
// A user is always linked to a user
// He has basic user to join
package datastores

import (
	// "fmt"
	. "github.com/smartystreets/goconvey/convey"
	. "models"
	// "strings"
	"testing"
	"time"
	u "utils"
)

func TestUserStore(t *testing.T) {
	ds := dbStore{}
	ds.InitConnection("root", "popcube_test", "popcube_dev")
	db := *ds.Db

	usi := UserStoreImpl{}
	rsi := RoleStoreImpl{}

	time.Sleep(100 * 100)

	ownerRole := *rsi.GetByName(Owner.RoleName, ds)
	adminRole := *rsi.GetByName(Admin.RoleName, ds)
	// standartRole := *rsi.GetByName(Standart.RoleName, ds)
	// guestRole := *rsi.GetByName(Guest.RoleName, ds)

	Convey("Testing save function", t, func() {
		dbError := u.NewLocAppError("userStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyExistError := u.NewLocAppError("userStoreImpl.Save", "save.transaction.create.already_exist", nil, "User Name: test")
		user := User{
			Username:  "TesT",
			Password:  "test",
			Email:     "test@popcube.fr",
			NickName:  "NickName",
			FirstName: "Test",
			LastName:  "L",
			Role:      ownerRole,
		}
		Convey("Given a correct user.", func() {
			appError := usi.Save(&user, ds)
			Convey("Trying to add it for the first time, should be accepted", func() {
				So(appError, ShouldBeNil)
				So(appError, ShouldNotResemble, dbError)
				So(appError, ShouldNotResemble, alreadyExistError)
			})
			Convey("Trying to add it a second time should return duplicate error", func() {
				appError2 := usi.Save(&user, ds)
				So(appError2, ShouldNotBeNil)
				So(appError2, ShouldResemble, alreadyExistError)
				So(appError2, ShouldNotResemble, dbError)
			})
		})
		db.Delete(&user)
	})

	Convey("Testing update function", t, func() {
		dbError := u.NewLocAppError("userStoreImpl.Save", "save.transaction.create.encounterError", nil, "")
		alreadyExistError := u.NewLocAppError("userStoreImpl.Save", "save.transaction.create.already_exist", nil, "User Name: electras")
		// empty := User{}
		user := User{
			Username:  "TesT",
			Password:  "test",
			Email:     "test@popcube.fr",
			NickName:  "NickName",
			FirstName: "Test",
			LastName:  "L",
			Role:      ownerRole,
		}
		userNew := User{
			Username:  "lucky",
			Password:  "lucke",
			Email:     "luckylucke@popcube.fr",
			NickName:  "LL",
			FirstName: "Luky",
			LastName:  "Luke",
			Locale:    "vn_VN",
			Role:      adminRole,
		}
		appError := usi.Save(&user, ds)
		So(appError, ShouldBeNil)
		So(appError, ShouldNotResemble, dbError)
		So(appError, ShouldNotResemble, alreadyExistError)

		Convey("Provided correct User to modify should not return errors", func() {
			appError := usi.Update(&user, &userNew, ds)
			userShouldResemble := userNew
			userShouldResemble.WebID = user.WebID
			userShouldResemble.UserID = user.UserID
			userShouldResemble.UpdatedAt = user.UpdatedAt
			So(appError, ShouldBeNil)
			So(appError, ShouldNotResemble, dbError)
			So(appError, ShouldNotResemble, alreadyExistError)
			So(user, ShouldResemble, userShouldResemble)
		})

		// Convey("Provided wrong old User to modify should result in old_user error", func() {
		// 	user.WebID = "TesT"
		// 	Convey("Incorrect ID user should return a message invalid id", func() {
		// 		appError := usi.Update(&user, &userNew, ds)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dbError)
		// 		So(appError, ShouldNotResemble, alreadyExistError)
		// 		So(appError, ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.WebID.app_error", nil, ""))
		// 	})
		// 	Convey("Incorrect username user should return error Invalid username", func() {
		// 		user1 := User{
		// 			Username:  "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone",
		// 			Password:  "test",
		// 			Email:     "test@popcube.fr",
		// 			NickName:  "NickName",
		// 			FirstName: "Test",
		// 			LastName:  "L",
		// 			Role:      Owner,
		// 		}
		// 		user1.PreSave()
		// 		So(user1.IsValid(false), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webID="+user1.WebID))
		// 		user2 := User{
		// 			WebID:     NewID(),
		// 			Password:  "test",
		// 			Email:     "test@popcube.fr",
		// 			NickName:  "NickName",
		// 			FirstName: "Test",
		// 			LastName:  "L",
		// 			Role:      Owner,
		// 		}
		// 		So(user2.IsValid(false), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webID="+user2.WebID))
		// 		user3 := User{
		// 			WebID:     NewID(),
		// 			Username:  "xD/",
		// 			Password:  "test",
		// 			Email:     "test@popcube.fr",
		// 			NickName:  "NickName",
		// 			FirstName: "Test",
		// 			LastName:  "L",
		// 			Role:      Owner,
		// 		}
		// 		So(user3.IsValid(false), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webID="+user3.WebID))
		// 		user3 = User{
		// 			WebID:     NewID(),
		// 			Username:  "xD\\",
		// 			Password:  "test",
		// 			Email:     "test@popcube.fr",
		// 			NickName:  "NickName",
		// 			FirstName: "Test",
		// 			LastName:  "L",
		// 			Role:      Owner,
		// 		}
		// 		So(user3.IsValid(false), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webID="+user3.WebID))
		// 		user3 = User{
		// 			WebID:     NewID(),
		// 			Username:  "xD*",
		// 			Password:  "test",
		// 			Email:     "test@popcube.fr",
		// 			NickName:  "NickName",
		// 			FirstName: "Test",
		// 			LastName:  "L",
		// 			Role:      Owner,
		// 		}
		// 		So(user3.IsValid(false), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webID="+user3.WebID))
		// 		user3 = User{
		// 			WebID:     NewID(),
		// 			Username:  "xD{",
		// 			Password:  "test",
		// 			Email:     "test@popcube.fr",
		// 			NickName:  "NickName",
		// 			FirstName: "Test",
		// 			LastName:  "L",
		// 			Role:      Owner,
		// 		}
		// 		So(user3.IsValid(false), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webID="+user3.WebID))
		// 		user3 = User{
		// 			WebID:     NewID(),
		// 			Username:  "xD}",
		// 			Password:  "test",
		// 			Email:     "test@popcube.fr",
		// 			NickName:  "NickName",
		// 			FirstName: "Test",
		// 			LastName:  "L",
		// 			Role:      Owner,
		// 		}
		// 		So(user3.IsValid(false), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webID="+user3.WebID))
		// 		user3 = User{
		// 			WebID:     NewID(),
		// 			Username:  "xD#",
		// 			Password:  "test",
		// 			Email:     "test@popcube.fr",
		// 			NickName:  "NickName",
		// 			FirstName: "Test",
		// 			LastName:  "L",
		// 			Role:      Owner,
		// 		}
		// 		So(user3.IsValid(false), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webID="+user3.WebID))
		// 		user3 = User{
		// 			WebID:     NewID(),
		// 			Username:  "xD_",
		// 			Password:  "test",
		// 			Email:     "test@popcube.fr",
		// 			NickName:  "NickName",
		// 			FirstName: "Test",
		// 			LastName:  "L",
		// 			Role:      Owner,
		// 		}
		// 		So(user3.IsValid(false), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Username.app_error", nil, "user_webID="+user3.WebID))
		// 	})

		// 	Convey("Incorrect Email user should return error Invalid email", func() {
		// 		user := User{
		// 			Password:  "test",
		// 			Email:     "testpopcube.fr",
		// 			NickName:  "NickName",
		// 			FirstName: "Test",
		// 			LastName:  "L",
		// 			Role:      Owner,
		// 		}
		// 		user.PreSave()
		// 		So(user.IsValid(false), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Email.app_error", nil, "user_webID="+user.WebID))
		// 		user = User{
		// 			Password:  "test",
		// 			Email:     "test/popcube.fr",
		// 			NickName:  "NickName",
		// 			FirstName: "Test",
		// 			LastName:  "L",
		// 			Role:      Owner,
		// 		}
		// 		user.PreSave()
		// 		So(user.IsValid(false), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Email.app_error", nil, "user_webID="+user.WebID))
		// 		user = User{
		// 			Password:  "test",
		// 			Email:     "CeNomDevraitJelespereEtreBeaucoupTropLongPourLatrailleMaximaleDemandeParcequelaJeSuiunPoilFeneantEtDeboussouleSansnuldouteilnyavaitpersone@popcube.fr",
		// 			NickName:  "NickName",
		// 			FirstName: "Test",
		// 			LastName:  "L",
		// 			Role:      Owner,
		// 		}
		// 		user.PreSave()
		// 		So(user.IsValid(false), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.Email.app_error", nil, "user_webID="+user.WebID))
		// 	})

		// 	Convey("NickName, FirstName: and Lastname should be less than 64 characters long", func() {
		// 		user := User{
		// 			Password:  "test",
		// 			Email:     "test@popcube.fr",
		// 			NickName:  "NickNameéèéééééééééééétroplongazdazdzadazdazdzadz_>_<azdazdzadazdazz",
		// 			FirstName: "Test",
		// 			LastName:  "L",
		// 			Role:      Owner,
		// 		}
		// 		user.PreSave()
		// 		So(user.IsValid(false), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.NickName.app_error", nil, "user_webID="+user.WebID))
		// 		user = User{
		// 			Password:  "test",
		// 			Email:     "test@popcube.fr",
		// 			NickName:  "NickName",
		// 			FirstName: "TestéèéèéèéèèéèéèéèéèéèèéèéèéèèéèéèNJnefiznfidsdfnpdsjfazddrfazdzadzadzadzadazd",
		// 			LastName:  "L",
		// 			Role:      Owner,
		// 		}
		// 		user.PreSave()
		// 		So(user.IsValid(false), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.first_name.app_error", nil, "user_webID="+user.WebID))
		// 		user = User{
		// 			Password:  "test",
		// 			Email:     "test@popcube.fr",
		// 			NickName:  "NickName",
		// 			FirstName: "Test",
		// 			LastName:  "TestéèéèéèéèèéèéèéèéèéèèéèéèéèèéèéèNJnefiznfidsdfdazdzadzadzadzadzadzadazdazdazdzadazdzanpdsjf",
		// 			Role:      Owner,
		// 		}
		// 		user.PreSave()
		// 		So(user.IsValid(false), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.last_name.app_error", nil, "user_webID="+user.WebID))
		// 	})

		// 	Convey("Password can]t be empty", func() {
		// 		user := User{
		// 			Email:     "test@popcube.fr",
		// 			NickName:  "NickName",
		// 			FirstName: "Test",
		// 			LastName:  "L",
		// 			Role:      Owner,
		// 		}
		// 		user.PreSave()
		// 		So(user.IsValid(false), ShouldResemble, u.NewLocAppError("user.IsValid", "model.user.is_valid.auth_data_pwd.app_error", nil, "user_webID="+user.WebID))
		// 	})
		// })

		// Convey("Provided wrong new User to modify should result in new_user error", func() {
		// 	userNew.UserName = strings.ToLower("ThisShouldBeAFreakingLongEnougthStringToRefuse.BahNon, pas tout seul. C'est long 64 caractères en vrai  ~#~")
		// 	Convey("Too long user name should return Too Long user name error", func() {
		// 		appError := usi.Update(&user, &userNew, ds)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dbError)
		// 		So(appError, ShouldNotResemble, alreadyExistError)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.user_name.app_error", nil, "id="+userNew.WebID))
		// 	})
		// 	Convey("Incorect Alpha Num user name should be refused", func() {
		// 		userNew.UserName = "?/+*"
		// 		appError = usi.Update(&user, &userNew, ds)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dbError)
		// 		So(appError, ShouldNotResemble, alreadyExistError)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.not_alphanum_user_name.app_error", nil, "id="+userNew.WebID))
		// 		userNew.UserName = "("
		// 		appError = usi.Update(&user, &userNew, ds)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dbError)
		// 		So(appError, ShouldNotResemble, alreadyExistError)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.not_alphanum_user_name.app_error", nil, "id="+userNew.WebID))
		// 		userNew.UserName = "{"
		// 		appError = usi.Update(&user, &userNew, ds)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dbError)
		// 		So(appError, ShouldNotResemble, alreadyExistError)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.not_alphanum_user_name.app_error", nil, "id="+userNew.WebID))
		// 		userNew.UserName = "}"
		// 		appError = usi.Update(&user, &userNew, ds)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dbError)
		// 		So(appError, ShouldNotResemble, alreadyExistError)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.not_alphanum_user_name.app_error", nil, "id="+userNew.WebID))
		// 		userNew.UserName = ")"
		// 		appError = usi.Update(&user, &userNew, ds)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dbError)
		// 		So(appError, ShouldNotResemble, alreadyExistError)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.not_alphanum_user_name.app_error", nil, "id="+userNew.WebID))
		// 		userNew.UserName = "["
		// 		appError = usi.Update(&user, &userNew, ds)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dbError)
		// 		So(appError, ShouldNotResemble, alreadyExistError)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.not_alphanum_user_name.app_error", nil, "id="+userNew.WebID))
		// 		userNew.UserName = "]"
		// 		appError = usi.Update(&user, &userNew, ds)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dbError)
		// 		So(appError, ShouldNotResemble, alreadyExistError)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.not_alphanum_user_name.app_error", nil, "id="+userNew.WebID))
		// 		userNew.UserName = " "
		// 		appError = usi.Update(&user, &userNew, ds)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dbError)
		// 		So(appError, ShouldNotResemble, alreadyExistError)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.not_alphanum_user_name.app_error", nil, "id="+userNew.WebID))
		// 	})
		// 	userNew.UserName = "electra"
		// 	userNew.Description = "Il Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face: Alors, la, c'était 250 en fait .... Du coup, on va multiplier par 4 un ? OK ? l Me faut beaucoup trop de character  ..... 1024, c'est grand. Très grand. Comme l'infini. C'est long. Surtout à la fin. Et puis même après tout ça, je suis pas sur que ce soit assez .... Compteur ??? Vous êtes la ? :p :'( :docker: :troll-face:"
		// 	Convey("Given a too long description, should return too long description error :p", func() {
		// 		appError = usi.Update(&user, &userNew, ds)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dbError)
		// 		So(appError, ShouldNotResemble, alreadyExistError)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.description.app_error", nil, "id="+userNew.WebID))
		// 	})
		// 	userNew.Description = "Stoppppppp"
		// 	userNew.Subject = "Encore beaucoup de caractere pour rien .... mais un peu moins cette fois. Il n'en faut que 250 ........... Fait dodo, cola mon p'tit frere. Fais dodo, j'ai pêté un cable. Swing du null, Swing du null, c'est le swing du null ..... :guitare: :singer: :music: Je suis un main troll :O"
		// 	Convey("Given a too long subject, should return too long description error :p", func() {
		// 		appError = usi.Update(&user, &userNew, ds)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dbError)
		// 		So(appError, ShouldNotResemble, alreadyExistError)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.subject.app_error", nil, "id="+userNew.WebID))
		// 	})
		// 	userNew.Subject = "Safe"
		// 	userNew.Type = "Outside of Range"
		// 	Convey("Providing a wrong type should not work", func() {
		// 		appError = usi.Update(&user, &userNew, ds)
		// 		So(appError, ShouldNotBeNil)
		// 		So(appError, ShouldNotResemble, dbError)
		// 		So(appError, ShouldNotResemble, alreadyExistError)
		// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Update.userNew.PreSave", "model.user.is_valid.type.app_error", nil, "id="+userNew.WebID))
		// 	})
		// })

		db.Delete(&user)
		db.Delete(&userNew)
	})

	// Convey("Testing Getters", t, func() {
	// 	user0 := User{
	// 		UserName:    "electra",
	// 		Type:        "video",
	// 		Private:     false,
	// 		Description: "Testing user description :O",
	// 		Subject:     "Sujet",
	// 		Avatar:      "jesuiscool.svg",
	// 	}
	// 	user1 := User{
	// 		UserName:    "mal",
	// 		Type:        "audio",
	// 		Private:     false,
	// 		Description: "Speaking on Malsdjisqnju BD song from R. Sechan",
	// 		Subject:     "Sujet1",
	// 		Avatar:      "cover_mal.svg",
	// 	}
	// 	user2 := User{
	// 		UserName: "lagaffesfantasio",
	// 		Type:     "direct",
	// 		Private:  false,
	// 		Avatar:   "gaston.svg",
	// 	}
	// 	user1New := User{
	// 		UserName:    "malheur",
	// 		Private:     true,
	// 		Description: "Let's speak about the BD Mal",
	// 		Subject:     "Mal",
	// 		Avatar:      "cover_mal_efique.svg",
	// 	}
	// 	user3 := User{
	// 		UserName:    "corsicarms",
	// 		Type:        "audio",
	// 		Private:     false,
	// 		Description: "Speaking on Corsic Arms song from R. Sechan",
	// 		Subject:     "Sujet",
	// 		Avatar:      "cover_csa.svg",
	// 	}

	// 	usi.Save(&user0, ds)
	// 	usi.Save(&user1, ds)
	// 	// usi.Update(&user1, &user1New, ds)
	// 	usi.Save(&user2, ds)
	// 	usi.Save(&user3, ds)

	// 	// Have to be after save so ID are up to date :O
	// 	userList := []User{
	// 		user0,
	// 		user1,
	// 		user2,
	// 		user3,
	// 	}

	// 	audioList := []User{user1, user3}
	// 	directList := []User{user2}
	// 	privateList := []User{user2}
	// 	publicList := []User{user0, user1, user3}
	// 	emptyList := []User{}

	// 	Convey("We have to be able to find all users in the db", func() {
	// 		users := usi.GetAll(ds)
	// 		So(users, ShouldNotResemble, &emptyList)
	// 		So(users, ShouldResemble, &userList)
	// 	})

	// 	Convey("We have to be able to find a user from is name", func() {
	// 		user := usi.GetByName(user0.UserName, ds)
	// 		So(user, ShouldNotResemble, &User{})
	// 		So(user, ShouldResemble, &user0)
	// 		user = usi.GetByName(user2.UserName, ds)
	// 		So(user, ShouldNotResemble, &User{})
	// 		So(user, ShouldResemble, &user2)
	// 		user = usi.GetByName(user3.UserName, ds)
	// 		So(user, ShouldNotResemble, &User{})
	// 		So(user, ShouldResemble, &user3)
	// 		Convey("Should also work from updated value", func() {
	// 			user = usi.GetByName(user1.UserName, ds)
	// 			So(user, ShouldNotResemble, &User{})
	// 			So(user, ShouldResemble, &user1)
	// 		})
	// 	})

	// 	Convey("We have to be able to find users from type", func() {
	// 		users := usi.GetByType("audio", ds)
	// 		So(users, ShouldNotResemble, &User{})
	// 		So(users, ShouldResemble, &audioList)
	// 		users = usi.GetByType("direct", ds)
	// 		So(users, ShouldNotResemble, &User{})
	// 		So(users, ShouldResemble, &directList)
	// 	})

	// 	Convey("We have to be able to find private or public users list", func() {
	// 		users := usi.GetPrivate(ds)
	// 		So(users, ShouldNotResemble, &User{})
	// 		So(users, ShouldResemble, &privateList)
	// 		users = usi.GetPublic(ds)
	// 		So(users, ShouldNotResemble, &User{})
	// 		So(users, ShouldResemble, &publicList)
	// 	})

	// 	Convey("Searching for non existent user should return empty", func() {
	// 		user := usi.GetByName("fantome", ds)
	// 		So(user, ShouldResemble, &User{})
	// 	})

	// 	db.Delete(&user0)
	// 	db.Delete(&user1)
	// 	db.Delete(&user1New)
	// 	db.Delete(&user2)
	// 	db.Delete(&user3)

	// 	Convey("Searching all in empty table should return empty", func() {
	// 		users := usi.GetAll(ds)
	// 		So(users, ShouldResemble, &[]User{})
	// 	})
	// })

	// Convey("Testing delete user", t, func() {
	// 	dberror := u.NewLocAppError("userStoreImpl.Delete", "update.transaction.delete.encounterError", nil, "")
	// 	user0 := User{
	// 		UserName:    "electra",
	// 		Type:        "video",
	// 		Private:     false,
	// 		Description: "Testing user description :O",
	// 		Subject:     "Sujet",
	// 		Avatar:      "jesuiscool.svg",
	// 	}
	// 	user1 := User{
	// 		UserName:    "mal",
	// 		Type:        "audio",
	// 		Private:     false,
	// 		Description: "Speaking on Malsdjisqnju BD song from R. Sechan",
	// 		Subject:     "Sujet1",
	// 		Avatar:      "cover_mal.svg",
	// 	}
	// 	user2 := User{
	// 		UserName: "lagaffesfantasio",
	// 		Type:     "direct",
	// 		Private:  false,
	// 		Avatar:   "gaston.svg",
	// 	}
	// 	user3 := User{
	// 		UserName:    "corsicarms",
	// 		Type:        "audio",
	// 		Private:     false,
	// 		Description: "Speaking on Corsic Arms song from R. Sechan",
	// 		Subject:     "Sujet",
	// 		Avatar:      "cover_csa.svg",
	// 	}

	// 	usi.Save(&user0, ds)
	// 	usi.Save(&user1, ds)
	// 	usi.Save(&user2, ds)
	// 	usi.Save(&user3, ds)

	// 	// Have to be after save so ID are up to date :O
	// 	// user3Old := user3
	// 	// userList1 := []User{
	// 	// 	user0,
	// 	// 	user1,
	// 	// 	user2,
	// 	// 	user3Old,
	// 	// }

	// 	Convey("Deleting a known user should work", func() {
	// 		appError := usi.Delete(&user2, ds)
	// 		So(appError, ShouldBeNil)
	// 		So(appError, ShouldNotResemble, dberror)
	// 		So(usi.GetByName("God", ds), ShouldResemble, &User{})
	// 	})

	// 	// Convey("Trying to delete from non conform user should return specific user error and should not delete users.", func() {
	// 	// 	user3.UserName = "Const"
	// 	// 	Convey("Too long or empty Name should return name error", func() {
	// 	// 		appError := usi.Delete(&user3, ds)
	// 	// 		So(appError, ShouldNotBeNil)
	// 	// 		So(appError, ShouldNotResemble, dberror)
	// 	// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Delete.user.PreSave", "model.user.username.app_error", nil, ""))
	// 	// 		So(usi.GetAll(ds), ShouldResemble, &userList1)
	// 	// 		user3.UserName = "+alpha"
	// 	// 		appError = usi.Delete(&user3, ds)
	// 	// 		So(appError, ShouldNotBeNil)
	// 	// 		So(appError, ShouldNotResemble, dberror)
	// 	// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Delete.user.PreSave", "model.user.username.app_error", nil, ""))
	// 	// 		So(usi.GetAll(ds), ShouldResemble, &userList1)
	// 	// 		user3.UserName = "alpha-numerique"
	// 	// 		appError = usi.Delete(&user3, ds)
	// 	// 		So(appError, ShouldNotBeNil)
	// 	// 		So(appError, ShouldNotResemble, dberror)
	// 	// 		So(appError, ShouldResemble, u.NewLocAppError("userStoreImpl.Delete.user.PreSave", "model.user.username.app_error", nil, ""))
	// 	// 		So(usi.GetAll(ds), ShouldResemble, &userList1)
	// 	// 	})
	// 	// })

	// 	db.Delete(&user0)
	// 	db.Delete(&user1)
	// 	db.Delete(&user2)
	// 	db.Delete(&user3)
	// })
}
