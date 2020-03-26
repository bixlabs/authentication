package token

import (
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/tools/rest"
)

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

func NewResponse(code int, user structures.User) Response {
	return newResponse(code, Result{User: user}, nil)
}

type Result struct {
	User structures.User `json:"user"`
}
