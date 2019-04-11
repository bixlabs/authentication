package structures

type User struct {
	ID               string `json:"id"` // I'm not sure we need this, I'm adding it for now and we might remove it later
	Email            string `json:"email"`
	Password         string `json:"password"`
	GivenName        string `json:"givenName"`
	SecondName       string `json:"secondName"`
	FamilyName       string `json:"familyName"`
	SecondFamilyName string `json:"secondFamilyName"`
}
