package controller

import (
	"apiaive/api/model"
)

func RegisterUser(user *model.User) error {
	db := GetDb()
	defer db.Close()
	err := user.HashPassword(user.Password)
	if err != nil {
		return err
	}

	record := db.Create(&user).Error
	if record != nil {
		return record
	}

	return nil
}
