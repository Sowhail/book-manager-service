package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	FirstName   string    `gorm:"not null"`
	LastName    string    `gorm:"not null"`
	Birthday    time.Time `gorm:"not null"`
	Nationality string    `gorm:"not null"`
	BookID      uint
}

type Book struct {
	gorm.Model
	Name            string    `gorm:"not null"`
	Author          Author    `gorm:"not null"`
	Category        string    `gorm:"not null"`
	Volume          int       `gorm:"not null"`
	PublishedAt     time.Time `gorm:"not null"`
	Summary         string    `gorm:"not null"`
	TableOfContents string    `gorm:"not null"` // contents are seperated by # in a string
	Publisher       string    `gorm:"not null"`
	UserID          uint
}
