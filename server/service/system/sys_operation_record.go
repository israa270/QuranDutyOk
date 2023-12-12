package system

import (
	"fmt"
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/system"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
	sysResp "github.com/ebedevelopment/next-gen-tms/server/model/system/response"
	repository "github.com/ebedevelopment/next-gen-tms/server/repository/system"
	useCase "github.com/ebedevelopment/next-gen-tms/server/usecase/excel"
	pdfCase "github.com/ebedevelopment/next-gen-tms/server/usecase/pdf"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
)

// OperationRecordService operation db
type OperationRecordService struct {
	operationRecordRepository repository.OperationRecordRepository
	excelUsecase              useCase.ExcelUseCase
	pdfUsecase      pdfCase.PDFUseCase
}

// CreateSysOperationRecord create record
func (p *OperationRecordService) CreateSysOperationRecord(sysOperationRecord system.SysOperationRecord) error {
	return p.operationRecordRepository.CreateSysOperationRecord(sysOperationRecord)
}

// GetSysOperationRecord item operations record
func (p *OperationRecordService) GetSysOperationRecord(id string) (sysResp.AuditLogResp, error) {
	return p.operationRecordRepository.GetSysOperationRecord(id)
}

// GetSysOperationRecordInfoList get operation record
func (p *OperationRecordService) GetSysOperationRecordInfoList(info sysReq.SysOperationRecordSearch) ([]sysResp.AuditLogResp, int64, error) {
	return p.operationRecordRepository.GetSysOperationRecordInfoList(info)
}

// ParseLogInfoList2Excel excel for Log Operation
func (p *OperationRecordService) ParseLogInfoList2Excel(info []sysResp.AuditLogResp, format string) (string, error) {

	header := []string{"Operation Status", "Status code", "Name", "IP", "Method", "Path", "Action Type",
		"Latency", "Agent", "resp", "Action By", "Created Date"}

	data := [][]string{}

	if format == utils.Excel {
		data = append(data, header)
	}

	for _, p := range info {
		var s []string
		s = append(s, p.StatusOp, fmt.Sprintf("%v", p.Status), p.Name, p.IP, p.Method, p.Path, p.ActionType,
			p.Latency, p.Agent, fmt.Sprintf("%v", p.Resp), p.UserEmail, p.CreatedAt)

		data = append(data, s)
	}

	var filePath string
	if format == utils.PDF {
		filePath = global.GvaConfig.Excel.Dir + "auditLogExport_" + fmt.Sprintf("%v", time.Now().Unix()) + ".pdf"
		return p.pdfUsecase.CreatePDF(filePath, header, data, "audit Log Data", pdfCase.Platform)
	}

	filePath = global.GvaConfig.Excel.Dir + "auditLogExport_" + fmt.Sprintf("%v", time.Now().Unix()) + ".xlsx"

	return p.excelUsecase.GenerateExcelSheet(data, filePath)
}
