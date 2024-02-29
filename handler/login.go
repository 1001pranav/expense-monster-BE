package handler

import (
	constant "expense-monster-BE/constants"
	d "expense-monster-BE/database"
	"expense-monster-BE/helper"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginData *constant.LoginAPIData
	loginResponse := constant.LoginResponse{}

	//Checking If JSON is valid or not
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status": "INVALID_REQUEST",
				"error":  err.Error(),
			},
		)
		return
	}

	if loginData.Email == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "INVALID_EMAIL",
				"message": "Required 'email'",
			},
		)
		return
	}

	log.Println(loginData)
	if loginData.Password == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status": "INVALID_PASSWORD",
				"error":  "Required 'password'",
			},
		)
		return
	}

	loginDBData, isError := d.GetUsersInfoByUsers(loginData.Email)
	if isError {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status": "INVALID_EMAIL_PASSWORD",
				"error":  "Email or Password is incorrect",
			},
		)
		return
	}

	hashedPassword := helper.CheckPasswordHash(loginData.Password, loginDBData.Password)
	if !hashedPassword {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status": "INVALID_EMAIL_PASSWORD",
				"error":  "Email or Password is incorrect",
			},
		)
		return
	}

	hashToken, err := helper.SignJWT(loginDBData.UserID)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status": "INTERNAL_SERVER_ERROR",
				"error":  err.Error(),
			},
		)
		return
	}

	loginDBData.AccessToken = hashToken
	d.UpdateUsers(loginDBData)

	loginResponse.AccessToken = loginDBData.AccessToken
	loginResponse.Email = loginDBData.Email
	loginResponse.UserID = loginDBData.UserID

	fmt.Printf("Email_id - %s \n password - %s \n accessToken - %s ", loginData.Email, loginData.Password, loginDBData.AccessToken)

	c.JSON(
		http.StatusAccepted,
		gin.H{
			"message": "Login successful",
			"data":    loginResponse,
		},
	)
}
