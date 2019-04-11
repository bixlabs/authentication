package login

import "github.com/bixlabs/authentication/authenticator/structures"

type Response struct {
	Token      string          `json:"jwt"`
	IssuedAt   int             `json:"iua"`
	Expiration int             `json:"exp"`
	User       structures.User `json:"user"`
}
