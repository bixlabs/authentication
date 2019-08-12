package providers

import (
	"context"
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/bixlabs/authentication/authenticator/provider/email/message"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	"github.com/mailgun/mailgun-go/v3"
	"time"
)

// MailgunSender is a Mailgun provider to send emails
type MailgunSender struct {
	mg     *mailgun.MailgunImpl
	Domain string `env:"MAILGUN_DOMAIN"`
	APIKey string `env:"MAILGUN_API_KEY"`
}

// NewMailgunSender returns an instance of the MailgunSender
func NewMailgunSender() email.Sender {
	sender := &MailgunSender{}
	err := env.Parse(sender)

	if err != nil {
		tools.Log().Panic("Parsing the env variables for the mailgun sender failed", err)
	}

	if sender.Domain == "" {
		tools.Log().Panic("A mailgun domain is required", err)
	}

	if sender.APIKey == "" {
		tools.Log().Panic("A mailgun api key is required", err)
	}

	sender.mg = mailgun.NewMailgun(sender.Domain, sender.APIKey)

	return sender
}

// Send is an implementation to send the emailMessage by email using Mailgun
func (ms MailgunSender) Send(emailMessage *message.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := ms.mg.Send(ctx, ms.fromEmailMessageToMailgunMessage(emailMessage))

	if err != nil {
		return err
	}

	return nil
}

func (ms MailgunSender) fromEmailMessageToMailgunMessage(message *message.Message) *mailgun.Message {
	from := message.From
	subject := message.Subject
	body := message.Text
	recipient := message.To

	mgMessage := ms.mg.NewMessage(from, subject, body, recipient)
	mgMessage.SetHtml(message.HTML)

	return mgMessage
}
