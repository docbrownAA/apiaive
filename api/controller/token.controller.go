package controller

import (
	"apiaive/api/model"
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Check if token is still valid
//
//Then valid the appointment related to and delete the token
//If not, return an error and delete the token
func ControlToken(generatedToken string) (bool, error) {
	db := GetDb()
	defer db.Close()

	var token model.Token

	err := db.Where("generated_token = ?", generatedToken).Find(&token).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("Token not found")
		return false, errors.New("token not found")
	} else {
		if time.Now().Before(token.ValidDate) {
			var appointment model.Appointment
			errApp := db.Where("Id = ?", token.Appid).Find(&appointment).Error
			if errors.Is(errApp, gorm.ErrRecordNotFound) {
				fmt.Println("Appointment not found")
				return false, errors.New("appointment not found")
			} else {
				appointment.Validated = true
				db.Save(&appointment)
				db.Delete(&token)
				return true, nil
			}
		} else {
			db.Delete(&token)
			return false, errors.New("token not valid any more")
		}

	}
}
