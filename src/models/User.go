package models

type User struct {
	Id string `json:"_id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Salt string `json."salt"`
	Hash string `json."hash"`
	Roles [string] `json."roles"`
	Status string `json."status"`
}