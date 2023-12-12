package service

import (
	"github.com/ebedevelopment/next-gen-tms/server/service/system"

)

// ServiceGroup
type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
