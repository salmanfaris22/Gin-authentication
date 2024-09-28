package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("token")

		fmt.Println(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
			ctx.Abort()
			return
		}

		// _, err = controllers.ValidateToken(token)
		// if err != nil {
		// 	ctx.JSON(http.StatusUnauthorized, err.Error())
		// 	ctx.Abort()
		// 	return
		// }

		ctx.Next()

	}
}
