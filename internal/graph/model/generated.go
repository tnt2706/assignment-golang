// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Response struct {
	IsSuccess *bool   `json:"isSuccess,omitempty"`
	Message   *string `json:"message,omitempty"`
}

type User struct {
	ID        string  `json:"id"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Email     *string `json:"email,omitempty"`
	Age       *int    `json:"age,omitempty"`
	Gender    *string `json:"gender,omitempty"`
}

type UserInput struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Email     *string `json:"email,omitempty"`
	Age       *int    `json:"age,omitempty"`
	Gender    *string `json:"gender,omitempty"`
}

type UserResponse struct {
	IsSuccess bool    `json:"isSuccess"`
	Message   *string `json:"message,omitempty"`
	User      *User   `json:"user,omitempty"`
}
