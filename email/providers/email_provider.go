package providers

import (
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"os"
	"strings"
)

// NewEmailProvider returns a concrete email provider depending on the selected configuration
func NewEmailProvider() email.Provider {
	emailProvider := strings.ToLower(os.Getenv("AUTH_SERVER_EMAIL_PROVIDER"))

	switch emailProvider {
	case "mailgun":
		return NewMailgunSender()
	case "sendgrid":
		return NewSengridSender()
	case "smtp":
		return NewSMTPSender()
	case "none":
		return NewNoneSender()
	default:
		return NewSMTPSender()
	}
}
