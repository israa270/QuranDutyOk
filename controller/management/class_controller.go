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

type ClassController struct {
	classService   management.ClassService
	teacherService management.TeacherService
	homeworkService management.HomeWorkService
}

func (m *ClassController) CreateClass(class model.Class, c *gin.Context) {
	//Get Username Created by
	tokenData, err := claimcase.GetBaseClaim(c)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].UserFail, zap.Error(err))
		response.FailWithDetailed(err.Error(), global.Translate("general.userFail"), http.StatusUnauthorized, "error", c)
		return
	}

	if tokenData.Role != utils.Admin {
		response.FailWithMessage("you not have permission", http.StatusUnauthorized, "error", c)
		return
	}

	// Check Class name
	if m.classService.CheckClassName(class.Name, class.VersionName) {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].DuplicateValueName)
		response.FailWithMessage(global.Translate("general.duplicateValueName"), http.StatusConflict, "warning", c)
		return
	}

	//Check Teacher already exist
	if !m.teacherService.CheckTeacherExist(class.TeacherID) {
		global.GvaLog.Error("teacher is not exist")
		response.FailWithMessage("teacher is not exist", http.StatusNotFound, "warning", c)
		return
	}

	class.CreatedBy = tokenData.Username

	if err := m.classService.CreateClass(class); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].CreationFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFail"), http.StatusInternalServerError, "error", c)
	} else {
		response.OkWithMessage(global.Translate("general.createSuccess"), http.StatusOK, "success", c)
	}
}


func (m *ClassController) GetClassList(info manReq.ClassSearch, c *gin.Context){
	if list, total, err := m.classService.GetClassList(info); err != nil {
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

func (m *ClassController) GetClassHomework(info manReq.GetHomeWorkQuery, c *gin.Context){
	if classHomeWork, err := m.homeworkService.GetClassHomework(info.ClassId); err != nil{
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GetDataFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFail"), http.StatusInternalServerError, "error", c)
		return
	}else{
		var homeworks []model.HomeWork
       for _, class := range classHomeWork{
            //get homework Data
			if homework , err := m.homeworkService.GetHomeWorkID(class.HomeWorkId); err != nil{
				global.GvaLog.Error(global.GvaLoggerMessage["log"].IdNotFound, zap.Error(err))
				response.FailWithMessage(global.Translate("general.idNotFound"), http.StatusNotFound, "error", c)
				return
			}else{
				homeworks = append(homeworks, homework)
			}
	   }

	   response.OkWithDetailed(homeworks, "get homework success", http.StatusOK, "success", c)
	   
	}
}

