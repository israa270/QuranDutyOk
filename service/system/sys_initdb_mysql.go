package system

import (
	"fmt"

	"github.com/ebedevelopment/next-gen-tms/server/config"
	"github.com/ebedevelopment/next-gen-tms/server/global"
	model "github.com/ebedevelopment/next-gen-tms/server/model/system"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlEmptyDsn(i config.Mysql) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", i.Username, i.Password, i.Path, i.WPort)
}

// initMysqlDB create database and initialization mysql
func (i *InitDBService) initMysqlDB(conf config.Mysql) error {
	dsn := MysqlEmptyDsn(conf)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.DbName)
	if err := i.createDatabase(dsn, "mysql", createSql); err != nil {
		return err
	} // create database

	mysqlConfig := conf
	if mysqlConfig.DbName == "" {
		global.GvaLog.Error("error db  name empty")
		return nil
	} // if not have database name, but jump initialization data

	if db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysqlConfig.Dsn(true), // DSN data source name
		DefaultStringSize:         191,                   // string type of field of default length
		SkipInitializeWithVersion: true,                  // know version automatic configure
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		global.GvaLog.Error("error configuration db ", zap.Error(err))
		return nil
	} else {

		global.GvaDB = db

	}

	if err := i.initTables(); err != nil {
		global.GvaDB = nil
		return err
	}

	// Create  user admin
	CreateAdmin()

	return nil
}

var adminService AdminService

func CreateAdmin() error {

	adminPassword := utils.BcryptHash("P@$$w0rD")

	admins := []model.Admin{
		{UserName: "admin*111", Password: adminPassword, Name: "admin" , Role: utils.Admin},
		{UserName: "admin*222", Password: adminPassword, Name: "admin" , Role: utils.Admin },
	}

	if err := adminService.CreateAdmin(admins); err != nil {
		global.GvaLog.Error("failed to create admins data", zap.Error(err))
		return err
	}

	return nil
}
