package structures

type User struct {
	ID               string `json:"id"` // I'm not sure we need this, I'm adding it for now and we might remove it later
	Email            string `json:"email"`
	Password         string `json:"password,omitempty"`
	GivenName        string `json:"givenName,omitempty"`
	SecondName       string `json:"secondName,omitempty"`
	FamilyName       string `json:"familyName,omitempty"`
	SecondFamilyName string `json:"secondFamilyName,omitempty"`
	ResetToken       string
}
