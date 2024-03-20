package handler

import (
	"expense-monster-BE/constants"
	"expense-monster-BE/database"
	"expense-monster-BE/helper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ResetPassword(c *gin.Context) {
	var requestData *constants.APIRequestResetPassword
	var response = constants.ResetPassword{
		Response: constants.Response{
			Status: constants.SUCCESS_STATUS,
			Error:  "",
		},
		Data: constants.Data{},
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		response.Status = constants.INTERNAL_SERVER_STATUS
		response.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	if requestData.PasswordType == 0 {
		response.Status = constants.STATUS_MISSING_PASSWORD_TYPE
		response.Error = ""
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if requestData.Email == "" {
		response.Error = ""
		response.Status = constants.REQUIRED_EMAIL_STATUS
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if requestData.PasswordType == constants.CHANGE_RESET_PASSWORD && requestData.OldPassword == nil {
		response.Status = constants.STATUS_MISSING_OLD_PASSWORD
		response.Error = ""
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else if requestData.PasswordType == constants.FORGOT_RESET_PASSWORD && requestData.OTP == nil {
		response.Status = constants.STATUS_MISSING_OTP
		response.Error = ""
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else if requestData.PasswordType < 1 || requestData.PasswordType > constants.FORGOT_RESET_PASSWORD {
		response.Status = constants.STATUS_INVALID_PASSWORD_TYPE
		response.Error = ""
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userData, isErr := database.GetUsersInfoByUsers(requestData.Email)

	if !isErr {
		response.Error = constants.USER_NOT_EXISTS_STATUS
		response.Status = ""
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if requestData.PasswordType == constants.FORGOT_RESET_PASSWORD {
		var otp uint = *requestData.OTP
		if userData.OTP == 0 {
			response.Status = constants.STATUS_OTP_EXPIRED
			response.Error = ""
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		if userData.OTP != otp {
			response.Status = constants.STATUS_INVALID_OTP
			response.Error = ""
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		if userData.OTPGeneratedOn.After(time.Now()) {
			response.Status = constants.STATUS_OTP_EXPIRED
			response.Error = ""
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		userData.OTP = 0
		userData.OTPGeneratedOn = time.Now()

	} else {
		var oldPassword string = *requestData.OldPassword
		if helper.CheckPasswordHash(oldPassword, userData.Password) {
			response.Status = constants.STATUS_PASSWORD_NOT_MATCH
			response.Error = ""
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}
	}

	password, err := helper.HashPassword(requestData.Password)

	if err != nil {
		response.Status = constants.INTERNAL_SERVER_STATUS
		response.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	userData.Password = password
	database.UpdateUsers(userData)

	c.JSON(http.StatusAccepted, response)
}
