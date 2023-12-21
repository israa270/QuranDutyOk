package data

import (
	"github.com/ebedevelopment/next-gen-tms/server/model/data"
	repository "github.com/ebedevelopment/next-gen-tms/server/repository/data"
	"mime/multipart"
)

// FileUploadAndDownloadService struct
type FileUploadAndDownloadService struct {
	fileUploadAndDownloadRepository repository.FileUploadAndDownloadRepository
}

// // Upload create file upload record
// func (e *FileUploadAndDownloadService) Upload(file tms.FileUploadAndDownload) error {
// 	return e.fileUploadAndDownloadRepository.Upload(file)
// }

// FindFile delete file chunk record
// func (e *FileUploadAndDownloadService) FindFile(id uint) (tms.FileUploadAndDownload, error) {
// 	return e.fileUploadAndDownloadRepository.FindFile(id)
// }

// DeleteFile delete file record
// func (e *FileUploadAndDownloadService) DeleteFile(file tms.FileUploadAndDownload) error {
// 	return e.fileUploadAndDownloadRepository.DeleteFile(file)
// }

// GetFileRecordInfoList get list
// func (e *FileUploadAndDownloadService) GetFileRecordInfoList(info request.PageInfo) (interface{}, int64, error) {
// 	return e.fileUploadAndDownloadRepository.GetFileRecordInfoList(info)
// }

// UploadFile  upload file
func (e *FileUploadAndDownloadService) UploadFile(header *multipart.FileHeader, noSave string) (*tms.FileUploadAndDownload, error) {
	return e.fileUploadAndDownloadRepository.UploadFile(header, noSave)
}

// EditFileName  edit file Name
// func (e *FileUploadAndDownloadService) EditFileName(file tms.FileUploadAndDownload) error {
// 	return e.fileUploadAndDownloadRepository.EditFileName(file)
// }
