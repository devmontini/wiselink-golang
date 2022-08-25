package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devmontini/wiselink-golang/database"
	"github.com/devmontini/wiselink-golang/models"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the SIGN IN USER!")
	var user models.Users

	//lo que me envian se guarda en user con modelo Users
	json.NewDecoder(r.Body).Decode(&user)

	userData := database.DB.Create(&user)
	err := userData.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		message := map[string]interface{}{
			"message": "User created successfully",
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(message)
	}
}

func LogInHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the LOG IN USER!")
	var user models.Users

	json.NewDecoder(r.Body).Decode(&user)

	userData := database.DB.Where("username = ?", user.Username).First(&user)
	err := userData.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
