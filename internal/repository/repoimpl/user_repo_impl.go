package repoimpl

import (
	"assignment/internal/graph/model"
	repo "assignment/internal/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewUserService(collection *mongo.Collection, ctx context.Context) repo.UserRepo {
	return &UserServiceImpl{
		collection: collection,
		ctx:        ctx,
	}
}

func (u *UserServiceImpl) FindById(id string) (*model.User, error) {
	var user model.User
	userObjId, _ := primitive.ObjectIDFromHex(id)
	err := u.collection.FindOne(u.ctx, bson.M{"_id": userObjId}).Decode(&user)
	return &user, err
}

func (u *UserServiceImpl) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.collection.FindOne(u.ctx, bson.M{"email": email}).Decode(&user)
	return &user, err
}
