package login

import (
	"github.com/bixlabs/authentication/authenticator/structures/login"
	"github.com/bixlabs/authentication/tools/rest"
)

func NewResponse(code int, result Result) RestResponse {
	return newResponse(code, result, nil)
}

func newResponse(code int, result Result, err error) RestResponse {
	r := RestResponse{}
	r.ResponseWrapper = rest.NewResponseWrapper(code, err)
	r.Result = result
	return r
}

func NewErrorResponse(code int, err error) RestResponse {
	return newResponse(code, Result{}, err)
}

type RestResponse struct {
	rest.ResponseWrapper
	Result Result `json:"result,omitempty"`
}

type Result struct {
	login.Response
}

type Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
