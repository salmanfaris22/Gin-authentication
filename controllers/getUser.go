package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Token string `json:"token"`
}

func GetUser(ctx *gin.Context) {
	var token TokenRequest
	err := ctx.ShouldBindJSON(&token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": err.Error(),
		})
		return
	}
	Tokenerr, user := useCheck(token.Token)
	if Tokenerr != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"messsage": Tokenerr.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"messsage": user,
	})
}
