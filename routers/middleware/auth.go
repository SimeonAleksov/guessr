package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
    "guessr.net/pkg/jwt"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwt.ValidateToken(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
