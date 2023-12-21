package data

import (
	"mime/multipart"
	"net/http"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	dataRes "github.com/ebedevelopment/next-gen-tms/server/model/data/response"
	"github.com/ebedevelopment/next-gen-tms/server/service/data"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)



type FileUploadController struct{
	fileUploadService data.FileUploadAndDownloadService

}




func (u *FileUploadController) UploadFile(header *multipart.FileHeader, noSave string ,c *gin.Context) {

	file, err := u.fileUploadService.UploadFile(header, noSave) // file upload take after arrive file path
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].LinkFilePath, zap.Error(err))
		response.FailWithMessage(global.Translate("file.linkFilePath"), http.StatusInternalServerError, "error", c)
		return
	}

	response.OkWithDetailed(dataRes.FileResponse{File: file.FileUploadAndDownloadDTO}, global.Translate("file.uploadSuccess"), http.StatusOK, "success", c)
}

