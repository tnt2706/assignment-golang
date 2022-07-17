package main

import (
	"assignment/src/configs"
	"assignment/src/routers"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	configs.ConnectDB()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})

	routers.UserRouter(router)

	router.Run("localhost:8080") 
}