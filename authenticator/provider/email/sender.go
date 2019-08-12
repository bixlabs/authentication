package email

import (
	"github.com/bixlabs/authentication/authenticator/provider/email/message"
)

type Sender interface {
	Send(emailMessage *message.Message) error
}
