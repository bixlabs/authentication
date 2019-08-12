package providers

import (
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/bixlabs/authentication/authenticator/provider/email/message"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	"gopkg.in/gomail.v2"
)

// SMTPSender is a SMTP provider to send emails
type SMTPSender struct {
	dialer   *gomail.Dialer
	Host     string `env:"SMTP_HOST" envDefault:"smtp.gmail.com"`
	Port     int    `env:"SMTP_PORT" envDefault:"587"`
	Username string `env:"SMTP_USERNAME"`
	Password string `env:"SMTP_PASSWORD"`
}

// NewSMTPSender returns an instance of the SMTPSender
func NewSMTPSender() email.Sender {
	sender := &SMTPSender{}

	err := env.Parse(sender)
	if err != nil {
		tools.Log().Panic("Parsing the env variables for the smtp sender failed", err)
	}

	if sender.Username == "" {
		tools.Log().Panic("A smtp username is required", err)
	}

	if sender.Password == "" {
		tools.Log().Panic("A smtp password is required", err)
	}

	sender.dialer = gomail.NewDialer(sender.Host, sender.Port, sender.Username, sender.Password)

	return sender
}

// Send is an implementation to send the emailMessage by email using SMTP
func (ss SMTPSender) Send(emailMessage *message.Message) error {
	if err := ss.dialer.DialAndSend(ss.fromEmailMessageToSMTPMessage(emailMessage)); err != nil {
		return err
	}

	return nil
}

func (ss SMTPSender) fromEmailMessageToSMTPMessage(message *message.Message) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", message.From)
	m.SetHeader("To", message.To)
	m.SetHeader("Subject", message.Subject)
	m.SetBody("text/html", message.HTML)

	return m
}
