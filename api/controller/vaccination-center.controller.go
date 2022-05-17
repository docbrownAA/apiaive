package controller

import "apiaive/api/model"

func GetVaccinationCenters() []model.VaccinationCenter {
	db := GetDb()
	defer db.Close()

	var vCenters []model.VaccinationCenter
	db.Find(&vCenters)

	return vCenters
}
