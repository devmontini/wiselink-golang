package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/devmontini/wiselink-golang/database"
	"github.com/devmontini/wiselink-golang/models"
	"github.com/gorilla/mux"
)

func filter(in []models.Events, Status bool) []models.Events {
	var out []models.Events
	for _, each := range in {
		if each.Status {
			out = append(out, each)
		}
	}
	return out
}

func CreateEventsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the CREATE EVENTS!")
	var event models.Events

	json.NewDecoder(r.Body).Decode(&event)

	eventData := database.DB.Create(&event)
	err := eventData.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		message := map[string]interface{}{
			"message": "Event created successfully",
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(message)
	}
}

type Event struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	ShortDesc string    `json:"short_desc"`
	Date      time.Time `json:"date"`
	Organizer string    `json:"organizer"`
	Place     string    `json:"place"`
	Status    bool      `json:"status"`
}

type Users struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

func GetEventsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the GET EVENTS!")
	event := []models.Events{}
	var user models.Users

	//AGARRAMOS EL ID Y LO PASAMOS A INT
	userid := mux.Vars(r)
	intId, err := strconv.Atoi(userid["userId"])
	if err != nil {
		fmt.Println(err)
	}

	//BUSCAMOS LA DATA DEL USUARIo
	userData := database.DB.Where("id = 1", intId).First(&user)
	err = userData.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	//BUSCAMOS TODOS LOS EVENTOS
	eventData := database.DB.Find(&event)
	err = eventData.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	//VEMOS SI ES ADMIN O NO
	if !user.Admin {
		filterEvents := filter(event, true)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(filterEvents)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(event)
	}
}

func GetEventHandlerById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the GET EVENT by ID!")
	params := mux.Vars(r)
	var event models.Events

	eventData := database.DB.Where("id = ?", params["id"]).First(&event)
	err := eventData.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(event)
	}
}

func DeleteEventHandlerById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the DELETE EVENT by ID!")
}

func EditEventHandlerById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the EDIT EVENT by ID!")
}
