package router

import (
	"github.com/ebedevelopment/next-gen-tms/server/router/data"
	"github.com/ebedevelopment/next-gen-tms/server/router/management"
	"github.com/ebedevelopment/next-gen-tms/server/router/system"
)

// GroupRouter  struct
type GroupRouter struct {
	System system.GroupRouter
	Data    data.GroupRouter
	Management  management.GroupRouter
}

var GroupRouterApp = new(GroupRouter)
