package management

import (
	"net/http"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	model "github.com/ebedevelopment/next-gen-tms/server/model/management"
	manReq "github.com/ebedevelopment/next-gen-tms/server/model/management/request"
	"github.com/ebedevelopment/next-gen-tms/server/service/management"
	claimcase "github.com/ebedevelopment/next-gen-tms/server/usecase/user/claim"
	"github.com/ebedevelopment/next-gen-tms/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TeacherController struct {
	teacherService management.TeacherService
}

func (m *TeacherController) CreateTeacher(name string, c *gin.Context) {
	//Get Username Created by
	tokenData, err := claimcase.GetBaseClaim(c)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GetOrganizationFail, zap.Error(err))
		response.FailWithDetailed(err.Error(), global.Translate("organization.getOrganizationFail"), http.StatusUnauthorized, "error", c)
		return
	}
	
	if tokenData.Role != utils.Admin{
       response.FailWithMessage("you not have permission", http.StatusUnauthorized, "error", c)
	   return 
	}
	// Check Teacher name
	if m.teacherService.CheckTeacherName(name) {
		global.GvaLog.Debug(global.GvaLoggerMessage["log"].DuplicateValueName)
		response.FailWithMessage(global.Translate("general.duplicateValueName"), http.StatusConflict, "warning", c)
		return
	}

	

	TeacherDB := model.Teacher{
        Name: name,
		Role: utils.Teacher,
		CreatedBy: tokenData.Username,
	}

	if err := m.teacherService.CreateTeacher(TeacherDB); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].CreationFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFail"), http.StatusInternalServerError, "error", c)
	} else {
		response.OkWithMessage(global.Translate("general.createSuccess"), http.StatusOK, "success", c)
	}
}




func (m *TeacherController) GetTeacherList(info manReq.ListSearch, c *gin.Context){
	if list, total, err := m.teacherService.GetTeacherList(info); err != nil {
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