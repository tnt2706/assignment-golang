package user

import (
	"assignment/src/configs"
	"assignment/src/graph/model"
	"assignment/src/graph/utils"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

var User *mongo.Collection = configs.GetCollection(configs.DB, "users")

func CreateUser(ctx context.Context, input model.CreateUserInput) (*model.UserResponse, error) {
	hash, errHash := utils.HashPassword(*input.Hash)
	if errHash != nil {
		log.Fatal("Hash password error")
	}

	newUser := &model.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Status:    input.Status,
		Hash:      &hash,
	}

	_, err := User.InsertOne(ctx, newUser)
	if err != nil {
		return &model.UserResponse{IsSuccess: false}, err
	}

	return &model.UserResponse{IsSuccess: true}, nil
}
