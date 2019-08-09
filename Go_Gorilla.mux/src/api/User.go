package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sheyguys/backendgin/config"
	"github.com/sheyguys/backendgin/src/models"
	"github.com/sheyguys/backendgin/src/repository"
)

func AddUser(w http.ResponseWriter, req *http.Request) {
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	userRepository := repository.NewUserRepository(db, "User")

	//get variable by path
	params := mux.Vars(req)
	var firstName, lastName, email string
	firstName = string(params["firstName"])
	lastName = string(params["lastName"])
	email = string(params["email"])
	// gender = string(params["gender"])
	// birth =string(params["birth"])

	var p models.User
	p.Firstname = firstName
	p.Lastname = lastName
	p.Email = email
	// p.Gender = gender
	// p.DateofBirth = birth
	userRepository.Save(&p)

}

func GetUserByEmail(w http.ResponseWriter, req *http.Request) {
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	userRepository := repository.NewUserRepository(db, "User")

	//get variable by path
	params := mux.Vars(req)
	var Email string
	Email = string(params["Email"])

	user, err2 := userRepository.FindByEmail(Email)
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(user)

}

//Default data Major
