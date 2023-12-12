package base

import (
	"errors"
	"net/http"
	
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
	"github.com/go-playground/validator/v10"


	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Login
// @Tags Base
// @Summary User login
// @Produce  application/json
// @Param data body sysReq.Login true "username, password, code"
// @Success 200 {object}  response.Response{} "return includes user info,token,expiration"
// @Failure 400 {object}  response.Response{}
// @Failure 403 {object}  response.Response{}
// @Failure 500 {object}  response.Response{}
// @Router /base/teacherLogin [post]
func (b *BaseApi) TeacherLogin(c *gin.Context) {

	var l sysReq.LoginDTO
	if err := c.ShouldBindJSON(&l); err != nil {
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

	b.baseController.TeacherLogin(l, c)
}




