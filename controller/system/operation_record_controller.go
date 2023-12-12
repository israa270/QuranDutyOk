package system

import (
	"net/http"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
	"github.com/ebedevelopment/next-gen-tms/server/service/system"

	// "github.com/ebedevelopment/next-gen-tms/server/service/system"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type OperationRecordController struct {
	operationRecordService system.OperationRecordService
}

func (p *OperationRecordController) FindSysOperationRecord(id string, c *gin.Context) {
	if operationRecord, err := p.operationRecordService.GetSysOperationRecord(id); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].QueryFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFail"), http.StatusInternalServerError, "error", c)
	} else {
		response.OkWithDetailed(operationRecord, global.Translate("general.querySuccess"), http.StatusOK, "success", c)
	}
}

func (p *OperationRecordController) GetSysOperationRecordList(info sysReq.SysOperationRecordSearch, c *gin.Context) {
	if list, total, err := p.operationRecordService.GetSysOperationRecordInfoList(info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GetDataFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFail"), http.StatusInternalServerError, "error", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     info.Page,
			PageSize: info.PageSize,
		}, global.Translate("general.getDataSuccess"), http.StatusOK, "success", c)

		// response.OkWithDetailed(list, global.Translate("general.getDataSuccess"), http.StatusOK, "success", c)
	}
}

func (p *OperationRecordController) ExportSysOperationExcel(info sysReq.SysOperationRecordSearch, c *gin.Context) {


	if list, _, err := p.operationRecordService.GetSysOperationRecordInfoList(info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GetDataFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFail"), http.StatusInternalServerError, "error", c)
	} else {

		if len(list) == 0 {
			response.FailWithMessage(global.Translate("general.emptyExportData"), http.StatusBadRequest, "error", c)
			return
		}

		if _, err := p.operationRecordService.ParseLogInfoList2Excel(list, info.Format); err != nil {
			global.GvaLog.Error(global.GvaLoggerMessage["log"].ExcelFail, zap.Error(err))
			response.FailWithMessage(global.Translate("excel.excelFail"), http.StatusInternalServerError, "error", c)
			return
		} else {
			// 	c.Writer.Header().Add("success", "true")
			// 	c.File(filePath)

		

			response.FailWithMessage("failed to get file", http.StatusInternalServerError, "error", c)

			// fileUrl = fileUrl[1:]
			// response.OkWithDetailed(gin.H{"fileUrl": global.GvaConfig.System.ServerPath + fileUrl}, global.Translate("general.getDataSuccess"), http.StatusOK, "success", c)
		}
	}
}
