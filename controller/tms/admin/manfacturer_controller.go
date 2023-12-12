package admin

import (
	"net/http"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	"github.com/ebedevelopment/next-gen-tms/server/model/tms"
	tmsReq "github.com/ebedevelopment/next-gen-tms/server/model/tms/request"
	"github.com/ebedevelopment/next-gen-tms/server/service/tms/admin"
	commoncase "github.com/ebedevelopment/next-gen-tms/server/usecase/common"
	// claimcase "github.com/ebedevelopment/next-gen-tms/server/usecase/user/claim"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type ManufacturerController struct {
	manufacturerService     admin.ManufacturerService
}

func (m *ManufacturerController) CreateManufacturer(manufacturer tms.ManufacturerDTO, c *gin.Context) {
	// Check manufacturer name
	if m.manufacturerService.CheckManufacturerByName(manufacturer.Name) {
		global.GvaLog.Debug(global.GvaLoggerMessage["log"].DuplicateValueName)
		response.FailWithMessage(global.Translate("general.duplicateValueName"), http.StatusConflict, "warning", c)
		return
	}

	// Validate phone
	if !commoncase.ValidatePhone(manufacturer.Phone) {
		response.FailWithMessage(global.Translate("general.validatePhone"), http.StatusBadRequest, "warning", c)
		return
	}

	// Check manufacturer phone
	if m.manufacturerService.CheckManufacturerPhone(manufacturer.Phone) {
		global.GvaLog.Debug(global.GvaLoggerMessage["log"].DuplicateValuePhone)
		response.FailWithMessage(global.Translate("general.duplicateValuePhone"), http.StatusConflict, "warning", c)
		return
	}

	// check manufacturer email & validate email in validator binding
	if m.manufacturerService.CheckManufacturerByEmail(manufacturer.Email) {
		global.GvaLog.Debug(global.GvaLoggerMessage["log"].DuplicatedEmail)
		response.FailWithMessage(global.Translate("general.duplicateValueEmail"), http.StatusConflict, "warning", c)
		return
	}

	//Get Username Created by
	// tokenData, err := claimcase.GetBaseClaim(c)
	// if err != nil {
	// 	global.GvaLog.Error(global.GvaLoggerMessage["log"].GetOrganizationFail, zap.Error(err))
	// 	response.FailWithDetailed(err.Error(), global.Translate("organization.getOrganizationFail"), http.StatusUnauthorized, "error", c)
	// 	return
	// }

	manufacturerDB := tms.Manufacturer{
		ManufacturerDTO: manufacturer,
		// CreatedBy:       tokenData.Email,
		Status:          true, // active when Create
	}

	if err := m.manufacturerService.CreateManufacturer(manufacturerDB); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].CreationFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFail"), http.StatusInternalServerError, "error", c)
	} else {
		response.OkWithMessage(global.Translate("general.createSuccess"), http.StatusOK, "success", c)
	}
}

func (m *ManufacturerController) DeleteManufacturer(manufacturerId int, c *gin.Context) {
	if customErr := m.manufacturerService.DeleteManufacturer(uint(manufacturerId)); customErr.Message != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].DeleteFail, zap.Error(customErr.Message))
		response.FailWithDetailed(customErr.Message.Error(), global.Translate("general.deleteFail"), customErr.StatusCode, "error", c)
	} else {
		response.OkWithMessage(global.Translate("general.deleteSuccess"), http.StatusOK, "success", c)
	}
}

func (m *ManufacturerController) UpdateManufacturer(manufacturer tms.ManufacturerDTO, manufID int, c *gin.Context) {
	//Get Username Created by
	// tokenData, err := claimcase.GetBaseClaim(c)
	// if err != nil {
	// 	global.GvaLog.Error(global.GvaLoggerMessage["log"].GetOrganizationFail, zap.Error(err))
	// 	response.FailWithDetailed(err.Error(), global.Translate("organization.getOrganizationFail"), http.StatusUnauthorized, "error", c)
	// 	return
	// }

	manufacturerDB := tms.Manufacturer{
		// UpdatedBy:       tokenData.Email,
		ManufacturerDTO: manufacturer,
	}

	if customErr := m.manufacturerService.UpdateManufacturer(uint(manufID), manufacturerDB); customErr.Message != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].UpdateFail, zap.Error(customErr.Message))
		response.FailWithDetailed(customErr.Message.Error(), global.Translate("general.updateFail"), customErr.StatusCode, "warning", c)
		return
	} else {
		response.OkWithMessage(global.Translate("general.updateSuccess"), http.StatusOK, "success", c)
	}
}

