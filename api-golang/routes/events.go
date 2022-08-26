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

func filter(in []Event, Status bool) []Event {
	var out []Event
	for _, each := range in {
		if each.Status {
			out = append(out, each)
		}
	}
	return out
}

func filterByTitle(ev []Event, title string) []Event {
	var out []Event
	for _, each := range ev {
		if each.Title == title {
			out = append(out, each)
		}
	}
	return out
}

type Event struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	ShortDesc string    `json:"short_desc"`
	Date      time.Time `json:"date"`
	Place     string    `json:"place"`
	Status    bool      `json:"status"`
}

type Users struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

func CreateEventsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the CREATE EVENTS!")
	var event models.Events
	var user models.Users

	//AGARRAMOS EL ID Y LO PASAMOS A INT
	v := r.URL.Query()
	userId := v.Get("userId")

	//BUSCAMOS LA DATA DEL USUARIO
	userData := database.DB.Where("id = ?", userId).First(&user)
	err := userData.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	//SI NO ES ADMIN LE MANDAMOS ERROR
	if !user.Admin {
		message := map[string]interface{}{
			"message": "Only admins have the power!",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}

	//Si ES ADMIN PROCEDEMOS A CREAR EL EVENTO
	json.NewDecoder(r.Body).Decode(&event)

	eventData := database.DB.Create(&event)
	err = eventData.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	message := map[string]interface{}{
		"message": "Event created successfully",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message)
}

func GetEventsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the GET EVENTS!")
	var event = []Event{}
	var user models.Users

	v := r.URL.Query()
	title := v.Get("title")
	userId := v.Get("userId")

	//BUSCAMOS LA DATA DEL USUARIo
	userData := database.DB.Where("id = ?", userId).First(&user)
	err := userData.Error
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

	if title != "" {
		//SI NO ES ADMIN
		if !user.Admin {
			filterEvents := filter(event, true)
			filterByTitle := filterByTitle(filterEvents, title)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(filterByTitle)
			return
		}

		//SI ES ADMIN
		filterByTitle := filterByTitle(event, title)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(filterByTitle)
		return
	}

	//SI NO ES ADMIN LE PASAMOS LOS EVENTOS PUBLICADOS
	if !user.Admin {
		filterEvents := filter(event, true)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(filterEvents)
		return
	}

	//SI ES ADMIN LES PASAMOS TODOS LOS EVENTOS
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}

func GetEventHandlerById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the GET EVENT by ID!")
	params := mux.Vars(r)
	var event models.Events

	eventData := database.DB.Where("id = ?", params["eventId"]).First(&event)
	err := eventData.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(event)
	}
}

func EditEventHandlerById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the EDIT EVENT by ID!")
	var event models.Events
	var user models.Users

	//AGARRAMOS EL ID DE USER Y LO PASAMOS A INT
	v := r.URL.Query()
	userId := v.Get("userId")

	//BUSCAMOS LA DATA DEL USUARIO
	userData := database.DB.Where("id = ?", userId).First(&user)
	err := userData.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	//SI NO ES ADMIN LE MANDAMOS ERROR
	if !user.Admin {
		message := map[string]interface{}{
			"message": "Only admins have the power!",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	} else {
		//SI ES ADMIN PROCEDE A UPDATEAR
		userids := mux.Vars(r)
		intIdevent, err := strconv.Atoi(userids["eventId"])
		if err != nil {
			fmt.Println(err)
			return
		}

		json.NewDecoder(r.Body).Decode(&event)

		database.DB.Where("id = ?", intIdevent).Updates(&event)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		message := map[string]interface{}{
			"message": "Event edited!",
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message)
	}
}

func DeleteEventHandlerById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the DELETE EVENT by ID!")
	var event models.Events
	var user models.Users

	//AGARRAMOS EL IDs
	v := r.URL.Query()
	userId := v.Get("userId")

	//BUSCAMOS LA DATA DEL USUARIo
	userData := database.DB.Where("id = ?", userId).First(&user)
	err := userData.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	//SI NO ES ADMIN NO PUEDE HACER DELETE
	if !user.Admin {
		message := map[string]interface{}{
			"message": "Only admins have the power!",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}

	//AGARRAMOS EL ID DEL EVENTO Y LO PASAMOS A INT
	userids := mux.Vars(r)
	intIdevent, err := strconv.Atoi(userids["eventId"])
	if err != nil {
		fmt.Println(err)
		return
	}

	//SE PROCEDE A BORRAR EL EVENTO
	database.DB.Where("id = ?", intIdevent).Delete(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	message := map[string]interface{}{
		"message": "Event deleted successfully",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}

type UsersEvents struct {
	UsersID  int `json:"usersid" gorm:"primaryKey"`
	EventsID int `json:"eventsid" gorm:"primaryKey"`
}

func InscEventHandlerById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the EVENT in USER by ID!")

	v := r.URL.Query()
	userId := v.Get("userId")
	intUserid, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println(err)
		return
	}

	eventId := mux.Vars(r)
	intIdevent, err := strconv.Atoi(eventId["eventId"])
	if err != nil {
		fmt.Println(err)
		return
	}

	var jointable = []UsersEvents{{UsersID: intUserid}, {EventsID: intIdevent}}
	if err != nil {
		fmt.Println(err)
		return
	}

	database.DB.Create(&jointable)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetUserEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the EVENT in USER by ID!")
	var jointable models.UsersEvents

	v := r.URL.Query()
	userId := v.Get("userId")
	intUserid, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println(err)
		return
	}

	database.DB.Where("UsersID = ?", intUserid).Find(&jointable)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jointable)
}
