package adapters

import (
	"os"

	"github.com/escaletech/go-escale/logger"
	"github.com/sirupsen/logrus"
)

var LogrusDataDogFormatter = &logrus.JSONFormatter{
	FieldMap: logrus.FieldMap{
		logrus.FieldKeyTime:  "timestamp",
		logrus.FieldKeyLevel: "level",
		logrus.FieldKeyMsg:   "message",
	},
}

var levelLog = map[logger.Level]logrus.Level{
	logger.DEBUG: logrus.DebugLevel,
	logger.INFO:  logrus.InfoLevel,
	logger.WARN:  logrus.WarnLevel,
	logger.ERROR: logrus.ErrorLevel,
	logger.FATAL: logrus.FatalLevel,
}

const FILE_ACCESS_LOG string = "access_log"

type logrusAdapter struct {
	log          *logrus.Logger
	accessOutput *os.File
}

func NewLogrusAdapter(env string, level logger.Level) logger.Adapter {
	var formatter logrus.Formatter = LogrusDataDogFormatter
	if env == "dev" {
		formatter = new(logrus.TextFormatter)
	}

	log := logrus.New()
	log.SetFormatter(formatter)
	log.SetLevel(levelLog[level])

	accessLogFile, err := os.OpenFile(FILE_ACCESS_LOG, os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Error("Failed to open file " + FILE_ACCESS_LOG + "message: " + err.Error())
	}

	return &logrusAdapter{
		log:          log,
		accessOutput: accessLogFile,
	}
}

func (la *logrusAdapter) Deprecated() *logrus.Logger {
	return la.log
}

func (la *logrusAdapter) Error(msg string) {
	la.log.SetOutput(os.Stderr)
	la.log.Error(msg)
}

func (la *logrusAdapter) Warn(msg string) {
	la.log.SetOutput(os.Stderr)
	la.log.Warn(msg)
}

func (la *logrusAdapter) Info(msg string) {
	la.log.SetOutput(os.Stdout)
	la.log.Info(msg)
}

func (la *logrusAdapter) Debug(msg string) {
	la.log.SetOutput(os.Stderr)
	la.log.Debug(msg)
}

func (la *logrusAdapter) Fatal(msg string) {
	la.log.SetOutput(os.Stderr)
	la.log.Fatal(msg)
}

func (la *logrusAdapter) Access(msg string) {
	la.log.SetOutput(la.accessOutput)
	la.log.Info(msg)
}
