package model

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	Id        uint      `gorm:"primaryKey AUTO_INCREMENT"`
	Email     string    `gorm:"not null" form:"email" json:"email"`
	Name      string    `gorm:"not null" form:"name" json:"name"`
	LastName  string    `gorm:"not null" form:"last_name" json:"last_name"`
	Vcid      int       `gorm:"not null" form:"vcid" json:"vcid"`
	Date      time.Time `gorm:"not null" form:"date" json:"date"`
	Validated bool      `gorm:"not null default false"`
}

func (a Appointment) ToString() string {
	return a.Email + a.Name + a.LastName + a.Date.String()
}

type VaccinationCenter struct {
	Id   int    `gorm:"AUTO_INCREMENT"`
	Name string `gorm:"not null"`
}

type Token struct {
	GeneratedToken uuid.UUID `gorm:"not null"`
	ValidDate      time.Time `gorm:"not null"`
	Appid          uint      `gorm:"not null"`
}

func New(Appid uint) Token {
	Token := Token{uuid.New(), time.Now().Add(4 * 60 * 1000), Appid}
	return Token
}

type Admin struct {
	UserName string `gorm:"not null"`
	Password string `gorm:"not null"`
	VcId     int    `gorm:"not null"`
}
