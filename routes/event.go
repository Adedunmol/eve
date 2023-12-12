package routes

import (
	"eve/handlers"
	"eve/middleware"
	"eve/util"
	"net/http"

	"github.com/gorilla/mux"
)

func EventRoutes(r *mux.Router) {
	u := r.PathPrefix("/events").Subrouter()

	// u.Handle("/", middleware.AuthMiddleware(http.HandlerFunc(handlers.CreateEventHandler))).Methods("POST")
	u.Handle("/", middleware.RoleAuthorization(middleware.AuthMiddleware(http.HandlerFunc(handlers.CreateEventHandler)), util.CREATE_EVENT)).Methods("POST")
}
