package controller

import (
	"apiaive/api/model"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
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

func CreateAppointment(appointment *model.Appointment) (model.Appointment, uuid.UUID, error) {

	// If fields are filled
	if appointment.Name != "" && appointment.Email != "" && appointment.LastName != "" && !appointment.Date.IsZero() && appointment.Vcid != 0 {
		//Convert date to remove seconds and nanosecond
		// then check if hour choose is correct
		appointment.Date = time.Date(appointment.Date.Year(), appointment.Date.Month(), appointment.Date.Day(), appointment.Date.Hour(), appointment.Date.Minute(), 0, 0, time.Local)
		hourly := appointment.Date

		switch hourly.Minute() {
		case 00:
		case 15:
		case 30:
		case 45:
		default:
			return *appointment, uuid.Nil, errors.New("hourly incorrect (minutes)")
		}

		if hourly.Hour() < 8 || hourly.Hour() > 18 {
			return *appointment, uuid.Nil, errors.New("appointment outside openning")
		}
		db := GetDb()
		defer db.Close()
		db.LogMode(true)

		//Check if the slot is available, if not return an error
		var checkAppointment model.Appointment
		errCheck := db.Where("vcid = ? AND date = ?", appointment.Vcid, appointment.Date).Find(&checkAppointment).Error
		if errors.Is(errCheck, gorm.ErrRecordNotFound) {
			db.Create(&appointment)
			// Create a token that will be send by email
			// may be a cron will be usefull to purge token
			token := GenerateToken(appointment.Id)
			db.Create(&token)
			fmt.Println(token)
			// Send email with the token
			return *appointment, token.GeneratedToken, nil

		} else {
			return *appointment, uuid.Nil, errors.New("slot not available")
		}
	} else {
		return *appointment, uuid.Nil, errors.New("fields are empty")
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
func GetAppointmentAvailables(vcid string, date time.Time) []model.Appointment {
	db := GetDb()
	db.LogMode(true)
	defer db.Close()
	var appointments []model.Appointment
	fmt.Println(date)
	fmt.Println(date.AddDate(0, 0, 5))
	db.Where("date BETWEEN ? AND ? AND vcid = ? AND validated = true", date, date.AddDate(0, 0, 5), vcid).Select([]string{"date", "vcid", "validated"}).Find(&appointments)

	return appointments
}
