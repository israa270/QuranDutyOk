package system

import (
	v1 "github.com/ebedevelopment/next-gen-tms/server/api/v1/system"
	"github.com/ebedevelopment/next-gen-tms/server/middleware"
	"github.com/gin-gonic/gin"
)

// OperationRecordRouter operation router
type OperationRecordRouter struct{
	auditLogApi v1.OperationRecordApi
}

// InitSysOperationRecordRouter  init operation router
func (s *OperationRecordRouter) InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("sysOperationRecord")
	operationRecordRouterM := Router.Group("sysOperationRecord").Use(middleware.OperationRecord())
	{
		operationRecordRouterM.GET("exportLogExcel", s.auditLogApi.ExportSysOperationExcel)
	}
	{
		operationRecordRouter.GET("findSysOperationRecord/:id", s.auditLogApi.FindSysOperationRecord)              // getByIDSysOperationRecord
		operationRecordRouter.GET("getSysOperationRecordList", s.auditLogApi.GetSysOperationRecordList)            // getSysOperationRecordList
	}
}
