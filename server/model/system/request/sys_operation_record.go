package request

import (
	"github.com/ebedevelopment/next-gen-tms/server/model/common/request"
)

// SysOperationRecordSearch operation search
type SysOperationRecordSearch struct {
	Name       string `json:"name" form:"name"`
	Ip         string `json:"ip" form:"ip"`
	Method     string `json:"method" form:"method"`
	StatusOp   string `json:"status" form:"status"`
	Path       string `json:"path" form:"path"`
	ActionTime string `json:"actionTime" form:"actionTime"`
	request.PageInfo
}
