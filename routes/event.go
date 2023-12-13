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
	// eventOrganizerChain := alice.New(middleware.AuthMiddleware, middleware.EventOrganizerRoute).Then(http.HandlerFunc(handlers.CreateEventHandler))
	// u.Handle("/", eventOrganizerChain).Methods("POST")
	u.Handle("/", middleware.AuthMiddleware(middleware.RoleAuthorization(http.HandlerFunc(handlers.CreateEventHandler), util.CREATE_EVENT))).Methods("POST")

}
