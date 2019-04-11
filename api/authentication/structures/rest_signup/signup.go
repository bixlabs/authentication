package rest_signup

//TODO: We could use nested struct promoted fields here but swaggo library is not generating correct documentation using that.
type Response struct {
	Status   string   `json:"status"`
	Code     int      `json:"code"`
	Messages []string `json:"messages"`
	Result   Result   `json:"result"`
}

type Result struct {
	Success bool `json:"success"`
}

type Request struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	GivenName        string `json:"givenName,omitempty"`
	SecondName       string `json:"secondName,omitempty"`
	FamilyName       string `json:"familyName,omitempty"`
	SecondFamilyName string `json:"secondFamilyName,omitempty"`
}
