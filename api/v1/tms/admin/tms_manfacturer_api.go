package admin

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	controller "github.com/ebedevelopment/next-gen-tms/server/controller/tms/admin"
	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	"github.com/ebedevelopment/next-gen-tms/server/model/tms"
	tmsReq "github.com/ebedevelopment/next-gen-tms/server/model/tms/request"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type ManufacturerApi struct {
	manufacturerController controller.ManufacturerController
}

// CreateManufacturer createManufacturer
// @Tags Manufacturer
// @Summary createManufacturer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tms.ManufacturerDTO true "createManufacturer"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Failure 400 {object} response.Response "{"code": 7,"message": "Bad Request","result": {},"type": "fail"}"
// @Failure 409 {object} response.Response "{"code": 7,"message": "conflict data","result": {},"type": "fail"}"
// @Failure 500 {object} response.Response "{"code": 7,"message": "create fail","result": {},"type": "fail"}"
// @Router /manufacturer/createManufacturer [post]
func (m *ManufacturerApi) CreateManufacturer(c *gin.Context) {
	var manufacturer tms.ManufacturerDTO

	if err := c.ShouldBindJSON(&manufacturer); err != nil {
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

	m.manufacturerController.CreateManufacturer(manufacturer, c)

}

// DeleteManufacturer deleteManufacturer
// @Tags Manufacturer
// @Summary deleteManufacturer
// @Security ApiKeyAuth
// @Produce application/json
// @Param id path int true "Manufacturer id"
// @Success 200 {object} response.Response "{"code": 0,"message": "delete Success","result": {},"type": "success"}"
// @Failure 404 {object} response.Response "{"code": 7,"message": "idNotFound","result": {},"type": "fail"}"
// @Failure 500 {object} response.Response "{"code": 7,"message": "Data Fail","result": {},"type": "fail"}"
// @Router /manufacturer/deleteManufacturer/{id}   [delete]
func (m *ManufacturerApi) DeleteManufacturer(c *gin.Context) {

	id := c.Param("id")
	manufacturerId, err := strconv.Atoi(id)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "warning", c)
		return
	}

	m.manufacturerController.DeleteManufacturer(manufacturerId, c)

}

// UpdateManufacturer updateManufacturer
// @Tags Manufacturer
// @Summary updateManufacturer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tms.ManufacturerDTO true "updateManufacturer"
// @Param  id path int true "Manufacturer id"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Router /manufacturer/updateManufacturer/{id}    [put]
func (m *ManufacturerApi) UpdateManufacturer(c *gin.Context) {
	var manufacturer tms.ManufacturerDTO

	id := c.Param("id")
	manufID, err := strconv.Atoi(id)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "warning", c)
		return
	}

	if err := c.ShouldBindJSON(&manufacturer); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := utils.HandleError(ve)
			response.FailWithDetailed(out, global.Translate("general.badRequest"), http.StatusBadRequest, "warning", c)
			return
		}
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "error", c)
		return
	}

	m.manufacturerController.UpdateManufacturer(manufacturer, manufID, c)
}

// UpdateManufacturerStatus UpdateManufacturerStatus
// @Tags Manufacturer
// @Summary UpdateManufacturerStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param  id path int true "Manufacturer id"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Failure 500 {object} response.Response "{"code": 7,"message": "Data Fail","result": {},"type": "fail"}"
// @Router /manufacturer/UpdateManufacturerStatus/{id}   [put]
func (m *ManufacturerApi) UpdateManufacturerStatus(c *gin.Context) {

	id := c.Param("id")
	manufID, err := strconv.Atoi(id)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "warning", c)
		return
	}

	m.manufacturerController.UpdateManufacturerStatus(manufID, c)
}

// FindManufacturer queryByIdManufacturer
// @Tags Manufacturer
// @Summary queryByIdManufacturer
// @Security ApiKeyAuth
// @Produce application/json
// @Param id path int true "queryByIdManufacturer"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Failure 400 {object} response.Response "{"code": 7,"message": "Bad Request","result": {},"type": "fail"}"
// @Failure 404 {object} response.Response "{"code": 7,"message": "idNotFound","result": {},"type": "fail"}"
// @Failure 500 {object} response.Response "{"code": 7,"message": "Data Fail","result": {},"type": "fail"}"
// @Router /manufacturer/findManufacturer/{id}   [get]
func (m *ManufacturerApi) FindManufacturer(c *gin.Context) {

	id := c.Param("id")
	manufID, err := strconv.Atoi(id)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "warning", c)
		return
	}

	m.manufacturerController.FindManufacturer(manufID, c)
}

// GetManufacturerList
// @Tags Manufacturer
// @Summary GetManufacturerList
// @Security ApiKeyAuth
// @accept Application/json
// @Produce Application/json
// @Param data query  tmsReq.ManufacturerSearch false "manufacturerList"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Failure 500
// @Router /manufacturer/getManufacturerList [get]
func (m *ManufacturerApi) GetManufacturerList(c *gin.Context) {
	var info tmsReq.ManufacturerSearch
	if err := c.ShouldBindQuery(&info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "error", c)
		return
	}

	m.manufacturerController.GetManufacturerList(info, c)
}

// exportManufacturerExcel
// @Tags Manufacturer
// @Summary exportManufacturerExcel
// @Security ApiKeyAuth
// @accept Application/json
// @Produce Application/json
// @Param data query  tmsReq.ManufacturerSearch false "manufacturerList"
// @Success 200 {object} response.Response "{"code": 0,"message": "DataSuccess","result": {},"type": "success"}"
// @Failure 500
// @Router /manufacturer/exportManufacturerExcel [get]
func (m *ManufacturerApi) ExportManufacturerExcel(c *gin.Context) {
	var info tmsReq.ManufacturerSearch
	if err := c.ShouldBindQuery(&info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(err))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "error", c)
		return
	}

	if info.Format != utils.Excel && info.Format != utils.PDF {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].BadRequest, zap.Error(fmt.Errorf("format not pdf or excel")))
		response.FailWithMessage(global.Translate("general.badRequest"), http.StatusBadRequest, "error", c)
		return
	}

	m.manufacturerController.ExportManufacturerExcel(info, c)
}
