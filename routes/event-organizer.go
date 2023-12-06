package routes

import (
	"eve/handlers"

	"github.com/gorilla/mux"
)

func EOrganizerRoutes(r *mux.Router) {
	e := r.PathPrefix("/event-organizers").Subrouter()

	e.HandleFunc("/register", handlers.CreateUserHandler).Methods("POST")
	e.HandleFunc("/login", handlers.LoginUserHandler).Methods("POST")
}
