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
}

func (m *StudentController) CreateStudent(student manReq.StudentDTO, c *gin.Context) {
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
	// Check Student name
	if m.studentService.CheckStudentName(student.Name) {
		global.GvaLog.Debug(global.GvaLoggerMessage["log"].DuplicateValueName)
		response.FailWithMessage(global.Translate("general.duplicateValueName"), http.StatusConflict, "warning", c)
		return
	}

	//Check Class Id 
	classID , err := m.classService.GetClassID(student.ClassName, student.Version)
	if err != nil{
		global.GvaLog.Error("class Name is not found")
		response.FailWithMessage("class Name is not found", http.StatusNotFound, "warning", c)
		return
	}

	StudentDB := model.Student{
        Name: student.Name,
		ClassID: classID,
		Role: utils.Student,
		CreatedBy: tokenData.Username,
	}

	if username, err := m.studentService.CreateStudent(StudentDB); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].CreationFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFail"), http.StatusInternalServerError, "error", c)
	} else {
		response.OkWithDetailed(gin.H{"username":username},global.Translate("general.createSuccess"), http.StatusOK, "success", c)
	}
}


func (m *StudentController) MoveStudent(student manReq.MoveStudent, c *gin.Context){
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
	// Check Student name
	if !m.studentService.CheckStudentClassExist(student.StudentId, student.OldClassId) {
		global.GvaLog.Debug(global.GvaLoggerMessage["log"].IdNotFound)
		response.FailWithMessage(global.Translate("general.idNotFound"), http.StatusNotFound, "warning", c)
		return
	}

	//check new Class id is exist 
	if !m.classService.CheckClassID(student.NewClassId){
		global.GvaLog.Error("new class Name is not found")
		response.FailWithMessage("new class Name is not found", http.StatusNotFound, "warning", c)
		return
	}

	if err := m.studentService.MoveStudent(student.StudentId, student.NewClassId); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].CreationFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFail"), http.StatusInternalServerError, "error", c)
	} else {
		response.OkWithMessage(global.Translate("general.createSuccess"), http.StatusOK, "success", c)
	}
}


func (m *StudentController) GetStudentList(info manReq.StudentSearch, c *gin.Context){
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