package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"assignment/src/configs"
	"assignment/src/models"
	"assignment/src/responses"
	res "assignment/src/responses"
	"assignment/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Sign up")
	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Sign up")
	}
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var user models.User

		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{IsSuccess: false, Message: err.Error()})
			return
		}

		if validateErr := validate.Struct(&user); validateErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{IsSuccess: false, Message: validateErr.Error()})
			return
		}

		hash, errHash := utils.HashPassword(user.Hash)
		if errHash != nil {
			log.Fatal("Hash password error")
		}

		newUser := models.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Hash:      hash,
			Salt:      14,
			Roles:     user.Roles,
			Status:    user.Status,
		}

		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, res.UserResponse{IsSuccess: false, Message: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, res.UserResponse{IsSuccess: true, Data: map[string]interface{}{"data": result}})

	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var user models.User
		defer cancel()

		userId, _ := primitive.ObjectIDFromHex(c.Param("userId"))

		err := userCollection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)

		fmt.Println(user)
		if err != nil {
			c.JSON(http.StatusOK, res.UserResponse{IsSuccess: false, Message: "User not found"})
		}

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusInternalServerError, res.UserResponse{IsSuccess: false, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, res.UserResponse{IsSuccess: true})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		userId := c.Param("userId")
		defer cancel()

		objUserId, _ := primitive.ObjectIDFromHex(userId)

		_, err := userCollection.DeleteOne(ctx, bson.M{"_id": objUserId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, res.UserResponse{IsSuccess: false, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, res.UserResponse{IsSuccess: true})
	}
}
