package models

type User struct {
	FirstName string   `json:"firstName,omitempty" validate:"required"`
	LastName  string   `json:"lastName,omitempty" validate:"required"`
	Email     string   `json:"email,omitempty" validate:"required"`
	Salt      int32    `json."salt,omitempty" validate:"required"`
	Hash      string   `json."hash,omitempty" validate:"required"`
	Roles     []string `json."roles,omitempty" validate:"required"`
	Status    string   `json."status,omitempty" validate:"required"`
}
