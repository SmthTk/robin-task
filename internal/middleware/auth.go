package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"robin-task/internal/auth"
	"robin-task/utils"
	"strings"
)

const BearerHead = "Bearer "

func AuthMiddleware(jwtService *auth.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, utils.ResponseData(false, "No token provided", nil))
			c.Abort()
			return
		}

		if !strings.HasPrefix(token, BearerHead) {
			c.JSON(http.StatusUnauthorized, utils.ResponseData(false, "Invalid token format", nil))
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, BearerHead)
		userID, role, err := jwtService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.ResponseData(false, "Invalid token", nil))
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Set("role", role)
		c.Next()
	}
}
