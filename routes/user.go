package routes

import (
	constant "expense-monster-BE/constants"
	"fmt"
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
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Login successful",
	})
}
