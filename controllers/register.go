package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
)

func Register(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	mu.Lock()
	lastId++
	user.ID = uint(lastId)
	mu.Unlock()

	UserDemo = append(UserDemo, user)
	ctx.JSON(200, gin.H{
		"messsage":  "user register done !",
		"user_info": user,
		"allUser":   UserDemo,
	})
}
