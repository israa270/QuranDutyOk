package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	modelReq "github.com/ebedevelopment/next-gen-tms/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


// UpdateCounterUserBlock update counter user block
func (b *UserUseCase) UpdateCounterUserBlock(loginAtObj *modelReq.SysLoginAttempt, c *gin.Context) (*modelReq.SysLoginAttempt,bool) {

	userBlockTime, err := time.ParseDuration(global.GvaConfig.Login.UserBlockTime)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].ParseDuration, zap.Error(err))
	}
	
	blockTime := time.Unix(loginAtObj.BlockTime, 0)

	tryTime, err := time.ParseDuration(global.GvaConfig.Login.LoginAttemptCountTime)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].ParseDuration, zap.Error(err))
	}

	lastLoginTime  := time.Unix(loginAtObj.LastLoginTime, 0)
	
	if loginAtObj.IsBlocked{
		if time.Now().Before(blockTime) {
			response.FailWithDetailed(gin.H{"message": global.Translate("sysUser.userBlock") + fmt.Sprintf("  %v Seconds", int64(userBlockTime/time.Second))}, "block", http.StatusForbidden, "warning", c)
			return loginAtObj,false
		}
	    
		// Unblock the user
		loginAtObj.Count = 0
		loginAtObj.BlockTime = 0
		loginAtObj.IsBlocked = false
        loginAtObj.LastLoginTime = 0  
		//Store in db
		if err := b.loginAttemptService.UpdateUserLoginAttempt(loginAtObj); err != nil{
			global.GvaLog.Error("failed to save login attempt ", zap.Error(err))
		}

		return  loginAtObj , true
	}else if loginAtObj.Count >= global.GvaConfig.Login.UserBlockNum {
		now := time.Now()
		if now.After(lastLoginTime) {
			difference := now.Sub(lastLoginTime)
			if difference > tryTime{
                // Block the user for 60 seconds
				loginAtObj.BlockTime = 0
				loginAtObj.LastLoginTime = time.Now().Unix()
				loginAtObj.IsBlocked = false
				loginAtObj.Count = 1
				//Store in db
				if err := b.loginAttemptService.UpdateUserLoginAttempt(loginAtObj); err != nil {
					global.GvaLog.Error("failed to save login attempt ", zap.Error(err))
				}
				return loginAtObj, true
			}
		}

       // Block the user for 60 seconds
	   loginAtObj.BlockTime = time.Now().Add(userBlockTime).Unix()
	   loginAtObj.LastLoginTime = time.Now().Unix()
	   loginAtObj.IsBlocked = true
		loginAtObj.Count += 1 
         //Store in db
		if err := b.loginAttemptService.UpdateUserLoginAttempt(loginAtObj); err != nil{
			global.GvaLog.Error("failed to save login attempt ", zap.Error(err))
		}
		response.FailWithDetailed(gin.H{"message": global.Translate("sysUser.userBlock") + fmt.Sprintf("   %v Seconds", int64(userBlockTime/time.Second))}, "block", http.StatusForbidden, "warning", c)
		return loginAtObj,false     
	}
	
	loginAtObj.Count += 1
	loginAtObj.LastLoginTime = time.Now().Unix()
	// Store in db 
	if err := b.loginAttemptService.UpdateUserLoginAttempt(loginAtObj); err != nil{
		global.GvaLog.Error("failed to save login attempt ", zap.Error(err))
	}


	return loginAtObj,true
}
