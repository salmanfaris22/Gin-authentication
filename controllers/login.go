package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/utils"
)

func LogineUser(ctx *gin.Context) {
	var user models.User

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	errUser, us := userCheck(user)
	if !errUser {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "this Mail catn find"})
		return
	}

	err = passCheck(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GeneretJWT(us.ID)
	// ctx.SetCookie("auth", token,70,)   set
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("token", token, 3600, "*/*", "localhost", false, true)
	err, TokenUser := TokenSet(us.ID, token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"messsage":  "welcome ",
		"logine":    true,
		"user_info": TokenUser,
		"token":     token,
		// "d":         UserDemo,
		// "allUser":   UserDemo,
	})
}
