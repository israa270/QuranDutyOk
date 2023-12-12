package router

import (
	"github.com/ebedevelopment/next-gen-tms/server/router/management"
	"github.com/ebedevelopment/next-gen-tms/server/router/system"
	"github.com/ebedevelopment/next-gen-tms/server/router/tms"
)

// GroupRouter  struct
type GroupRouter struct {
	System system.GroupRouter
	Tms    tms.GroupRouter
	Management  management.GroupRouter
}

var GroupRouterApp = new(GroupRouter)
