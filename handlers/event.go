package handlers

import (
	"encoding/json"
	"eve/database"
	"eve/models"
	"eve/util"
	"fmt"
	"net/http"
)

type CreateEventDto struct {
	Name     string `json:"name"`
	About    string `json:"about"`
	Tickets  int    `json:"tickets"`
	Price    int    `json:"price"`
	Location string `json:"location"`
	Category string `json:"category"`
	// Date     time.Time `json:"date"`
}

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var eventDto CreateEventDto
	err := json.NewDecoder(r.Body).Decode(&eventDto)

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

	username := r.Context().Value("username")

	var foundUser models.User

	result := database.Database.Db.Where(models.User{Username: username.(string)}).First(&foundUser)

	if result.Error != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, APIResponse{Message: "user does not exist", Data: nil, Status: "error"})
		return
	}

	event := models.Event{
		Name:        eventDto.Name,
		About:       eventDto.About,
		Tickets:     eventDto.Tickets,
		Price:       eventDto.Price,
		Location:    eventDto.Location,
		Category:    eventDto.Category,
		OrganizerID: foundUser.ID,
	}

	result = database.Database.Db.Create(&event)

	if result.Error != nil {
		fmt.Println(result.Error)
		util.RespondWithJSON(w, http.StatusInternalServerError, APIResponse{Message: "error creating event", Data: nil, Status: "error"})
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, APIResponse{Message: "", Data: event, Status: "success"})
}
