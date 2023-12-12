package common

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/ebedevelopment/next-gen-tms/server/global"
)

func  ValidateScreenShot(url string) error {
	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())
	// open the uploaded file
	file, err := os.Open(url)
	if err != nil {
		return err
	}

	buff := make([]byte, 512) // why 512 bytes
	_, err = file.Read(buff)
	if err != nil {
		return err
	}

	filetype := http.DetectContentType(buff)

	if filetype == "image/jpeg" || filetype == "image/jpg" || filetype == "image/png" {
		return nil
	}

	return fmt.Errorf(global.Translate("application.screenShotFormat"))
}