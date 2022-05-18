package controller

import (
	"apiaive/api/model"
)

func GetUsers() []model.Appointment {
	db := GetDb()
	defer db.Close()
	var users []model.Appointment
	// SELECT * FROM users
	db.Find(&users)
	// Affichage des données
	return users
}
