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

	u.Handle("/", middleware.AuthMiddleware(middleware.RoleAuthorization(http.HandlerFunc(handlers.CreateEventHandler), util.CREATE_EVENT))).Methods("POST")
	u.Handle("/", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetAllEventsHandler))).Methods("GET")
	u.Handle("/{id}", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetEventHandler))).Methods("GET")
	u.Handle("/{id}", middleware.AuthMiddleware(middleware.RoleAuthorization(http.HandlerFunc(handlers.UpdateEventHandler), util.MODIFY_EVENT))).Methods("PATCH")
	u.Handle("/{id}", middleware.AuthMiddleware(middleware.RoleAuthorization(http.HandlerFunc(handlers.DeleteEventHandler), util.DELETE_EVENT))).Methods("DELETE")

	u.Handle("/{id}/reserve", middleware.AuthMiddleware(http.HandlerFunc(handlers.ReserveEvent))).Methods("POST")
}
