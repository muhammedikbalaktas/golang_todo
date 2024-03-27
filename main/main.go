package main

import (
	"todo/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/create_user", controllers.CreateUser)
	router.POST("/get_user", controllers.GetUser)
	router.Run("localhost:8080")
}
