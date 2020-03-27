package tools

import (
	"github.com/sirupsen/logrus"
	"os"
)

// For more information: https://github.com/sirupsen/logrus/blob/master/example_basic_test.go
// If we want to use a specific logging service there
// are multiple formatters implementations here: https://github.com/sirupsen/logrus#formatters

var log *logrus.Logger //nolint

func InitializeLogger() {
	if isProduction() {
		initializeLoggerForProduction()
	} else {
		initializeLogger()
	}
}

func isProduction() bool {
	return os.Getenv("AUTH_SERVER_APP_ENV") == "production"
}

func initializeLoggerForProduction() {
	initializeLogger()
	log.Level = logrus.InfoLevel
	log.SetReportCaller(false)
}

func initializeLogger() {
	log = logrus.New()
	log.Formatter = &logrus.TextFormatter{FullTimestamp: true}
	log.SetReportCaller(true)
	log.Level = logrus.TraceLevel
	log.Out = os.Stdout
}

func Log() *logrus.Logger {
	return log
}
