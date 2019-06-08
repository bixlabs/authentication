package rest_login

import (
	"github.com/bixlabs/authentication/authenticator/structures/login"
	"net/http"
)

//TODO: We could use nested struct promoted fields here but the swaggo library is not generating correct documentation using that.
type Response struct {
	Status   string          `json:"status"`
	Code     int             `json:"code"`
	Messages []string        `json:"messages,omitempty"`
	Result   *login.Response `json:"result,omitempty"`
}

func NewErrorResponse(code int, err error) Response {
	return Response{Status: http.StatusText(code), Code: code, Messages: []string{err.Error()}}
}

func NewResponse(code int, result *login.Response) Response {
	return Response{Status: http.StatusText(code), Code: code, Result: result}
}

type Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
