package main

import (
	"eve/database"
	"eve/routes"
	"eve/util"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func Initializers() database.DbInstance {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("An error occurred while connecting to db: ", err)
	}

	util.InsertRoles()

	return db
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Hello")
}

func main() {

	db := Initializers()

	r := mux.NewRouter()

	rNew := routes.RoutesSetup(r.PathPrefix("/api/v1").Subrouter())

	rNew.HandleFunc("/", testHandler).Methods("GET")

	srv := &http.Server{
		Addr:         "127.0.0.1:3500",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      rNew,
	}

	fmt.Println(db)
	fmt.Printf("Server in running on port %v\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
