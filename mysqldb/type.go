package mysqldb

import (
	"time"

	"github.com/sirupsen/logrus"
	gormLogger "gorm.io/gorm/logger"
)

type ConnectionParams struct {
	DSN               string
	DefaultStringSize uint
	Identifier        string
	ConnMaxLifetime   time.Duration
	MaxOpenConns      int
	MaxIdleConns      int
	ParseTime         bool
}

type GormConnParams struct {
	ConnectionParams ConnectionParams
	Log              *logrus.Logger
	GormLoggerConfig *gormLogger.Config
}
