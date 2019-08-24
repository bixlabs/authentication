package providers

import (
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/bixlabs/authentication/authenticator/provider/email/message"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sirupsen/logrus"
)

// SendgridSender is the Sendgrid provider to send emails
type SendgridSender struct {
	sg     *sendgrid.Client
	APIKey string `env:"AUTH_SERVER_SENDGRID_API_KEY"`
}

// NewSengridSender returns an instance of the SendgridSender
func NewSengridSender() email.Provider {
	sender := &SendgridSender{}

	contextLogger := sender.getLogger()
	contextLogger.Info("email provider is initializing")

	err := env.Parse(sender)

	if err != nil {
		contextLogger.Panic("parsing the env variables for the email provider failed", err)
	}

	if sender.APIKey == "" {
		contextLogger.Panic("an api key is required", err)
	}

	sender.sg = sendgrid.NewSendClient(sender.APIKey)

	contextLogger.Info("email provider was initialized")

	return sender
}

// Send is an implementation to send the emailMessage by email using Sendgrid
func (ss SendgridSender) Send(emailMessage *message.Message) error {
	contextLogger := ss.getLogger()
	sendgridMessage := ss.fromEmailMessageToSendgridMessage(emailMessage)

	if _, err := ss.sg.Send(sendgridMessage); err != nil {
		logFields := logrus.Fields{"err": err, "message_type": emailMessage.Type}
		contextLogger.WithFields(logFields).Error("there was an error sending the email")

		return err
	}

	return nil
}

func (ss SendgridSender) fromEmailMessageToSendgridMessage(message *message.Message) *mail.SGMailV3 {
	from := mail.NewEmail(message.FromName, message.From)
	to := mail.NewEmail(message.ToName, message.To)

	return mail.NewSingleEmail(from, message.Subject, to, message.Text, message.HTML)
}

func (ss SendgridSender) getLogger() *logrus.Entry {
	return tools.Log().WithField("email_provider", "sendgrid")
}
