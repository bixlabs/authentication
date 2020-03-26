package login

import (
	"github.com/bixlabs/authentication/authenticator/structures/login"
	"github.com/bixlabs/authentication/tools/rest"
)

type RestResponse struct {
	rest.ResponseWrapper
	Result Result `json:"result,omitempty"`
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

func NewResponse(code int, result Result) RestResponse {
	return newResponse(code, result, nil)
}

type Result struct {
	login.Response
}

type Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// We need this because go-swag library doesn't support embedded struct and doesn't show all the attributes in
// the documentation.
type SwaggerResponse struct {
	Status   string          `json:"status"`
	Code     int             `json:"code"`
	Messages []string        `json:"messages"`
	Result   *login.Response `json:"result,omitempty"`
}
