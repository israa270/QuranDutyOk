package internal

import (
	"log"
	"os"
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbBASE interface {
	GetLogMode() string
}

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm customize configure
func (g *_gorm) Config() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})

	var logMode DbBASE
	
	switch global.GvaConfig.System.DbType {
		case "mysql":
			logMode = &global.GvaConfig.Mysql

		// case "pgsql":
		// 	logMode = &global.GvaConfig.Pgsql

		default:
			logMode = &global.GvaConfig.Mysql
	}

	switch logMode.GetLogMode() {
		case "silent", "Silent":
			config.Logger = _default.LogMode(logger.Silent)
		case "error", "Error":
			config.Logger = _default.LogMode(logger.Error)
		case "warn", "Warn":
			config.Logger = _default.LogMode(logger.Warn)
		case "info", "Info":
			config.Logger = _default.LogMode(logger.Info)
		default:
			config.Logger = _default.LogMode(logger.Info)
	}

	return config
}
