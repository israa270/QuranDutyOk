package system

import (
	"fmt"
	"net/http"

	controller "github.com/ebedevelopment/next-gen-tms/server/controller/system"
	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
	"github.com/ebedevelopment/next-gen-tms/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// OperationRecordApi struct
type OperationRecordApi struct {
	operationController controller.OperationRecordController
}

// FindSysOperationRecord
// @Tags SysOperationRecord
// @Summary queryByIdSysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data path string true "ID"
// @Success 200 {object}  response.Response{data=map[string]interface{},msg=string} "queryByIdSysOperationRecord"
// @Failure 500 {object}  response.Response{}
// @Router /sysOperationRecord/findSysOperationRecord/{id} [get]
func (s *OperationRecordApi) FindSysOperationRecord(c *gin.Context) {
	id := c.Param("id")

	s.operationController.FindSysOperationRecord(id, c)

}

// GetSysOperationRecordList
// @Tags SysOperationRecord
// @Summary pagingSysOperationRecordList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.SysOperationRecordSearch false "page number, page size, submit condition"
// @Success 200 {object}  response.Response{data=response.PageResult,msg=string} "pagingSysOperationRecordList,return includes list,total,page number,Quantity per page"
// @Failure 500 {object}  response.Response{}
// @Router /sysOperationRecord/getSysOperationRecordList [get]
func (s *OperationRecordApi) GetSysOperationRecordList(c *gin.Context) {
	var info sysReq.SysOperationRecordSearch
	if err := c.ShouldBindQuery(&info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "error", c)
		return
	}

	s.operationController.GetSysOperationRecordList(info, c)
}

// ExportSysOperationExcel
// @Tags sysOperationRecord
// @Summary exportLogExcel
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/json
// @Param data query sysReq.SysOperationRecordSearch false "exportLogExcelFileInfo"
// @Success 200
// @Failure 500 {object}  response.Response{}
// @Router /sysOperationRecord/exportLogExcel [get]
func (s *OperationRecordApi) ExportSysOperationExcel(c *gin.Context) {
	var info sysReq.SysOperationRecordSearch
	if err := c.ShouldBindQuery(&info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "error", c)
		return
	}

	if info.Format != utils.Excel && info.Format != utils.PDF {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(fmt.Errorf("format not pdf or excel")))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "error", c)
		return
	}

	s.operationController.ExportSysOperationExcel(info, c)

}
