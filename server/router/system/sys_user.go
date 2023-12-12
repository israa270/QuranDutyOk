package system

import (
	v1 "github.com/ebedevelopment/next-gen-tms/server/api/v1/system/user"
	"github.com/ebedevelopment/next-gen-tms/server/middleware"
	"github.com/gin-gonic/gin"
)

// UserRouter struct
type UserRouter struct{
	userApi v1.UserApi
}

// InitUserRouter init user router
func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	{
		userRouter.POST("register", s.userApi.Register)                //  user register
		userRouter.DELETE("deleteUser/:id", s.userApi.DeleteUser)      // delete User
		userRouter.PUT("setUserInfo", s.userApi.SetUserInfo)           // setup user info
		userRouter.PUT("disableUser/:email", s.userApi.UpdateUserStatus) //status Activate
        userRouter.GET("exportUsersExcel", s.userApi.ExportUsersExcel)
	}
	{
		userRouterWithoutRecord.GET("getUserList", s.userApi.GetUserList) // paging User list
		userRouterWithoutRecord.GET("getUserInfo", s.userApi.GetUserInfo) // get User info
		userRouterWithoutRecord.GET("getUserById/:id", s.userApi.GetUserById)
		userRouterWithoutRecord.GET("logout", s.userApi.Logout)
	}
}
