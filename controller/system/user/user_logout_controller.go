package user

import (
	"net/http"
	// "time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	"github.com/ebedevelopment/next-gen-tms/server/model/system"
	// claimcase "github.com/ebedevelopment/next-gen-tms/server/usecase/user/claim"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *UserController) Logout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	//get Email to store in operation log
	// tokenData,err := claimcase.GetBaseClaim(c)
	// if err != nil{
	// 	global.GvaLog.Error(global.GvaLoggerMessage["log"].UserData,zap.Error(err))
	// 	response.FailWithDetailed(err.Error(),global.Translate("general.userData"),http.StatusUnauthorized, "error", c)
	// 	return
	// }

	// ipAddress := c.ClientIP()
	// userAgent := c.Request.UserAgent()

	jwt := system.JwtBlacklist{Jwt: token}
	if err := u.jwtService.JsonInBlacklist(jwt); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].UserLogoutFail, zap.Error(err))
		response.FailWithMessage(global.Translate("sysUser.UserLogoutFail"), http.StatusInternalServerError, "error", c)
	} else {
		response.OkWithMessage(global.Translate("sysUser.UserLogoutSuccess"), http.StatusOK, "success", c)
	}
}