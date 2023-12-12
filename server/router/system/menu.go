package system

import (

"github.com/gin-gonic/gin"
v1 "github.com/ebedevelopment/next-gen-tms/server/api/v1/system"
)

type MenuRouter struct {
	menuApi v1.MenuApi
}

// InitUserRouter init user router
func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	// userRouter := Router.Group("menu").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("menu")
	{
		menuRouterWithoutRecord.GET("getMenu", s.menuApi.GetMenu)

	}
}