func (m *ManufacturerController) UpdateManufacturerStatus(manufID int, c *gin.Context) {
	// //Get Username
	// tokenData, err := claimcase.GetBaseClaim(c)
	// if err != nil {
	// 	global.GvaLog.Error(global.GvaLoggerMessage["log"].GetOrganizationFail, zap.Error(err))
	// 	response.FailWithDetailed(err.Error(), global.Translate("organization.getOrganizationFail"), http.StatusUnauthorized, "error", c)
	// 	return
	// }

	// if customErr := m.manufacturerService.UpdateManufacturerStatus(uint(manufID), tokenData.Email); customErr.Message != nil {
	// 	global.GvaLog.Error(global.GvaLoggerMessage["log"].UpdateFail, zap.Error(customErr.Message))
	// 	response.FailWithDetailed(customErr.Message.Error(), global.Translate("general.updateFail"), customErr.StatusCode, "error", c)
	// } else {
	// 	response.OkWithMessage(global.Translate("general.updateSuccess"), http.StatusOK, "success", c)
	// }

}

func (m *ManufacturerController) FindManufacturer(manufID int, c *gin.Context) {
	if manufacturer, err := m.manufacturerService.GetManufacturer(uint(manufID)); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].QueryFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFail"), http.StatusNotFound, "error", c)
	} else {
		// For UI design
		// var manDetails tmsResp.ManufacturerDetails
		// manDetails.ManufacturerDTO = manufacturer.ManufacturerDTO

		// for _, model := range manufacturer.Models {
		// 	manDetails.Models = append(manDetails.Models, model.Name)
		// }

		// manDetails.CreateAt = manufacturer.CreatedAt.Format("2006-01-02 15:04:05")
		// manDetails.UpdateAt = manufacturer.UpdatedAt.Format("2006-01-02 15:04:05")
		// response.OkWithData(manDetails, http.StatusOK, "success", c)

		response.OkWithData(manufacturer, http.StatusOK, "success", c)
	}
}

func (m *ManufacturerController) GetManufacturerList(info tmsReq.ManufacturerSearch, c *gin.Context) {
	if list, total, err := m.manufacturerService.GetManufacturerInfoList(info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GetDataFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFail"), http.StatusInternalServerError, "error", c)
	} else {

		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     info.Page,
			PageSize: info.PageSize,
		}, global.Translate("general.getDataSuccess"), http.StatusOK, "success", c)

		// For UI Response
		// var Manufacturers []tmsResp.ManufacturerDetails
		// for _, Manufacturer := range list {
		// 	var manDetails tmsResp.ManufacturerDetails
		// 	manDetails.ManufacturerDTO = Manufacturer.ManufacturerDTO
		// 	for _, model := range Manufacturer.Models {
		// 		manDetails.Models = append(manDetails.Models, model.Name)
		// 	}
		// 	manDetails.CreateAt = Manufacturer.CreatedAt.Format("2006-01-02 15:04:05")
		// 	manDetails.UpdateAt = Manufacturer.UpdatedAt.Format("2006-01-02 15:04:05")

		// 	Manufacturers = append(Manufacturers, manDetails)
		// }

		// response.OkWithDetailed(Manufacturers, global.Translate("general.getDataSuccess"), http.StatusOK, "success", c)
	}
}

func (m *ManufacturerController) ExportManufacturerExcel(info tmsReq.ManufacturerSearch, c *gin.Context) {
	if list, _, err := m.manufacturerService.GetManufacturerInfoList(info); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GetDataFail, zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFail"), http.StatusInternalServerError, "error", c)
	} else {

		if len(list) == 0 {
			response.FailWithMessage(global.Translate("general.emptyExportData"), http.StatusBadRequest, "error", c)
			return
		}

		if _, err := m.manufacturerService.ParseManufacturerInfoList2Excel(list, info.Format); err != nil {
			global.GvaLog.Error(global.GvaLoggerMessage["log"].ExcelFail, zap.Error(err))
			response.FailWithMessage(global.Translate("excel.excelFail"), http.StatusInternalServerError, "error", c)
			return
		} else {
			// 	c.Writer.Header().Add("success", "true")
			// 	c.File(filePath)

		
			response.FailWithMessage("failed to get file", http.StatusInternalServerError, "error", c)
		}
	}
}
