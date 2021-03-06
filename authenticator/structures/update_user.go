package structures

// UpdateUser represents the object to update a user
type UpdateUser struct {
	Email            string `json:"email"`
	Password         string `json:"-"`
	GivenName        string `json:"givenName,omitempty"`
	SecondName       string `json:"secondName,omitempty"`
	FamilyName       string `json:"familyName,omitempty"`
	SecondFamilyName string `json:"secondFamilyName,omitempty"`
}
