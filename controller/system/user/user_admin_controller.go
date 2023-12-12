package user

import (
	"net/http"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
	claimcase "github.com/ebedevelopment/next-gen-tms/server/usecase/user/claim"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *UserController) GetUserById(userId int, c *gin.Context) {

	if user, err := u.userService.GetUserInfo(uint(userId)); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GetDataFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFail"), http.StatusInternalServerError, "error", c)
	} else {
		response.OkWithDetailed(user, global.Translate("general.getDataSuccess"), http.StatusOK, "success", c)
	}
}



func (u *UserController) DeleteUser(userID int, c *gin.Context) {

	tokenData, err := claimcase.GetBaseClaim(c)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].UserData, zap.Error(err))
		response.FailWithDetailed(err.Error(), global.Translate("general.userData"), http.StatusUnauthorized, "error", c)
		return
	}
	// not delete yourself
	if tokenData.ID == uint(userID) {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].DeleteUserFail)
		response.FailWithMessage(global.Translate("sysUser.deleteUserFail"), http.StatusConflict, "error", c)
		return
	}

	//get user email to force logout this user
	_, exist := u.userService.GetUserEmail(uint(userID))
	if !exist {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].UserExist)
		response.FailWithMessage(global.Translate("sysUser.userExist"), http.StatusNotFound, "error", c)
		return
	}

	if err := u.userService.DeleteUser(userID); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].DeleteFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.deleteFail"), http.StatusInternalServerError, "error", c)
	} else {
		response.OkWithMessage(global.Translate("general.deleteSuccess"), http.StatusOK, "success", c)
	}
}



func (u *UserController) UpdateUserStatus(userEmail string, c *gin.Context) {
	//check Email exist or not
	userExist, err := u.userService.GetUserByEmail(userEmail)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].EmailNotFound, zap.Error(err))
		response.FailWithMessage(global.Translate("sysUser.emailNotFound"), http.StatusNotFound, "warning", c)
		return
	}

	//user not able to disable or enable self
	// tokenData, err := claimcase.GetBaseClaim(c)
	// if err != nil {
	// 	global.GvaLog.Error(global.GvaLoggerMessage["log"].UserData, zap.Error(err))
	// 	response.FailWithDetailed(err.Error(), global.Translate("general.userData"), http.StatusUnauthorized, "error", c)
	// 	return
	// }

	// only admin can disable user
	// if tokenData.Email == userEmail {
	// 	global.GvaLog.Error(global.GvaLoggerMessage["log"].DisableYourSelf)
	// 	response.FailWithMessage(global.Translate("sysUser.disableYourSelf"), http.StatusConflict, "warning", c)
	// 	return
	// }

	if userExist.Status != utils.StatusActive && userExist.Status != utils.StatusDisable {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].UserNotRegister)
		response.FailWithMessage(global.Translate("sysUser.userNotRegister"), http.StatusConflict, "warning", c)
		return
	}

	// userExist.Status = "disable"
	userExist.UserStatus = !userExist.UserStatus
	if userExist.UserStatus {
		userExist.Status = utils.StatusActive
	} else {
		userExist.Status = utils.StatusDisable
	}

	// userExist.UpdatedBy = tokenData.Email
	if err := u.userService.UpdateUserStatus(userExist); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].DisableUserFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.disableUserFail"), http.StatusInternalServerError, "error", c)
	} else {
		// if userExist.Status == utils.StatusDisable {
		// 	// Force Logout this User
		// 	// if global.GvaConfig.System.UseMultiPoint && global.GvaRedis != nil {
		// 	// 	RedisJwtToken, err := u.jwtService.GetRedisJWT(userExist.Email)
		// 	// 	if err != nil {
		// 	// 		global.GvaLog.Error("get redis jwt failed", zap.Error(err))
		// 	// 	} else { // when of before of get success time just enter row json operations
		// 	// 		if err = u.jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken}); err != nil{
		// 	// 		    global.GvaLog.Error("failed to create key token in blacklist ", zap.Error(err))
		// 	// 		}
		// 	// 	}
		// 	// }
		// }

		response.OkWithMessage(global.Translate("general.disableUser"), http.StatusOK, "success", c)
	}
}

func (u *UserController) GetUserList(info sysReq.UserSearch, c *gin.Context) {

	if list, total, err := u.userService.GetUserInfoList(info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GetDataFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFail"), http.StatusInternalServerError, "error", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     info.Page,
			PageSize: info.PageSize,
		}, global.Translate("general.getDataSuccess"), http.StatusOK, "success", c)
	}
}

func (u *UserController) ExportUsersExcel(info sysReq.UserSearch, c *gin.Context) {


	if list, _, err := u.userService.GetUserInfoList(info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GetDataFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFail"), http.StatusInternalServerError, "error", c)
	} else {
		if len(list) == 0 {
			response.FailWithMessage(global.Translate("general.emptyExportData"), http.StatusBadRequest, "error", c)
			return
		}

		if _, err := u.userService.ParseUsersInfoList2Excel(list, info.Format); err != nil {
			global.GvaLog.Error(global.GvaLoggerMessage["log"].ExcelFail, zap.Error(err))
			response.FailWithMessage(global.Translate("excel.excelFail"), http.StatusInternalServerError, "error", c)
			return
		} else {
			// 	c.Writer.Header().Add("success", "true")
			// 	c.File(filePath)

		
	

			response.FailWithMessage("failed to get file", http.StatusInternalServerError, "error", c)

			// fileUrl = fileUrl[1:]
			// response.OkWithDetailed(gin.H{"fileUrl": global.GvaConfig.System.ServerPath + fileUrl}, global.Translate("general.getDataSuccess"), http.StatusOK, "success", c)
		}
	}
}

