package implementation

import (
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/bixlabs/authentication/authenticator/provider/email/message"
	"github.com/bixlabs/authentication/authenticator/provider/email/template"
	"github.com/bixlabs/authentication/authenticator/provider/email/template/forgotpassword"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
)

// AuthSender builds all the different kinds of email messages
// and use the emailSender to send them.
type sender struct {
	From     string `env:"AUTH_SERVER_EMAIL_FROM" envDefault:"test@example.com"`
	FromName string `env:"AUTH_SERVER_EMAIL_FROM_NAME"`
	provider email.Provider
}

func NewSender(provider email.Provider) email.Sender {
	s := &sender{provider: provider}
	err := env.Parse(s)

	if err != nil {
		tools.Log().Panic("Parsing the env variables for the auth sender failed", err)
	}

	return s
}

// builds a forgot email message and send it using the emailSender
// this message contains the code to reset the password.
func (s sender) ForgotPasswordRequest(user structures.User, code string) error {
	defaultTemplate := forgotpassword.NewTemplateHTML()
	defaultTemplateParam := forgotpassword.NewTempateParam(code)
	templateBuilder := template.NewTemplateBuilder(defaultTemplate)
	htmlMessage, err := templateBuilder.Build(defaultTemplateParam)

	if err != nil {
		return err
	}

	emailMessage := &message.Message{
		From:     s.From,
		FromName: s.FromName,
		To:       user.Email,
		ToName:   "",
		Subject:  "Reset your Password",
		HTML:     htmlMessage,
		Type:     "Forgot Password",
	}

	return s.provider.Send(emailMessage)
}
