package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `json:"_id"`
	Name        string             `json:"usr_name"`
	Password    string             `json:"usr_password"`
	Salt        string             `json:"usr_salt"`
	Email       string             `json:"usr_email"`
	Phone       string             `json:"usr_phone"`
	Sex         string             `json:"usr_sex"`
	Avatar      string             `json:"usr_avatar"`
	DateOfBirth string             `json:"usr_date_of_birth"`
	Roles       []string           `json:"usr_roles"`
	Status      string             `json:"usr_status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
