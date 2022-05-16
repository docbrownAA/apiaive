package controller

import (
	"gduvinage/api/model"

	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./data.db")
	db.LogMode(true)

	if !db.HasTable(&model.Appointment{}) {
		db.CreateTable(&model.Appointment{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.Appointment{})
	}

	if !db.HasTable(&model.VaccinationCenter{}) {
		db.CreateTable(&model.VaccinationCenter{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable((&model.VaccinationCenter{}))
		var vCenter = model.VaccinationCenter{Name: "Portet"}
		db.Create(&vCenter)

		vCenter = model.VaccinationCenter{Name: "Toulouse"}
		db.Create(&vCenter)
		vCenter = model.VaccinationCenter{Name: "Lyon"}
		db.Create(&vCenter)

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

	return db
}
