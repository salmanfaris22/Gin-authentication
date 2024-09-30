package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pusher/pusher-http-go/v5"
)

func ChatApp(ctx *gin.Context) {
	var data map[string]string
	pusherClient := pusher.Client{
		AppID:   "1872587",
		Key:     "4c95e51a448a7ba3b55e",
		Secret:  "624a2493c42a299a9aa3",
		Cluster: "mt1",
		Secure:  true,
	}
	err := ctx.BindJSON(&data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = pusherClient.Trigger("channel", "message", data)
	if err != nil {
		fmt.Println(err.Error())
	}
	ctx.JSON(200, gin.H{
		"message": data,
	})
}
