package email

import "github.com/bixlabs/authentication/authenticator/structures"

type Sender interface {
	SendEmailPasswordRequest(user structures.User, code string) error
}
