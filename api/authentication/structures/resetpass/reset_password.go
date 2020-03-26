package resetpass

import "github.com/bixlabs/authentication/tools/rest"

type Response struct {
	rest.ResponseWrapper
	Result Result `json:"result,omitempty"`
}

func NewResponse() Response {
	return Response{Result: Result{Success: true}}
}

func NewUnsuccessfulResponse() Response {
	return Response{Result: Result{Success: false}}
}

type Result struct {
	Success bool `json:"success"`
}

type Request struct {
	Email       string `json:"email"`
	Code        string `json:"code"`
	NewPassword string `json:"newPassword"`
}
