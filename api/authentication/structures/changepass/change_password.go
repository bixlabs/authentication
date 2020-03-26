package changepass

import "github.com/bixlabs/authentication/tools/rest"

func NewResponse(code int, success bool) Response {
	return newResponse(code, Result{Success: success}, nil)
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

type Response struct {
	rest.ResponseWrapper
	Result Result `json:"result,omitempty"`
}

type Result struct {
	Success bool `json:"success"`
}

type Request struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
	Email       string `json:"email"`
}
