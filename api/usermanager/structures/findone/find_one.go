package findone

import (
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/tools/rest"
)

type Request struct {
	Email string `json:"email"`
}

type Result struct {
	User structures.User `json:"user"`
}

type Response struct {
	rest.ResponseWrapper
	Result Result `json:"result"`
}

func NewResponse(code int, user structures.User) Response {
	return newResponse(code, Result{User: user}, nil)
}

func NewErrorResponse(code int, err error) Response {
	return newResponse(code, Result{}, err)
}

func newResponse(code int, result Result, err error) Response {
	r := Response{}
	r.ResponseWrapper = rest.NewResponseWrapper(code, err)
	r.Result = result

	return r
}
