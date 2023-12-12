package data

import (
	"net/http"

	controller "github.com/ebedevelopment/next-gen-tms/server/controller/data"
	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileUploadAndDownloadApi struct {
	fileUploadController controller.FileUploadController
}

// UploadFile
// @Tags FileUploadAndDownload
// @Summary uploadFile
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "uploadFile"
// @Success 200 {object}  response.Response{} "uploadFile,return includes file details"
// @Failure 500 {object}  response.Response{}
// @Failure 400 {object}  response.Response{}
// @Router /fileUploadAndDownload/upload [post]
func (u *FileUploadAndDownloadApi) UploadFile(c *gin.Context) {

	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].ReceiveFile, zap.Error(err))
		response.FailWithMessage(global.Translate("file.receiveFile"), http.StatusBadRequest, "error", c)
		return
	}
	//limit file size
	// APK the value of 200 MB 8 == 200000000
	//TODO: read this value from config with unit and convert it, don't it fixed value
	if header.Size > 200000000 {
		global.GvaLog.Debug(global.GvaLoggerMessage["log"].LimitFileSize)
		response.FailWithMessage(global.Translate("file.limitFileSize"), http.StatusBadRequest, "warning", c)
		return
	}

	u.fileUploadController.UploadFile(header, noSave, c)

}

// UploadFirmwareFile
// @Tags UploadFirmwareFile
// @Summary UploadFirmwareFile
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "UploadFirmwareFile"
// @Success 200 {object}  response.Response{} "upload file,return includes file details"
// @Failure 500 {object}  response.Response{}
// @Failure 400 {object}  response.Response{}
// @Router /fileUploadAndDownload/uploadFirmware [post]
func (u *FileUploadAndDownloadApi) UploadFirmwareFile(c *gin.Context) {
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].ReceiveFile, zap.Error(err))
		response.FailWithMessage(global.Translate("file.receiveFile"), http.StatusInternalServerError, "error", c)
		return
	}
	//limit file size
	// Firmware the value of 1.5 GB 9 == 1500000000
	//TODO: read this value from config with unit and convert it, don't it fixed value
	if header.Size > 1500000000 {
		global.GvaLog.Debug(global.GvaLoggerMessage["log"].LimitFileSize)
		response.FailWithMessage(global.Translate("file.limitFileSize"), http.StatusBadRequest, "warning", c)
		return
	}

	u.fileUploadController.UploadFile(header, noSave, c)
}
