package providers

import (
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"os"
	"strings"
)

// NewEmailProvider returns a concrete email provider depending on the selected configuration
func NewEmailProvider() email.Sender {
	emailProvider := strings.ToLower(os.Getenv("EMAIL_PROVIDER"))

	switch emailProvider {
	case "mailgun":
		return NewMailgunSender()
	case "sendgrid":
		return NewSengridSender()
	case "smtp":
		return NewSMTPSender()
	default:
		return NewSMTPSender()
	}
}
