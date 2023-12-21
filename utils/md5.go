package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//MD5V
//md5 encryption
//str []byte
//string
func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
