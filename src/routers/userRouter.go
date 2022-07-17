package routers

import (
	"assignment/src/controllers"

	"github.com/gin-gonic/gin"
)


func UserRouter(router *gin.Engine){
	router.GET("/user" , controllers.GetUser())
	router.GET("/users" , controllers.GetUsers())
	router.POST("/sign-up" , controllers.CreateUser())
	router.PUT("/update-profile" , controllers.UpdateUser())
	router.DELETE("/delete-user" , controllers.DeleteUser())
}