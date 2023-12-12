// generateManufacturer
package tms

import (
	"github.com/ebedevelopment/next-gen-tms/server/global"
)

// Manufacturer
type Manufacturer struct {
	global.GvaModel

	ManufacturerDTO
	Status  bool   `json:"status" form:"status" gorm:"column:status;comment:;"`
	
	// Models []Model `json:"model,omitempty" form:"model" gorm:"ForeignKey:ManufacturerID"`

	CreatedBy string `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:;size:255;"`
	UpdatedBy string `json:"updatedBy,omitempty" form:"updatedBy" gorm:"column:updated_by;comment:;size:255;"`
}

// TableName Manufacturer table Name
func (Manufacturer) TableName() string {
	return "tms_manufacturer"
}

// ManufacturerDTO
type ManufacturerDTO struct {
	Name        string `json:"name" form:"name" gorm:"column:name;comment:;size:255;" binding:"required,min=3,max=50"`
	Email       string `json:"email" form:"email" gorm:"column:email;comment:;" binding:"required,email"`
	ContactName string `json:"contactName" form:"contactName" gorm:"column:contact_name;comment:;" binding:"required,min=3,max=50"`
	Phone       string `json:"phone" form:"phone" gorm:"column:phone;comment:;" binding:"required"`

	
	Country string `json:"country" form:"country" gorm:"column:country" binding:"required"` //TODO:"binding:"required,iso3166_1_alpha2"

	//Optional data
	City        string `json:"city,omitempty" form:"city" gorm:"column:city" binding:"max=100"`
	PostalCode  string `json:"postalCode,omitempty" form:"postalCode" gorm:"column:postal_code;comment:;"` //binding:"postcode_iso3166_alpha2=GB"
	Address     string `json:"address,omitempty" form:"address" gorm:"column:address;comment:;type:text" binding:"max=100"`
	Description string `json:"description,omitempty" form:"description" gorm:"column:description;comment:;type:text" binding:"max=255"`

	SignatureProvider string `json:"signatureProvider,omitempty" form:"signatureProvider" gorm:"column:signature_provider" binding:"max=255"`
}