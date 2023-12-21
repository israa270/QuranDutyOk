package system

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/system"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
	sysResp "github.com/ebedevelopment/next-gen-tms/server/model/system/response"
	"go.uber.org/zap"
)

// OperationRecordRepository operation db
type OperationRecordRepository struct{}

// CreateSysOperationRecord create record
func (p *OperationRecordRepository) CreateSysOperationRecord(sysOperationRecord system.SysOperationRecord) (err error) {
	err = global.GvaDB.Create(&sysOperationRecord).Error
	return err
}

// // DeleteSysOperationRecordByIds batch delete  record
// func (p *OperationRecordRepository) DeleteSysOperationRecordByIds(ids request.IdsReq) (err error) {
// 	err = global.GvaDB.Delete(&[]system.SysOperationRecord{}, "id in (?)", ids.Ids).Error
// 	return err
// }

// // DeleteSysOperationRecord delete operations record
// func (p *OperationRecordRepository) DeleteSysOperationRecord(id string) (err error) {
// 	var sysOperationRecord system.SysOperationRecord
// 	err = global.GvaDB.Where("id=? ", id).Delete(&sysOperationRecord).Error
// 	return err
// }

// GetSysOperationRecord getByIDone item operations record
func (p *OperationRecordRepository) GetSysOperationRecord(id string) (op sysResp.AuditLogResp, err error) {
	var operation system.SysOperationRecord
	err = global.GvaDB.Where("id = ?", id).Preload("User").First(&operation).Error

	op = sysResp.AuditLogResp{
		ID:       operation.ID,
		Method:   operation.Method,
		IP:       operation.IP,
		Path:     operation.Path,
		Status:   operation.Status,
		StatusOp: operation.StatusOp,
		Latency:  fmt.Sprintf("%v", time.Duration(operation.Latency.Microseconds())),
		Agent:    operation.Agent,
		Body:     operation.Body,
		Resp:     operation.Resp,
		// ErrorMessage: operation.ErrorMessage,

		UserEmail:  operation.User.Email,
		CreatedAt:  operation.CreatedAt.Format("2006-01-02 15:04:05"),
		Name:       operation.Name,
		ActionType: operation.ActionType,
	}
	var m interface{}
	if err := json.Unmarshal([]byte(operation.Body), &m); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GetDataFail, zap.Error(err))
	}

	op.Body = m

	if err := json.Unmarshal([]byte(operation.Resp), &m); err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GetDataFail, zap.Error(err))
	}
	op.Resp = m

	return
}

// GetSysOperationRecordInfoList get operation record
func (p *OperationRecordRepository) GetSysOperationRecordInfoList(info sysReq.SysOperationRecordSearch) (list []sysResp.AuditLogResp, total int64, err error) {
	if info.PageSize == 0 || info.PageSize > global.GvaConfig.Mysql.LimitRecords {
		info.PageSize = global.GvaConfig.Mysql.LimitRecords
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// create db
	db := global.GvaDB.Model(&system.SysOperationRecord{})
	var sysOperationRecords []system.SysOperationRecord
	var operationRecords []sysResp.AuditLogResp
	// If conditional search create search
	if info.Name != "" {
		db = db.Where("name= ? ", info.Name)
	}

	if info.Ip != "" {
		db = db.Where("ip =? ", info.Ip)
	}

	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.StatusOp != "" {
		db = db.Where("status_operation = ?", info.StatusOp)
	}

	t := time.Time{}
	if info.CreatedBefore != t && info.CreatedAfter != t {
		db.Where("created_at  >= ? AND created_at  <= ?", info.CreatedBefore, info.CreatedAfter)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	//Operation For Sub Org and parent
	err = db.Order("id " + info.Order).Limit(limit).Offset(offset).Preload("User").Find(&sysOperationRecords).Error

	for _, operation := range sysOperationRecords {
		op := sysResp.AuditLogResp{
			ID:       operation.ID,
			Method:   operation.Method,
			IP:       operation.IP,
			Path:     operation.Path,
			Status:   operation.Status,
			StatusOp: operation.StatusOp,
			Latency:  fmt.Sprintf("%v", time.Duration(operation.Latency.Microseconds())),
			Agent:    operation.Agent,
			Body:     operation.Body,
			Resp:     operation.Resp,
			// ErrorMessage: operation.ErrorMessage,

			UserEmail:  operation.User.Email,
			CreatedAt:  operation.CreatedAt.Format("2006-01-02 15:04:05"),
			Name:       operation.Name,
			ActionType: operation.ActionType,
		}

		var m interface{}
		if operation.Body != "" {
			if err := json.Unmarshal([]byte(operation.Body), &m); err != nil {
				global.GvaLog.Error(global.GvaLoggerMessage["log"].GetDataFail, zap.Error(err))
			}
			op.Body = m
		}

		var mResp interface{}
		if operation.Resp != "" {
			if err := json.Unmarshal([]byte(operation.Resp), &mResp); err != nil {
				global.GvaLog.Error(global.GvaLoggerMessage["log"].GetDataFail, zap.Error(err))
			}
			op.Resp = mResp
		}

		operationRecords = append(operationRecords, op)
	}
	//sort desc operation
	sort.SliceStable(operationRecords, func(i, j int) bool {
		return operationRecords[i].ID > operationRecords[j].ID
	})

	return operationRecords, total, err
}
