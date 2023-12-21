package management

import (
	"net/http"
	"time"

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

type HomeWorkController struct {
	homeWorkService   management.HomeWorkService
	classService      management.ClassService
	studentHomeworkService management.StudentHomeWorkService
	studentService    management.StudentService
}

func (m *HomeWorkController) CreateHomeWork(homeWork model.HomeWork, c *gin.Context) {
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

     //Validate Expire Date
	 expireTime := time.Unix(0, homeWork.ExpireDate)
	 if expireTime.Before(time.Now()){
		global.GvaLog.Error("expire date must be after now")
		response.FailWithMessage("expire date must be after now", http.StatusInternalServerError, "error", c)
		return
	 }

	homeWork.CreatedBy = tokenData.Username

	if err := m.homeWorkService.CreateHomeWork(homeWork); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].CreationFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFail"), http.StatusInternalServerError, "error", c)
	} else {
		response.OkWithMessage(global.Translate("general.createSuccess"), http.StatusOK, "success", c)
	}
}


func (m *HomeWorkController) GetHomeWorkList(info manReq.HomeWorkSearch, c *gin.Context){
	if list, total, err := m.homeWorkService.GetHomeWorkList(info); err != nil {
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


func (m *HomeWorkController) AssignHomeWorkToClass(info manReq.AssignHomeWorkToClassesDTO, c *gin.Context){
        
	//get Homework data

    homework, err := m.homeWorkService.GetHomeWorkID(info.HomeWorkId)
	if err != nil{
		global.GvaLog.Error(global.GvaLoggerMessage["log"].IdNotFound, zap.Error(err))
		response.FailWithMessage(global.Translate("general.idNotFound"), http.StatusNotFound, "error", c)
		return
	}

	var studentHomeWork []model.StudentHomeWorks
	for _, id := range info.ClassIds {
		if class, err := m.classService.GetClassID(id); err != nil {
			global.GvaLog.Error(global.GvaLoggerMessage["log"].IdNotFound, zap.Error(err))
			response.FailWithMessage(global.Translate("general.idNotFound"), http.StatusNotFound, "error", c)
			return
		} else {
			homework.Class = append(homework.Class, class)

			var infoStudent manReq.StudentSearch
			infoStudent.ClassId = id
			//get all students in class
			if list, _, err := m.studentService.GetStudentList(infoStudent); err != nil {
				global.GvaLog.Error(global.GvaLoggerMessage["log"].GetDataFail, zap.Error(err))
			}else{
				if len(list) != 0{
					for _, student := range list {
                       stHomework := model.StudentHomeWorks{
						StudentId: student.ID,
						HomeworkId: info.HomeWorkId,
					   }
					   studentHomeWork = append(studentHomeWork, stHomework)
					}
				}
			}
		}
	}

	if err := m.homeWorkService.AssignHomeWorkToClass(homework); err != nil{
		global.GvaLog.Error(global.GvaLoggerMessage["log"].CreationFail, zap.Error(err))
		response.FailWithMessage("fail to assign homework", http.StatusNotFound, "error", c)
		return
	}

    
	// Create Homework Student to be able to change homework status
	if len(studentHomeWork) != 0{
     	if err := m.studentHomeworkService.CreateStudentHomeworks(studentHomeWork); err != nil{
			global.GvaLog.Error(global.GvaLoggerMessage["log"].CreationFail, zap.Error(err))
			response.FailWithMessage("fail to assign homework to student", http.StatusNotFound, "error", c)
			return
		}

	}
 
	response.OkWithMessage("assign homework success", http.StatusOK,"success", c )
}
