package data_stores

import (
	_ "github.com/go-sql-driver/mysql"
	"models"
	. "utils"
)

type OrganisationStoreImpl struct {
	OrganisationStore
}

// Use to save data in BB
func (osi OrganisationStoreImpl) Save(organisation *models.Organisation, ds DataStore) *AppError {
	db := *ds.Db
	transaction := db.Begin()
	organisation.PreSave()
	if appError := organisation.IsValid(); appError != nil {
		transaction.Rollback()
		return NewLocAppError("organisation_store_impl.Save.organisation.PreSave", appError.Id, nil, appError.DetailedError)
	}
	if !transaction.NewRecord(organisation) {
		transaction.Rollback()
		return NewLocAppError("organisation_store_impl.Save", "save.transaction.create.already_exist", nil, "Organisation Name: "+organisation.OrganisationName)
	}
	if err := transaction.Create(&organisation).Error; err != nil {
		transaction.Rollback()
		return NewLocAppError("organisation_store_impl.Save", "save.transaction.create.encounter_error: "+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Used to update data in DB
func (osi OrganisationStoreImpl) Update(organisation *models.Organisation, new_organisation *models.Organisation, ds DataStore) *AppError {
	db := *ds.Db
	transaction := db.Begin()
	new_organisation.PreSave()
	if appError := organisation.IsValid(); appError != nil {
		transaction.Rollback()
		return NewLocAppError("organisation_store_impl.Update.organisation_old.PreSave", appError.Id, nil, appError.DetailedError)
	}
	if appError := new_organisation.IsValid(); appError != nil {
		transaction.Rollback()
		return NewLocAppError("organisation_store_impl.Update.organisation_new.PreSave", appError.Id, nil, appError.DetailedError)
	}
	if err := transaction.Model(&organisation).Updates(&new_organisation).Error; err != nil {
		transaction.Rollback()
		return NewLocAppError("organisation_store_impl.Update", "update.transaction.updates.encounter_error: "+err.Error(), nil, "")
	}
	transaction.Commit()
	return nil
}

// Used to get organisation from DB
func (osi OrganisationStoreImpl) Get(ds DataStore) *models.Organisation {
	db := *ds.Db
	organisation := models.Organisation{}
	db.First(&organisation)
	return &organisation
}
