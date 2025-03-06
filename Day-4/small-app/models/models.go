package models

// contains all the struct to model data

type User struct {
	FName    string `json:"f_name"`
	LName    string `json:"l_name"`
	Password string `json:"-"`
	Email    string `json:"email"`
}
