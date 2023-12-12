package request

import "github.com/ebedevelopment/next-gen-tms/server/model/common/request"

type NameRequest struct {
	Name string `json:"name"`
	request.PageInfo
}

type ClassSearch struct {
	Name        string `json:"name" form:"name"`
	VersionName string `json:"versionName" form:"versionName"`
	request.PageInfo
}

type ListSearch struct {
	Name     string `json:"name" form:"name"`
	UserName string `json:"userName" form:"userName"`
	request.PageInfo
}



type StudentSearch struct{
	Name     string `json:"name" form:"name"`
	UserName string `json:"userName" form:"userName"`
	ClassID  uint    `json:"classId" form:"classId"`

	request.PageInfo
}


type StudentDTO struct {
	Name string `json:"name"`
	ClassName string  `json:"className"`
	Version   string   `json:"version"`
}


type HomeWorkSearch struct {
	Name string `json:"name"`
	ExpireStatus bool `json:"status"`
	request.PageInfo
}

type MoveStudent struct{
	StudentId uint   `json:"studentId"`
	OldClassId uint  `json:"oldClassId"`
	NewClassId uint  `json:"newClassId"`
}
