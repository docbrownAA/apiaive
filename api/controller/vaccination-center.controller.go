package controller

import "gduvinage/api/model"

func GetVaccinationCenters() []model.VaccinationCenter {
	db := InitDB()
	defer db.Close()

	var vCenters []model.VaccinationCenter
	db.Find(&vCenters)

	return vCenters
}
