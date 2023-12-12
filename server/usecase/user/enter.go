package user

import (
	"github.com/ebedevelopment/next-gen-tms/server/service/system"
	"github.com/ebedevelopment/next-gen-tms/server/usecase/user/claim"
)



type UserUseCase struct{
	userService      system.UserService
	loginAttemptService  system.LoginAttemptService
	ClaimUseCase claim.ClaimUseCase
}