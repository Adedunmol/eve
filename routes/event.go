package routes

import (
	"eve/handlers"
	"eve/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func EventRoutes(r *mux.Router) {
	u := r.PathPrefix("/events").Subrouter()

	u.Handle("/", middleware.AuthMiddleware(http.HandlerFunc(handlers.CreateEventHandler))).Methods("POST")
}
