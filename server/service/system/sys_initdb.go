package system

import (
	"database/sql"
	"fmt"

	"github.com/ebedevelopment/next-gen-tms/server/config"
	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/management"
	"github.com/ebedevelopment/next-gen-tms/server/model/system"

	// "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
	"github.com/ebedevelopment/next-gen-tms/server/model/tms"
)

// InitDBService init service
type InitDBService struct{}

// InitDB create database and initialization total
// func (initDBService *InitDBService) InitDB(conf request.InitDB) error {
// 	switch conf.DBType {
// 	case "mysql":
// 		return initDBService.initMysqlDB(conf)
// 	// case "pgsql":
// 	// 	return initDBService.initPgsqlDB(conf)
// 	default:
// 		return initDBService.initMysqlDB(conf)
// 	}
// }

func (initDBService *InitDBService) InitDB(conf config.Mysql) error {
	// switch conf.DBType {
	// case "mysql":
		return initDBService.initMysqlDB(conf)
	// case "pgsql":
	// 	return initDBService.initPgsqlDB(conf)
	// default:
	// 	return initDBService.initMysqlDB(conf)
	// }
}



// initTables initialization table
func (i *InitDBService) initTables() error {
	if err := global.GvaDB.AutoMigrate(

		system.SysUser{},
	
		system.JwtBlacklist{},
		system.SysOperationRecord{},
		system.SysLoginAttempt{},
       
		system.Admin{},


		tms.FileUploadAndDownload{},
		management.Class{},
		management.HomeWork{},
		management.Student{},
		management.Teacher{},



	); err != nil {
		return err
	}

	return nil
}

// createDatabase create database(mysql)
func (i *InitDBService) createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}
