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

type StudentRepository struct {
}

// CreateStudent createStudentRecord
func (m *StudentRepository) CreateStudent(student model.Student) (username string, err error) {
	if err = global.GvaDB.Create(&student).Error; err != nil {
		return "", err
	}

	//Update Username and password
	student.UserName = "student" + fmt.Sprintf("%v", student.ID)
	student.Password = utils.BcryptHash("asd" + student.UserName)

	err = global.GvaDB.Where("id",student.ID).Updates(&student).Error
	return student.UserName, err
}

func (m *StudentRepository) MoveStudent(studentId uint, newClassId uint) error {
	var student model.Student
	err := global.GvaDB.Where("id = ?", studentId).First(&student).Update("class_id",newClassId).Error
	return err
}





func (m *StudentRepository) CheckStudentName(name string) bool {
	var student model.Student
	err := global.GvaDB.Where("name = ?", name).First(&student).Error
	return err == nil
}




func (m *StudentRepository) CheckStudentExist(id uint) bool {
	var student model.Student
	err := global.GvaDB.Where("id = ?", id).First(&student).Error
	return err == nil
}

func (m *StudentRepository) CheckStudentClassExist(studentId uint, classId uint) bool {
	var student model.Student
	err := global.GvaDB.Where("id = ? AND class_id ", studentId, classId).First(&student).Error
	return err == nil
}



func (m *StudentRepository) GetStudentList(info request.StudentSearch) ([]model.Student, int64, error) {
	if info.PageSize == 0 || info.PageSize > global.GvaConfig.Mysql.LimitRecords {
		info.PageSize = global.GvaConfig.Mysql.LimitRecords
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// create db
	db := global.GvaDB.Model(&model.Student{})
	var students []model.Student
	// If conditional search create search
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.UserName != "" {
		db = db.Where("username = ?", "%"+info.UserName+"%")
	}

	if info.ClassID != 0 {
		db = db.Where("class_id = ?", info.ClassID)
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

	err = db.Order("id " + info.Order).Limit(limit).Offset(offset).Find(&students).Error

	return students, total, err
}

func (m *StudentRepository) CheckStudentByUserName(name string) bool {
	var student model.Student
	err := global.GvaDB.Where("userName = ?", name).First(&student).Error
	return err == nil
}

// Login User login
func (s *StudentRepository) Login(u *model.Student) (*model.Student, error) {
	var student model.Student
	if err := global.GvaDB.Where("username = ?", u.UserName).First(&student).Error; err != nil {
		return &student, errors.New("failed to get username")
	}

	if ok := utils.BcryptCheck(u.Password, student.Password); !ok {
		return nil, errors.New(global.Translate("sysUser.wrongPassword"))
	}

	student.LastLoginTime = time.Now().Format("2006-01-02 15:04:05")

	//Update last login data , Activate date
	if err := global.GvaDB.Model(&student).Where("username = ?", u.UserName).Save(&student).Error; err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].UpdateLastLoginFail, zap.Error(err))
	}

	return &student, nil
}
