package user

import (
	"github.com/ebedevelopment/next-gen-tms/server/model/system"
	sysResp "github.com/ebedevelopment/next-gen-tms/server/model/system/response"
)


// ConvertUserToResponse  response return
func (s *UserUseCase) ConvertUserToResponse(u *system.SysUser) *sysResp.SysUserResponse {
	user := &sysResp.SysUserResponse{
		ID:            u.ID,
		Email:         u.Email,
		AuthorityId:   u.AuthorityId,
		Phone:         u.Phone,
		RegisterTime:  u.RegisterTime.Format("2006-01-02 15:04:05"),
		LastLoginTime: u.LastLoginTime,
		Status:        u.Status,
		UserStatus:    u.UserStatus,
	}
	return user
}