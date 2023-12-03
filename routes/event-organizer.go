package routes

import (
	"fmt"

	"github.com/gorilla/mux"
)

func EOrganizerRoutes(r *mux.Router) {
	e := r.PathPrefix("/event-organizers").Subrouter()

	fmt.Println(e)
}
