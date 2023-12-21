package management

import "github.com/ebedevelopment/next-gen-tms/server/global"


type Student  struct{
	global.GvaModel
    
	Name string   `json:"name"  gorm:"column:name"`
	Role string   `json:"role"  gorm:"column:role"`

	ClassID uint  `json:"classID" gorm:"column:class_id"`
	Class  Class  `json:"-" gorm:"foreignKey:ClassID;references:ID;comment:student class"`

	// HomeWork []HomeWork  `json:"-"  gorm:"many2many:student_homework;ForeignKey:id;References:id"`
	
	UserName  string `json:"userName" gorm:"column:username"`
	Password  string  `json:"-" gorm:"column:password"`

	LastLoginTime string    `json:"lastLoginTime" gorm:"column:last_login_time"`

	CreatedBy string `json:"createdBy" gorm:"column:created_by"`
	UpdatedBy string `json:"updatedBy" gorm:"column:updated_by"`
}

func (Student) TableName() string {
	return "student"
}