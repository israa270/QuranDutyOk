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

type StudentController struct {
	studentService management.StudentService
	classService   management.ClassService
	studentHomeWorkService  management.StudentHomeWorkService
	homeWorkService  management.HomeWorkService
}

func (m *StudentController) CreateStudent(student manReq.StudentDTO, c *gin.Context) {
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
	// Check Student name Unique in all classes
	if m.studentService.CheckStudentName(student.Name) {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].DuplicateValueName)
		response.FailWithMessage(global.Translate("general.duplicateValueName"), http.StatusConflict, "warning", c)
		return
	}

	//Check Class Id
	if !m.classService.CheckClassID(student.ClassId) {
		global.GvaLog.Error("class is not found")
		response.FailWithMessage("class is not found", http.StatusNotFound, "warning", c)
		return
	}

	StudentDB := model.Student{
		Name:      student.Name,
		ClassID:   student.ClassId,
		Role:      utils.Student,
		CreatedBy: tokenData.Username,
	}

	if err := m.studentService.CreateStudent(StudentDB); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].CreationFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFail"), http.StatusInternalServerError, "error", c)
	} else {
        //Updates Count of Student in Class


		response.OkWithMessage(global.Translate("general.createSuccess"), http.StatusOK, "success", c)
	}
}

func (m *StudentController) MoveStudent(student manReq.MoveStudentDTO, c *gin.Context) {
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
	// Check Student name
	studentExist, err := m.studentService.CheckStudentExistWithClass(student.StudentId, student.OldClassId)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].IdNotFound)
		response.FailWithMessage(global.Translate("general.idNotFound"), http.StatusNotFound, "warning", c)
		return
	}

	//Check Class Id
	if !m.classService.CheckClassID(student.NewClassId) {
		global.GvaLog.Error("class is not found")
		response.FailWithMessage("class is not found", http.StatusNotFound, "warning", c)
		return
	}

	studentExist.ClassID = student.NewClassId
	studentExist.UpdatedBy = tokenData.Username

	if err := m.studentService.MoveStudent(studentExist); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].CreationFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.moveSuccess"), http.StatusInternalServerError, "error", c)
	} else {
		response.OkWithMessage(global.Translate("general.moveSuccess"), http.StatusOK, "success", c)
	}
}

func (m *StudentController) GetStudentList(info manReq.StudentSearch, c *gin.Context) {
	if list, total, err := m.studentService.GetStudentList(info); err != nil {
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


func (m *StudentController) GetStudentHomeWork(info manReq.StudentHomeWorkSearch, c *gin.Context) {
	if list, total, err := m.studentHomeWorkService.GetStudentHomeWorks(info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GetDataFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFail"), http.StatusInternalServerError, "error", c)
	} else {

		var homeWorks  []model.HomeWork
		for _, stHomeWork := range list{
			//get homeWork info 
			if homeWork , err := m.homeWorkService.GetHomeWorkID(stHomeWork.HomeworkId); err != nil{
				global.GvaLog.Error("failed to get homeWork data", zap.Error(err))
			}else{
				homeWork.StudentHomeWorkStatus = stHomeWork.StatusChanged
				homeWork.UpdatedStatus = stHomeWork.StatusChangeDate
				
				homeWorks = append(homeWorks, homeWork)
			}
		}

		response.OkWithDetailed(response.PageResult{
			List:     homeWorks,
			Total:    total,
			Page:     info.Page,
			PageSize: info.PageSize,
		}, global.Translate("general.getDataSuccess"), http.StatusOK, "success", c)
	}

}

func (m *StudentController) UpdateStudentHomeWork(homeWorkStatus manReq.UpdateStudentHomeWork, c *gin.Context) {
    tokenData, err := claimcase.GetBaseClaim(c)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].UserFail, zap.Error(err))
		response.FailWithDetailed(err.Error(), global.Translate("general.userFail"), http.StatusUnauthorized, "error", c)
		return
	}

	if err := m.studentHomeWorkService.UpdateStudentHomeworks(homeWorkStatus, tokenData.Username); err != nil{
		global.GvaLog.Error(global.GvaLoggerMessage["log"].UpdateFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.updateFail"), http.StatusInternalServerError, "error", c)
	}else{
       response.OkWithMessage("update success", http.StatusOK, "success", c)
	}
}
