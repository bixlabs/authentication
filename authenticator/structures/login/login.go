package login

import "github.com/bixlabs/authentication/authenticator/structures"

type Response struct {
	Token      string          `json:"jwt"`
	IssuedAt   int64           `json:"iat"`
	Expiration int64           `json:"exp"`
	User       structures.User `json:"user"`
}
