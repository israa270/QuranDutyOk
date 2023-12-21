package request

import "github.com/ebedevelopment/next-gen-tms/server/model/common/request"

// AuthorityRequest struct
type AuthorityRequest struct {
	AuthorityId   string `json:"authorityId" form:"authorityId" binding:"required"`
	AuthorityName string `json:"authorityName" form:"authorityName" binding:"required"`
	PrivilegeIds  []uint `json:"privilegeIds" form:"privilegeIds"`
}

// AuthoritySearch struct
type AuthoritySearch struct {
	AuthorityName string `json:"authorityName" form:"authorityName"`
	request.PageInfo
}
