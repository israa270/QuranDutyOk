package common

import (
	"regexp"
	"github.com/ebedevelopment/next-gen-tms/server/global"
)

func ValidatePhone(phone string) bool {
	re := regexp.MustCompile(`^\s*(?:\+?(\d{1,3}))?[-. (]*(\d{3})[-. )]*(\d{3})[-. ]*(\d{4})(?: *x(\d+))?\s*$`)
	if !re.MatchString(phone) {
		global.GvaLog.Debug(global.GvaLoggerMessage["log"].ValidatePhone)
		return false
	}
    return true
}


func ValidateEmail(email string) bool {
	reEmail := regexp.MustCompile(`^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$`)
	if !reEmail.MatchString(email) {
		global.GvaLog.Debug(global.GvaLoggerMessage["log"].ValidateEmail)
		return false
	}
    return true
}