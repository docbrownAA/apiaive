package controller

import (
	"apiaive/api/auth"
	"apiaive/api/model"
	"errors"

	"github.com/jinzhu/gorm"
)

func GeneratedToken(token *model.TokenRequest) (string, error) {
	var user model.User
	db := GetDb()
	defer db.Close()
	err := db.Where("email = ?", token.Email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}
	credentialError := user.CheckPassword(token.Password)
	if credentialError != nil {
		return "", errors.New("invalid credential")
	}
	tokenString, err := auth.GenerateJWT(user.Email, user.Password)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
