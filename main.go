package main

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
	"main.go/middlewares"
)

func main() {
	router := gin.Default()
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.LogineUser)

	protuct := router.Group("/")
	protuct.Use(middlewares.JWTAuthMiddleware())
	{
		protuct.GET("/getUser", controllers.GetUser)
		protuct.GET("/logout", controllers.LogOut)

	}
	router.Run()
}
