package models

type User struct {
	Id string `json:"_id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json."password`
}