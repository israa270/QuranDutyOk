package user

import (
	"net/http"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


// SetUserInfo
// @Tags SysUser
// @Summary setup user info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sysReq.ChangeUserInfo true "phone, headerImg"
// @Success 200 {object}  response.Response{data=map[string]interface{},msg=string} "setup user info"
// @Failure 400 {object}  response.Response{}
// @Failure 404 {object}  response.Response{}
// @Failure 500 {object}  response.Response{}
// @Router /user/setUserInfo [put]
func (u *UserApi) SetUserInfo(c *gin.Context) {

	var user sysReq.ChangeUserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "error", c)
		return
	}

	u.userController.SetUserInfo(user, c)
}

// GetUserInfo
// @Tags SysUser
// @Summary get user info
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {object}  response.Response{data=map[string]interface{},msg=string} "get user info"
// @Failure 500 {object}  response.Response{}
// @Router /user/getUserInfo [get]
func (u *UserApi) GetUserInfo(c *gin.Context) {
	u.userController.GetUserInfo(c)
}
