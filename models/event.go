package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name      string
	Tickets   int
	Organizer User
}
