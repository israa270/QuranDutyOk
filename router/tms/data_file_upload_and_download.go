package tms

import (
	v1 "github.com/ebedevelopment/next-gen-tms/server/api/v1/tms/data"
	"github.com/gin-gonic/gin"
)

// FileUploadAndDownloadRouter struct
type FileUploadAndDownloadRouter struct{
	FileUploadAndDownloadApi v1.FileUploadAndDownloadApi
}

// InitFileUploadAndDownloadRouter init file upload and download
func (e *FileUploadAndDownloadRouter) InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	fileUploadAndDownloadRouter := Router.Group("fileUploadAndDownload")
	{
		fileUploadAndDownloadRouter.POST("upload", e.FileUploadAndDownloadApi.UploadFile)                 // upload file
		fileUploadAndDownloadRouter.POST("uploadFirmware", e.FileUploadAndDownloadApi.UploadFirmwareFile) // upload file

		// fileUploadAndDownloadRouter.GET("getFileList", FileUploadAndDownloadApi.GetFileList)  // get upload file list
		// fileUploadAndDownloadRouter.DELETE("deleteFile", FileUploadAndDownloadApi.DeleteFile) // delete specify file

		// fileUploadAndDownloadRouter.PUT("editFileName", FileUploadAndDownloadApi.EditFileName)
	}
}
