package routes

import (
	"eve/handlers"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	u := r.PathPrefix("/users").Subrouter()

	u.HandleFunc("/register", handlers.CreateUserHandler).Methods("POST")
}
