package system

import "github.com/ebedevelopment/next-gen-tms/server/global"

type Admin struct {
	global.GvaModel

	Name     string `json:"name" gorm:"column:name"`
	Role     string `json:"role" gorm:"column:role"`
	UserName string `json:"userName" gorm:"column:username"`
	Password string `json:"-" gorm:"column:password"`

	LastLoginTime string    `json:"lastLoginTime" gorm:"column:last_login_time"`
}

func (Admin) TableName() string {
	return "admin"
}



