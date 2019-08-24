package email

import (
	"github.com/bixlabs/authentication/authenticator/provider/email/message"
	"github.com/bixlabs/authentication/authenticator/structures"
)

type Sender interface {
	ForgotPasswordRequest(user structures.User, code string) error
}

type Provider interface {
	Send(emailMessage *message.Message) error
}
