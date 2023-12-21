package tms

import "github.com/ebedevelopment/next-gen-tms/server/global"

// AuditOperation
type AuditOperation struct {
	global.GvaModel

	ResellerId   uint   `json:"resellerId" form:"resellerId" gorm:"column:reseller_id"`
	EntityID     int    `json:"entityId" form:"entityId" gorm:"column:entity_id"`
	ClientIP     int    `json:"clientIp" form:"clientIp" gorm:"column:client_ip"`
	RequestURI   string `json:"requestUri" form:"requestUri" gorm:"column:request_uri;type:text"`
	AuditType    string `json:"auditType" form:"auditType" gorm:"column:audit_type"`
	AuditAction  string `json:"auditAction" form:"auditAction" gorm:"column:audit_action"`
	AuditAction2 string `json:"auditAction2" form:"auditAction2" gorm:"column:audit_action2"`
	HTTPCode     int    `json:"httpCode" form:"httpCode" gorm:"column:http_code"`
	BizCode      int    `json:"bizCode" form:"bizCode" gorm:"column:biz_code"`
	BizMessage   string `json:"bizMessage" form:"bizMessage" gorm:"column:biz_message"`
	CreatedBy    string `json:"createdBy" form:"createdBy" gorm:"column:created_by"`
}

// TableName audit operation
func (AuditOperation) TableName() string {
	return "tms_audit_operation"
}
