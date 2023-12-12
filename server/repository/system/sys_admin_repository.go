package system

import (
	"errors"
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	model "github.com/ebedevelopment/next-gen-tms/server/model/system"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
	"go.uber.org/zap"
)


type AdminRepository struct{

}




func (m *AdminRepository) CreateAdmin(admins []model.Admin) error{
	err := global.GvaDB.Create(&admins).Error
	return err
}

// func (m *AdminRepository) GetAdminByUserName(name  string) (model.Admin, error) {
// 	var admin model.Admin
// 	err := global.GvaDB.Where("userName = ?", name).First(&admin).Error
// 	return   admin, err
// }


func (m *AdminRepository) CheckAdminByUserName(name  string) bool{
	var admin model.Admin
	err := global.GvaDB.Where("userName = ?", name).First(&admin).Error
	return   err == nil
}




// Login User login
func (s *AdminRepository) Login(u *model.Admin) (*model.Admin, error) {
	var admin model.Admin
	if err := global.GvaDB.Where("username = ?", u.UserName).First(&admin).Error; err != nil {
		return &admin, errors.New("failed to get username")
	}

	if ok := utils.BcryptCheck(u.Password, admin.Password); !ok {
		return nil, errors.New(global.Translate("sysUser.wrongPassword"))
	}

	admin.LastLoginTime = time.Now().Format("2006-01-02 15:04:05")
	
	//Update last login data , Activate date
	if err := global.GvaDB.Model(&admin).Where("username = ?", u.UserName).Save(&admin).Error; err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].UpdateLastLoginFail, zap.Error(err))
	}

	return &admin, nil
}