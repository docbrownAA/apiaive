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
//
func GetAppointmentsByCenterId(username string) ([]model.Appointment, error) {
	db := GetDb()
	defer db.Close()
	db.LogMode(true)
	var user model.User
	startDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8, 0, 0, 0, time.Local)
	endDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 18, 0, 0, 0, time.Local)
	err := db.Where("email = ? OR user_name = ? ", username, username, startDate, endDate).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	fmt.Println(user)
	var appointments []model.Appointment
	db.Where("vcid = ? AND date BETWEEN ? AND ? AND validated = 1", user.Vcid, startDate, endDate).Find(&appointments)

	return appointments, nil
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
		var checkAppointment []model.Appointment
		var vCenter model.VaccinationCenter
		db.Where("vcid = ? AND date = ? AND ", appointment.Vcid, appointment.Date).Find(&checkAppointment)
		db.Where("id = ? ", appointment.Vcid).First(&vCenter)

		if len(checkAppointment) < vCenter.Slots {

			db.Create(&appointment)

			//db.Commit()
			// Create a token that will be send by email
			// may be a cron will be usefull to purge token
			token := GenerateAppointmentToken(appointment.Id)
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

func GenerateAppointmentToken(appId int) model.TokenAppointment {
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
	fmt.Println("appointment.controller", date)
	fmt.Println(date.AddDate(0, 0, 5))
	var startDate = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
	db.Where("date BETWEEN ? AND ? AND vcid = ? AND validated = true", startDate, startDate.AddDate(0, 0, 5), vcid).Select([]string{"date", "vcid", "validated"}).Find(&appointments)

	return appointments
}

// Check if token is still valid
//
//Then valid the appointment related to and delete the token
//If not, return an error and delete the token
func ControlToken(generatedToken string) (bool, error) {
	db := GetDb()
	defer db.Close()
	db.LogMode(true)
	var token model.TokenAppointment

	err := db.Where("generated_token = ?", generatedToken).Find(&token).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("Token not found")
		return false, errors.New("token not found")
	} else {
		var appointment model.Appointment
		errApp := db.Where("id = ?", token.Appid).Find(&appointment).Error
		if time.Now().Before(token.ValidDate) {
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
			db.Delete(&appointment)
			db.Delete(&token)
			return false, errors.New("token not valid any more")
		}

	}
}
