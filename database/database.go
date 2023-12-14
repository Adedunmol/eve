package database

import (
	"eve/models"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func InitDB() (DbInstance, error) {
	var err error

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the db: ", err)
	}
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Event{})

	Database = DbInstance{Db: db}

	return Database, nil
}
