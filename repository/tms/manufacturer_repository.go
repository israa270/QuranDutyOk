package tms

import (
	"errors"
	"fmt"

	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common"
	"github.com/ebedevelopment/next-gen-tms/server/model/tms"
	tmsReq "github.com/ebedevelopment/next-gen-tms/server/model/tms/request"
	commoncase "github.com/ebedevelopment/next-gen-tms/server/usecase/common"
)

// ManufacturerRepository
type ManufacturerRepository struct {
}

// CreateManufacturer createManufacturerRecord
func (m *ManufacturerRepository) CreateManufacturer(manufacturer tms.Manufacturer) (err error) {
	err = global.GvaDB.Create(&manufacturer).Error
	return err
}

// DeleteManufacturer deleteManufacturerRecord when status is active refuse delete
// and when status is disable make delete to all model and terminals belong.
func (m *ManufacturerRepository) DeleteManufacturer(id uint) *common.CustomError {

	manufacturer, err := m.GetManufacturer(id)
	if err != nil {
		return &common.CustomError{
			Message:    err,
			StatusCode: 404,
		}
	}

	// can not delete active status
	if manufacturer.Status {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].StatusActive + " " + manufacturer.Name)
		return &common.CustomError{
			Message:    fmt.Errorf(global.Translate("general.statusActive")),
			StatusCode: 409,
		}
	}


	err = global.GvaDB.Model(&manufacturer).Delete(&manufacturer).Error
	return &common.CustomError{Message: err, StatusCode: 500}
}

// UpdateManufacturer updateManufacturerRecord
func (m *ManufacturerRepository) UpdateManufacturer(id uint, manufReq tms.Manufacturer) *common.CustomError {
	var manufacturer tms.Manufacturer

	//check manufacturer exist
	manufExist, err := manufacturerRepository.GetManufacturer(id)
	if err != nil {
		return &common.CustomError{
			Message:    errors.New(global.Translate("manufacturer.manufacturerExist")),
			StatusCode: 404,
		}
	}

	if manufExist.Name != manufReq.Name {
		if err := global.GvaDB.Where("name =? ", manufReq.Name).First(&manufacturer).Error; err == nil {
			return &common.CustomError{
				Message:    errors.New(global.Translate("general.duplicateValueName")),
				StatusCode: 409,
			}
		}
	}

	if manufExist.Phone != manufReq.Phone {
		// Validate phone
		if !commoncase.ValidatePhone(manufReq.Phone) {
			return &common.CustomError{
				Message:    errors.New(global.Translate("general.validatePhone")),
				StatusCode: 400,
			}

		}

		if err := global.GvaDB.Where("phone =? ", manufReq.Phone).First(&manufacturer).Error; err == nil {
			return &common.CustomError{
				Message:    errors.New(global.Translate("general.duplicateValuePhone")),
				StatusCode: 409,
			}
		}
	}

	if manufExist.Email != manufReq.Email {

		if m.CheckManufacturerByEmail(manufReq.Email) {
			return &common.CustomError{
				Message:    errors.New(global.Translate("general.duplicateValueEmail")),
				StatusCode: 409,
			}
		}
	}

	//For empty values
	manufExist.ManufacturerDTO = manufReq.ManufacturerDTO

	//updates
	err = global.GvaDB.Where("id =? ", id).Save(&manufExist).Error
	return &common.CustomError{Message: err, StatusCode: 500}
}

// UpdateManufacturerStatus update manufacture status from active to disable and disable all models and disable all terminal
func (m *ManufacturerRepository) UpdateManufacturerStatus(id uint, userEmail string) common.CustomError {
	manufacturer, err := manufacturerRepository.GetManufacturer(id)
	if err != nil {
		return common.CustomError{
			Message:    err,
			StatusCode: 404,
		}
	}

	manufacturer.UpdatedBy = userEmail
	manufacturer.Status = !manufacturer.Status
	manufacturer.UpdatedAt = time.Now()

	if err = global.GvaDB.Where("id =? ", id).Save(&manufacturer).Error; err != nil {
		return common.CustomError{
			Message:    err,
			StatusCode: 500,
		}
	}

	return common.CustomError{}
}

