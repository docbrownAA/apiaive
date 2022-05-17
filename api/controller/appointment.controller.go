package controller

import (
	"apiaive/api/model"
)

func GetAppointments() []model.Appointment {
	db := GetDb()
	defer db.Close()

	var appointments []model.Appointment
	db.Find(&appointments)

	return appointments
}

func GetAppointmentsByCenterId(vCId string) []model.Appointment {
	db := GetDb()
	defer db.Close()

	var appointments []model.Appointment
	db.Where("vc_id = ?", vCId).Find(&appointments)

	return appointments
}
