package request

import (
	"github.com/ebedevelopment/next-gen-tms/server/model/common/request"
)

// ManufacturerSearch
type ManufacturerSearch struct {
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Phone string `json:"phone" form:"phone"`
	Country string `json:"country" form:"country"`
	Status  string `json:"status" form:"status"`

	Platform  string `json:"platform" form:"platform"`  //platform=Android|monitor|prolin
	request.PageInfo
}


// ManufacturerModel
type ManufacturerModel struct {
	ManufacturerName string `json:"ManufacturerName"`
	ModelName        string `json:"modelName"`
}
