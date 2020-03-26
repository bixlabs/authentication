package changepass

import "github.com/bixlabs/authentication/tools/rest"

type Response struct {
	rest.ResponseWrapper
	Result Result `json:"result,omitempty"`
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

func NewResponse(code int, success bool) Response {
	return newResponse(code, Result{Success: success}, nil)
}

type Result struct {
	Success bool `json:"success"`
}

type Request struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
	Email       string `json:"email"`
}

// We need this because go-swag library doesn't support embedded struct and doesn't show all the attributes in
// the documentation.
type SwaggerResponse struct {
	Status   string    `json:"status"`
	Code     int       `json:"code"`
	Messages []string  `json:"messages"`
	Result   *Response `json:"result,omitempty"`
}
