package response

import "github.com/ebedevelopment/next-gen-tms/server/model/tms"

// ManufacturerDetails
type ManufacturerDetails struct {
	tms.ManufacturerDTO
	Models   []string `json:"models"`
	CreateAt string   `json:"createdAt"`
	UpdateAt string   `json:"updateAt"`
}


type ManufactureDetails struct {
	ManufactureId   uint   `json:"id,omitempty" form:"id"`
	ManufactureName string `json:"name" form:"name"`
}
