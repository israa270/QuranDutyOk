package user

import (
	"errors"
	"net/http"

	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/gin-gonic/gin"
)

// Register
// @Tags SysUser
// @Summary User registered account
// @Produce  application/json
// @Param data body sysReq.Register true "username, password, roleID"
// @Success 200 {object}  response.Response{} "User registered account,return includes user info"
// @Failure 400 {object}  response.Response{}
// @Failure 404 {object}  response.Response{}
// @Failure 409 {object}  response.Response{}
// @Failure 500 {object}  response.Response{}
// @Router /user/register [post]
func (u *UserApi) Register(c *gin.Context) {
	var r sysReq.Register
	if err := c.ShouldBindJSON(&r); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := utils.HandleError(ve)
			response.FailWithDetailed(out, global.Translate("general.badRequest"), http.StatusBadRequest, "warning", c)
			return
		}
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "warning", c)
		return
	}

	u.userController.Register(r,c)
}

