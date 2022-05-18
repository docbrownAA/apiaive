package controller

import (
	"apiaive/api/model"
	"errors"
	"fmt"
	"time"
)

func GetAppointments() []model.Appointment {
	db := GetDb()
	defer db.Close()

	var appointments []model.Appointment
	db.Find(&appointments)

	return appointments
}

// Get appointments for a specific center
//
// Need to be secured
func GetAppointmentsByCenterId(vCId string) []model.Appointment {
	db := GetDb()
	defer db.Close()

	var appointments []model.Appointment
	db.Where("vc_id = ?", vCId).Find(&appointments)

	return appointments
}

func PostAppointment(jsonAppointment *model.Appointment) (model.Appointment, error) {
	db := GetDb()
	defer db.Close()

	// If fields are filled
	if jsonAppointment.Name != "" && jsonAppointment.Email != "" && jsonAppointment.LastName != "" && !jsonAppointment.Date.IsZero() && jsonAppointment.Vcid != 0 {
		db.Create(&jsonAppointment)
		// Create a token that will be send by email
		// may be a cron will be usefull to purge token
		token := GenerateToken(jsonAppointment.Id)
		db.Create(&token)
		fmt.Println(token)
		// Send email with the token
		return *jsonAppointment, nil
	} else {
		return *jsonAppointment, errors.New("fields are empty")
	}
}

func GenerateToken(appId uint) model.Token {
	token := model.New(appId)

	return token
}

// Return appointments took for a vaccine center
// between the date in params and the date plus 5 days
//
// This away it allows front-end to handle the calendar
func GetAppointmentsAvaibles(vcid string, date time.Time) []model.Appointment {
	db := GetDb()
	db.LogMode(true)
	defer db.Close()
	var appointments []model.Appointment
	fmt.Println(date)
	fmt.Println(date.AddDate(0, 0, 5))
	db.Where("date BETWEEN ? AND ? AND vcid = ? AND validated = true", date, date.AddDate(0, 0, 5), vcid).Select([]string{"date", "vcid", "validated"}).Find(&appointments)

	return appointments
}
