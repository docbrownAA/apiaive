package controller

import (
	"apiaive/api/model"
	"errors"

	"github.com/jinzhu/gorm"
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

func PostAppointment(jsonAppointment *model.Appointment) (*gorm.DB, error) {
	db := GetDb()
	defer db.Close()

	// Si le champ est bien saisi
	if jsonAppointment.Name != "" && jsonAppointment.Email != "" && jsonAppointment.LastName != "" && !jsonAppointment.Date.IsZero() && jsonAppointment.VcId != 0 {
		// INSERT INTO "user" (name,LastName,email) VALUES (json.name,json.last_name,json.email);

		return db.Create(&jsonAppointment), nil
		// Affichage des données saisies
	} else {
		// Affichage de l'erreur
		return nil, errors.New("fields are empty")
	}
}
