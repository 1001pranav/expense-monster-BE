package routes

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
	var loginData constant.LoginAPIData

	if err := c.BindJSON(&loginData); err != nil {
		fmt.Println("*** Error On Login ***", err.Error())

		if loginData.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Required 'email'"})
			return
		}

		if loginData.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Required 'password'"})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	fmt.Printf("Email_id - %s \n password - %s ", loginData.Email, loginData.Password)

func Register(c *gin.Context) {
	var registerData *constant.LoginAPIData

	//Check JSON is valid or not.
	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status": "INVALID_REQUEST",
				"error":  err.Error(),
			},
		)
		return
	}

	// Checking if email and password are present
	if registerData.Email == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "INVALID_EMAIL",
				"message": "Required 'email'",
			},
		)
		return
	}
	if registerData.Password == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "INVALID_PASSWORD",
				"message": "Required 'password'",
			},
		)
		return
	}

	// Check if user with email already registered, Throw error if registered
	_, userError := d.GetUsersInfoByUsers(registerData.Email)
	if !userError {
		c.JSON(
			http.StatusNotAcceptable,
			gin.H{
				"status":  "INVALID_EMAIL",
				"message": "Email already exists",
			},
		)
		return
	}

	// Hashing Password while storing in database.
	hashedPassword, err := helper.HashPassword(registerData.Password)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "SOMETHING_WENT_WRONG",
				"message": err,
			})
		return
	}

	registerData.Password = hashedPassword

	errData := d.CreateUsers(*registerData)

	if errData != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "INTERNAL_SERVER_ERROR",
				"message": errData,
			},
		)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": registerData,
	})
}
