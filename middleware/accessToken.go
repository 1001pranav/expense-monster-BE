package middleware

import (
	"net/http"

	"expense-monster-BE/helper"

	"github.com/gin-gonic/gin"
)

func ValidateAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "UNAUTHORIZED",
			})
			c.Abort()
			return
		}
		isVerify := helper.VerifyJWToken(token)

		if isVerify == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "INVALID_REQUEST",
				"error":  isVerify,
			})
		}
		c.Next()
	}
}
