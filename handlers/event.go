package handlers

import (
	"encoding/json"
	"eve/models"
	"eve/util"
	"fmt"
	"net/http"
	"time"
)

type CreateEventDto struct {
	Name     string    `json:"name"`
	About    string    `json:"about"`
	Tickets  int       `json:"tickets"`
	Price    int       `json:"price"`
	Location string    `json:"location"`
	Category string    `json:"category"`
	Date     time.Time `json:"date"`
}

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event CreateEventDto
	err := json.NewDecoder(r.Body).Decode(&event)

	if _, ok := err.(*json.InvalidUnmarshalError); ok {
		fmt.Println(err)
		util.RespondWithError(w, http.StatusInternalServerError, "Unable to format the request body")
		return
	}

	if err != nil {
		fmt.Println(err)
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	var foundUser models.User

	username := r.Context().Value("username")

	fmt.Println(foundUser)
	fmt.Println(username)
}
