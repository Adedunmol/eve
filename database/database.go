package database

import (
	"eve/models"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type DbInstance struct {
// 	Db *gorm.DB
// }

func InitDb() (*gorm.DB, error) {
	var err error

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		return nil, err
	} else {
		db.AutoMigrate(&models.User{})
		return db, nil
	}
}
