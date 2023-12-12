package response

import (
	"github.com/ebedevelopment/next-gen-tms/server/model/system"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
)

// SysUserResponse struct
type SysUserResponse struct {
	ID            uint   `json:"ID"`
	Username      string `json:"userName"`
	Email         string `json:"email"`
	AuthorityId   string `json:"authorityId,omitempty"`
	Authority     string `json:"authority,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Status        string `json:"status,omitempty"`
	UserStatus    bool   `json:"userStatus"`
	RegisterTime  string `json:"registerTime,omitempty"`
	LastLoginTime string `json:"lastLoginTime,omitempty"`
	CreatedAt     string `json:"createdAt,omitempty"`
	UpdatedAt     string `json:"updatedAt,omitempty"`
}

// LoginResponse login response
type LoginResponse struct {
	User      SysUserResponse `json:"user"`
	Token     string          `json:"token"`
	ExpiresAt int64           `json:"expiresAt"`
}

// UserData user data for token
type UserData struct {
	ID            uint   `json:"id"`
	UserEmail     string `json:"userEmail"`

	AuthorityCode string `json:"authorityCode"`
}

// For UI response.
type UserCasbin struct {
	User          *system.SysUser
	CollectionMap map[string][]sysReq.CasbinInfo
}

type UserPasswordCommand struct{
	Email string 
	Password string
}
