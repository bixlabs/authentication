package structures

import "time"

type User struct {
	ID                string     `json:"id"`
	Email             string     `json:"email"`
	Password          string     `json:"-"`
	GivenName         string     `json:"givenName,omitempty"`
	SecondName        string     `json:"secondName,omitempty"`
	FamilyName        string     `json:"familyName,omitempty"`
	SecondFamilyName  string     `json:"secondFamilyName,omitempty"`
	ResetToken        string     `json:"-"`
	GeneratedPassword string     `json:"generatedPassword,omitempty"`
	DeletedAt         *time.Time `json:"-"`
}
