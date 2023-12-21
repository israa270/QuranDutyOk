package data

import (
	v1 "github.com/ebedevelopment/next-gen-tms/server/api/v1/data"
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
	}
}
