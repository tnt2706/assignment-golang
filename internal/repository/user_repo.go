package repository

import "assignment/internal/model"

type UserRepo interface {
	FindById(id string) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
}
