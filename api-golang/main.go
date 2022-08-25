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

	r.HandleFunc("/", HomeRoute)

	r.HandleFunc("/login", routes.LogInHandler).Methods("GET")
	r.HandleFunc("/signup", routes.SignUpHandler).Methods("POST")

	r.HandleFunc("/events/{userId}", routes.CreateEventsHandler).Methods("POST")
	r.HandleFunc("/events/{userId}", routes.GetEventsHandler).Methods("GET")
	r.HandleFunc("/events/{userId}/{eventId}", routes.GetEventHandlerById).Methods("GET")
	r.HandleFunc("/events/{userId}/{eventId}", routes.DeleteEventHandlerById).Methods("DELETE")
	r.HandleFunc("/events/{userId}/{eventId}", routes.EditEventHandlerById).Methods("PUT")

	http.ListenAndServe(":8080", r)
}
