package response

import "github.com/ebedevelopment/next-gen-tms/server/model/tms"

// FileResponse struct
type FileResponse struct {
	File tms.FileUploadAndDownloadDTO `json:"file"`
}
