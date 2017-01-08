package models

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	ORGANISATION_DISPLAY_NAME_MAX_RUNES = 64
	ORGANISATION_NAME_MAX_LENGTH        = 64
	ORGANISATION_DESCRIPTION_MAX_RUNES  = 1024
	ORGANISATION_SUBJECT_MAX_RUNES      = 250
)

type Organisation struct {
	IdOrganisation   uint64 `gorm:"primary_key;column:idOrganisation;AUTO_INCREMENT" json:"-"`
	DockerStack      int    `gorm:"column:dockerStack;not null;unique" json:"docker_stack"`
	OrganisationName string `gorm:"column:organisationName;not null;unique" json:"display_name"`
	Description      string `gorm:"column:desciption" json:"description,omitempty"`
	Avatar           string `gorm:"column:avatar" json:"avatar,omitempty"`
	Domain           string `gorm:"column:domain" json:"avatar,omitempty"`
}

func (organisation *Organisation) toJson() string {
	b, err := json.Marshal(organisation)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func organisationFromJson(data io.Reader) *Organisation {
	decoder := json.NewDecoder(data)
	var organisation Organisation
	err := decoder.Decode(&organisation)
	if err == nil {
		return &organisation
	} else {
		return nil
	}
}

func (organisation *Organisation) isValid() *AppError {

	if len(organisation.OrganisationName) == 0 || utf8.RuneCountInString(organisation.OrganisationName) > ORGANISATION_DISPLAY_NAME_MAX_RUNES {
		return NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IdOrganisation, 10))
	}

	if !IsValidOrganisationIdentifier(organisation.OrganisationName) {
		return NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.not_alphanum_organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IdOrganisation, 10))
	}

	if utf8.RuneCountInString(organisation.Description) > ORGANISATION_DESCRIPTION_MAX_RUNES {
		return NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.description.app_error", nil, "id="+strconv.FormatUint(organisation.IdOrganisation, 10))
	}

	return nil
}

func (organisation *Organisation) preSave() {
	organisation.OrganisationName = strings.ToLower(organisation.OrganisationName)

	if organisation.Avatar == "" {
		organisation.Avatar = "default_organisation_avatar.svg"
	}
}
