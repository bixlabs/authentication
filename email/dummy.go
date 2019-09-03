package email

import (
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/bixlabs/authentication/authenticator/structures"
)

type sender struct{}

func NewDummySender() email.Sender {
	return &sender{}
}

func (s sender) ForgotPasswordRequest(user structures.User, code string) error {
	return nil
}
