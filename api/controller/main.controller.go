package controller

import (
	"apiaive/api/model"
	"time"

	"github.com/jinzhu/gorm"
)

func InitDB() {
	db, err := gorm.Open("sqlite3", "./data.db")
	db.LogMode(true)

	if !db.HasTable(&model.VaccinationCenter{}) {
		db.CreateTable(&model.VaccinationCenter{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable((&model.VaccinationCenter{}))
		var vCenter = model.VaccinationCenter{Name: "Portet"}
		db.Create(&vCenter)

		var vCenter2 = model.VaccinationCenter{Name: "Toulouse"}
		db.Create(&vCenter2)
		var vCenter3 = model.VaccinationCenter{Name: "Lyon"}
		db.Create(&vCenter3)

	}

	if !db.HasTable(&model.Appointment{}) {
		db.CreateTable(&model.Appointment{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.Appointment{})
		var app = model.Appointment{}
		var vaccinationCenters []model.VaccinationCenter
		db.Find(&vaccinationCenters)
		for _, center := range vaccinationCenters {
			app = model.Appointment{Email: "gduvinage@gmail.com", Name: "Gaël", LastName: "Duvinage", VcId: int(center.Id), Date: time.Now(), Validated: true}
			db.Create(&app)
		}
	}
	if !db.HasTable(&model.Token{}) {
		db.CreateTable(&model.Token{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable((&model.Token{}))
	}

	if !db.HasTable(&model.Admin{}) {
		db.CreateTable(&model.Admin{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable((&model.Admin{}))
	}

	if err != nil {
		panic(err)
	}
}

func GetDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./data.db")

	if err != nil {
		panic(err)
	}

	return db
}
