package signup

import "github.com/bixlabs/authentication/tools/rest"

//TODO: We could use nested struct promoted fields here but swaggo library is not generating correct documentation using that.
type Response struct {
	rest.ResponseWrapper
	Result *Result `json:"result"`
}

func newResponse(code int, result *Result, err error) Response {
	r := Response{}
	r.ResponseWrapper = rest.NewResponseWrapper(code, err)
	r.Result = result
	return r
}

func NewErrorResponse(code int, err error) Response {
	return newResponse(code, nil, err)
}

func NewResponse(code int, result *Result) Response {
	return newResponse(code, result, nil)
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

// We need this because go-swag library doesn't support embedded struct and doesn't show all the attributes in
// the documentation.
type SwaggerResponse struct {
	Status   string          `json:"status"`
	Code     int             `json:"code"`
	Messages []string        `json:"messages"`
	Result   *Response `json:"result,omitempty"`
}
