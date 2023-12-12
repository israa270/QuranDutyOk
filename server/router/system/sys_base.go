package system

import (
	// v1 "github.com/ebedevelopment/next-gen-tms/server/api/v1/system"
	"github.com/ebedevelopment/next-gen-tms/server/api/v1/system/base"
	"github.com/gin-gonic/gin"
)

// BaseRouter struct
type BaseRouter struct{
	baseApi base.BaseApi
}

// InitBaseRouter init base router
func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("login", s.baseApi.Login)
		baseRouter.POST("studentLogin", s.baseApi.StudentLogin)
		baseRouter.POST("teacherLogin", s.baseApi.TeacherLogin)
		baseRouter.GET("captcha", s.baseApi.Captcha)
	}
	
	return baseRouter
}
