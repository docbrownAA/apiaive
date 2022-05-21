package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Appointment struct {
	Id        int       `gorm:"primaryKey AUTO_INCREMENT"`
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
	Id    int    `gorm:"AUTO_INCREMENT"`
	Name  string `gorm:"not null"`
	Slots int    `gorm:"not null"`
}

type TokenAppointment struct {
	GeneratedToken uuid.UUID `gorm:"not null"`
	ValidDate      time.Time `gorm:"not null"`
	Appid          int       `gorm:"not null"`
}

func New(Appid int) TokenAppointment {
	Token := TokenAppointment{uuid.New(), time.Now().Add(time.Minute * time.Duration(10)), Appid}
	return Token
}

type User struct {
	UserName string `gorm:"unique;not null"`
	Email    string `grom:"unique;not null"`
	Password string `gorm:"not null"`
	Vcid     int    `gorm:"not null" json:"vcid"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
