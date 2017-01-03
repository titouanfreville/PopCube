package models

import (
	"encoding/json"
	"io"
	"regexp"
)

type Role struct {
	RoleId        uint   `gorm:"primary_key;column:roleId;AUTO_INCREMENT" json:"-"`
	RoleName      string `gorm:"column:roleName;unique_index" json:"name"`
	CanUsePrivate bool   `gorm:"column:canUsePrivate" json:"canUsePrivate"`
	CanModerate   bool   `gorm:"column:canModerate" json:"canModerate"`
	CanArchive    bool   `gorm:"column:canArchive" json:"canArchive"`
	CanInvite     bool   `gorm:"column:canInvite" json:"canInvite"`
	CanManage     bool   `gorm:"column:canManage" json:"canManage"`
	CanManageUser bool   `gorm:"column:canManageUser" json:"canManageUser"`
}

var (
	OWNER = Role{
		RoleName:      "owner",
		CanUsePrivate: true,
		CanModerate:   true,
		CanArchive:    true,
		CanInvite:     true,
		CanManage:     true,
		CanManageUser: true,
	}
	ADMIN = Role{
		RoleName:      "admin",
		CanUsePrivate: true,
		CanModerate:   true,
		CanArchive:    true,
		CanInvite:     true,
		CanManage:     true,
		CanManageUser: true,
	}
	STANDART = Role{
		RoleName:      "standart",
		CanUsePrivate: true,
		CanModerate:   true,
		CanArchive:    true,
		CanInvite:     false,
		CanManage:     false,
		CanManageUser: false,
	}
	GUEST = Role{
		RoleName:      "guest",
		CanUsePrivate: false,
		CanModerate:   false,
		CanArchive:    false,
		CanInvite:     false,
		CanManage:     false,
		CanManageUser: false,
	}
	BASICS_ROLES = []*Role{
		&OWNER,
		&ADMIN,
		&STANDART,
		&GUEST,
	}
	restrictedRoleNames = []string{
		"owner",
		"admin",
		"standart",
		"guest",
	}
	validRoleNameChars = regexp.MustCompile(`^[a-z]+$`)
)

func (role *Role) isValid() *AppError {
	if !isValidRoleName(role.RoleName) {
		return NewLocAppError("Role.IsValid", "model.role.rolename.app_error", nil, "")
	}

	return nil
}

func (role *Role) preSave() {
	if role.RoleName == "" {
		role.RoleName = NewId()
	}
}

func (role *Role) toJson() string {
	b, err := json.Marshal(role)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func isValidRoleName(u string) bool {
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

func roleFromJson(data io.Reader) *Role {
	decoder := json.NewDecoder(data)
	var role Role
	err := decoder.Decode(&role)
	if err == nil {
		return &role
	} else {
		return nil
	}
}

func roleListToJson(roleList []*Role) string {
	b, err := json.Marshal(roleList)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func roleListFromJson(data io.Reader) []*Role {
	decoder := json.NewDecoder(data)
	var roleList []*Role
	err := decoder.Decode(&roleList)
	if err == nil {
		return roleList
	} else {
		return nil
	}
}
