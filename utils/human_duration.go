package utils

import (
	"strconv"
	"strings"
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"go.uber.org/zap"
)

func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, err := strconv.Atoi(d[:index])     
		if err != nil{
           global.GvaLog.Error("failed to parse hours ", zap.Error(err))
		}

		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
