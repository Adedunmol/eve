package routes

import "github.com/gorilla/mux"

func RoutesSetup(r *mux.Router) *mux.Router {

	EventRoutes(r)
	UserRoutes(r)
	EOrganizerRoutes(r)

	return r
}
