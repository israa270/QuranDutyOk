package core

import (
	"fmt"
	"os"

	"github.com/ebedevelopment/next-gen-tms/server/core/internal"
	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GvaConfig.Zap.Director); !ok { 
		fmt.Printf("create %v directory\n", global.GvaConfig.Zap.Director)
		_ = os.Mkdir(global.GvaConfig.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.GvaConfig.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// // Zap
// func Zap() (logger *zap.Logger) {
// 	if ok, _ := utils.PathExists(global.GvaConfig.Zap.Director); !ok { // yes no have Director file folder
// 		fmt.Printf("create %v directory\n", global.GvaConfig.Zap.Director)
// 		_ = os.Mkdir(global.GvaConfig.Zap.Director, os.ModePerm)
// 	}
// 	// debug level
// 	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
// 		return lev == zap.DebugLevel
// 	})
// 	// log level
// 	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
// 		return lev == zap.InfoLevel
// 	})
// 	//warn level
// 	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
// 		return lev == zap.WarnLevel
// 	})
// 	// mistake level
// 	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
// 		return lev >= zap.ErrorLevel
// 	})

// 	cores := [...]zapcore.Core{
// 		getEncoderCore(fmt.Sprintf("./%s/server_debug.log", global.GvaConfig.Zap.Director), debugPriority),
// 		getEncoderCore(fmt.Sprintf("./%s/server_info.log", global.GvaConfig.Zap.Director), infoPriority),
// 		getEncoderCore(fmt.Sprintf("./%s/server_warn.log", global.GvaConfig.Zap.Director), warnPriority),
// 		getEncoderCore(fmt.Sprintf("./%s/server_error.log", global.GvaConfig.Zap.Director), errorPriority),
// 	}
// 	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

// 	if global.GvaConfig.Zap.ShowLine {
// 		logger = logger.WithOptions(zap.AddCaller())
// 	}
// 	return logger
// }

// // getEncoderConfig get zapcore.EncoderConfig
// func getEncoderConfig() (config zapcore.EncoderConfig) {
// 	config = zapcore.EncoderConfig{
// 		MessageKey:     "message",
// 		LevelKey:       "level",
// 		TimeKey:        "time",
// 		NameKey:        "logger",
// 		CallerKey:      "caller",
// 		StacktraceKey:  global.GvaConfig.Zap.StacktraceKey,
// 		LineEnding:     zapcore.DefaultLineEnding,
// 		EncodeLevel:    zapcore.LowercaseLevelEncoder,
// 		EncodeTime:     CustomTimeEncoder,
// 		EncodeDuration: zapcore.SecondsDurationEncoder,
// 		EncodeCaller:   zapcore.FullCallerEncoder,
// 	}
// 	switch {
// 	case global.GvaConfig.Zap.EncodeLevel == "LowercaseLevelEncoder": // lower case edit code device (default)
// 		config.EncodeLevel = zapcore.LowercaseLevelEncoder
// 	case global.GvaConfig.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // lower case edit code device with color
// 		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
// 	case global.GvaConfig.Zap.EncodeLevel == "CapitalLevelEncoder": // capital letter edit code device
// 		config.EncodeLevel = zapcore.CapitalLevelEncoder
// 	case global.GvaConfig.Zap.EncodeLevel == "CapitalColorLevelEncoder": // capital letter edit code device with color
// 		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
// 	default:
// 		config.EncodeLevel = zapcore.LowercaseLevelEncoder
// 	}
// 	return config
// }

// // getEncoder get zapcore.Encoder
// func getEncoder() zapcore.Encoder {
// 	if global.GvaConfig.Zap.Format == "json" {
// 		return zapcore.NewJSONEncoder(getEncoderConfig())
// 	}
// 	return zapcore.NewConsoleEncoder(getEncoderConfig())
// }

// // getEncoderCore get Encoder of zapcore.Core
// func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
// 	writer := utils.GetWriteSyncer(fileName) // use file-rotate log setter row log separate
// 	return zapcore.NewCore(getEncoder(), writer, level)
// }

// // customize log format time between format
// func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
// 	enc.AppendString(t.Format(global.GvaConfig.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
// }
