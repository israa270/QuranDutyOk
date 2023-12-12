package user

import (
	"net/http"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	"github.com/ebedevelopment/next-gen-tms/server/model/system"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
	claimcase "github.com/ebedevelopment/next-gen-tms/server/usecase/user/claim"
	commoncase "github.com/ebedevelopment/next-gen-tms/server/usecase/common"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


func (u *UserController) SetUserInfo(user sysReq.ChangeUserInfo, c *gin.Context) {
	//Get Username Created by
	// tokenData, err := claimcase.GetBaseClaim(c)
	// if err != nil {
	// 	global.GvaLog.Error(global.GvaLoggerMessage["log"].GetOrganizationFail, zap.Error(err))
	// 	response.FailWithDetailed(err.Error(), global.Translate("organization.getOrganizationFail"), http.StatusUnauthorized, "error", c)
	// 	return
	// }

	//Check Email exist or not
	// if !u.userService.CheckUserByEmail(tokenData.Email) {
	// 	global.GvaLog.Error(global.GvaLoggerMessage["log"].TokenNotValid, zap.Error(err))

	// 	response.FailWithMessage(global.Translate("init.TokenNotValid"), http.StatusNotFound, "warning", c)
	// 	return
	// }
	//Validate On MobileNumber egyptian
	// re := regexp.MustCompile(`^01[0125][0-9]{8}$`)
	// validate any mobile phone
	if user.Phone != "" {
		if !commoncase.ValidatePhone(user.Phone) {
			response.FailWithMessage(global.Translate("general.validatePhone"), http.StatusBadRequest, "warning", c)
			return
		}
	}

	//image optional
	if !utils.CheckFilePath(user.HeaderImg) && user.HeaderImg != "" {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].ImageProfile)
		response.FailWithMessage(global.Translate("sysUser.imageProfile"), http.StatusNotFound, "error", c)
		return
	}

	if err := u.userService.SetUserInfo(system.SysUser{
		// Email:     tokenData.Email,
		Phone:     user.Phone,
	}); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].ModifyFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.modifyFail"), http.StatusInternalServerError, "error", c)
	} else {
		response.OkWithMessage(global.Translate("general.modifySuccess"), http.StatusOK, "success", c)
	}
}

func (u *UserController) GetUserInfo(c *gin.Context) {
	tokenData, err := claimcase.GetBaseClaim(c)
	if err != nil {
		global.GvaLog.Error(err.Error())
		response.FailWithMessage(err.Error(), http.StatusUnauthorized, "error", c)
		return
	}

	if global.GvaDB == nil {
		response.FailWithMessage(global.Translate("sysInitDB.db"), http.StatusServiceUnavailable, "warning", c)
		return
	}

	if user, err := u.userService.GetUserInfo(tokenData.ID); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GetDataFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFail"), http.StatusInternalServerError, "error", c)
	} else {

		response.OkWithDetailed(user, global.Translate("general.getDataSuccess"), http.StatusOK, "success", c)
	}
}
