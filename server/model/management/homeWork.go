package management

import "github.com/ebedevelopment/next-gen-tms/server/global"



type HomeWork struct{
	global.GvaModel
	
	Name string  `json:"name" gorm:"column:name"`
    Description string `json:"description" gorm:"column:description"`

	ExpireDate  int64  `json:"expireDate" gorm:"column:expire_date"`
	ExpireStatus bool  `json:"expireStatus" gorm:"column:expire_status"`

	CreatedBy   string  `json:"createdBy" gorm:"column:created_by"`
    UpdatedBy   string  `json:"updatedBy" gorm:"column:updated_by"`
}

//TODO: Assign Homework To Classes

func (HomeWork) TableName() string {
	return "homework"
}