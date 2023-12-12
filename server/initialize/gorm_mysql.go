package initialize

import (
	"fmt"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/initialize/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql initialization Mysql database
func GormMysql(writer  bool) *gorm.DB {
	m := global.GvaConfig.Mysql
	if m.DbName == "" {
		global.GvaLog.Error("db name is empty "+ fmt.Sprintf("%#v",m.DbName))
		return nil
	}

	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(writer), // DSN data source name
		DefaultStringSize:         191,     // string type of field of default length
		SkipInitializeWithVersion: false,   // know version automatic configure
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)

		
		return db
	}
}

// GormMysqlByConfig initialization Mysql database used delete configure
// func GormMysqlByConfig(m config.DB) *gorm.DB {
// 	if m.DbName == "" {
// 		return nil
// 	}
// 	mysqlConfig := mysql.Config{
// 		DSN:                       m.Dsn(), // DSN data source name
// 		DefaultStringSize:         191,     // string type of field of default length
// 		SkipInitializeWithVersion: false,   // know version automatic configure
// 	}
// 	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
// 		panic(err)
// 	} else {
// 		sqlDB, _ := db.DB()
// 		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
// 		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
// 		return db
// 	}
// }
