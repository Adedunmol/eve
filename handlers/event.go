package handlers

import (
	"encoding/json"
	"eve/database"
	"eve/models"
	"eve/util"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
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

type ReserveSpot struct {
	Tickets int `json:"tickets"`
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

func GetEventHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	if vars["id"] == "" {
		util.RespondWithJSON(w, http.StatusBadRequest, APIResponse{Message: "no event id sent in the url param", Data: nil, Status: "error"})
		return
	}

	var event models.Event

	result := database.Database.Db.First(&event, vars["id"])

	if result.Error != nil {
		fmt.Println(result.Error)
		util.RespondWithJSON(w, http.StatusNotFound, APIResponse{Message: "event not found", Data: nil, Status: "success"})
		return
	}

	util.RespondWithJSON(w, http.StatusOK, APIResponse{Message: "", Data: event, Status: "success"})
}

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	if vars["id"] == "" {
		util.RespondWithJSON(w, http.StatusBadRequest, APIResponse{Message: "no event id sent in the url param", Data: nil, Status: "error"})
		return
	}

	var event models.Event

	result := database.Database.Db.Where("deleted_at is null").First(&event, vars["id"])

	if result.Error != nil {
		fmt.Println(result.Error)
		util.RespondWithJSON(w, http.StatusNotFound, APIResponse{Message: "event not found", Data: nil, Status: "error"})
		return
	}

	result = database.Database.Db.Delete(&event)

	if result.Error != nil {
		fmt.Println(result.Error)
		util.RespondWithJSON(w, http.StatusNotFound, APIResponse{Message: "error deleting event", Data: nil, Status: "error"})
		return
	}

	util.RespondWithJSON(w, http.StatusOK, APIResponse{Message: "", Data: event, Status: "success"})
}

func GetAllEventsHandler(w http.ResponseWriter, r *http.Request) {
	var events []models.Event

	database.Database.Db.Where("deleted_at is null").Find(&events)

	util.RespondWithJSON(w, http.StatusOK, APIResponse{Message: "", Data: events, Status: "success"})
}

func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	if vars["id"] == "" {
		util.RespondWithJSON(w, http.StatusBadRequest, APIResponse{Message: "no event id sent in the url param", Data: nil, Status: "error"})
		return
	}

	var event models.Event

	result := database.Database.Db.Where("deleted_at is null").First(&event, vars["id"])

	if result.Error != nil {
		fmt.Println(result.Error)
		util.RespondWithJSON(w, http.StatusNotFound, APIResponse{Message: "event not found", Data: nil, Status: "success"})
		return
	}

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

	result = database.Database.Db.Model(&event).Updates(models.Event{
		Name:     eventDto.Name,
		About:    eventDto.About,
		Tickets:  eventDto.Tickets,
		Price:    eventDto.Price,
		Location: eventDto.Location,
		Category: eventDto.Category,
	})

	if result.Error != nil {
		fmt.Println(err)
		util.RespondWithError(w, http.StatusInternalServerError, "Error updating event")
		return
	}

	util.RespondWithJSON(w, http.StatusOK, APIResponse{Message: "", Data: event, Status: "success"})
}

func ReserveEvent(w http.ResponseWriter, r *http.Request) {

	wg := new(sync.WaitGroup)

	vars := mux.Vars(r)

	if vars["id"] == "" {
		util.RespondWithJSON(w, http.StatusBadRequest, APIResponse{Message: "no event id sent in the url param", Data: nil, Status: "error"})
		return
	}

	var event models.Event

	result := database.Database.Db.Where("deleted_at is null").First(&event, vars["id"])

	if result.Error != nil {
		fmt.Println(result.Error)
		util.RespondWithJSON(w, http.StatusNotFound, APIResponse{Message: "event not found", Data: nil, Status: "error"})
		return
	}

	if event.Tickets == 0 {
		util.RespondWithJSON(w, http.StatusBadRequest, APIResponse{Message: "event sold out", Data: nil, Status: "success"})
		return
	}

	var eventDto ReserveSpot
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

	if eventDto.Tickets > eventDto.Tickets {
		util.RespondWithJSON(w, http.StatusBadRequest, APIResponse{Message: "tickets not enough", Data: nil, Status: "success"})
		return
	}

	username := r.Context().Value("username")

	var foundUser models.User

	result = database.Database.Db.Where(models.User{Username: username.(string)}).First(&foundUser)

	if result.Error != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, APIResponse{Message: "user does not exist", Data: nil, Status: "error"})
		return
	}

	purchase := models.Purchase{
		EventID: uint8(event.ID),
		BuyerID: uint8(foundUser.ID),
	}

	result = database.Database.Db.Create(&purchase)

	if result.Error != nil {
		fmt.Println(result.Error)
		util.RespondWithJSON(w, http.StatusInternalServerError, APIResponse{Message: "error creating a purchase", Data: nil, Status: "error"})
		return
	}

	event.Tickets -= eventDto.Tickets

	result = database.Database.Db.Model(&event).Update("tickets", event.Tickets)

	if result.Error != nil {
		fmt.Println(err)
		util.RespondWithError(w, http.StatusInternalServerError, "Error updating event")
		return
	}

	message := fmt.Sprintf(`
You just purchased %d ticket(s) to attend %s

If you didn't make the purchase, kindly ignore
	`, eventDto.Tickets, event.Name)

	mux := &sync.Mutex{}

	fileStr := ""

	wg.Add(2)
	go util.GeneratePdf(event, foundUser, &fileStr, mux, wg)

	go util.SendMail(foundUser.Email, "Tickect purchase", []byte(message), wg)

	util.RespondWithJSON(w, http.StatusCreated, APIResponse{Message: "", Data: purchase, Status: "success"})

	wg.Wait()

	err = os.Remove(fileStr)

	if err != nil {
		fmt.Println(err)
		return
	}
}
