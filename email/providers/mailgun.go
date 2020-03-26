package providers

import (
	"context"
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/bixlabs/authentication/authenticator/provider/email/message"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	"github.com/mailgun/mailgun-go/v3"
	"github.com/sirupsen/logrus"
	"time"
)

const sendEmailContextTimeoutSeconds = 10

type mailgunSender struct {
	mg     *mailgun.MailgunImpl
	Domain string `env:"AUTH_SERVER_MAILGUN_DOMAIN"`
	APIKey string `env:"AUTH_SERVER_MAILGUN_API_KEY"`
}

func NewMailgunSender() email.Provider {
	sender := &mailgunSender{}

	contextLogger := sender.getLogger()
	contextLogger.Info("email provider is initializing")

	err := env.Parse(sender)

	if err != nil {
		contextLogger.Panic("parsing the env variables for the email provider failed", err)
	}

	if sender.Domain == "" {
		contextLogger.Panic("a domain is required", err)
	}

	if sender.APIKey == "" {
		contextLogger.Panic("an api key is required", err)
	}

	sender.mg = mailgun.NewMailgun(sender.Domain, sender.APIKey)

	contextLogger.Info("email provider was initialized")

	return sender
}

func (ms mailgunSender) Send(emailMessage *message.Message) error {
	contextLogger := ms.getLogger()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*sendEmailContextTimeoutSeconds)
	defer cancel()

	mailgunMessage := ms.fromEmailMessageToMailgunMessage(emailMessage)
	if _, _, err := ms.mg.Send(ctx, mailgunMessage); err != nil {
		logFields := logrus.Fields{"err": err, "message_type": emailMessage.Type}
		contextLogger.WithFields(logFields).Error("there was an error sending the email")

		return err
	}

	return nil
}

func (ms mailgunSender) fromEmailMessageToMailgunMessage(message *message.Message) *mailgun.Message {
	mgMessage := ms.mg.NewMessage(message.From, message.Subject, message.Text, message.To)
	mgMessage.SetHtml(message.HTML)

	return mgMessage
}

func (ms mailgunSender) getLogger() *logrus.Entry {
	return tools.Log().WithField("email_provider", "mailgun")
}
