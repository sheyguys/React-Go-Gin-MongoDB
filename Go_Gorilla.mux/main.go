package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sheyguys/backendgin/src/api"
)

func main() {

	//default USER
	router := mux.NewRouter()
	//User
	router.HandleFunc("/user/{firstName}/{lastName}/{email}", api.AddUser).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", handlers.CORS(handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}), handlers.AllowedOrigins([]string{"*"}))(router)))

}
