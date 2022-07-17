package models


type User struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName string `json:"lastName" validate:"required"`
	Email string `json:"email" validate:"required"`
	Salt string `json."salt" validate:"required`
	Hash int32 `json."hash" validate:"required`
	Roles []string `json."roles" validate:"required`
	Status string `json."status" validate:"required`
}