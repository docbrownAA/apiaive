package model

import (
	"time"
)

type Appointment struct {
	Id        uint      `gorm:"primaryKey AUTO_INCREMENT"`
	Email     string    `gorm:"not null" form:"email" json:"email"`
	Name      string    `gorm:"not null" form:"name" json:"name"`
	LastName  string    `gorm:"not null" form:"last_name" json:"last_name"`
	VcId      int       `gorm:"not null" form:"vcid" json:"vcid"`
	Date      time.Time `gorm:"not null" form:"date" json:"date"`
	Validated bool      `gorm:"not null"`
}

type VaccinationCenter struct {
	Id   int    `gorm:"AUTO_INCREMENT"`
	Name string `gorm:"not null"`
}

type Token struct {
	Email          string    `gorm:"not null"`
	GeneratedToken string    `gorm:"not null"`
	ValidDate      time.Time `gorm:"not null"`
}

type Admin struct {
	UserName string `gorm:"not null"`
	Password string `gorm:"not null"`
	VcId     int    `gorm:"not null"`
}
