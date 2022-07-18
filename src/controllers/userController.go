package controllers

import (
  "context"
  "time"
	"net/http"
  "log"

  "assignment/src/configs"
  "assignment/src/models"
  "assignment/src/responses"
  "assignment/src/utils"

  "github.com/gin-gonic/gin"
  "github.com/go-playground/validator/v10"
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

    defer cancel();

    if err := c.BindJSON(&user); err != nil {
      c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
      return
    }

    if validateErr := validate.Struct(&user); validateErr != nil {
      c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validateErr.Error()}})
      return
    }

    hash, errHash := utils.HashPassword(user.Hash)
    if errHash != nil{
      log.Fatal("Hash password error")
    }

    newUser := models.User{
      FirstName: user.FirstName,
      LastName: user.LastName,
      Email: user.Email,
      Hash :hash,
      Salt : 14,
      Roles: user.Roles,
      Status: user.Status,
    }

    result, err := userCollection.InsertOne(ctx, newUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
        return
    }

    c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})

  }
}

func UpdateUser() gin.HandlerFunc {
  return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

    defer cancel();

  }
}

func DeleteUser() gin.HandlerFunc {
  return func(c *gin.Context) {
		_,cancel := context.WithTimeout(context.Background() , 10*time.Second)

    defer cancel()

    userId := c.Param("userId")
    log.Fatal(userId)
  }
}

