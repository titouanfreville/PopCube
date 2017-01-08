package models

import (
	"encoding/json"
	"io"
)

const (
	DEFAULT_LOCALE     = "fr_FR"
	DEFAULT_TIMEZONE   = "UTC-0"
	LOCAL_MAX_SIZE     = 5
	TIME_ZONE_MAX_SIZE = 6
	MAX_TIME           = 1440
)

type Parameter struct {
	IdParameter uint64 `gorm:"primary_key;column:idParameter;AUTO_INCREMENT" json:"-"`
	Local       string `gorm:"column:local;not null; unique" json:"local"`
	TimeZone    string `gorm:"column:timeZone;not null; unique;" json:"time_zone"`
	SleepStart  int    `gorm:"column:sleepStart;not null;unique" json:"sleep_start"`
	SleepEnd    int    `gorm:"column:sleepEnd;not null;unique" json:"sleep_end"`
}

func (parameter *Parameter) toJson() string {
	b, err := json.Marshal(parameter)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func parameterFromJson(data io.Reader) *Parameter {
	decoder := json.NewDecoder(data)
	var parameter Parameter
	err := decoder.Decode(&parameter)
	if err == nil {
		return &parameter
	} else {
		return nil
	}
}

func (parameter *Parameter) isValid() *AppError {

	if len(parameter.Local) == 0 || len(parameter.Local) > LOCAL_MAX_SIZE {
		return NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_local.app_error", nil, "")
	}

	if len(parameter.TimeZone) == 0 || len(parameter.TimeZone) > TIME_ZONE_MAX_SIZE {
		return NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_timezone.app_error", nil, "")
	}

	if parameter.SleepStart < 0 || parameter.SleepStart > MAX_TIME {
		return NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_sleep_start.app_error", nil, "")
	}

	if parameter.SleepEnd < 0 || parameter.SleepEnd > MAX_TIME {
		return NewLocAppError("Parameter.IsValid", "model.parameter.is_valid.parameter_sleep_end.app_error", nil, "")
	}

	return nil
}

func (parameter *Parameter) preSave() {
	if parameter.Local == "" {
		parameter.Local = DEFAULT_LOCALE
	}
	if parameter.TimeZone == "" {
		parameter.TimeZone = DEFAULT_TIMEZONE
	}
}
