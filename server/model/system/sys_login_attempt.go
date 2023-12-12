package system

import (
	// "time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
)

// SysLoginAttempt store login attempt
type SysLoginAttempt struct {
	global.GvaModel
	Email         string     `gorm:"column:email"`
	IpAddress     string     `gorm:"column:ip_address"`
	UserAgent     string     `gorm:"column:user_agent"`
	// LastLoginTime time.Time  `gorm:"last_login_time"`
	LastLoginTime int64  `gorm:"last_login_time"`
	Count         int        `gorm:"column:count"`
	// BlockTime     *time.Time `gorm:"block_time"`
	BlockTime     int64      `gorm:"block_time"`
	IsBlocked     bool       `gorm:"is_blocked"`
}

// TableName SysLoginAttempt table Name
func (SysLoginAttempt) TableName() string {
	return "sys_login_attempts"
}

// // SysIpAddress struct
// type SysIpAddress struct {
// 	global.GvaModel
// 	IP        string     `gorm:"column:ip_address"`
// 	Count     int        `gorm:"column:count"`
// 	BlockTime *time.Time `gorm:"column:block_time"`   //Block time for count attempts increase count in config
// }

// // TableName SysIpAddress table Name
// func (SysIpAddress) TableName() string {
// 	return "sys_login_ip_address"
// }
