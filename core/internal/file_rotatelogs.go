package internal

import (
	"os"
	"path"
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

var FileRotateLogs = new(fileRotateLogs)

type fileRotateLogs struct{}

// GetWriteSyncer  zapcore.WriteSyncer
func (r *fileRotateLogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotateLogs.New(
		path.Join(global.GvaConfig.Zap.Director, "%Y-%m-%d", level+".log"),
		rotateLogs.WithClock(rotateLogs.Local),
		rotateLogs.WithMaxAge(time.Duration(global.GvaConfig.Zap.MaxAge)*24*time.Hour), 
		rotateLogs.WithRotationTime(time.Hour*24),
	)
	if global.GvaConfig.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
