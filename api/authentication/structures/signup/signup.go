package signup

import "github.com/bixlabs/authentication/tools/rest"

type Response struct {
	rest.ResponseWrapper
	Result Result `json:"result"`
}

func newResponse(code int, result Result, err error) Response {
	r := Response{}
	r.ResponseWrapper = rest.NewResponseWrapper(code, err)
	r.Result = result
	return r
}

func NewErrorResponse(code int, err error) Response {
	return newResponse(code, Result{}, err)
}

func NewResponse(code int) Response {
	return newResponse(code, Result{Success: true}, nil)
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
