package system

import (
	"github.com/ebedevelopment/next-gen-tms/server/global"
)

// JwtBlacklist struct
type JwtBlacklist struct {
	global.GvaModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
