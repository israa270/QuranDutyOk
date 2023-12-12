package utils

import (
	"fmt"
	"net/url"

	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"go.uber.org/zap"
)

// ByteCountSI size of file in string
func ByteCountSI(b int64) string {
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
func GetFileExtension(fileName string) string {
	return filepath.Ext(fileName)
}

// GetFileSize get file size
func GetFileSize(fileUrl string) (int64, error) {
	info, err := os.Stat(fileUrl)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// Check file Path exist or not
func CheckFilePath(fileUrl string) bool {
	_, err := os.Stat(fileUrl)
	return err == nil
}


// GetFileName get file name from path
func GetFileName(url string) string{
	return path.Base(url)
}

func MoveFile(source string) string{
	// Check if the source file exists
	_, err := os.Stat(source)
	if os.IsNotExist(err) {
		global.GvaLog.Error("model icon is not exist")
		return ""
	}

	parsedURL, err := url.Parse(source)
	if err != nil {
		global.GvaLog.Error("Error parsing URL: ", zap.Error(err))
		return ""
	}

	fileName := path.Base(parsedURL.Path)

	destination := "assets/icons/" +fileName
	
	err = os.Rename(source, destination)
	if err != nil {
		fmt.Println("Error moving model icon:", err)
		return ""
	}

	return destination
}

func FileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

//FileExists check file exist
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
	   return false
	}
	return !info.IsDir()
 }

 func HashFileName(fName string) string{
	// read file extension
	ext := path.Ext(fName)

	// read file name
	name := strings.TrimSuffix(fName, ext)

	name = MD5V([]byte(name))
	// new filename
	return name + "_" + time.Now().Format("20060102150405") + ext
 }
