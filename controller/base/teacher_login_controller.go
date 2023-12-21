package base

import (
	"net/http"
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	common "github.com/ebedevelopment/next-gen-tms/server/model/common/request"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	model "github.com/ebedevelopment/next-gen-tms/server/model/management"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (b *BaseController) TeacherLogin(l sysReq.LoginDTO, c *gin.Context) {
	//put it in every fn because if pass from token can make panic
	if global.GvaDB == nil  {  //|| global.GvaDB == nil
		response.FailWithMessage(global.Translate("sysInitDB.db"), http.StatusServiceUnavailable, "warning", c)
		return
	}

	ipAddress := c.ClientIP()
	userAgent := c.Request.UserAgent()

	loginAttObj, err := b.loginAttemptService.CheckUserBlock(l.UserName)
	if err == nil {
		if loginAttObj.IsBlocked {
			if loginAtObj, block := b.UserUseCase.UpdateCounterUserBlock(&loginAttObj, c); !block {
				return
			} else {
				loginAttObj.Count = loginAtObj.Count
			}
		}
	}

	loginAttObj.Email = l.UserName
	loginAttObj.IpAddress = ipAddress
	loginAttObj.UserAgent = userAgent

	//check Email exist or not
	if !b.teacherService.CheckTeacherByUserName(l.UserName) {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].EmailNotFound, zap.Error(err))

		if _, block := b.UserUseCase.UpdateCounterUserBlock(&loginAttObj, c); !block {
			return
		}

		//Generate Captcha
		b.GenerateCaptchaResp(global.Translate("sysUser.emailNotFound"), c)
		return
	}

	//For return Captcha with login
	if err == nil && loginAttObj.Count != 0 {
		lastTime := time.Unix(loginAttObj.LastLoginTime, 0)
		if time.Now().Before(lastTime) {

			if l.Captcha == "" && l.CaptchaId == "" {
				response.FailWithMessage(global.Translate("sysCaptcha.vCodeFail"), 500, global.Translate("sysCaptcha.vCodeFail"), c)
				return
			}
			if !b.Verify(l.CaptchaId, l.Captcha, true) {
			
				//Update Counter in login Attempt
				if _, block := b.UserUseCase.UpdateCounterUserBlock(&loginAttObj, c); !block {
					return
				}

				//Generate Captcha
				b.GenerateCaptchaResp(global.Translate("sysCaptcha.vCodeFail"), c)
				return
			}
		}
	}

	u := &model.Teacher{UserName: l.UserName, Password: l.Password}
	// if !userExist.Ldap {
	if user, err := b.teacherService.Login(u); err != nil {
		//Update Counter in login Attempt
		if _, block := b.UserUseCase.UpdateCounterUserBlock(&loginAttObj, c); !block {
			return
		}

		//Generate Captcha
		b.GenerateCaptchaResp(global.Translate("sysUser.userNameOrPasswordError"), c)
		return
	} else {

		u := &common.UserToken{
			ID: user.ID,
			UserName: user.UserName,
			Role: user.Role,
			LastLoginTime: user.LastLoginTime,
		 }
		 //Check User Login Before to prevent login more than time from two browser
		 b.tokenNext(c, *u)
	}
}