// GetManufacturer getByIDManufacturerRecord
func (m *ManufacturerRepository) GetManufacturerDetails(id uint) (manufacturer *tms.Manufacturer, err error) {
	err = global.GvaDB.Where("id = ?", id).Preload("Models").First(&manufacturer).Error
	return
}

func (m *ManufacturerRepository) GetManufacturer(id uint) (manufacturer *tms.Manufacturer, err error) {
	err = global.GvaDB.Where("id = ?", id).First(&manufacturer).Error
	return
}

// CheckManufacturerIdExist   check manufacturer exist status
func (m *ManufacturerRepository) CheckManufacturerIdExist(id uint) bool {
	var manufacturer *tms.Manufacturer
	err := global.GvaDB.Where("id=? AND status =?", id, true).First(&manufacturer).Error
	return err == nil
}

// // CheckManufacturerIdExistWithoutStatus  check manufacturer exist or not
// func (m *ManufacturerRepository) CheckManufacturerIdExistWithoutStatus(id uint) bool {
// 	var manufacturer *tms.Manufacturer
// 	err := global.GvaDB.Where("id=?", id).First(&manufacturer).Error
// 	return err == nil
// }

// GetManufacturerName
func (m *ManufacturerRepository) GetManufacturerByName(name string) (manufacturer *tms.Manufacturer, err error) {
	err = global.GvaDB.Where("name = ?", name).First(&manufacturer).Error
	return
}

// GetManufacturerIDByName get manufacturer ID by manufacturer Name used in excel
func (m *ManufacturerRepository) GetManufacturerIDByName(name string) (uint, error) {
	var manufacturer *tms.Manufacturer
	if err := global.GvaDB.Where("name = ?", name).First(&manufacturer).Error; err != nil {
		return 0, err
	}
	return manufacturer.ID, nil
}

// CheckManufacturerByName check manufacturer by name
func (m *ManufacturerRepository) CheckManufacturerByName(name string) bool {
	var manufacturer *tms.Manufacturer
	err := global.GvaDB.Where("name = ?", name).First(&manufacturer).Error
	return err == nil
}

// GetManufacturerByPhone
// func (m *ManufacturerRepository) GetManufacturerByPhone(phone string) (manufacturer *tms.Manufacturer, err error) {
// 	err = global.GvaDB.Where("phone = ?", phone).First(&manufacturer).Error
// 	return
// }

// CheckManufacturerPhone
func (m *ManufacturerRepository) CheckManufacturerPhone(phone string) bool {
	var manufacturer *tms.Manufacturer
	err := global.GvaDB.Where("phone = ?", phone).First(&manufacturer).Error
	return err == nil
}

// GetManufacturerByEmail
// func (m *ManufacturerRepository) GetManufacturerByEmail(email string) (manufacturer *tms.Manufacturer, err error) {
// 	err = global.GvaDB.Where("email = ?", email).First(&manufacturer).Error
// 	return
// }

// CheckManufacturerByEmail
func (m *ManufacturerRepository) CheckManufacturerByEmail(email string) bool {
	var manufacturer *tms.Manufacturer
	err := global.GvaDB.Where("email = ?", email).First(&manufacturer).Error
	return err == nil
}

