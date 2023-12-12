package management

import (
	"github.com/ebedevelopment/next-gen-tms/server/middleware"
	v1 "github.com/ebedevelopment/next-gen-tms/server/api/v1/management"
	"github.com/gin-gonic/gin"
)

// TeacherRouter
type TeacherRouter struct {
	TeacherApi v1.TeacherApi
}

// InitTeacherRouter initialization Teacher routing info
func (s *TeacherRouter) InitTeacherRouter(Router *gin.RouterGroup) {
	TeacherRouter := Router.Group("teacher").Use(middleware.OperationRecord())
	TeacherRouterWithoutRecord := Router.Group("teacher")
	{
		TeacherRouter.POST("createTeacher", s.TeacherApi.CreateTeacher)  
	}
	{
		TeacherRouterWithoutRecord.GET("getTeacherList", s.TeacherApi.GetTeacherList)
	
	}
}