package request

import "github.com/ebedevelopment/next-gen-tms/server/model/common/request"

type CreateTeacherDTO struct {
	Name string `json:"name"`
}

type ClassSearch struct {
	Name        string `json:"name" form:"name"`
	VersionName string `json:"versionName" form:"versionName"`
	TeacherId   uint   	`json:"teacherId" form:"teacherId"`
	request.PageInfo
}


type GetHomeWorkQuery struct {
	ClassId uint  `json:"classId" form:"classId"`

	//TODO: student 
	Student   bool  `json:"student" form:"student"`
	StudentId uint   `json:"studentId" form:"studentId"`
}


type ListSearch struct {
	Name     string `json:"name" form:"name"`
	UserName string `json:"userName" form:"userName"`
	request.PageInfo
}

type ClassRequest struct{
	ClassId uint `json:"classId"`
}

type StudentSearch struct{
	Name     string `json:"name" form:"name"`
	UserName string `json:"userName" form:"userName"`
	ClassId  uint    `json:"classId" form:"classId"`

	request.PageInfo
}

type StudentHomeWorkSearch struct{
	StudentId  uint `json:"studentId"`
	Status     string `json:"status"`
	request.PageInfo
}

type UpdateStudentHomeWork struct{
	StudentId uint `json:"studentId"`
	HomeWorkId uint `json:"homeWorkId"`
	HomeWorkStatus string `json:"status"`
}

type StudentDTO struct {
	Name string `json:"name"`
	ClassId uint `json:"classId"`
}


type MoveStudentDTO struct{
	StudentId  uint `json:"studentId"`
	OldClassId uint `json:"oldClassId"`
	NewClassId uint `json:"newClassId"`	
}

type AssignHomeWorkToClassesDTO struct{
    HomeWorkId   uint `json:"homeworkId"`
	ClassIds     []uint  `json:"classIds"`
}

type HomeWorkSearch struct {
	Name string `json:"name"`
	ExpireStatus bool `json:"status"`
	request.PageInfo
}


