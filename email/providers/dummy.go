package providers

import (
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/bixlabs/authentication/authenticator/provider/email/message"
)

// DummySender is a Dummy provider to send emails
type DummySender struct {
	EmailMessage *message.Message
}

// NewDummySender returns an instance of the DummySender
func NewDummySender() email.Sender {
	return &DummySender{}
}

func (sender *DummySender) GetEmailMessage() *message.Message {
	return sender.EmailMessage
}

// Send is an implementation to send the emailMessage by email using Dummy
func (sender *DummySender) Send(emailMessage *message.Message) error {
	sender.EmailMessage = emailMessage

	return nil
}
