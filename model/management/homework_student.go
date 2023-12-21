package management

import "github.com/ebedevelopment/next-gen-tms/server/global"



type StudentHomeWorks struct{
	global.GvaModel


	StudentId uint `json:"studentId"  gorm:"column:student_id"`
	HomeworkId uint `json:"homeworkId" gorm:"column:home_work_id"`
	
	StatusChangeDate  string  `json:"statusChangeDate" gorm:"column:status_change_date"`
	StatusChanged      string   `json:"statusChanged" gorm:"column:status_changed"`

	HomeWorkStatus     string   `json:"homeWorkStatus" gorm:"column:home_work_status"`  //white| yellow| red

	UpdatedBy  string `json:"updatedBy" gorm:"column:updated_by"`
}


func (StudentHomeWorks) TableName() string {
	return "student_homeworks"
}

