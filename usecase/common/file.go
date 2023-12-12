package common

import (
	"fmt"
	// "net/http"
	"os"
	"path"
	"strings"

	// "io"
	// "net/http"
	"path/filepath"
	"time"
)

// RemoveFile remove file from folder if useless
func (c *CommonUsecase) RemoveFile(fileUrl string) (err error) {
	err = os.RemoveAll(fileUrl)
	if err != nil {
		return err
	}
	return nil
}

// FilenameWithoutExtension  return file name without extension
func (c *CommonUsecase) FilenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}


// ByteCountSI size of file in string
func (c *CommonUsecase) ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

// GetFileExtension return file extension
func (c *CommonUsecase) GetFileExtension(fileName string) string {
	return filepath.Ext(fileName)
}

// GetFileSize get file size
func (c *CommonUsecase) GetFileSize(fileUrl string) (int64, error) {
	info, err := os.Stat(fileUrl)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// Check file Path exist or not
func (c *CommonUsecase) CheckFilePath(fileUrl string) bool {
	_, err := os.Stat(fileUrl)
	return err == nil
}


// GetFileName get file name from path
func (c *CommonUsecase) GetFileName(url string) string{
	return path.Base(url)
}


func (c *CommonUsecase) FileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

//FileExists check file exist
func (c *CommonUsecase) FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
	   return false
	}
	return !info.IsDir()
 }

 func (c *CommonUsecase) HashFileName(fName string) string{
	// read file extension
	ext := path.Ext(fName)

	// read file name
	name := strings.TrimSuffix(fName, ext)

	name = MD5V([]byte(name))
	// new filename
	return name + "_" + time.Now().Format("20060102150405") + ext
 }



// APK the value of 200 MB  == 200000000
// Firmware the value of 1.5 GB  == 1500000000


// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
// func DownloadFile(filepath string, url string) (string,error) {

// 	// Get the data
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	// Create the file
// 	out, err := os.Create(filepath)
// 	if err != nil {
// 		return "",err
// 	}
// 	defer out.Close()

// 	fileName := GetFileName(url)
// 	newPath := filepath + fileName
// 	// Write the body to file
// 	_, err = io.Copy(out, resp.Body)
// 	return newPath, err
// }