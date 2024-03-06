package handler

import (
	constant "expense-monster-BE/constants"
	d "expense-monster-BE/database"
	"expense-monster-BE/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  constant.LoginResponse
// @Failure      400  {object}  {"status": "INVALID_REQUEST"}
// @Failure      400  {object}  http.StatusBadRequest
// @Failure      400  {object}  http.StatusBadRequest
// @Failure      400  {object}  http.StatusBadRequest
// @Failure      400  {object}  http.StatusBadRequest
// @Failure      500  {object}  http.StatusInternalServerError
// @Router       /user/login [post]
func Login(c *gin.Context) {

	var loginData *constant.LoginAPIData
	var response = constant.LoginResponse{}
	loginResponse := constant.LoginResponseData{}

	//Checking If JSON is valid or not
	if err := c.ShouldBindJSON(&loginData); err != nil {
		response.Status = constant.INVALID_REQUEST_STATUS
		response.Error = err.Error()
		c.JSON(
			http.StatusBadRequest,
			response,
		)
		return
	}

	if loginData.Email == "" {
		response.Status = constant.REQUIRED_EMAIL_STATUS
		response.Error = constant.REQUIRED_EMAIL_MESSAGE
		c.JSON(
			http.StatusBadRequest,
			response,
		)
		return
	}

	if loginData.Password == "" {
		response.Status = constant.REQUIRED_PASSWORD_STATUS
		response.Error = constant.REQUIRED_PASSWORD_MESSAGE
		c.JSON(
			http.StatusBadRequest,
			response,
		)
		return
	}

	loginDBData, isError := d.GetUsersInfoByUsers(loginData.Email)
	if isError {
		response.Status = constant.INVALID_EMAIL_PASSWORD_STATUS
		response.Error = constant.INVALID_EMAIL_PASSWORD_MESSAGE
		c.JSON(
			http.StatusBadRequest,
			response,
		)
		return
	}

	hashedPassword := helper.CheckPasswordHash(loginData.Password, loginDBData.Password)
	if !hashedPassword {
		response.Status = constant.INVALID_EMAIL_PASSWORD_STATUS
		response.Error = constant.INVALID_EMAIL_PASSWORD_MESSAGE
		c.JSON(
			http.StatusBadRequest,
			response,
		)
		return
	}

	hashToken, err := helper.SignJWT(loginDBData.UserID)

	if err != nil {
		response.Status = constant.INTERNAL_SERVER_STATUS
		response.Error = err.Error()
		c.JSON(
			http.StatusInternalServerError,
			response,
		)
		return
	}

	loginDBData.AccessToken = hashToken
	d.UpdateUsers(loginDBData)
	response.Status = constant.SUCCESS_STATUS
	response.Data = &loginResponse
	loginResponse.AccessToken = loginDBData.AccessToken
	loginResponse.Email = loginDBData.Email
	loginResponse.UserID = loginDBData.UserID

	fmt.Printf("Email_id - %s \n password - %s \n accessToken - %s ", loginData.Email, loginData.Password, loginDBData.AccessToken)

	c.JSON(
		http.StatusAccepted,
		response,
	)
}
