package config

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

// Zap struct
type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                           // level
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                        // format
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                        // log before prefix
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                 // log file folder
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"show-line"`                 // show row
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`       // edit code level
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"` // stack name
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`  // format console
	MaxAge        int    `mapstructure:"max-age" json:"max-age" yaml:"max-age"`
}

// ZapEncodeLevel  EncodeLevel  zapcore.LevelEncoder
func (z *Zap) ZapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": 
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": 
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": 
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": 
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// TransportLevel  zapcore.Level
func (z *Zap) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
