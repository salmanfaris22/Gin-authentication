package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := ctx.Cookie("token")

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
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
