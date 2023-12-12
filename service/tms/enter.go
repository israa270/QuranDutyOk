package tms

import (
	"github.com/ebedevelopment/next-gen-tms/server/service/tms/admin"
	"github.com/ebedevelopment/next-gen-tms/server/service/tms/data"
)

// ServiceGroup  struct
type ServiceGroup struct {
	admin.ServiceGroupAdmin
	data.ServiceGroupData
}
