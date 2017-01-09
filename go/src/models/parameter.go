package models

import (
	"encoding/json"
	"io"
	u "utils"
)

const (
	DefaultLocale   = "fr_FR"
	DefaultTimezone = "UTC-0"
	localMaxSize    = 5
	timeZoneMaxSize = 6
	maxTime         = 1440
)

type Parameter struct {
	IDParameter uint64 `gorm:"primary_key;column:idParameter;AUTO_INCREMENT" json:"-"`
	Local       string `gorm:"column:local;not null; unique" json:"local"`
	TimeZone    string `gorm:"column:timeZone;not null; unique;" json:"time_zone"`
	SleepStart  int    `gorm:"column:sleepStart;not null;unique" json:"sleep_start"`
	SleepEnd    int    `gorm:"column:sleepEnd;not null;unique" json:"sleep_end"`
}

func (parameter *Parameter) ToJSON() string {
	b, err := json.Marshal(parameter)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func ParameterFromJSON(data io.Reader) *Parameter {
	decoder := json.NewDecoder(data)
	var parameter Parameter
	err := decoder.Decode(&parameter)
	if err == nil {
		return &parameter
	} else {
		return nil
	}
}

func (parameter *Parameter) IsValid() *u.AppError {

	if len(parameter.Local) == 0 || len(parameter.Local) > localMaxSize {
		return u.NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_local.app_error", nil, "")
	}

	if len(parameter.TimeZone) == 0 || len(parameter.TimeZone) > timeZoneMaxSize {
		return u.NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_timezone.app_error", nil, "")
	}

	if parameter.SleepStart < 0 || parameter.SleepStart > maxTime {
		return u.NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_sleep_start.app_error", nil, "")
	}

	if parameter.SleepEnd < 0 || parameter.SleepEnd > maxTime {
		return u.NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_sleep_end.app_error", nil, "")
	}

	return nil
}

func (parameter *Parameter) PreSave() {
	if parameter.Local == "" {
		parameter.Local = DefaultLocale
	}
	if parameter.TimeZone == "" {
		parameter.TimeZone = DefaultTimezone
	}
}
