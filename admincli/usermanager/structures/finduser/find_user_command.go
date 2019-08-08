package finduser

// Command received by the find user command
type Command struct {
	Email            string
	Password         string
	GivenName        string
	SecondName       string
	FamilyName       string
	SecondFamilyName string
}

// Result to the find user command
type Result struct {
	ID               string
	Email            string
	Password         string
	GivenName        string
	SecondName       string
	FamilyName       string
	SecondFamilyName string
}
