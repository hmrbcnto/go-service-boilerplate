package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName  string
	MiddleName string
	LastName   string
	Email      string
}
