package main

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
)

func main() {
	router := gin.Default()
	router.POST("/register", controllers.Register)
	router.Run()
}
