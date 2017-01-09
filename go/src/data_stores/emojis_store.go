package data_stores

import (
	// l4g "github.com/alecthomas/log4go"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
	"models"
	. "utils"
	// "time"
)

type EmojiStoreImpl struct {
	EmojiStore
}

// Use to save emoji in BB
func (asi EmojiStoreImpl) Save(emoji *models.Emoji, ds DataStore) *AppError {
	db := *ds.Db
	transaction := db.Begin()
	if appError := emoji.IsValid(); appError != nil {
		transaction.Rollback()
		return NewLocAppError("emoji_store_impl.Save.emoji.PreSave", appError.Id, nil, appError.DetailedError)
	}
	if !transaction.NewRecord(emoji) {
		transaction.Rollback()
		return NewLocAppError("emoji_store_impl.Save", "save.transaction.create.already_exist", nil, "Emoji Name: "+emoji.Name)
	}
	if err := transaction.Create(&emoji).Error; err != nil {
		transaction.Rollback()
		return NewLocAppError("emoji_store_impl.Save", "save.transaction.create.encounter_error :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Used to update emoji in DB
func (asi EmojiStoreImpl) Update(emoji *models.Emoji, new_emoji *models.Emoji, ds DataStore) *AppError {
	db := *ds.Db
	transaction := db.Begin()
	if appError := emoji.IsValid(); appError != nil {
		transaction.Rollback()
		return NewLocAppError("emoji_store_impl.Update.emoji_old.PreSave", appError.Id, nil, appError.DetailedError)
	}
	if appError := new_emoji.IsValid(); appError != nil {
		transaction.Rollback()
		return NewLocAppError("emoji_store_impl.Update.emoji_new.PreSave", appError.Id, nil, appError.DetailedError)
	}
	if err := transaction.Model(&emoji).Updates(&new_emoji).Error; err != nil {
		transaction.Rollback()
		return NewLocAppError("emoji_store_impl.Update", "update.transaction.updates.encounter_error :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Used to get emoji from DB
func (asi EmojiStoreImpl) GetAll(ds DataStore) *[]models.Emoji {
	db := *ds.Db
	emojis := []models.Emoji{}
	db.Find(&emojis)
	return &emojis
}

// Used to get emoji from DB
func (asi EmojiStoreImpl) GetByName(emojiName string, ds DataStore) *models.Emoji {
	db := *ds.Db
	emoji := models.Emoji{}
	db.Where("name = ?", emojiName).First(&emoji)
	return &emoji
}

// Used to get emoji from DB
func (asi EmojiStoreImpl) GetByShortcut(EmojiShortcut string, ds DataStore) *models.Emoji {
	db := *ds.Db
	emoji := models.Emoji{}
	db.Where("shortcut = ?", EmojiShortcut).First(&emoji)
	return &emoji
}

// Used to get emoji from DB
func (asi EmojiStoreImpl) GetByLink(emojiLink string, ds DataStore) *models.Emoji {
	db := *ds.Db
	emoji := models.Emoji{}
	db.Where("link = ?", emojiLink).First(&emoji)
	return &emoji
}

// Used to get emoji from DB
func (asi EmojiStoreImpl) Delete(emoji *models.Emoji, ds DataStore) *AppError {
	db := *ds.Db
	transaction := db.Begin()
	if appError := emoji.IsValid(); appError != nil {
		transaction.Rollback()
		return NewLocAppError("emoji_store_impl.Delete.emoji.PreSave", appError.Id, nil, appError.DetailedError)
	}
	if err := transaction.Delete(&emoji).Error; err != nil {
		transaction.Rollback()
		return NewLocAppError("emoji_store_impl.Delete", "update.transaction.delete.encounter_error :"+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}
