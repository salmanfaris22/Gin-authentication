package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/controllers"
	"main.go/middlewares"
)

type Router interface {
	Start()
}

type impel struct {
	gin *gin.Engine
}

func (i *impel) Start() {

	i.gin.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},

		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * 3600,
		AllowCredentials: true,
	}))
	i.gin.POST("/register", controllers.Register)
	i.gin.POST("/login", controllers.LogineUser)

	protuct := i.gin.Group("/")
	protuct.Use(middlewares.JWTAuthMiddleware())
	{
		// protuct.POST("/messages", controllers.ChatApp)
		protuct.POST("/getUser", controllers.GetUser)
		protuct.GET("/logout", controllers.LogOut)

	}
	i.gin.Run()
}

func NewRouter() Router {
	return &impel{
		gin: gin.New(),
	}
}
