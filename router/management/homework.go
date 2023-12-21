package management



import (
	"github.com/ebedevelopment/next-gen-tms/server/middleware"
	v1 "github.com/ebedevelopment/next-gen-tms/server/api/v1/management"
	"github.com/gin-gonic/gin"
)

// HomeWorkRouter
type HomeWorkRouter struct {
	homeWorkApi v1.HomeWorkApi
}

// InitHomeWorkRouter initialization HomeWork routing info
func (s *HomeWorkRouter) InitHomeWorkRouter(Router *gin.RouterGroup) {
    homeWorkRouter := Router.Group("homeWork").Use(middleware.OperationRecord())
	homeWorkRouterWithoutRecord := Router.Group("homeWork")
	{
		homeWorkRouter.POST("createHomeWork", s.homeWorkApi.CreateHomeWork)  
		homeWorkRouter.PUT("assignHomeWorkToClass", s.homeWorkApi.AssignHomeWorkToClass)
		 
	}
	{
		homeWorkRouterWithoutRecord.GET("getHomeWorkList", s.homeWorkApi.GetHomeWorkList)  

	}
	
}