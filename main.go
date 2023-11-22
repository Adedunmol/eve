package main

import (
	"eve/database"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type DBClient struct {
	db *gorm.DB
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("An error occurred while connecting to db: ", err)
	}

	dbClient := &DBClient{db: db}

	fmt.Println("connection: ", dbClient)
}
