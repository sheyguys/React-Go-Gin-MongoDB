package models

type User struct {

	// User_ID 	string `json:"_id,omitempty"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	// Gender		string `json:"gender" binding:"required"`
	// DateofBirth	string `json:"dateofbirth" binding:"required"`

}
type Users []User
