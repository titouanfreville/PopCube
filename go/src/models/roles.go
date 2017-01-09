package models

import (
	"encoding/json"
	"io"
	"regexp"
	u "utils"
)

type Role struct {
	IDRole        uint64 `gorm:"primary_key;column:idRole;AUTO_INCREMENT" json:"-"`
	RoleName      string `gorm:"column:roleName;unique_index;not null;unique" json:"name"`
	CanUsePrivate bool   `gorm:"column:canUsePrivate;not null" json:"canUsePrivate"`
	CanModerate   bool   `gorm:"column:canModerate;not null" json:"canModerate"`
	CanArchive    bool   `gorm:"column:canArchive;not null" json:"canArchive"`
	CanInvite     bool   `gorm:"column:canInvite;not null" json:"canInvite"`
	CanManage     bool   `gorm:"column:canManage;not null" json:"canManage"`
	CanManageUser bool   `gorm:"column:canManageUser;not null" json:"canManageUser"`
}

var (
	Owner = Role{
		RoleName:      "owner",
		CanUsePrivate: true,
		CanModerate:   true,
		CanArchive:    true,
		CanInvite:     true,
		CanManage:     true,
		CanManageUser: true,
	}
	Admin = Role{
		RoleName:      "admin",
		CanUsePrivate: true,
		CanModerate:   true,
		CanArchive:    true,
		CanInvite:     true,
		CanManage:     true,
		CanManageUser: true,
	}
	Standart = Role{
		RoleName:      "standart",
		CanUsePrivate: true,
		CanModerate:   true,
		CanArchive:    true,
		CanInvite:     false,
		CanManage:     false,
		CanManageUser: false,
	}
	Guest = Role{
		RoleName:      "guest",
		CanUsePrivate: false,
		CanModerate:   false,
		CanArchive:    false,
		CanInvite:     false,
		CanManage:     false,
		CanManageUser: false,
	}
	BasicsRoles = []*Role{
		&Owner,
		&Admin,
		&Standart,
		&Guest,
	}
	restrictedRoleNames = []string{
		"owner",
		"admin",
		"standart",
		"guest",
	}
	validRoleNameChars = regexp.MustCompile(`^[a-z]+$`)
)

func (role *Role) IsValid() *u.AppError {
	if !IsValidRoleName(role.RoleName) {
		return u.NewLocAppError("Role.IsValid", "model.role.rolename.app_error", nil, "")
	}

	return nil
}

func (role *Role) PreSave() {
	if role.RoleName == "" {
		role.RoleName = NewID()
	}
}

func (role *Role) ToJSON() string {
	b, err := json.Marshal(role)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func IsValidRoleName(u string) bool {
	if len(u) == 0 || len(u) > 64 {
		return false
	}

	if !validRoleNameChars.MatchString(u) {
		return false
	}

	for _, restrictedRoleName := range restrictedRoleNames {
		if u == restrictedRoleName {
			return false
		}
	}

	return true
}

func RoleFromJSON(data io.Reader) *Role {
	decoder := json.NewDecoder(data)
	var role Role
	err := decoder.Decode(&role)
	if err == nil {
		return &role
	} else {
		return nil
	}
}

func RoleListToJSON(roleList []*Role) string {
	b, err := json.Marshal(roleList)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func RoleListFromJSON(data io.Reader) []*Role {
	decoder := json.NewDecoder(data)
	var roleList []*Role
	err := decoder.Decode(&roleList)
	if err == nil {
		return roleList
	} else {
		return nil
	}
}
