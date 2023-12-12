package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetUserById
// @Tags SysUser
// @Summary get user info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id path int true "id"
// @Success 200 {object}  response.Response{data=map[string]interface{},msg=string} "get user info"
// @Failure 500 {object}  response.Response{}
// @Router /user/getUserById/{id} [get]
func (u *UserApi) GetUserById(c *gin.Context) {
	//TODO: admin in -> user in org or subOrg  who see this api
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "warning", c)
		return
	}

	u.userController.GetUserById(userId, c)
}



// DeleteUser
// @Tags SysUser
// @Summary deleteUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data path string true "User ID"
// @Success 200 {object}  response.Response{} "deleteUser"
// @Failure 500 {object}  response.Response{}
// @Router /user/deleteUser/{id}   [delete]
func (u *UserApi) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "warning", c)
		return
	}

	u.userController.DeleteUser(userId, c)
}


// DisableUser
// @Tags SysUser
// @Summary User update password
// @Security ApiKeyAuth
// @Produce  application/json
// @Param email path string true "email"
// @Success 200 {object}  response.Response{} "User update password"
// @Failure 500 {object}  response.Response{}
// @Router /user/disableUser/{email} [put]
func (u *UserApi) UpdateUserStatus(c *gin.Context) {
	userEmail := c.Param("email")
	u.userController.UpdateUserStatus(userEmail, c)
}

// GetUserList
// @Tags SysUser
// @Summary pagingUser list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query  sysReq.UserSearch false "user search"
// @Failure 500 {object}  response.Response{}
// @Router /user/getUserList [get]
func (u *UserApi) GetUserList(c *gin.Context) {
	var info sysReq.UserSearch
	if err := c.ShouldBindQuery(&info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "error", c)
		return
	}

	u.userController.GetUserList(info, c)
}

// ExportUsersExcel
// @Tags SysUser
// @Summary exportUsersExcel
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/json
// @Param data query sysReq.UserSearch false "export User Excel file info"
// @Success 200
// @Failure 500 {object}  response.Response{}
// @Router /user/exportUsersExcel [get]
func (u *UserApi) ExportUsersExcel(c *gin.Context) {

	var info sysReq.UserSearch
	if err := c.ShouldBindQuery(&info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "error", c)
		return
	}

	if info.Format != utils.Excel && info.Format != utils.PDF {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(fmt.Errorf("format not pdf or excel")))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "error", c)
		return
	}

	u.userController.ExportUsersExcel(info, c)
}

