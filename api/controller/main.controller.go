package controller

import (
	"apiaive/api/model"
	"time"

	"github.com/jinzhu/gorm"
)

//Init the database with fake data
func InitDB() {
	db, err := gorm.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)

	if !db.HasTable(&model.VaccinationCenter{}) {
		db.CreateTable(&model.VaccinationCenter{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable((&model.VaccinationCenter{}))
		var vCenter = model.VaccinationCenter{Name: "Portet", Slots: 2}
		db.Create(&vCenter)

		var vCenter2 = model.VaccinationCenter{Name: "Toulouse", Slots: 4}
		db.Create(&vCenter2)
		var vCenter3 = model.VaccinationCenter{Name: "Lyon", Slots: 4}
		db.Create(&vCenter3)

		var vCenter4 = model.VaccinationCenter{Name: "Paris", Slots: 10}
		db.Create(&vCenter4)

	}

	if !db.HasTable(&model.Appointment{}) {
		db.CreateTable(&model.Appointment{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.Appointment{})
		var app = model.Appointment{}
		var vaccinationCenters []model.VaccinationCenter
		db.Find(&vaccinationCenters)
		for _, center := range vaccinationCenters {
			appDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8+center.Id, 15, 0, 0, time.Local)
			app = model.Appointment{Email: "gduvinage@gmail.com", Name: "Gaël", LastName: "Duvinage", Vcid: int(center.Id), Date: appDate, Validated: true}
			db.Create(&app)
		}
	}
	if !db.HasTable(&model.TokenAppointment{}) {
		db.CreateTable(&model.TokenAppointment{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable((&model.TokenAppointment{}))
	}

	if !db.HasTable(&model.User{}) {
		db.CreateTable(&model.User{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable((&model.User{}))
	}

}

func GetDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./data.db")

	if err != nil {
		panic(err)
	}

	return db
}
