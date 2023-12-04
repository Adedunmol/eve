package handlers

import (
	"encoding/json"
	"eve/util"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if _, ok := err.(*json.InvalidUnmarshalError); ok {
		util.RespondWithError(w, http.StatusInternalServerError, "Unable to format the request body")
		return
	}

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
}
