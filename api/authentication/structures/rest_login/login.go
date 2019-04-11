package rest_login

import (
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/authenticator/structures/login"
)

//TODO: We could use nested struct promoted fields here but the swaggo library is not generating correct documentation using that.
type Response struct {
	Status   string   `json:"status"`
	Code     int      `json:"code"`
	Messages []string `json:"messages"`
	Result   Result   `json:"result"`
}

type Result struct {
	Token      string          `json:"jwt"`
	IssuedAt   int             `json:"iua"`
	Expiration int             `json:"exp"`
	User       structures.User `json:"user"`
}

type Request struct {
	Email    string
	Password string
}

// Example of how we must create translators to communicate between business and this layer.
func LoginResponseToRest(login login.Response) Response {
	result := Response{}
	result.Result.Token = login.Token
	result.Result.IssuedAt = login.IssuedAt
	result.Result.Expiration = login.Expiration
	result.Result.User = login.User
	return result
}

func RestResponseToLogin(rest Response) login.Response {
	result := login.Response{}
	result.Token = rest.Result.Token
	result.IssuedAt = rest.Result.IssuedAt
	result.Expiration = rest.Result.Expiration
	result.User = rest.Result.User
	return result
}
