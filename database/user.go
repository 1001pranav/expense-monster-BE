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

func CreateUsers(userData c.LoginAPIData) error {
	db := h.Connection()

	user := c.ModelUsers{
		Email:    userData.Email,
		Password: userData.Password,
		Status:   c.STATUS_ACTIVE,
	}
	result := db.Create(&user)

	return result.Error
}

func UpdateUsers(userData c.ModelUsers) error {
	db := h.Connection()
	db.Save(&userData)
	return nil
}
