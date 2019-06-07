package sqlite

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email            string
	Password         string
	GivenName        string
	SecondName       string
	FamilyName       string
	SecondFamilyName string
	ResetToken       string
}
