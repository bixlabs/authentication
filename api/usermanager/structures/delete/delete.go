package delete

import (
	"github.com/bixlabs/authentication/tools/rest"
)

type Request struct {
	Email string `json:"email"`
}

type Result struct {
	Success bool `json:"sucess"`
}

type Response struct {
	rest.ResponseWrapper
	Result Result `json:"result"`
}

func NewResponse(code int) Response {
	return newResponse(code, Result{Success: true}, nil)
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
