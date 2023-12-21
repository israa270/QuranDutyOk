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

type HomeWorkApi struct {
	homeWorkController management.HomeWorkController
}

// CreateHomeWork createHomeWork
// @Tags HomeWork
// @Summary createHomeWork
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tms.HomeWorkDTO true "createHomeWork"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Failure 400 {object} response.Response "{"code": 7,"message": "Bad Request","result": {},"type": "fail"}"
// @Failure 409 {object} response.Response "{"code": 7,"message": "conflict data","result": {},"type": "fail"}"
// @Failure 500 {object} response.Response "{"code": 7,"message": "create fail","result": {},"type": "fail"}"
// @Router /homeWork/createHomeWork [post]
func (m *HomeWorkApi) CreateHomeWork(c *gin.Context) {
	var homeWork model.HomeWork

	if err := c.ShouldBindJSON(&homeWork); err != nil {
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

	m.homeWorkController.CreateHomeWork(homeWork, c)
}

// GetHomeWorkList
// @Tags HomeWork
// @Summary GetHomeWorkList
// @Security ApiKeyAuth
// @accept Application/json
// @Produce Application/json
// @Param data query  tmsReq.HomeWorkSearch false "HomeWorkList"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Failure 500
// @Router /homeWork/getHomeWorkList [get]
func (m *HomeWorkApi) GetHomeWorkList(c *gin.Context) {
	var info manReq.HomeWorkSearch
	if err := c.ShouldBindQuery(&info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "error", c)
		return
	}

	m.homeWorkController.GetHomeWorkList(info, c)
}


// AssignHomeWorkToClass
// @Tags HomeWork
// @Summary AssignHomeWorkToClass
// @Security ApiKeyAuth
// @accept Application/json
// @Produce Application/json
// @Param data query  tmsReq.AssignHomeWorkToClassesDTO false "HomeWorkList"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Failure 500
// @Router /homeWork/assignHomeWorkToClass [put]
func (m *HomeWorkApi) AssignHomeWorkToClass(c *gin.Context) {
	var homework manReq.AssignHomeWorkToClassesDTO
	if err := c.ShouldBindJSON(&homework); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "error", c)
		return
	}

	m.homeWorkController.AssignHomeWorkToClass(homework, c)
}