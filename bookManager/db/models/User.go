package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName     string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	EmailAddress string `gorm:"unique;not null"`
	PhoneNumber  string `gorm:"unique;not null"`
	Gender       string
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	Books        []Book
}
