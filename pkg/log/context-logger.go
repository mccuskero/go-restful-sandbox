package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

type ContextLogger struct {
	logger *logrus.Entry
}

func NewContextLogger(fields map[string]interface{}) *ContextLogger {
	logger := logrus.WithFields(fields)

	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)
	
	// Only log the warning severity or above.
	logrus.SetLevel(logrus.WarnLevel)

	return &ContextLogger{
		logger: logger,
	}
}

func (log *ContextLogger) Info(msg string) {
	log.logger.Info(msg)
}

func (log *ContextLogger) Warn(msg string) {
	log.logger.Warn(msg)
}

func (log *ContextLogger) Error(msg string) {
	log.logger.Error(msg)
}

func (log *ContextLogger) Fatal(msg string) {
	log.logger.Fatal(msg)
}

func (log *ContextLogger) Debug(msg string) {

	log.logger.Debug(msg)
}


