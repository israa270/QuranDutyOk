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
// @Router /student/moveStudent [post]
func (m *StudentApi) MoveStudent(c *gin.Context) {
	var Student model.MoveStudent

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


