package logger

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

var dataDogFormatter = &logrus.JSONFormatter{
	FieldMap: logrus.FieldMap{
		logrus.FieldKeyTime:  "timestamp",
		logrus.FieldKeyLevel: "level",
		logrus.FieldKeyMsg:   "message",
	},
}

func New(env string) *logrus.Logger {
	var formatter logrus.Formatter = new(logrus.TextFormatter)
	if env != "dev" {
		formatter = dataDogFormatter
	}

	logger := logrus.New()
	logger.SetFormatter(formatter)

	if env == "test" {
		logger.Out = ioutil.Discard
	}

	return logger
}
