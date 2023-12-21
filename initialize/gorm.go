package initialize

import (
	"os"
	"strings"

	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/ebedevelopment/next-gen-tms/server/global"
	tms "github.com/ebedevelopment/next-gen-tms/server/model/data"
	"github.com/ebedevelopment/next-gen-tms/server/model/management"
	"github.com/ebedevelopment/next-gen-tms/server/model/system"
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

	err := db.AutoMigrate(
		system.SysUser{},
		system.JwtBlacklist{},
		system.SysOperationRecord{},
		system.SysLoginAttempt{},

		adapter.CasbinRule{},

		tms.FileUploadAndDownload{},

		system.Admin{},
		
		management.Class{},
		management.HomeWork{},
		management.Student{},
		management.Teacher{},
        management.StudentHomeWorks{},
	)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].TableFail, zap.Error(err))
		os.Exit(0)
	}

	global.GvaLog.Error(global.GvaLoggerMessage["log"].TableSuccess)
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
	global.GvaLog.Error(global.GvaLoggerMessage["log"].LoadJson)
	return nil
}
