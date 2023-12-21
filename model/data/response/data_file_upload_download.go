package response

import "github.com/ebedevelopment/next-gen-tms/server/model/data"

// FileResponse struct
type FileResponse struct {
	File tms.FileUploadAndDownloadDTO `json:"file"`
}
