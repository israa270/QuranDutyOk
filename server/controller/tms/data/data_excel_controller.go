package data

import (
	"net/http"
	"os"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


type ExcelController struct{

}

func (e *ExcelController) DownloadTemplate(fileName string, c *gin.Context) {

	filePath := global.GvaConfig.Excel.DirTemplate + fileName

	fi, err := os.Stat(filePath)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].NotFoundFile, zap.Error(err))
		response.FailWithMessage(global.Translate("general.notFoundFile"), http.StatusNotFound, "error", c)
		return
	}
	if fi.IsDir() {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].FileFolder, zap.Error(err))
		response.FailWithMessage(global.Translate("general.fileFolder"), http.StatusInternalServerError, "error", c)
		return
	}
	// c.Writer.Header().Add("success", "true")
	// c.File(filePath)

	filePath = filePath[1:]
	response.OkWithDetailed(gin.H{"fileUrl": global.GvaConfig.System.ServerPath + filePath}, global.Translate("general.getDataSuccess"), http.StatusOK, "success", c)
}