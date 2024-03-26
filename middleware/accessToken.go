package middleware

import (
	"net/http"

	"expense-monster-BE/constants"
	"expense-monster-BE/helper"

	"github.com/gin-gonic/gin"
)

func ValidateAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := constants.Response{}
		token := c.GetHeader("Authorization")
		if token == "" {
			response.Status = constants.STATUS_MISSING_AUTH
			response.Error = constants.MESSAGE_MISSING_AUTH

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		isVerify := helper.VerifyJWToken(token)

		if isVerify == nil {
			response.Status = constants.STATUS_FAILED_AUTH
			response.Error = constants.MESSAGE_FAILED_AUTH

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Next()
	}
}
