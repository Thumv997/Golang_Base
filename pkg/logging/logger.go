package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

func Init() {
	log = logrus.New()

	// Set output to a log file
	file, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to open log file, using default stderr")
	}

	// Optionally, you can configure the log format, level, and other options
	log.SetFormatter(&logrus.TextFormatter{})
	log.SetLevel(logrus.DebugLevel)
}

func GetLogger() *logrus.Logger {
	return log
}