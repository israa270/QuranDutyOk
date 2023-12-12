package management

import (
	"errors"
	"fmt"
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	model "github.com/ebedevelopment/next-gen-tms/server/model/management"
	"github.com/ebedevelopment/next-gen-tms/server/model/management/request"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
	"go.uber.org/zap"
)

type TeacherRepository struct{}

// CreateTeacher createTeacherRecord
func (m *TeacherRepository) CreateTeacher(teacher model.Teacher) (err error) {
	if err = global.GvaDB.Create(&teacher).Error; err != nil {
		return err
	}

	//Update Username and password
	teacher.UserName = "teacher" + fmt.Sprintf("%v", teacher.ID)
	teacher.Password = utils.BcryptHash("asd" + teacher.UserName)

	err = global.GvaDB.Where("id", teacher.ID).Updates(&teacher).Error
	return err
}

func (m *TeacherRepository) CheckTeacherName(name string) bool {
	var teacher model.Teacher
	err := global.GvaDB.Where("name = ?", name).First(&teacher).Error
	return err == nil
}

func (m *TeacherRepository) CheckTeacherExist(id uint) bool {
	var teacher model.Teacher
	err := global.GvaDB.Where("id = ?", id).First(&teacher).Error
	return err == nil
}

func (m *TeacherRepository) GetTeacherList(info request.ListSearch) ([]model.Teacher, int64, error) {
	if info.PageSize == 0 || info.PageSize > global.GvaConfig.Mysql.LimitRecords {
		info.PageSize = global.GvaConfig.Mysql.LimitRecords
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// create db
	db := global.GvaDB.Model(&model.Teacher{})
	var teachers []model.Teacher
	// If conditional search create search
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.UserName != "" {
		db = db.Where("username = ?", "%"+info.UserName+"%")
	}

	t := time.Time{}
	if info.CreatedBefore != t && info.CreatedAfter != t {
		db.Where("created_at  >= ? AND created_at  <= ?", info.CreatedBefore, info.CreatedAfter)
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}

	err = db.Order("id " + info.Order).Limit(limit).Offset(offset).Find(&teachers).Error

	return teachers, total, err
}

func (m *TeacherRepository) CheckTeacherByUserName(name string) bool {
	var teacher model.Teacher
	err := global.GvaDB.Where("userName = ?", name).First(&teacher).Error
	return err == nil
}

// Login User login
func (s *TeacherRepository) Login(u *model.Teacher) (*model.Teacher, error) {
	var teacher model.Teacher
	if err := global.GvaDB.Where("username = ?", u.UserName).First(&teacher).Error; err != nil {
		return &teacher, errors.New("failed to get username")
	}

	if ok := utils.BcryptCheck(u.Password, teacher.Password); !ok {
		return nil, errors.New(global.Translate("sysUser.wrongPassword"))
	}

	teacher.LastLoginTime = time.Now().Format("2006-01-02 15:04:05")

	//Update last login data , Activate date
	if err := global.GvaDB.Model(&teacher).Where("username = ?", u.UserName).Save(&teacher).Error; err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].UpdateLastLoginFail, zap.Error(err))
	}

	return &teacher, nil
}
