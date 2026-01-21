package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *logrus.Logger // Global logger instance

// Init initializes the logger with rotation and formatting
func Init(serviceName string) {
	Logger = logrus.New()

	// Set up log rotation
	logFile := &lumberjack.Logger{
		Filename:   "../../logging-service/logger/log/" + serviceName + ".log",
		MaxSize:    10, // Max megabytes before log is rotated
		MaxBackups: 3,  // Max number of old log files to retain
		MaxAge:     28, // Max number of days to retain old logs
		Compress:   true,
	}

	// Set output to both file and console
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	Logger.SetOutput(multiWriter)
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetLevel(logrus.InfoLevel)
}
