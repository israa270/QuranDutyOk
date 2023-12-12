package management

import "github.com/ebedevelopment/next-gen-tms/server/global"

type Class struct {
	global.GvaModel

	Name        string `json:"name" gorm:"column:name"`
	VersionName string `json:"versionName" gorm:"column:version_name"`

	StudentCount int `json:"count" gorm:"column:count"`

	TeacherID uint    `json:"teacherID" gorm:"column:teacher_id"`
	Teacher   Teacher `json:"-" gorm:"foreignKey:TeacherID;references:ID;comment:teacher class"`
	CreatedBy string  `json:"createdBy" gorm:"column:created_by"`
}

func (Class) TableName() string {
	return "class"
}
