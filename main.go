package main

import (
	"eve/database"
	"eve/util"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("An error occurred while connecting to db: ", err)
	}

	util.InsertRoles()

	srv := &http.Server{
		Addr:         "127.0.0.1:3500",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println(db)
	fmt.Printf("Server in running on port %v\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
