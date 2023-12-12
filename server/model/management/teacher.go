package management

import "github.com/ebedevelopment/next-gen-tms/server/global"

type Teacher struct {
	global.GvaModel

	Name string `json:"name"  gorm:"column:name"`
	Role string `json:"role"  gorm:"column:role"`

	Class []Class `json:"-"  gorm:"many2many:teacher_class;ForeignKey:id;References:id"`

	UserName string `json:"userName" gorm:"column:username"`
	Password string `json:"-" gorm:"column:password"`

	LastLoginTime string    `json:"lastLoginTime" gorm:"column:last_login_time"`

	CreatedBy string `json:"createdBy" gorm:"column:created_by"`
}

func (Teacher) TableName() string {
	return "teacher"
}
