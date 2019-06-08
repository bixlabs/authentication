package rest_reset_password_request

import (
	"net/http"
)

//TODO: We could use nested struct promoted fields here but the swaggo library is not generating correct documentation using that.
type Response struct {
	Status   string   `json:"status"`
	Code     int      `json:"code"`
	Messages []string `json:"messages,omitempty"`
	Result   *Result  `json:"result,omitempty"`
}

func NewErrorResponse(code int, err error) Response {
	return Response{Status: http.StatusText(code), Code: code, Messages: []string{err.Error()}}
}

func NewResponse(code int, result *Result) Response {
	return Response{Status: http.StatusText(code), Code: code, Result: result}
}

type Result struct {
	Success bool `json:"success"`
}

type Request struct {
	Email string `json:"email"`
}
