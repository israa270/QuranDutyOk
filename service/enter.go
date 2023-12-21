package service

import (
	"github.com/ebedevelopment/next-gen-tms/server/service/system"
	"github.com/ebedevelopment/next-gen-tms/server/service/data"
)

// ServiceGroup
type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
    DataServiceGroup    data.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
