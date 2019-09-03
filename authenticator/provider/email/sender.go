package email

import (
	"github.com/bixlabs/authentication/authenticator/provider/email/message"
	"github.com/bixlabs/authentication/authenticator/structures"
)

// Sender is how business layer uses emails.
type Sender interface {
	ForgotPasswordRequest(user structures.User, code string) error
}

// Provider represents different email manager platforms
type Provider interface {
	Send(emailMessage *message.Message) error
}
