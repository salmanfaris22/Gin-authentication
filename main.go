package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/controllers"
	"main.go/middlewares"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},

		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * 3600,
		AllowCredentials: true,
	}))
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.LogineUser)

	protuct := router.Group("/")
	protuct.Use(middlewares.JWTAuthMiddleware())
	{

		protuct.POST("/getUser", controllers.GetUser)
		protuct.GET("/logout", controllers.LogOut)

	}
	router.Run()
}
