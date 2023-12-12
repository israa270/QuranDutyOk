package system

import (
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/system"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
)

// FindUserById  get user by id  without details
func (s *UserRepository) FindUserById(id int) (user *system.SysUser, err error) {
	var u system.SysUser
	err = global.GvaDB.Where("id= ?", id).First(&u).Error
	return &u, err
}

func (s *UserRepository) CheckUserById(id uint) bool {
	var u *system.SysUser
	err := global.GvaDB.Where("id= ?", id).First(&u).Error
	return err == nil
}

// GetUserEmail
func (s *UserRepository) GetUserEmail(id uint) (string, bool) {
	var u *system.SysUser
	err := global.GvaDB.Where("id= ?", id).First(&u).Error
	return u.Email, err == nil
}

// GetUserByEmail
func (s *UserRepository) GetUserByEmail(email string) (u *system.SysUser, err error) {
	err = global.GvaDB.Where("email= ?", email).First(&u).Error
	return
}

// GetUserInfo get user info
func (s *UserRepository) GetUserInfo(id uint) (*system.SysUser, error) {
	var user *system.SysUser
	if err := global.GvaDB.Preload("Authority").First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// CheckUserByEmail
func (s *UserRepository) CheckUserByEmail(email string) bool {
	var u *system.SysUser
	err := global.GvaDB.Where("email= ?", email).First(&u).Error
	return err == nil
}

// GetUserInfoList paging data
func (s *UserRepository) GetUserInfoList(info sysReq.UserSearch) (list []system.SysUser, total int64, err error) {
	if info.PageSize == 0 || info.PageSize > global.GvaConfig.Mysql.LimitRecords {
		info.PageSize = global.GvaConfig.Mysql.LimitRecords
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GvaDB.Model(&system.SysUser{})
	var userList []system.SysUser

	if info.Email != "" {
		db = db.Where("email = ?", info.Email)
	}

	if info.Status != "" {
		db = db.Where("status =? ", info.Status)
	}

	if info.AuthorityId != "" {
		db = db.Where("authority_id = ?", info.AuthorityId)
	}

	t := time.Time{}
	if info.CreatedBefore != t && info.CreatedAfter != t {
		db.Where("created_at  >= ? AND created_at  <= ?", info.CreatedBefore, info.CreatedAfter)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Order("id "+info.Order).Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error

	return userList, total, err
}
