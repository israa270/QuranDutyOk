package management

import (
	"errors"
	"net/http"

	"github.com/ebedevelopment/next-gen-tms/server/controller/management"
	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	model "github.com/ebedevelopment/next-gen-tms/server/model/management"
	manReq "github.com/ebedevelopment/next-gen-tms/server/model/management/request"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type ClassApi struct {
	classController management.ClassController
}

// CreateClass createClass
// @Tags Class
// @Summary createClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tms.ClassDTO true "createClass"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Failure 400 {object} response.Response "{"code": 7,"message": "Bad Request","result": {},"type": "fail"}"
// @Failure 409 {object} response.Response "{"code": 7,"message": "conflict data","result": {},"type": "fail"}"
// @Failure 500 {object} response.Response "{"code": 7,"message": "create fail","result": {},"type": "fail"}"
// @Router /class/createClass [post]
func (m *ClassApi) CreateClass(c *gin.Context) {
	var class model.Class

	if err := c.ShouldBindJSON(&class); err != nil {
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

	m.classController.CreateClass(class, c)
}


// GetClassList
// @Tags Class
// @Summary GetClassList
// @Security ApiKeyAuth
// @accept Application/json
// @Produce Application/json
// @Param data query  tmsReq.ClassSearch false "ClassList"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Failure 500
// @Router /class/getClassList [get]
func (m *ClassApi) GetClassList(c *gin.Context) {
	var info manReq.ClassSearch
	if err := c.ShouldBindQuery(&info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "error", c)
		return
	}

	m.classController.GetClassList(info, c)
}