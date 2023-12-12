package system

import "github.com/ebedevelopment/next-gen-tms/server/model/system"



// UpdateExpireFlag in expire password
func (s *UserService) UpdateExpireFlag(id uint, expireEmail string) error {
	return s.userRepository.UpdateExpireFlag(id, expireEmail)
}


// UpdateResetPasswordFlag update reset password flag
func (s *UserService) UpdateResetPasswordFlag(u *system.SysUser) error {
	return s.userRepository.UpdateResetPasswordFlag(u)
}

// ResetPassword
func (s *UserService) ResetPassword(user *system.SysUser) error {
	return s.userRepository.ResetPassword(user)
}

func (s *UserService) GetUserByResetToken(resetToken string, resetAt int64) (u *system.SysUser, err error) {
	return s.userRepository.GetUserByResetToken(resetToken, resetAt)
}


func (s *UserService) GetUserByVerifyCode(verifyCode string) (u *system.SysUser, err error) {
	return s.userRepository.GetUserByVerifyCode(verifyCode)
}
