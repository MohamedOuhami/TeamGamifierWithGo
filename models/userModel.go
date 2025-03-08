package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"unique"`
	LastName  string
	Username  string `gorm:"unique"`
	Email     string
	Password  string
}
