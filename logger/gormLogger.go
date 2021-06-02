package logger

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

func NewGormLogger(log *logrus.Logger, gormLoggerConfig *logger.Config) *logger.Interface {
	newLogger := logger.New(
		log,
		*gormLoggerConfig,
	)

	return &newLogger
}
