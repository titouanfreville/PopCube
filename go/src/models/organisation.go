package models

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"
	u "utils"
)

const (
	organisationDisplayNameMaxRunes = 64
	organisationNameMaxLength       = 64
	organisationDescriptionMaxRunes = 1024
	organisationSubjectMaxRunes     = 250
)

type Organisation struct {
	IDOrganisation   uint64 `gorm:"primary_key;column:idOrganisation;AUTO_INCREMENT" json:"-"`
	DockerStack      int    `gorm:"column:dockerStack;not null;unique" json:"docker_stack"`
	OrganisationName string `gorm:"column:organisationName;not null;unique" json:"display_name"`
	Description      string `gorm:"column:desciption" json:"description,omitempty"`
	Avatar           string `gorm:"column:avatar" json:"avatar,omitempty"`
	Domain           string `gorm:"column:domain" json:"avatar,omitempty"`
}

func (organisation *Organisation) ToJSON() string {
	b, err := json.Marshal(organisation)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func OganisationFromJSON(data io.Reader) *Organisation {
	decoder := json.NewDecoder(data)
	var organisation Organisation
	err := decoder.Decode(&organisation)
	if err == nil {
		return &organisation
	} else {
		return nil
	}
}

func (organisation *Organisation) IsValid() *u.AppError {

	if len(organisation.OrganisationName) == 0 || utf8.RuneCountInString(organisation.OrganisationName) > organisationDisplayNameMaxRunes {
		return u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10))
	}

	if !IsValidOrganisationIDentifier(organisation.OrganisationName) {
		return u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.not_alphanum_organisation_name.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10))
	}

	if utf8.RuneCountInString(organisation.Description) > organisationDescriptionMaxRunes {
		return u.NewLocAppError("Organisation.IsValid", "model.organisation.is_valid.description.app_error", nil, "id="+strconv.FormatUint(organisation.IDOrganisation, 10))
	}

	return nil
}

func (organisation *Organisation) PreSave() {
	organisation.OrganisationName = strings.ToLower(organisation.OrganisationName)

	if organisation.Avatar == "" {
		organisation.Avatar = "default_organisation_avatar.svg"
	}
}