// GetManufacturerInfoList pagingManufacturerRecord
func (m *ManufacturerRepository) GetManufacturerInfoList(info tmsReq.ManufacturerSearch) (list []tms.Manufacturer, total int64, err error) {
	if info.PageSize == 0 || info.PageSize > global.GvaConfig.Mysql.LimitRecords {
		info.PageSize = global.GvaConfig.Mysql.LimitRecords
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// create db
	db := global.GvaDB.Model(&tms.Manufacturer{})
	var manufacturers []tms.Manufacturer
	// If conditional search create search
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Email != "" {
		db = db.Where("email LIKE ?", "%"+info.Email+"%")
	}

	if info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	}

	if info.Country != "" {
		db = db.Where("country LIKE ?", "%"+info.Country+"%")
	}

	if info.Status == global.GvaConfig.Login.StatusTrue {
		db = db.Where("status=?", true)
	} else if info.Status == global.GvaConfig.Login.StatusFalse {
		db = db.Where("status =?", false)
	}

	t := time.Time{}
	if info.CreatedBefore != t && info.CreatedAfter != t {
		db.Where("created_at  >= ? AND created_at  <= ?", info.CreatedBefore, info.CreatedAfter)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if info.Platform != "" {
		err = db.Order("id " + info.Order).Limit(limit).Offset(offset).Preload("Models").Find(&manufacturers).Error
		return manufacturers, total, err
	}

	err = db.Order("id " + info.Order).Limit(limit).Offset(offset).Find(&manufacturers).Error

	return manufacturers, total, err
}

// GetManufacturerNamesList   used in terminal download Template in excel
func (m *ManufacturerRepository) GetManufacturerNamesList() (manufacturerNames []tmsReq.ManufacturerModel, err error) {
	// create db
	var manufacturers []tms.Manufacturer
	var manufacturerModels []tmsReq.ManufacturerModel

	err = global.GvaDB.Model(&tms.Manufacturer{}).Where("status=?", true).Find(&manufacturers).Error
	for _, manufacturer := range manufacturers {

		err = global.GvaDB.Where("id = ?", manufacturer.ID).Preload("Models").First(&manufacturer).Error
		// if len(manufacturer.Models) != 0 {
		// 	for _, model := range manufacturer.Models {
		// 		// Check Status active
		// 		if model.Status {
		// 			ManufacturerModel := tmsReq.ManufacturerModel{
		// 				ManufacturerName: manufacturer.Name,
		// 				ModelName:        model.Name,
		// 			}
		// 			manufacturerModels = append(manufacturerModels, ManufacturerModel)
		// 		}
		// 	}
		// }
	}
	return manufacturerModels, err
}

// DeleteManufacturerByIds batch delete ManufacturerRecord
// func (m *ManufacturerRepository) DeleteManufacturerByIds(ids request.IdsReq) common.CustomError {
// 	var manufacturers []tms.Manufacturer
// 	if err := global.GvaDB.Where("id in ?", ids.Ids).Find(&manufacturers).Error; err != nil{
//          return common.CustomError{Message: err, StatusCode: 500}
// 	}

// 	for _, manufacturer := range manufacturers {

// 		//  can not delete active status
// 		if manufacturer.Status{
// 			global.GvaLog.Error(global.GvaLoggerMessage["log"].StatusActive + " " + manufacturer.Name)
// 			return common.CustomError{
// 				Message:    fmt.Errorf(global.Translate("general.statusActive")+ " " + manufacturer.Name),
// 				StatusCode: 409,
// 			}
// 		}

// 		// status is disable
// 		if err := global.GvaDB.Where("id = ?",manufacturer.ID).Preload("Models").First(&manufacturer).Error; err != nil {
// 			return common.CustomError{
// 				Message:    err,
// 				StatusCode: 404,
// 			}
// 		}

// 		// Manufacturer has models don't delete but change status
// 		if len(manufacturer.Models) != 0 {
// 			// Delete Models
// 			for _, model := range manufacturer.Models{
// 				// Delete model
// 				if customErr := modelRepository.DeleteModel(model.ID); customErr.Message != nil{
// 					return customErr
// 				}
// 			}
// 		}

// 		if err := global.GvaDB.Model(&manufacturer).Delete(&manufacturer).Error; err != nil{
// 		     return common.CustomError{Message: err,StatusCode:500}
// 		}
// 	}

// 	// err := global.GvaDB.Delete(&[]tms.Manufacturer{}, "id in ?", ids.Ids).Error
// 	return common.CustomError{Message: nil, StatusCode: 500}
// }
