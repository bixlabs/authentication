package email

import (
	"github.com/bixlabs/authentication/authenticator/provider/email/message"
	"github.com/bixlabs/authentication/authenticator/provider/email/template"
	"github.com/bixlabs/authentication/authenticator/provider/email/template/forgotpassword"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
)

// AuthSender builds all the different kinds of email messages
// and use the emailSender to send them.
type AuthSender struct {
	From        string `env:"AUTH_SERVER_EMAIL_FROM" envDefault:"test@example.com"`
	FromName    string `env:"AUTH_SERVER_EMAIL_FROM_NAME"`
	emailSender Sender
}

// NewAuthSender returns an instance of the AuthSender
func NewAuthSender(emailSender Sender) *AuthSender {
	authSender := &AuthSender{emailSender: emailSender}
	err := env.Parse(authSender)

	if err != nil {
		tools.Log().Panic("Parsing the env variables for the auth sender failed", err)
	}

	return authSender
}

// SendEmailForgotPassword builds a forgot email message and send it using the emailSender,
// this forgot email message contains the code to reset the password.
func (as AuthSender) SendEmailForgotPassword(user structures.User, code string) error {
	templateBuilder := template.NewTemplateBuilder()
	htmlMessage, textMessage, err := templateBuilder.Build("forgot_password", &forgotpassword.TemplateParam{Code: code})
	if err != nil {
		return err
	}

	emailMessage := &message.Message{
		From:     as.From,
		FromName: as.FromName,
		To:       user.Email,
		ToName:   "",
		Subject:  "Reset your Password",
		HTML:     htmlMessage,
		Text:     textMessage,
	}

	return as.emailSender.Send(emailMessage)
}
