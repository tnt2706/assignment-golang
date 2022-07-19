package routers

import (
	"assignment/src/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.GET("/user", controllers.GetUser())
	router.GET("/users", controllers.GetUsers())
	router.POST("/user", controllers.CreateUser())
	router.PUT("/user/:userId", controllers.UpdateUser())
	router.DELETE("/user/:userId", controllers.DeleteUser())
}
