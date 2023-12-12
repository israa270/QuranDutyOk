package management



import (
	"github.com/ebedevelopment/next-gen-tms/server/middleware"
	v1 "github.com/ebedevelopment/next-gen-tms/server/api/v1/management"
	"github.com/gin-gonic/gin"
)

// ClassRouter
type ClassRouter struct {
	classApi v1.ClassApi
}

// InitClassRouter initialization Class routing info
func (s *ClassRouter) InitClassRouter(Router *gin.RouterGroup) {
    classRouter := Router.Group("class").Use(middleware.OperationRecord())
	ClassRouterWithoutRecord := Router.Group("class")
	{
		classRouter.POST("createClass", s.classApi.CreateClass)  
	}
	{
		ClassRouterWithoutRecord.GET("getClassList", s.classApi.GetClassList)  

	}
	
}