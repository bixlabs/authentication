package tools

import (
	"github.com/sirupsen/logrus"
	"os"
)

// For more information: https://github.com/sirupsen/logrus/blob/master/example_basic_test.go

var log *logrus.Logger //nolint

func InitializeLogger() {
	log = logrus.New()
	log.SetLevel(logrus.TraceLevel)

	if os.Getenv("AUTH_SERVER_APP_ENV") == "production" {
		log.SetLevel(logrus.InfoLevel)
	}

	log.SetOutput(os.Stdout)
}

func Log() *logrus.Logger {
	return log
}
