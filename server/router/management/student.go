package management

import (
	"github.com/ebedevelopment/next-gen-tms/server/middleware"
	v1 "github.com/ebedevelopment/next-gen-tms/server/api/v1/management"
	"github.com/gin-gonic/gin"
)

// StudentRouter
type StudentRouter struct {
	studentApi v1.StudentApi
}

// InitStudentRouter initialization Student routing info
func (s *StudentRouter) InitStudentRouter(Router *gin.RouterGroup) {
	studentRouter := Router.Group("student").Use(middleware.OperationRecord())
	studentRouterWithoutRecord := Router.Group("student")
	{
		studentRouter.POST("createStudent", s.studentApi.CreateStudent)  

		studentRouter.PUT("moveStudent", s.studentApi.MoveStudent)
	}
	{
		studentRouterWithoutRecord.GET("getStudentList", s.studentApi.GetStudentList)
	
	}
}