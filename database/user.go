package sqlFunction

import (
	c "expense-monster-BE/constants"
	h "expense-monster-BE/helper"
	"log"

	"gorm.io/gorm"
)

func GetUsersInfoByUsers(emailID string) ([]c.ModelUsers, *gorm.DB) {
	db := h.Connection()

	var findByUserEmail []c.ModelUsers
	if err := db.Find(&findByUserEmail); err != nil {
		log.Fatalln("Error Pulling Users from database", err)
		return nil, err
	}

	return findByUserEmail, nil
}
