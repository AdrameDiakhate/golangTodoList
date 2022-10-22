package middleware

import (
	"strings"

	"github.com/AdrameDiakhate/golangTodoList.git/utils"
	"github.com/gin-gonic/gin"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenAuth := context.GetHeader("Authorization")

		if tokenAuth == "" {
			context.JSON(400, gin.H{"error": "No token provided"})
			context.Abort()
			return
		}

		tokenString := strings.Split(tokenAuth, "Bearer ")[1]

		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		Email, Username, Role, err := utils.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Set("Email", Email)
		context.Set("Usernamme", Username)
		context.Set("Role", Role)
		context.Next()
	}
}
