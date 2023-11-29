package models

import "gorm.io/gorm"

const (
	CREATE_EVENT uint8 = 1
	MODIFY_EVENT uint8 = 2
	DELETE_EVENT uint8 = 4
	MODERATE     uint8 = 8
	ADMIN        uint8 = 16
)

type Role struct {
	gorm.Model
	Name        string
	Permissions uint8
	Default     bool
}
