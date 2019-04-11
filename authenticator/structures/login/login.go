package login

import "github.com/bixlabs/authentication/authenticator/structures"

type Response struct {
	Token      string          `json:"token"`
	IssuedAt   int             `json:"iat"`
	Expiration int             `json:"exp"`
	User       structures.User `json:"user"`
}
