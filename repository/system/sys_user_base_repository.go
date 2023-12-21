package system

import (
	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/system"
)

// UpdateExpireFlag
func (s *UserRepository) UpdateExpireFlag(id uint, expireEmail string) error {
	var u *system.SysUser
	err := global.GvaDB.Where("id = ?", id).First(&u).Update("expire_flag", expireEmail).Error
	return err
}

// UpdateResetPasswordFlag update reset password flag
func (s *UserRepository) UpdateResetPasswordFlag(u *system.SysUser) error {
	err := global.GvaDB.Where("id = ?", u.ID).Save(&u).Error
	return err
}

// ResetPassword
func (s *UserRepository) ResetPassword(user *system.SysUser) (err error) {
	err = global.GvaDB.Where("id=?", user.ID).Save(&user).Error

	return err
}

func (s *UserRepository) GetUserByResetToken(resetToken string, resetAt int64) (u *system.SysUser, err error) {
	err = global.GvaDB.Where("password_reset_token = ? AND password_reset_at > ? ", resetToken, resetAt).First(&u).Error
	return
}

func (s *UserRepository) GetUserByVerifyCode(verifyCode string) (u *system.SysUser, err error) {
	err = global.GvaDB.Where("verification_code = ?", verifyCode).First(&u).Error
	return
}
