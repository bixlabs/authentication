package tools

import (
	"github.com/sirupsen/logrus"
	"os"
)

// For more information: https://github.com/sirupsen/logrus/blob/master/example_basic_test.go
// If we want to use a specific logging service there are multiple formatters implementations here: https://github.com/sirupsen/logrus#formatters

var logger *logrus.Logger

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
	logger.Level = logrus.InfoLevel
	logger.SetReportCaller(false)

}

func initializeLogger() {
	logger = logrus.New()
	logger.Formatter = &logrus.TextFormatter{FullTimestamp: true}
	logger.SetReportCaller(true)
	logger.Level = logrus.TraceLevel
	logger.Out = os.Stdout
}

func Log() *logrus.Logger {
	return logger
}
