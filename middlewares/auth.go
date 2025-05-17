package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaustubsingh09/go-auth-gin/utils"
)

func Authentication(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		// abort other requests with AbortWithStatusJSON() method
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token verification failed"})
		return
	}

	context.Set("userId", userId)
	context.Next()

}
