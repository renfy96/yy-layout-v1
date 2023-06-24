package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model

	Account  string
	Password string
	Salt     string
	Phone    string
	Email    string
}
