package base


import (
	"net/http"
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	common "github.com/ebedevelopment/next-gen-tms/server/model/common/request"
	"github.com/ebedevelopment/next-gen-tms/server/model/system"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
	claimcase "github.com/ebedevelopment/next-gen-tms/server/usecase/user/claim"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (b *BaseController) AdminLogin(l sysReq.LoginDTO, c *gin.Context) {
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
	if !b.adminService.CheckAdminByUserName(l.UserName) {
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

	u := &system.Admin{UserName: l.UserName, Password: l.Password}
	// if !userExist.Ldap {
	if user, err := b.adminService.Login(u); err != nil {
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

// tokenNext jwt
func (b *BaseController) tokenNext(c *gin.Context, user common.UserToken) {
	j := &claimcase.JWT{SigningKey: []byte(global.GvaConfig.JWT.SigningKey)} // unique signature
	claims := j.CreateClaims(sysReq.BaseClaims{
		ID:          user.ID,
		Username:    user.UserName,
		Role:        user.Role,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GetTokenFail, zap.Error(err))
		response.FailWithMessage(global.Translate("sysUser.getTokenFail"), http.StatusInternalServerError, "error", c)
		return
	}
	
	response.OkWithDetailed(common.LoginDTOResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, global.Translate("sysUser.loginSuccess"), http.StatusOK, "success", c)
}