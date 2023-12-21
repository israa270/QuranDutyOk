package management


import (
	"errors"
	"net/http"

	"github.com/ebedevelopment/next-gen-tms/server/controller/management"
	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	model "github.com/ebedevelopment/next-gen-tms/server/model/management/request"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type StudentApi struct {
	studentController management.StudentController
}

// CreateStudent createStudent
// @Tags Student
// @Summary createStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tms.StudentDTO true "createStudent"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Failure 400 {object} response.Response "{"code": 7,"message": "Bad Request","result": {},"type": "fail"}"
// @Failure 409 {object} response.Response "{"code": 7,"message": "conflict data","result": {},"type": "fail"}"
// @Failure 500 {object} response.Response "{"code": 7,"message": "create fail","result": {},"type": "fail"}"
// @Router /student/createStudent [post]
func (m *StudentApi) CreateStudent(c *gin.Context) {
	var Student model.StudentDTO

	if err := c.ShouldBindJSON(&Student); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := utils.HandleError(ve)
			response.FailWithDetailed(out, global.Translate("general.badRequest"), http.StatusBadRequest, "warning", c)
			return
		}
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "warning", c)
		return
	}

	m.studentController.CreateStudent(Student, c)

}


// GetStudentList
// @Tags Student
// @Summary GetStudentList
// @Security ApiKeyAuth
// @accept Application/json
// @Produce Application/json
// @Param data query  tmsReq.StudentSearch false "StudentList"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Failure 500
// @Router /student/getStudentList [get]
func (m *StudentApi) GetStudentList(c *gin.Context) {
	var info model.StudentSearch
	if err := c.ShouldBindQuery(&info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "error", c)
		return
	}

	m.studentController.GetStudentList(info, c)
}

// MoveStudent MoveStudent
// @Tags Student
// @Summary MoveStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tms.StudentDTO true "MoveStudent"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Failure 400 {object} response.Response "{"code": 7,"message": "Bad Request","result": {},"type": "fail"}"
// @Failure 409 {object} response.Response "{"code": 7,"message": "conflict data","result": {},"type": "fail"}"
// @Failure 500 {object} response.Response "{"code": 7,"message": "create fail","result": {},"type": "fail"}"
// @Router /student/moveStudent [put]
func (m *StudentApi) MoveStudent(c *gin.Context) {
	var Student model.MoveStudentDTO

	if err := c.ShouldBindJSON(&Student); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := utils.HandleError(ve)
			response.FailWithDetailed(out, global.Translate("general.badRequest"), http.StatusBadRequest, "warning", c)
			return
		}
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "warning", c)
		return
	}

	m.studentController.MoveStudent(Student, c)

}


// GetStudentHomeWork
// @Tags Student
// @Summary GetStudentHomeWork
// @Security ApiKeyAuth
// @accept Application/json
// @Produce Application/json
// @Param data query  tmsReq.StudentSearch false "StudentList"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Failure 500
// @Router /student/getStudentHomeWork [get]
func (m *StudentApi) GetStudentHomeWork(c *gin.Context) {
	var info model.StudentHomeWorkSearch
	if err := c.ShouldBindQuery(&info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "error", c)
		return
	}

	m.studentController.GetStudentHomeWork(info, c)
}


// updateStudentHomeWork updateStudentHomeWork
// @Tags Student
// @Summary updateStudentHomeWork
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tms.StudentDTO true "updateStudentHomeWork"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Failure 400 {object} response.Response "{"code": 7,"message": "Bad Request","result": {},"type": "fail"}"
// @Failure 409 {object} response.Response "{"code": 7,"message": "conflict data","result": {},"type": "fail"}"
// @Failure 500 {object} response.Response "{"code": 7,"message": "create fail","result": {},"type": "fail"}"
// @Router /student/updateStudentHomeWork [put]
func (m *StudentApi) UpdateStudentHomeWork(c *gin.Context) {
	var homeWorkStatus model.UpdateStudentHomeWork

	if err := c.ShouldBindJSON(&homeWorkStatus); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := utils.HandleError(ve)
			response.FailWithDetailed(out, global.Translate("general.badRequest"), http.StatusBadRequest, "warning", c)
			return
		}
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "warning", c)
		return
	}

	m.studentController.UpdateStudentHomeWork(homeWorkStatus, c)

}
