package main

import (
	"fmt"
	"net/http"

	"github.com/devmontini/wiselink-golang/database"
	"github.com/devmontini/wiselink-golang/routes"
	"github.com/gorilla/mux"
)

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func HomeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the API!"))
	fmt.Println("Hit in the HOME PAGE!")
}

func main() {
	r := mux.NewRouter()

	database.Connect()

	r.Use(ContentTypeMiddleware)

	//HOME
	r.HandleFunc("/", HomeRoute)

	//SIGN UP
	r.HandleFunc("/signup", routes.SignUpHandler).Methods("POST")

	//LOG IN
	r.HandleFunc("/login", routes.LogInHandler).Methods("GET")

	//CREAR EVENTOS
	r.HandleFunc("/events", routes.CreateEventsHandler).
		Queries(
			"userId", "{userId}",
		).Methods("POST")

	//BUSCAR TODOS LOS EVENTOS
	r.HandleFunc("/events", routes.GetEventsHandler).
		Queries(
			"userId", "{userId}",
			"title", "{title}",
		).Methods("GET")

	//BUSCAR UN EVENTO
	r.HandleFunc("/events/info/{eventId}", routes.GetEventHandlerById).Methods("GET")

	//EDITAR EL EVENTO
	r.HandleFunc("/events/info/{eventId}", routes.EditEventHandlerById).
		Queries(
			"userId", "{userId}",
		).Methods("PUT")

	//BORRAR EL EVENTO
	r.HandleFunc("/events/info/{eventId}", routes.DeleteEventHandlerById).
		Queries(
			"userId", "{userId}",
		).Methods("DELETE")

	//INSCRIBIRSE AL EVENTO
	r.HandleFunc("/user/{eventId}", routes.InscEventHandlerById).
		Queries(
			"userId", "{userId}",
		).Methods("POST")

	//TRAER LAS INSCRIPCIONES DE USUARIOS
	r.HandleFunc("/user/{userId}", routes.GetUserEvents).Methods("GET")

	http.ListenAndServe(":8080", r)
}
