package middleware

import (
	"errors"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/service"
)


var(
	
	 jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
	 operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService

     ErrTokenExpired = errors.New(global.Translate("init.tokenExpire"))
)