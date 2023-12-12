package system

import (
	"github.com/ebedevelopment/next-gen-tms/server/model/system"
	repository "github.com/ebedevelopment/next-gen-tms/server/repository/system"
)


type LoginAttemptService struct{
	loginAttemptRepository  repository.LoginAttemptRepository
} 

// CheckUserIsBlocked or not
func (s *LoginAttemptService) CheckUserBlock(email string) (system.SysLoginAttempt, error) {
	return s.loginAttemptRepository.CheckUserBlock(email)
}

// func (userService *UserService) CreateUserLoginAttempt(loginAttempt system.SysLoginAttempt) (err error) {
// 	err = global.GvaDB.Create(&loginAttempt).Error
// 	return err
// }

func (s *LoginAttemptService) UpdateUserLoginAttempt(loginObj *system.SysLoginAttempt)  error {
	return s.loginAttemptRepository.UpdateUserLoginAttempt(loginObj)
}
