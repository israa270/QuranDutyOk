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

type TeacherApi struct {
	teacherController management.TeacherController
}

// CreateTeacher createTeacher
// @Tags Teacher
// @Summary createTeacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tms.TeacherDTO true "createTeacher"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Failure 400 {object} response.Response "{"code": 7,"message": "Bad Request","result": {},"type": "fail"}"
// @Failure 409 {object} response.Response "{"code": 7,"message": "conflict data","result": {},"type": "fail"}"
// @Failure 500 {object} response.Response "{"code": 7,"message": "create fail","result": {},"type": "fail"}"
// @Router /teacher/createTeacher [post]
func (m *TeacherApi) CreateTeacher(c *gin.Context) {
	var teacher model.CreateTeacherDTO

	if err := c.ShouldBindJSON(&teacher); err != nil {
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

	m.teacherController.CreateTeacher(teacher.Name, c)

}


// GetTeacherList
// @Tags Teacher
// @Summary GetTeacherList
// @Security ApiKeyAuth
// @accept Application/json
// @Produce Application/json
// @Param data query  tmsReq.TeacherSearch false "TeacherList"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Failure 500
// @Router /teacher/getTeacherList [get]
func (m *TeacherApi) GetTeacherList(c *gin.Context) {
	var info model.ListSearch
	if err := c.ShouldBindQuery(&info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "error", c)
		return
	}

	m.teacherController.GetTeacherList(info, c)
}


