package handlers

import (
	"encoding/json"
	"net/http"

	"eduscale/internal/models"
	"eduscale/internal/services"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	err := services.RegisterUser(user)

	if err != nil {

		http.Error(w, err.Error(), 500)

		return
	}

	w.Write([]byte("User Registered"))

}

func Login(w http.ResponseWriter, r *http.Request) {

	var request models.User

	json.NewDecoder(r.Body).Decode(&request)

	user, err := services.LoginUser(
		request.Email,
		request.Password,
	)

	if err != nil {

		http.Error(w, err.Error(), 401)

		return
	}

	json.NewEncoder(w).Encode(user)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := services.GetAllUsers()

	if err != nil {

		http.Error(w, err.Error(), 500)

		return
	}

	json.NewEncoder(w).Encode(users)

}