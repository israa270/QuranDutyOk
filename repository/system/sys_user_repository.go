package system

import (
	"errors"
	"fmt"
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/system"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
	"go.uber.org/zap"
)

// UserRepository user db fn
type UserRepository struct {
}

// Register User register
func (s *UserRepository) Register(r sysReq.Register, code string, createdBy string) (system.SysUser, error) {
	var u system.SysUser

	u.Email = r.Email
	u.AuthorityId = r.AuthorityId
	u.CreatedBy = createdBy
	u.RegisterTime = time.Now()
	u.Status = utils.UserPending
	u.UserStatus = true


	err := global.GvaDB.Create(&u).Error
	return u, err
}

// Login User login
func (s *UserRepository) Login(u *system.SysUser, ipAddress string) (*system.SysUser, error) {
	var user system.SysUser
	if err := global.GvaDB.Where("email = ?", u.Email).Preload("Authority").First(&user).Error; err != nil {
		return &user, errors.New("failed to get username")
	}

	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New(global.Translate("sysUser.wrongPassword"))
	}

	user.LastLoginTime = time.Now().Format("2006-01-02 15:04:05")
	user.Status = utils.StatusUserActive

	//Update last login data , Activate date
	if err := global.GvaDB.Model(&user).Where("email = ?", user.Email).Save(&user).Error; err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].UpdateLastLoginFail, zap.Error(err))
	}

	return &user, nil
}

// ChangePassword updateUser password
func (s *UserRepository) ChangePassword(u *system.SysUser) error {
	err := global.GvaDB.Save(&u).Error
	return err
}


func (s *UserRepository) ChangeUserImageProfile(email string,imgProfile string) error {
	err := global.GvaDB.Model(&system.SysUser{}).Where("email = ?", email).Updates(map[string]interface{}{
		"header_img":      imgProfile,
	}).Error

	return err
}

// ConfirmPassword updateUser password
func (s *UserRepository) ConfirmPassword(user *system.SysUser) error {
	user.Status = utils.StatusRegister
	user.UserStatus = true
	user.IsVerified = true
	user.ActivateDate = time.Now().Format("2006-01-02 15:04:05")
	user.Password = utils.BcryptHash(user.Password)

	err := global.GvaDB.Save(&user).Error
	return err
}



// DeleteUser delete User
func (s *UserRepository) DeleteUser(id int) (err error) {
	var user system.SysUser
	err = global.GvaDB.Where("id = ?", id).Delete(&user).Error
	return err
}

// SetUserInfo setup User info
func (s *UserRepository) SetUserInfo(req system.SysUser) error {
	var u system.SysUser
	if req.Phone != "" {
		if err := global.GvaDB.Where("phone = ?", req.Phone).First(&u).Error; err == nil {
			if req.ID != u.ID {
				return fmt.Errorf(global.Translate("general.duplicateValuePhone"))
			}
		}
	}
	return global.GvaDB.Where("email=?", req.Email).Updates(&req).Error
}

// UpdateUserStatus
func (s *UserRepository) UpdateUserStatus(req *system.SysUser) error {
	return global.GvaDB.Where("id=?", req.ID).Updates(&req).Error
}

