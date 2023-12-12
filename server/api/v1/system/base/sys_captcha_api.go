package base

import (
	"net/http"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Captcha
// @Tags Base
// @Summary generate code
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object}  response.Response{} "generate code,return includes random number id,base64,code length"
// @Failure 500 {object} response.Response "{"code": 7,"message": "Data Fail","result": {},"type": "fail"}"
// @Router /base/captcha [get]
func (b *BaseApi) Captcha(c *gin.Context) {
	if resp, err := b.baseController.GenerateCaptcha(); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].VCodeFail, zap.Error(err))
		response.FailWithMessage(global.Translate("sysCaptcha.vCodeFail"), http.StatusInternalServerError, "error", c)
	} else {
		response.OkWithDetailed(resp, global.Translate("sysCaptcha.vCodeSuccess"), http.StatusOK, "success", c)
	}
}
