package tms

import (
	// "errors"
	"mime/multipart"
	"strings"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	// "github.com/ebedevelopment/next-gen-tms/server/model/common/request"
	"github.com/ebedevelopment/next-gen-tms/server/model/tms"
	"github.com/ebedevelopment/next-gen-tms/server/utils/upload"
)

// FileUploadAndDownloadRepository struct
type FileUploadAndDownloadRepository struct{}

// Upload create file upload record
func (e *FileUploadAndDownloadRepository) upload(file tms.FileUploadAndDownload) error {
	return global.GvaDB.Create(&file).Error
}

// FindFile delete file chunk record
// func (e *FileUploadAndDownloadRepository) FindFile(id uint) (tms.FileUploadAndDownload, error) {
// 	var file tms.FileUploadAndDownload
// 	err := global.GvaDB.Where("id = ?", id).First(&file).Error
// 	return file, err
// }

// DeleteFile delete file record
// func (e *FileUploadAndDownloadRepository) DeleteFile(file tms.FileUploadAndDownload) (err error) {
// 	var fileFromDb tms.FileUploadAndDownload
// 	fileFromDb, err = e.FindFile(file.ID)
// 	if err != nil {
// 		return
// 	}
// 	oss := upload.NewOss()
// 	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
// 		return errors.New("file delete failed")
// 	}
// 	err = global.GvaDB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
// 	return err
// }

// // GetFileRecordInfoList get list
// func (e *FileUploadAndDownloadRepository) GetFileRecordInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
// 	if info.PageSize == 0 || info.PageSize > global.GvaConfig.Mysql.LimitRecords {
// 		info.PageSize = global.GvaConfig.Mysql.LimitRecords
// 	}
// 	limit := info.PageSize
// 	offset := info.PageSize * (info.Page - 1)
// 	db := global.GvaDB.Model(&tms.FileUploadAndDownload{})
// 	var fileLists []tms.FileUploadAndDownload
// 	err = db.Count(&total).Error
// 	if err != nil {
// 		return
// 	}
// 	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
// 	return fileLists, total, err
// }

// UploadFile  upload file
func (e *FileUploadAndDownloadRepository) UploadFile(header *multipart.FileHeader, noSave string) (file *tms.FileUploadAndDownload, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		return nil, err
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := tms.FileUploadAndDownload{
			FileUploadAndDownloadDTO: tms.FileUploadAndDownloadDTO{
				Url:  filePath,
				Name: header.Filename,
				Tag:  s[len(s)-1],
				Key:  key,
			},
		}
		return &f, e.upload(f)
	}
	return
}

// EditFileName  edit file Name
// func (e *FileUploadAndDownloadRepository) EditFileName(file tms.FileUploadAndDownload) (err error) {
// 	var fileFromDb tms.FileUploadAndDownload
// 	return global.GvaDB.Where("id = ?", file.ID).First(&fileFromDb).Update("name", file.Name).Error
// }
