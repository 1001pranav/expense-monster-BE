package database

import (
	c "expense-monster-BE/constants"
	h "expense-monster-BE/helper"
)

func GetUsersInfoByUsers(emailID string) (c.ModelUsers, bool) {
	db := h.Connection()

	var findByUserEmail c.ModelUsers

	if err := db.First(&findByUserEmail, "Email=?", emailID); err != nil {
		return findByUserEmail, false
	}

	return findByUserEmail, true
}

func CreateUsers(userData c.LoginAPIData) (uint, error) {
	db := h.Connection()

	user := c.ModelUsers{
		Email:    userData.Email,
		Password: userData.Password,
		Status:   c.STATUS_ACTIVE,
	}
	result := db.Create(&user)

	return user.UserID, result.Error
}

func UpdateUsers(userData c.ModelUsers) error {
	db := h.Connection()
	db.Save(&userData)
	return nil
}

func GetByAccessToken(accessToken string, userID uint) bool {
	db := h.Connection()
	var userData c.ModelUsers
	if err := db.First(&userData, "AccessToken=? AND UserID=?", accessToken, userID); err != nil {
		return false
	}
	return true
}
