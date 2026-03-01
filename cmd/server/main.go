package main

import (
	"fmt"
	"net/http"

	"eduscale/configs"
	"eduscale/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {

	configs.ConnectDB()

	router := mux.NewRouter()

	router.HandleFunc(
		"/register",
		handlers.Register,
	).Methods("POST")

	router.HandleFunc(
		"/login",
		handlers.Login,
	).Methods("POST")

	router.HandleFunc(
		"/users",
		handlers.GetUsers,
	).Methods("GET")

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", router)

}