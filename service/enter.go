package service

import (
	"github.com/ebedevelopment/next-gen-tms/server/service/system"
	"github.com/ebedevelopment/next-gen-tms/server/service/tms"
)

// ServiceGroup
type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	TmsServiceGroup    tms.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
