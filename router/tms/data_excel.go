package tms

import (
	v1 "github.com/ebedevelopment/next-gen-tms/server/api/v1/tms/data"
	"github.com/gin-gonic/gin"
)

// ExcelRouter struct
type ExcelRouter struct{
	excelApi v1.ExcelApi
}

// InitDataExcelRouter init excel router
func (e *ExcelRouter) InitDataExcelRouter(Router *gin.RouterGroup) {
	excelRouter := Router.Group("excel")
	{
		excelRouter.GET("downloadTemplate", e.excelApi.DownloadTemplate) // load template file
	}
}
