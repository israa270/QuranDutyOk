package response

import "github.com/ebedevelopment/next-gen-tms/server/config"

// SysConfigResponse config response
type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
