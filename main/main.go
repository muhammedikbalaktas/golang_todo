package main

import (
	"todo/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/create_user", controllers.CreateUser)
	router.Run("localhost:8080")
}
