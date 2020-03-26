package resetpass

import "github.com/bixlabs/authentication/tools/rest"

func NewResponse(code int) Response {
	return newResponse(code, Result{Success: true}, nil)
}

func newResponse(code int, result Result, err error) Response {
	r := Response{}
	r.ResponseWrapper = rest.NewResponseWrapper(code, err)
	r.Result = result
	return r
}

func NewErrorResponse(code int, err error) Response {
	return newResponse(code, Result{Success: false}, err)
}

type Response struct {
	rest.ResponseWrapper
	Result Result `json:"result,omitempty"`
}

type Result struct {
	Success bool `json:"success"`
}

type Request struct {
	Email       string `json:"email"`
	Code        string `json:"code"`
	NewPassword string `json:"newPassword"`
}
