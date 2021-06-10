package mysqldb

import (
	"github.com/escaletech/go-escale/messages"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/escaletech/go-escale/logger"
)

// Connect to a MySQL instance using Gorm
func GormConnect(params GormConnParams) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               setDSN(params.ConnectionParams.DSN, params.ConnectionParams.ParseTime),
		DefaultStringSize: params.ConnectionParams.DefaultStringSize,
	}), &gorm.Config{
		Logger: *logger.NewGormLogger(params.Log, params.GormLoggerConfig),
	})

	if err != nil {
		params.Log.Errorf(messages.DBConnectionError(params.ConnectionParams.Identifier, err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(params.ConnectionParams.ConnMaxLifetime)
	sqlDB.SetMaxOpenConns(params.ConnectionParams.MaxOpenConns)
	sqlDB.SetMaxIdleConns(params.ConnectionParams.MaxIdleConns)

	return db
}

func setDSN(dsn string, parseTime bool) string {
	if parseTime {
		dsn += "?parseTime=true"
	}

	return dsn
}
