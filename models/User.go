package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName		string
	LastName		string
	EmailAddress	string
	PhoneNumber		string
	Username		string
	Password		string
}
