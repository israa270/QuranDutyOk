package base

import (
	"fmt"
	"net/http"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	sysResp "github.com/ebedevelopment/next-gen-tms/server/model/system/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

// when enabling multiServer deployment，use of configure，use redis storage code
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

// GenerateCaptcha generate captcha
func (b *BaseController) GenerateCaptcha() (captchaResp sysResp.SysCaptchaResponse, err error) {

	driver := base64Captcha.NewDriverDigit(global.GvaConfig.Captcha.ImgHeight, global.GvaConfig.Captcha.ImgWidth, global.GvaConfig.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   //use redis
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		return captchaResp, err
	} else {
		return sysResp.SysCaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: global.GvaConfig.Captcha.KeyLong,
		}, nil
	}
}

// Verify captcha's answer directly
func (b *BaseController) Verify(id, answer string, clear bool) bool {
	return store.Verify(id, answer, clear)
}




// generate captcha for login fail
func (b *BaseController) GenerateCaptchaResp(msg string, c *gin.Context) {
	if captchaResp, err := b.GenerateCaptcha(); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].LoginFailWithCaptcha, zap.Error(err))
		response.FailWithMessage(msg, http.StatusInternalServerError, "error", c)
	} else {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].LoginFail, zap.Error(fmt.Errorf(msg)))
		response.FailWithDetailed(sysResp.SysCaptchaResponse{
			CaptchaId:     captchaResp.CaptchaId,
			PicPath:       captchaResp.PicPath,
			CaptchaLength: global.GvaConfig.Captcha.KeyLong,
			LoginStatus:   "captcha",
		}, msg, http.StatusInternalServerError, "error", c)
	}
}