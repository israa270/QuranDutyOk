package tms

import (
	v1 "github.com/ebedevelopment/next-gen-tms/server/api/v1/tms/admin"
	"github.com/ebedevelopment/next-gen-tms/server/middleware"
	"github.com/gin-gonic/gin"
)

// ManufacturerRouter
type ManufacturerRouter struct {
	manufacturerApi v1.ManufacturerApi
}

// InitManufacturerRouter initialization Manufacturer routing info
func (s *ManufacturerRouter) InitManufacturerRouter(Router *gin.RouterGroup) {
	manufacturerRouter := Router.Group("manufacturer").Use(middleware.OperationRecord())
	manufacturerRouterWithoutRecord := Router.Group("manufacturer")
	{
		manufacturerRouter.POST("createManufacturer", s.manufacturerApi.CreateManufacturer)             // CreateManufacturer
		manufacturerRouter.DELETE("deleteManufacturer/:id", s.manufacturerApi.DeleteManufacturer)       // deleteManufacturer
		manufacturerRouter.PUT("updateManufacturer/:id", s.manufacturerApi.UpdateManufacturer)          // updateManufacturer
		manufacturerRouter.PUT("UpdateManufacturerStatus/:id", s.manufacturerApi.UpdateManufacturerStatus)
		// excel
		manufacturerRouter.GET("exportManufacturerExcel", s.manufacturerApi.ExportManufacturerExcel)
	}
	{
		manufacturerRouterWithoutRecord.GET("findManufacturer/:id", s.manufacturerApi.FindManufacturer)   // getByIDManufacturer
		manufacturerRouterWithoutRecord.GET("getManufacturerList", s.manufacturerApi.GetManufacturerList) // get Manufacturer list
	}
}
