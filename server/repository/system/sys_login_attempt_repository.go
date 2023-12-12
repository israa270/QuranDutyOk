package system

import (


	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/system"

)

type LoginAttemptRepository struct{

}

// // UpdateUserBlockTime update user block time
func (s *LoginAttemptRepository) UpdateUserBlockTime(loginObj system.SysLoginAttempt) (err error) {
	err = global.GvaDB.Where("email = ?", loginObj.Email).Save(&loginObj).Error
	return
}

// CheckUserIsBlocked or not
func (s *LoginAttemptRepository) CheckUserBlock(email string) (system.SysLoginAttempt, error) {
	var loginObj system.SysLoginAttempt
	err := global.GvaDB.Where("email = ?", email).First(&loginObj).Error
	return loginObj, err
}

// func (userService *UserService) CreateUserLoginAttempt(loginAttempt system.SysLoginAttempt) (err error) {
// 	err = global.GvaDB.Create(&loginAttempt).Error
// 	return err
// }

func (s *LoginAttemptRepository) UpdateUserLoginAttempt(loginObj *system.SysLoginAttempt) error {
	err := global.GvaDB.Where("email = ?", loginObj.Email).Save(&loginObj).Error
	return err
}
