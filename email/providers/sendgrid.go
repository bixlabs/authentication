package providers

import (
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/bixlabs/authentication/authenticator/provider/email/message"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendgridSender is the Sendgrid provider to send emails
type SendgridSender struct {
	sg     *sendgrid.Client
	APIKey string `env:"SENDGRID_API_KEY"`
}

// NewSengridSender returns an instance of the SendgridSender
func NewSengridSender() email.Sender {
	sender := &SendgridSender{}
	err := env.Parse(sender)

	if err != nil {
		tools.Log().Panic("Parsing the env variables for the sendgrid sender failed", err)
	}

	if sender.APIKey == "" {
		tools.Log().Panic("A sendgrid api key is required", err)
	}

	sender.sg = sendgrid.NewSendClient(sender.APIKey)

	return sender
}

// Send is an implementation to send the emailMessage by email using Sendgrid
func (ss SendgridSender) Send(emailMessage *message.Message) error {
	_, err := ss.sg.Send(ss.fromEmailMessageToSendgridMessage(emailMessage))

	if err != nil {
		return err
	}

	return nil
}

func (ss SendgridSender) fromEmailMessageToSendgridMessage(message *message.Message) *mail.SGMailV3 {
	from := mail.NewEmail(message.FromName, message.From)
	to := mail.NewEmail(message.ToName, message.To)

	return mail.NewSingleEmail(from, message.Subject, to, message.Text, message.HTML)
}
