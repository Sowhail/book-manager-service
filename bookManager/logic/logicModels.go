package logic

import (
	"time"
)

type User struct {
	UserName     string `json:"userName"`
	Password     string `json:"password"`
	EmailAddress string `json:"emailAddress"`
	PhoneNumber  string `json:"phoneNumber"`
	Gender       string `json:"gender"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

type author struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Birthday    time.Time `json:"birthday"`
	Nationality string    `json:"nationality"`
}

type Book struct {
	Name            string    `json:"name"`
	Author          author    `json:"author"`
	Category        string    `json:"category"`
	Volume          int       `json:"volume"`
	PublishedAt     time.Time `json:"published_at"`
	Summary         string    `json:"summary"`
	TableOfContents []string  `json:"table_of_contents"`
	Publisher       string    `json:"publisher"`
}
