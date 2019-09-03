package providers

import (
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/bixlabs/authentication/authenticator/provider/email/message"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

// SMTPSender is a SMTP provider to send emails
type SMTPSender struct {
	dialer   *gomail.Dialer
	Host     string `env:"AUTH_SERVER_SMTP_HOST" envDefault:"smtp.gmail.com"`
	Port     int    `env:"AUTH_SERVER_SMTP_PORT" envDefault:"587"`
	Username string `env:"AUTH_SERVER_SMTP_USERNAME"`
	Password string `env:"AUTH_SERVER_SMTP_PASSWORD"`
}

func NewSMTPSender() email.Provider {
	sender := &SMTPSender{}

	contextLogger := sender.getLogger()
	contextLogger.Info("email provider is initializing")

	err := env.Parse(sender)
	if err != nil {
		contextLogger.Panic("parsing the env variables for the email provider failed", err)
	}

	if sender.Username == "" {
		contextLogger.Panic("a username is required", err)
	}

	if sender.Password == "" {
		contextLogger.Panic("a password is required", err)
	}

	sender.dialer = gomail.NewDialer(sender.Host, sender.Port, sender.Username, sender.Password)

	contextLogger.Info("email provider was initialized")

	return sender
}

func (ss SMTPSender) Send(emailMessage *message.Message) error {
	contextLogger := ss.getLogger()

	SMTPMessage := ss.fromEmailMessageToSMTPMessage(emailMessage)
	if err := ss.dialer.DialAndSend(SMTPMessage); err != nil {
		logFields := logrus.Fields{"err": err, "message_type": emailMessage.Type}
		contextLogger.WithFields(logFields).Error("there was an error sending the email")

		return err
	}

	return nil
}

func (ss SMTPSender) fromEmailMessageToSMTPMessage(message *message.Message) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", message.From)
	m.SetHeader("FromName", message.FromName)
	m.SetHeader("To", message.To)
	m.SetHeader("ToName", message.ToName)
	m.SetHeader("Subject", message.Subject)
	m.SetBody("text/html", message.HTML)

	return m
}

func (ss SMTPSender) getLogger() *logrus.Entry {
	return tools.Log().WithField("email_provider", "smtp")
}
