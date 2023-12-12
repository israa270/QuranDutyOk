package user

import (
	"net/http"
    "github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	"github.com/gin-gonic/gin"
)

// Logout
// @Tags Base
// @Summary User logout
// @Produce  application/json
// @Success 200 {object}  response.Response{} "return includes user info,token,expiration"
// @Failure 500 {object}  response.Response{}
// @Router /user/logout [get]
func (u *UserApi) Logout(c *gin.Context) {

	if global.GvaDB == nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].DB)
		response.FailWithMessage(global.Translate("sysInitDB.db"), http.StatusServiceUnavailable, "warning", c)
		return
	}
    
	u.userController.Logout(c)
}