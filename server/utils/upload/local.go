package upload

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
	"go.uber.org/zap"
)

type Local struct{}

// UploadFile upload file
func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {

	// read file extension
	ext := path.Ext(file.Filename)

	// read file name
	name := strings.TrimSuffix(file.Filename, ext)

	name = utils.MD5V([]byte(name))
	// new filename
	filename := name + "_" + time.Now().Format("20060102150405") + ext

	// filename := file.Filename
	mkdirErr := os.MkdirAll(global.GvaConfig.Local.Path, os.ModePerm)
	if mkdirErr != nil {
		global.GvaLog.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}

	// path and filename
	p := global.GvaConfig.Local.Path + "/" + filename

	f, openError := file.Open() // read file
	if openError != nil {
		global.GvaLog.Error("function file.Open() Failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			global.GvaLog.Error("close file Failed", zap.Any("err", err.Error()))
		}
	}(f) // create file defer close

	out, createErr := os.Create(p)
	if createErr != nil {
		global.GvaLog.Error("function os.Create() Failed", zap.Any("err", createErr.Error()))

		return "", "", errors.New("function os.Create() Failed, err:" + createErr.Error())
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			global.GvaLog.Error("close file Failed", zap.Any("err", err.Error()))
		}
	}(out) // create file defer close

	_, copyErr := io.Copy(out, f) //（copy）file
	if copyErr != nil {
		global.GvaLog.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return p, filename, nil
}

// DeleteFile delete file
func (*Local) DeleteFile(key string) error {
	p := global.GvaConfig.Local.Path + "/" + key
	if strings.Contains(p, global.GvaConfig.Local.Path) {
		if err := os.Remove(p); err != nil {
			return errors.New("local file delete failed, err:" + err.Error())
		}
	}
	return nil
}
