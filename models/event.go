package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name      string    `json:"name"`
	About     string    `json:"about"`
	Tickets   int       `json:"tickets"`
	Price     int       `json:"price"`
	Location  string    `json:"location"`
	Category  string    `json:"category"`
	Date      time.Time `json:"date"`
	Organizer User
}
