package models

type User struct {
	Id        uint   `json:id`
	Firstname string `json:"first_name"`
	Lastname  string `json:last_name`
	Email     string `json:email`
	Password  string `json:"-"`
}
