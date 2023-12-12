package response


// AuditLogResp audit log resp
type AuditLogResp struct {
	ID           uint        `json:"id" form:"id"`
	IP           string      `json:"ip" form:"ip"`
	Method       string      `json:"method" form:"method"`
	Path         string      `json:"path" form:"path"`
	Status       int         `json:"status" form:"status"`
	StatusOp     string      `json:"statusOperation" form:"statusOperation"`
	Latency      string      `json:"latency" form:"latency"`
	Agent        string      `json:"agent" form:"agent"`
	Body         interface{} `json:"body,omitempty" form:"body"`
	Resp         interface{} `json:"resp,omitempty" form:"resp"`
	UserName     string      `json:"userName" form:"userName"`
	UserEmail    string      `json:"userEmail" form:"userEmail"`
	CreatedAt    string      `json:"createdAt" form:"createdAt"`
	Name         string      `json:"name" form:"name"`
	ActionType   string      `json:"actionType" form:"actionType"`
}

type RecordOperation struct {
	Ip           string        `json:"ip" form:"ip"`
	Method       string        `json:"method" form:"method"`
	Path         string        `json:"path" form:"path"`
	Status       int           `json:"status" form:"status"`
	StatusOp     string        `json:"statusOperation" form:"statusOperation"`
	Latency      string `json:"latency,omitempty" form:"latency"`
	Agent        string        `json:"agent" form:"agent"`
	// ErrorMessage string        `json:"errorMessage,omitempty" form:"errorMessage"`
	Body         interface{}   `json:"body,omitempty" form:"body"`
	Resp         interface{}   `json:"resp,omitempty" form:"resp"`
	UserName     string        `json:"userName,omitempty" form:"userName"`
	UserEmail    string        `json:"userEmail" form:"userEmail"`
	Name         string        `json:"name,omitempty" form:"name"`
	ActionType   string        `json:"actionType" form:"actionType"`
	AuthorityId  string        `json:"authorityId" form:"authorityId"`
}
