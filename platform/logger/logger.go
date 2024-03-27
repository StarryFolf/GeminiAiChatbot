package logger

import (
	"fiber/pkg/config"
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	*logrus.Logger
}

var logger = &Logger{}

func SetUpLogger() {
	logger = &Logger{logrus.New()}
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetOutput(os.Stdout)

	if config.AppConfig().Debug {
		logger.SetLevel(logrus.DebugLevel)
	}
}

func GetLogger() *Logger {
	return logger
}
