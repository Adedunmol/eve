package models

import (
	"fmt"

	"gorm.io/gorm"
)

const (
	USER         uint8 = 1
	CREATE_EVENT uint8 = 2
	MODIFY_EVENT uint8 = 4
	DELETE_EVENT uint8 = 8
	MODERATE     uint8 = 16
	ADMIN        uint8 = 32
)

type Role struct {
	gorm.Model
	Name        string
	Permissions uint8
	Default     bool
}

func (r *Role) AddPermission(perm uint8) {
	if !r.HasPermission(perm) {
		r.Permissions += perm
	}
}

func (r *Role) RemovePermission(perm uint8) {
	if r.HasPermission(perm) {
		r.Permissions -= perm
	}
}

func (r *Role) ResetPermissions(perm uint8) {
	r.Permissions = 0
}

func (r *Role) HasPermission(perm uint8) bool {
	return (r.Permissions & perm) == perm
}

func InsertRoles() {
	roles := make(map[string][]uint8)

	roles["User"] = []uint8{USER}
	roles["Event-Organizer"] = []uint8{USER, CREATE_EVENT, MODIFY_EVENT, DELETE_EVENT}
	roles["Moderator"] = []uint8{USER, CREATE_EVENT, MODIFY_EVENT, DELETE_EVENT, MODERATE}
	roles["Admin"] = []uint8{USER, CREATE_EVENT, MODIFY_EVENT, DELETE_EVENT, MODERATE, ADMIN}

	default_role := "User"
	// for r := range roles {
	// 	role := db
	// }

	fmt.Println(roles, default_role)
}
