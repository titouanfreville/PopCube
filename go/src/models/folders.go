package models

import (
	"encoding/json"
	"io"
)

type Folder struct {
	IdFolder uint64  `gorm:"primary_key;column:idFolder;AUTO_INCREMENT" json:"-"`
	Link     string  `gorm:"column:link;not null;unique" json:"link"`
	Name     string  `gorm:"column:name;not null;unique" json:"name"`
	Type     string  `gorm:"column:type;not null;" json:"type"`
	Message  Message `gorm:"column:message; not null;ForeignKey:IdMessage;" json:"-"`
}

func (folder *Folder) isValid() *AppError {
	if len(folder.Name) == 0 {
		return NewLocAppError("Folder.IsValid", "model.folder.name.app_error", nil, "")
	}

	if len(folder.Link) == 0 {
		return NewLocAppError("Folder.IsValid", "model.folder.link.app_error", nil, "")
	}
	if len(folder.Type) == 0 {
		return NewLocAppError("Folder.IsValid", "model.folder.type.app_error", nil, "")
	}
	if folder.Message == (Message{}) {
		return NewLocAppError("Folder.IsValid", "model.folder.message.app_error", nil, "")
	}
	return nil
}

func (folder *Folder) toJson() string {
	b, err := json.Marshal(folder)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func folderFromJson(data io.Reader) *Folder {
	decoder := json.NewDecoder(data)
	var folder Folder
	err := decoder.Decode(&folder)
	if err == nil {
		return &folder
	} else {
		return nil
	}
}

func folderListToJson(folderList []*Folder) string {
	b, err := json.Marshal(folderList)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func folderListFromJson(data io.Reader) []*Folder {
	decoder := json.NewDecoder(data)
	var folderList []*Folder
	err := decoder.Decode(&folderList)
	if err == nil {
		return folderList
	} else {
		return nil
	}
}
