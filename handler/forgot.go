package handler

import (
	"crypto/rand"
	"expense-monster-BE/constants"
	"expense-monster-BE/database"
	"expense-monster-BE/helper"
	"strconv"
	"strings"
	"time"

	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ForgotPassword(c *gin.Context) {
	var requestData *constants.APIRequestForgotPassword
	var response = constants.ForgotPasswordResponse{
		Response: constants.Response{
			Status: "",
			Error:  "",
		},
		Data: &constants.ForgotPasswordResponseData{},
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		response.Status = constants.INVALID_REQUEST_STATUS
		response.Error = err.Error()
		c.JSON(
			http.StatusBadRequest,
			response,
		)
		return
	}

	if requestData.Email == "" {
		response.Status = constants.REQUIRED_EMAIL_STATUS
		response.Error = constants.REQUIRED_EMAIL_MESSAGE
		c.JSON(
			http.StatusBadRequest,
			response,
		)
		return
	}
	userData, isError := database.GetUsersInfoByUsers(requestData.Email)

	if isError {
		response.Error = constants.USER_NOT_EXISTS_MESSAGE
		response.Status = constants.USER_NOT_EXISTS_STATUS
		c.JSON(http.StatusBadRequest, response)
		return
	}

	randomOTP, randomOTPErr := rand.Int(rand.Reader, big.NewInt(900000))
	if randomOTPErr != nil {
		response.Status = constants.INTERNAL_SERVER_STATUS
		response.Error = randomOTPErr.Error()
		c.JSON(
			http.StatusInternalServerError,
			response,
		)
		return
	}

	randomOTP = randomOTP.Add(randomOTP, big.NewInt(100000))
	mailMessage := constants.REPLACE_STRINGS(constants.FORGOT_MAIL, strings.Split(userData.Email, "@")[0], "[USER]")
	mailMessage = constants.REPLACE_STRINGS(mailMessage, strconv.FormatInt(randomOTP.Int64(), 10), "[OTP]")
	mailMessage = constants.REPLACE_STRINGS(mailMessage, strconv.Itoa(constants.MAX_ATTEMPT_FORGOT), "[ATTEMPT]")

	userData.OTP = uint(randomOTP.Int64())
	userData.OTPGeneratedOn = time.Now().Local().Add(time.Minute * time.Duration(constants.MAX_TIME_FORGOT_MINS))

	database.UpdateUsers(userData)
	// const MinInStrings = strconv.Itoa(constants.MAX_TIME_FORGOT_MIN)
	// mailMessage = constants.REPLACE_STRINGS()
	helper.SendMailSMTP(userData.Email, []byte(mailMessage))
	response.Data.OTP = userData.OTP
	response.Data.OTPExpiresOn = userData.OTPGeneratedOn

	c.JSON(http.StatusOK, response)
}
