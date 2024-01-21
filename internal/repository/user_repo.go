package repository

import (
	"assignment/internal/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo interface {
	FindById(string) (*model.User, error)
	FindUserByEmail(string) (*model.User, error)
	CreateUser(*model.User) (*model.User, error)
}

type userRepoImpl struct {
	DB *mongo.Collection
}

func NewUserRepo(DB *mongo.Collection) UserRepo {
	return &userRepoImpl{DB: DB}
}

func (u *userRepoImpl) FindById(id string) (*model.User, error) {
	var user *model.User
	userObjId, _ := primitive.ObjectIDFromHex(id)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := u.DB.FindOne(ctx, bson.M{"_id": userObjId}).Decode(&user)
	return user, err
}

func (u *userRepoImpl) FindUserByEmail(email string) (*model.User, error) {
	var user *model.User

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := u.DB.FindOne(ctx, bson.M{"usr_email": email}).Decode(&user)
	return user, err
}

func (u *userRepoImpl) CreateUser(user *model.User) (*model.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	user.HashPassword(user.Password)
	result, err := u.DB.InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)
	user.ID = oid.Hex()

	return user, err
}
