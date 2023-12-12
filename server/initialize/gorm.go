package initialize

import (
	"os"
	"strings"

	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/management"
	"github.com/ebedevelopment/next-gen-tms/server/model/system"
	"github.com/ebedevelopment/next-gen-tms/server/model/tms"
	"github.com/ebedevelopment/next-gen-tms/server/utils/translate"
	"go.uber.org/zap"

	"gorm.io/gorm"
)

// Gorm initialization database
func Gorm(writer bool) *gorm.DB {
	switch global.GvaConfig.System.DbType {
	case "mysql":
		return GormMysql(writer)
	// case "pgsql":
	// 	return GormPgSql(writer)
	default:
		return GormMysql(writer)
	}
}

// RegisterTables register database table
func RegisterTables(db *gorm.DB) {
	//    _ = db.Exec("ALTER TABLE tms_model ADD CONSTRAINT model_info UNIQUE (name, manufacturer_id)").Error
	//TODO Unique in terminal SN -> model id + serial no
	// _ = db.Exec("ALTER TABLE tms_model ADD CONSTRAINT model_info UNIQUE (name, manufacturer_id)").Error

	err := db.AutoMigrate(
		system.SysUser{},
		system.JwtBlacklist{},
		system.SysOperationRecord{},
		system.SysLoginAttempt{},


		adapter.CasbinRule{},

		tms.FileUploadAndDownload{},

		tms.Manufacturer{},

		system.Admin{},
		management.Class{},
		management.HomeWork{},
		management.Student{},
		management.Teacher{},
	)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].TableFail, zap.Error(err))
		os.Exit(0)
	}

	global.GvaLog.Debug(global.GvaLoggerMessage["log"].TableSuccess)
}

// MenuTranslator menu translator English and Ar
func MenuTranslator() error {
	langPath := global.GvaConfig.MenuTranslator.Dir
	langFiles, err := os.ReadDir(langPath)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].FilePath, zap.Error(err))
		return err
	}
	global.GvaMenu = make(map[string]translate.Menu)
	for _, langFile := range langFiles {
		langFilePath := langPath + langFile.Name()
		menu, err := translate.ParseJsonFile(langFilePath)
		if err != nil {
			global.GvaLog.Error(global.GvaLoggerMessage["log"].FilePath, zap.Error(err))
		}
		if strings.Contains(langFile.Name(), "ar") {
			global.GvaMenu["ar"] = *menu
		} else {
			global.GvaMenu["en"] = *menu
		}
	}
	global.GvaLog.Debug(global.GvaLoggerMessage["log"].LoadJson)
	return nil
}
