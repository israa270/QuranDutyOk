package data

import (
	"net/http"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	controller "github.com/ebedevelopment/next-gen-tms/server/controller/data"
	"github.com/gin-gonic/gin"
)

type ExcelApi struct{
	excelController  controller.ExcelController
}

// DownloadTemplate
// @Tags excel
// @Summary loadTemplate
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param fileName query string true "templateName"
// @Success 200
// @Failure 500
// @Router /excel/downloadTemplate [get]
func (e *ExcelApi) DownloadTemplate(c *gin.Context) {

	fileName := c.Query("fileName")
	if fileName == ""{
		global.GvaLog.Error(global.GvaLoggerMessage["log"].NotFoundFile)
		response.FailWithMessage(global.Translate("general.notFoundFile"), http.StatusBadRequest, "error", c)
		return
	}
	
	e.excelController.DownloadTemplate(fileName, c)
}
