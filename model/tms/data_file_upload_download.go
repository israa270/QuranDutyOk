package tms

import (
	"github.com/ebedevelopment/next-gen-tms/server/global"
)

// FileUploadAndDownload file upload struct
type FileUploadAndDownload struct {
	global.GvaModel
	FileUploadAndDownloadDTO
}

// TableName FileUploadAndDownload table Name
func (FileUploadAndDownload) TableName() string {
	return "tms_file_upload_and_downloads"
}


type FileUploadAndDownloadDTO struct{
	Name string `json:"name" gorm:"comment:filename"`     // file name
	Url  string `json:"url" gorm:"comment:fileAddress"`   // file address
	Tag  string `json:"tag" gorm:"comment:fileLabel"`     // file label
	Key  string `json:"key" gorm:"comment:serial number"` // serial number
}