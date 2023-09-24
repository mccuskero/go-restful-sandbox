package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

type NormalLogger struct {
	logger *logrus.Logger
}

func NewNormalLogger() *NormalLogger {
	logger := logrus.New()
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.WarnLevel)

	return &NormalLogger{
		logger: logger,
	}
}

func (log *NormalLogger) Info(msg string) {
	log.logger.Info(msg)
}

func (log *NormalLogger) Warn(msg string) {
	log.logger.Warn(msg)
}

func (log *NormalLogger) Error(msg string) {
	log.logger.Error(msg)
}

func (log *NormalLogger) Fatal(msg string) {
	log.logger.Fatal(msg)
}

func (log *NormalLogger) Debug(msg string) {

	log.logger.Debug(msg)
}
