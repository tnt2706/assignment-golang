package repository

import (
	"assignment/internal/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepo interface {
	CreateToDo(*model.Todo) (*model.Todo, error)
}

type todoRepoImpl struct {
	DB *mongo.Collection
}

func NewTodoRepo(DB *mongo.Collection) UserRepo {
	return &userRepoImpl{DB: DB}
}

func (u *todoRepoImpl) CreateToDo(todo *model.Todo) (*model.Todo, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := u.DB.InsertOne(ctx, todo)

	if err != nil {
		return nil, err
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)
	todo.ID = oid.Hex()

	return todo, err
}
