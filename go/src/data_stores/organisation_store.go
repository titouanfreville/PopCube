package data_stores

import (
	// l4g "github.com/alecthomas/log4go"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
	"models"
	. "utils"
	// "time"
)

func Save(organisation *models.Organisation, ds DataStore) *AppError {
	db := *ds.Db
	transaction := db.Begin()
	organisation.PreSave()
	appError := organisation.IsValid()
	if appError == nil {
		if transaction.NewRecord(organisation) {
			transaction.Create(&organisation)
			return nil
		} else {
			return nil
		}
	}
	return nil
}

// func Update() {

// }

// func Get() {

// }

// func GetAll() {

// }
