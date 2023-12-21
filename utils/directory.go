package utils

import (
	"errors"
	"os"
)

// PathExists file content yes no exist
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("exist same name file")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
