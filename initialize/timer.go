package initialize

import (
	"github.com/ebedevelopment/next-gen-tms/server/config"
	"github.com/ebedevelopment/next-gen-tms/server/global"
	"go.uber.org/zap"
)

// Timer
func Timer() {
	if global.GvaConfig.Timer.Start {
		for i := range global.GvaConfig.Timer.Detail {
			go func(detail config.Detail) {
				global.GvaTimer.AddTaskByFunc("ClearDB", global.GvaConfig.Timer.Spec, func() {
					err := ClearTable(global.GvaDB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						global.GvaLog.Error("error in Timer fn to delete  tables", zap.Error(err))
					}
				})
			}(global.GvaConfig.Timer.Detail[i])
		}
	}
}

