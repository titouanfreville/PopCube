package data_stores

import (
	// l4g "github.com/alecthomas/log4go"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
	"models"
	. "utils"
	// "time"
)

type AvatarStoreImpl struct {
	AvatarStore
}

// Use to save avatar in BB
func (asi AvatarStoreImpl) Save(avatar *models.Avatar, ds DataStore) *AppError {
	db := *ds.Db
	transaction := db.Begin()
	if appError := avatar.IsValid(); appError != nil {
		transaction.Rollback()
		return NewLocAppError("avatar_store_impl.Save.avatar.PreSave", appError.Id, nil, appError.DetailedError)
	}
	if !transaction.NewRecord(avatar) {
		transaction.Rollback()
		return NewLocAppError("avatar_store_impl.Save", "save.transaction.create.already_exist", nil, "Avatar Name: "+avatar.Name)
	}
	if err := transaction.Create(&avatar).Error; err != nil {
		transaction.Rollback()
		return NewLocAppError("avatar_store_impl.Save", "save.transaction.create.encounter_error", nil, "")
	}
	transaction.Commit()
	return nil
}

// Used to update avatar in DB
func (asi AvatarStoreImpl) Update(avatar *models.Avatar, new_avatar *models.Avatar, ds DataStore) *AppError {
	db := *ds.Db
	transaction := db.Begin()
	if appError := avatar.IsValid(); appError != nil {
		transaction.Rollback()
		return NewLocAppError("avatar_store_impl.Update.avatar_old.PreSave", appError.Id, nil, appError.DetailedError)
	}
	if appError := new_avatar.IsValid(); appError != nil {
		transaction.Rollback()
		return NewLocAppError("avatar_store_impl.Update.avatar_new.PreSave", appError.Id, nil, appError.DetailedError)
	}
	if err := transaction.Model(&avatar).Updates(&new_avatar).Error; err != nil {
		transaction.Rollback()
		return NewLocAppError("avatar_store_impl.Update", "update.transaction.updates.encounter_error", nil, "")
	}
	transaction.Commit()
	return nil
}

// Used to get avatar from DB
func (asi AvatarStoreImpl) GetAll(ds DataStore) *[]models.Avatar {
	db := *ds.Db
	avatars := []models.Avatar{}
	db.Find(&avatars)
	return &avatars
}

// Used to get avatar from DB
func (asi AvatarStoreImpl) GetByName(avatarName string, ds DataStore) *models.Avatar {
	db := *ds.Db
	avatar := models.Avatar{}
	db.Where("name = ?", avatarName).First(&avatar)
	return &avatar
}

// Used to get avatar from DB
func (asi AvatarStoreImpl) GetByLink(avatarLink string, ds DataStore) *models.Avatar {
	db := *ds.Db
	avatar := models.Avatar{}
	db.Where("link = ?", avatarLink).First(&avatar)
	return &avatar
}

// Used to get avatar from DB
func (asi AvatarStoreImpl) Delete(avatar *models.Avatar, ds DataStore) *AppError {
	db := *ds.Db
	transaction := db.Begin()
	if appError := avatar.IsValid(); appError != nil {
		transaction.Rollback()
		return NewLocAppError("avatar_store_impl.Delete.avatar.PreSave", appError.Id, nil, appError.DetailedError)
	}
	if err := transaction.Delete(&avatar).Error; err != nil {
		transaction.Rollback()
		return NewLocAppError("avatar_store_impl.Delete", "update.transaction.delete.encounter_error", nil, "")
	}
	transaction.Commit()
	return nil
}
