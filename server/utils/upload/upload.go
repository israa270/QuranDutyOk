package upload

import (
	"mime/multipart"

	"github.com/ebedevelopment/next-gen-tms/server/global"
)

// OSS storage
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

// NewOss OSS method
func NewOss() OSS {
	switch global.GvaConfig.System.OssType {
	case "local":
		return &Local{}
	default:
		return &Local{}
	}
}
