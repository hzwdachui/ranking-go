package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func InitLogger(logLevel string) {
	logger = logrus.New()

	level := logrus.DebugLevel
	logger.SetLevel(logrus.ErrorLevel)
	switch {
	case logLevel == "debug":
		level = logrus.DebugLevel
	case logLevel == "warn":
		level = logrus.WarnLevel
	case logLevel == "info":
		level = logrus.InfoLevel
	case logLevel == "error":
		level = logrus.ErrorLevel
	default:
		level = logrus.DebugLevel
	}
	logPath, _ := os.Getwd()
	logFile := fmt.Sprintf("%s/log/log.log", logPath)
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		logger.SetOutput(os.Stdout)
		logger.Warn(err)
	} else {
		logger.SetOutput(f)
	}
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(level)
	logger.Info("Init logger successfully")
}

// Info
func Info(v ...interface{}) {
	logger.Info(v...)
}

// Infof
func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

// Warnf
func Warnf(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

// Warn
func Warn(v ...interface{}) {
	logger.Warn(v...)
}

// Error
func Error(v ...interface{}) {
	logger.Error(v...)
}

// Debug
func Debug(v ...interface{}) {
	logger.Debug(v...)
}

// Fatal
func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}
