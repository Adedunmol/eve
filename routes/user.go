package routes

import (
	"fmt"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	u := r.PathPrefix("/users").Subrouter()

	fmt.Println(u)

}
