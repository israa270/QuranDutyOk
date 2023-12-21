package management

import "github.com/ebedevelopment/next-gen-tms/server/global"



type HomeWork struct{
	global.GvaModel
	
	Name string  `json:"name" gorm:"column:name"`
    Description string `json:"description" gorm:"column:description"`

	ExpireDate  int64  `json:"expireDate" gorm:"column:expire_date"`
	ExpireStatus bool  `json:"expireStatus" gorm:"column:expire_status"`

	Class []Class `json:"-" gorm:"many2many:homework_classes;ForeignKey:id;References:id"` //home_class

	//For Response 
	StudentHomeWorkStatus string `json:"studentHomeWorkStatus,omitempty"  gorm:"-"`
	UpdatedStatus     string  `json:"UpdatedStatus,omitempty" gorm:"-"`

	CreatedBy   string  `json:"createdBy" gorm:"column:created_by"`
    UpdatedBy   string  `json:"updatedBy" gorm:"column:updated_by"`
}


func (HomeWork) TableName() string {
	return "homework"
}


type HomeworkClasses struct{
	ClassId  uint 
	HomeWorkId uint
 }