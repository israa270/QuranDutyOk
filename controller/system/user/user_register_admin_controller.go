package user

import (
	"errors"
	"net/http"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
	// claimcase "github.com/ebedevelopment/next-gen-tms/server/usecase/user/claim"
    "github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

func (u *UserController) Register(r sysReq.Register, c *gin.Context) {

	// // check authority id - this role exist
	// if !u.authorityService.CheckAuthorityInfo(r.AuthorityId) {
	// 	global.GvaLog.Error(global.GvaLoggerMessage["log"].RoleExist)
	// 	response.FailWithMessage(global.Translate("sysAuthority.roleExist"), http.StatusNotFound, "error", c)
	// 	return
	// }

	//Get Username Created by
	// tokenData, err := claimcase.GetBaseClaim(c)
	// if err != nil {
	// 	global.GvaLog.Error(global.GvaLoggerMessage["log"].GetOrganizationFail, zap.Error(err))
	// 	response.FailWithDetailed(err.Error(), global.Translate("organization.getOrganizationFail"), http.StatusUnauthorized, "error", c)
	// 	return
	// }

	//check User email
	if u.userService.CheckUserByEmail(r.Email) {
		e := errors.New(global.Translate("sysUser.duplicatedEmail"))
		global.GvaLog.Error(global.GvaLoggerMessage["log"].RegistrationFail, zap.Error(e))
		response.FailWithDetailed(e, global.Translate("sysUser.registrationFail"), http.StatusConflict, "error", c)
		return
	}



	if userReturn, err := u.userService.Register(r, "" ,"tokenData.Email"); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].RegistrationFail, zap.Error(err))
		response.FailWithDetailed(err.Error(), global.Translate("sysUser.registrationFail"), http.StatusInternalServerError, "error", c)
	} else {
		response.OkWithMessage(global.Translate("mail.sendEmailSuccess")+"to "+userReturn.Email, http.StatusOK, "success", c)
	}
}
