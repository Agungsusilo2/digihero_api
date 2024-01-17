package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		providedKey := c.GetHeader("X-API-KEY")

		if providedKey != apiKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key", "code": http.StatusUnauthorized})
			c.Abort()
			return
		}

		c.Next()
	}
}
