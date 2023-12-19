package models

import (
	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	BuyerID uint8
	EventID uint8
}
