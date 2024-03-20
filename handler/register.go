package handler

import (
	constant "expense-monster-BE/constants"
	d "expense-monster-BE/database"
	"expense-monster-BE/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerData *constant.LoginAPIData
	var responseData = constant.RegisterResponseData{}

	var response = constant.RegisterResponse{
		Response: constant.Response{
			Status: constant.SUCCESS_STATUS,
			Error:  "",
		},
		Data: &responseData,
	}

	//Check JSON is valid or not.
	if err := c.ShouldBindJSON(&registerData); err != nil {
		response.Status = constant.INVALID_REQUEST_STATUS
		response.Error = err.Error()
		c.JSON(
			http.StatusBadRequest,
			response,
		)
		return
	}

	// Checking if email and password are present
	if registerData.Email == "" {
		response.Status = constant.REQUIRED_EMAIL_STATUS
		response.Error = constant.REQUIRED_EMAIL_MESSAGE
		c.JSON(
			http.StatusBadRequest,
			response,
		)
		return
	}

	if registerData.Password == "" {
		response.Status = constant.REQUIRED_PASSWORD_STATUS
		response.Error = constant.REQUIRED_PASSWORD_MESSAGE
		c.JSON(
			http.StatusBadRequest,
			response,
		)
		return
	}

	// Check if user with email already registered, Throw error if registered
	userData, userError := d.GetUsersInfoByUsers(registerData.Email)

	//Returns userError false if userError Exists.
	if userData.UserID != 0 || userError {
		response.Status = constant.EMAIL_EXISTS_STATUS
		response.Error = constant.EMAIL_EXISTS_MESSAGE
		c.JSON(
			http.StatusNotAcceptable,
			response,
		)
		return
	}

	// Hashing Password while storing in database.
	hashedPassword, err := helper.HashPassword(registerData.Password)
	if err != nil {
		response.Status = constant.INTERNAL_SERVER_STATUS
		response.Error = err.Error()
		c.JSON(
			http.StatusInternalServerError,
			response,
		)
		return
	}

	registerData.Password = hashedPassword

	userID, errData := d.CreateUsers(*registerData)

	if errData != nil {
		response.Status = constant.INTERNAL_SERVER_STATUS
		response.Error = errData.Error()
		c.JSON(
			http.StatusInternalServerError,
			response,
		)
		return
	}

	responseData.UserID = userID
	responseData.Email = registerData.Email
	c.JSON(http.StatusAccepted, response)
}
