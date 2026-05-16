package middleware

import (
	"e-commerce-api/internal/constants"
	"e-commerce-api/internal/infrastructure/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	tokenManager auth.TokenManager
}

func NewMiddleware(tokenManager auth.TokenManager) *Middleware {
	return &Middleware{
		tokenManager: tokenManager,
	}
}

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		userID, err := m.tokenManager.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.Set(constants.UserIDKey, userID)
		c.Next()

	}
}
