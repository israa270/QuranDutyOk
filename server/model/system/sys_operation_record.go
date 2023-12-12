// Package system SysOperation Record
package system

import (
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
)

// SysOperationRecord operation record for apis
type SysOperationRecord struct {
	global.GvaModel
	Name         string        `json:"name" form:"name" gorm:"column:name"`
	IP           string        `json:"ip" form:"ip" gorm:"column:ip;comment:please request ip"`                 // please request ip
	Method       string        `json:"method" form:"method" gorm:"column:method;comment:please request method"` // please request method
	Path         string        `json:"path" form:"path" gorm:"column:path;comment:please request path"`         // please request path
	ActionType   string        `json:"actionType" form:"actionType" gorm:"column:action_type"`
	Status       int           `json:"status" form:"status" gorm:"column:status;comment:please request latency"` // please request latency
	StatusOp     string        `json:"statusOperation" gorm:"column:status_operation"`
	Latency      time.Duration `json:"latency" form:"latency" gorm:"column:latency;comment:Delay" swaggertype:"string"`   // Delay
	Agent        string        `json:"agent" form:"agent" gorm:"column:agent;comment:agent"`                              // agent
	// ErrorMessage string        `json:"errorMessage" form:"errorMessage" gorm:"column:error_message;comment:mistake info"` // mistake info
	Body         string        `json:"body" form:"body" gorm:"type:text;column:body;comment:please requestBody"`          // please request Body
	Resp         string        `json:"resp" form:"resp" gorm:"type:text;column:resp;comment:resp Body"`                   // resp Body
	UserID       int           `json:"userId" form:"userId" gorm:"column:user_id;comment:User id"`                        // User id
	User         SysUser       `json:"-" form:"user" gorm:"foreignKey:UserID;references:ID"`
}


// TableName SysOperationRecord table Name
func (SysOperationRecord) TableName() string {
	return "sys_operation_records"
}
