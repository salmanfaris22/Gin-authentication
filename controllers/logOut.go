package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogOut(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully!"})
}
