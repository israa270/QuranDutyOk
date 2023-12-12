package internal

import (
	"fmt"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"gorm.io/gorm/logger"
)

// writer
type writer struct {
	logger.Writer
}

// NewWriter writer constructor
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf format printing log
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.GvaConfig.System.DbType {
		case "mysql":
			logZap = global.GvaConfig.Mysql.LogZap
		// case "pgsql":
		// 	logZap = global.GvaConfig.Pgsql.LogZap
	}
	
	if logZap {
		global.GvaLog.Debug(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
