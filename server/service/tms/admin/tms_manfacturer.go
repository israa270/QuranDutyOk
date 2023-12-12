package admin

import (
	"fmt"
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common"
	"github.com/ebedevelopment/next-gen-tms/server/model/tms"
	tmsReq "github.com/ebedevelopment/next-gen-tms/server/model/tms/request"
	repository "github.com/ebedevelopment/next-gen-tms/server/repository/tms"
	commoncase "github.com/ebedevelopment/next-gen-tms/server/usecase/common"
	useCase "github.com/ebedevelopment/next-gen-tms/server/usecase/excel"
	pdfCase "github.com/ebedevelopment/next-gen-tms/server/usecase/pdf"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
)

// ManufacturerService
type ManufacturerService struct {
	manufacturerRepository repository.ManufacturerRepository
	excelUsecase           useCase.ExcelUseCase
	pdfUsecase      pdfCase.PDFUseCase
}

// CreateManufacturer createManufacturerRecord
func (m *ManufacturerService) CreateManufacturer(manufacturer tms.Manufacturer) error {
	return m.manufacturerRepository.CreateManufacturer(manufacturer)
}

// DeleteManufacturer deleteManufacturerRecord when status is active refuse delete
// and when status is disable make delete to all model and terminals belong.
func (m *ManufacturerService) DeleteManufacturer(id uint) *common.CustomError {
	return m.manufacturerRepository.DeleteManufacturer(id)
}

// UpdateManufacturer updateManufacturerRecord
func (m *ManufacturerService) UpdateManufacturer(id uint, manufReq tms.Manufacturer) *common.CustomError {
	return m.manufacturerRepository.UpdateManufacturer(id, manufReq)
}

// UpdateManufacturerStatus update manufacture status from active to disable and disable all models and disable all terminal
func (m *ManufacturerService) UpdateManufacturerStatus(id uint, userEmail string) common.CustomError {
	return m.manufacturerRepository.UpdateManufacturerStatus(id, userEmail)
}

// GetManufacturer getByIDManufacturerRecord
func (m *ManufacturerService) GetManufacturer(id uint) (*tms.Manufacturer, error) {
	return m.manufacturerRepository.GetManufacturer(id)
}

// CheckManufacturerIdExist   check manufacturer exist status
func (m *ManufacturerService) CheckManufacturerIdExist(id uint) bool {
	return m.manufacturerRepository.CheckManufacturerIdExist(id)
}

// GetManufacturerName
func (m *ManufacturerService) GetManufacturerByName(name string) (*tms.Manufacturer, error) {
	return m.manufacturerRepository.GetManufacturerByName(name)
}

// GetManufacturerIDByName get manufacturer ID by manufacturer Name used in excel
func (m *ManufacturerService) GetManufacturerIDByName(name string) (uint, error) {
	return m.manufacturerRepository.GetManufacturerIDByName(name)
}

// CheckManufacturerByName check manufacturer by name
func (m *ManufacturerService) CheckManufacturerByName(name string) bool {
	return m.manufacturerRepository.CheckManufacturerByName(name)
}

// CheckManufacturerPhone
func (m *ManufacturerService) CheckManufacturerPhone(phone string) bool {
	return m.manufacturerRepository.CheckManufacturerPhone(phone)
}

// CheckManufacturerByEmail
func (m *ManufacturerService) CheckManufacturerByEmail(email string) bool {
	return m.manufacturerRepository.CheckManufacturerByEmail(email)
}

// GetManufacturerInfoList pagingManufacturerRecord
func (m *ManufacturerService) GetManufacturerInfoList(info tmsReq.ManufacturerSearch) ([]tms.Manufacturer, int64, error) {
	return m.manufacturerRepository.GetManufacturerInfoList(info)
}

// excel for Manufacturer export
func (m *ManufacturerService) ParseManufacturerInfoList2Excel(info []tms.Manufacturer, format string) (string, error) {

	header := []string{"Name", "Email", "Contact Name", "Phone", "country", "Status", "Created By", "Created Date", "Updated By", "Updated Date"}
	data := [][]string{}

	if format == utils.Excel {
		data = append(data, header)
	}

	for _, i := range info {
		var s []string

		s = append(s, i.Name, i.Email, i.ContactName, i.Phone, i.Country, commoncase.UserStatus(i.Status), i.CreatedBy, i.CreatedAt.Format("2006-01-02 15:04:05"), i.UpdatedBy, i.UpdatedAt.Format("2006-01-02 15:04:05"))
		data = append(data, s)
	}

	var filePath string
	if format == utils.PDF {
		filePath = global.GvaConfig.Excel.Dir + "manufacturerExport_" + fmt.Sprintf("%v", time.Now().Unix()) + ".pdf"
		return m.pdfUsecase.CreatePDF(filePath, header, data, "Manufacturer Data", pdfCase.Manufacturer)
	}

	filePath = global.GvaConfig.Excel.Dir + "manufacturerExport_" + fmt.Sprintf("%v", time.Now().Unix()) + ".xlsx"

	return m.excelUsecase.GenerateExcelSheet(data, filePath)

}

// GetManufacturerNamesList
func (m *ManufacturerService) GetManufacturerNamesList() ([]tmsReq.ManufacturerModel, error) {
	return m.manufacturerRepository.GetManufacturerNamesList()
}
