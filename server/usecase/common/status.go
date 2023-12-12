package common

import "github.com/ebedevelopment/next-gen-tms/server/global"

// UserStatus status is active and disable
func UserStatus(status bool) string{
	if status{
		return global.GvaConfig.Login.StatusTrue
	}
	return global.GvaConfig.Login.StatusFalse
}