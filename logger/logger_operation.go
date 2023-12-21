package logger

import (
	"encoding/json"
	"io"
	"os"

	"github.com/ebedevelopment/next-gen-tms/server/config"
	"github.com/ebedevelopment/next-gen-tms/server/global"
	"go.uber.org/zap"
)

func ParseJsonFile(filePath string) (*config.MessageLogger, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			global.GvaLog.Error("error in parse json", zap.Error(err))
		}
	}(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)

	var menu config.MessageLogger
	err = json.Unmarshal(byteValue, &menu)
	if err != nil {
		global.GvaLog.Error("error in parse json", zap.Error(err))
	}

	return &menu, nil
}

// LoggerPath log path
func LoggerPath() error {
	logPath := global.GvaConfig.LoggerPath.Dir
	logFiles, err := os.ReadDir(logPath)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].FilePath, zap.Error(err))
		return err
	}

	for _, langFile := range logFiles {
		langFilePath := logPath + langFile.Name()
		menu, err := ParseJsonFile(langFilePath)
		if err != nil {
			global.GvaLog.Error(global.GvaLoggerMessage["log"].FilePath, zap.Error(err))
		}
		global.GvaLoggerMessage["log"] = menu
	}

	global.GvaLog.Error(global.GvaLoggerMessage["log"].LoadJson)
	return nil
}
