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

type HomeWorkController struct {
	homeWorkService   management.HomeWorkService
}

func (m *HomeWorkController) CreateHomeWork(homeWork model.HomeWork, c *gin.Context) {
	//Get Username Created by
	tokenData, err := claimcase.GetBaseClaim(c)
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GetOrganizationFail, zap.Error(err))
		response.FailWithDetailed(err.Error(), global.Translate("organization.getOrganizationFail"), http.StatusUnauthorized, "error", c)
		return
	}

	if tokenData.Role != utils.Admin {
		response.FailWithMessage("you not have permission", http.StatusUnauthorized, "error", c)
		return
	}

     //TODO : Validate Expire Date

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