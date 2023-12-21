package user

import (
	"github.com/ebedevelopment/next-gen-tms/server/service/system"
)

type UserController struct {
	
	userService system.UserService
	jwtService  system.JwtService
}