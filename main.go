package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/ebedevelopment/next-gen-tms/server/core"
	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/initialize"
	"github.com/ebedevelopment/next-gen-tms/server/logger"
	"github.com/ebedevelopment/next-gen-tms/server/service/system"
	"github.com/ebedevelopment/next-gen-tms/server/utils/translate"
	"go.uber.org/zap"
)


var initDBService  system.InitDBService
//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	global.GvaVP = core.Viper() // initialization Viper

	initialize.OtherInit()

	global.GvaLog = core.Zap() // initialization zap log library
	zap.ReplaceGlobals(global.GvaLog)

	logger.LoggerPath() //initialization log message

	//load menu ar-en files
	if global.GvaDB == nil {
		err := initialize.MenuTranslator()
		if err != nil {
			global.GvaLog.Error("failed to get menu translator", zap.Error(err))
		}
	}

	global.GvaDB = initialize.Gorm(true) // gorm connection database for writer
	// global.GvaDB = initialize.Gorm(false) // gorm connection database for reader
    
    if global.GvaDB == nil {
		if err := initDBService.InitDB(global.GvaConfig.Mysql); err != nil {
			global.GvaLog.Error(global.GvaLoggerMessage["log"].AutoCreateDBFail + " in write", zap.Error(err))
			log.Fatal(err)
		}
	}
	


	initialize.Timer()
	// initialize.DBList()

	global.GvaTranslator = translate.Translator{} // create translator instance  here
	// global.GvaTranslator.InitTranslator(global.GvaConfig.Language.Language, global.GvaConfig.Language.Dir)
	global.GvaTranslator.InitTranslator(global.GvaConfig.Language.Dir)

	if global.GvaDB != nil {
		initialize.RegisterTables(global.GvaDB) // initialization table
		db, _ := global.GvaDB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				global.GvaLog.Error("failed to initialize Register Tables", zap.Error(err))
			}
		}(db)
	}



	err := os.MkdirAll(global.GvaConfig.Local.Path, os.ModePerm)
	if err != nil {
		// Handle the error if the directory could not be created
		global.GvaLog.Error("error in creating directory uploads ", zap.Error(err))
		// return 
	}

	if err := core.RunWindowsServer(); err !=nil{
		log.Fatal(err)
	}
}
