package providers

import (
	"fmt"
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/bixlabs/authentication/authenticator/provider/email/message"
	"github.com/bixlabs/authentication/tools"
	"github.com/sirupsen/logrus"
)

type NoneSender struct {
}

func NewNoneSender() email.Provider {
	sender := &NoneSender{}

	contextLogger := sender.getLogger()
	contextLogger.Info("email provider is initializing")
	contextLogger.Info("email provider was initialized")

	return sender
}

func (n NoneSender) Send(emailMessage *message.Message) error {
	contextLogger := n.getLogger()

	logFields := logrus.Fields{
		"message_type": emailMessage.Type,
		"message":      emailMessage.HTML,
		"to":           fmt.Sprintf("%s <%s>", emailMessage.ToName, emailMessage.To),
	}

	contextLogger.WithFields(logFields).Info("none email will be send")

	return nil
}

func (n NoneSender) getLogger() *logrus.Entry {
	return tools.Log().WithField("email_provider", "none")
}
