package models

type User struct {
	FirstName string   `json:"firstName" validate:"required"`
	LastName  string   `json:"lastName" validate:"required"`
	Email     string   `json:"email" validate:"required"`
	Salt      int32    `json."salt" validate:"required"`
	Hash      string   `json."hash" validate:"required"`
	Roles     []string `json."roles" validate:"required"`
	Status    string   `json."status" validate:"required"`
}
