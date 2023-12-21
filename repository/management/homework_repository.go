package management

import (
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	model "github.com/ebedevelopment/next-gen-tms/server/model/management"
	"github.com/ebedevelopment/next-gen-tms/server/model/management/request"
)

type HomeWorkRepository struct{}

// CreateHomeWork createHomeWorkRecord
func (m *HomeWorkRepository) CreateHomeWork(homeWork model.HomeWork) (err error) {
	err = global.GvaDB.Create(&homeWork).Error

	return err
}


func (m *HomeWorkRepository) GetHomeWorkID(id uint) (model.HomeWork, error) {
	var homeWork model.HomeWork
	err := global.GvaDB.Where("id =?", id).First(&homeWork).Error
	return homeWork, err
}

func (m *HomeWorkRepository) AssignHomeWorkToClass(homeworkUpdated model.HomeWork) error{
	err := global.GvaDB.Where("id = ?", homeworkUpdated.ID).Updates(&homeworkUpdated).Error
	return err
}



func (m *HomeWorkRepository) GetClassHomework(classId uint) ([]model.HomeworkClasses, error){
	var classHomework  []model.HomeworkClasses
	err := global.GvaDB.Where("class_id = ?", classId).Find(&classHomework).Error
	return classHomework, err
}



func (m *HomeWorkRepository) GetHomeWorkList(info request.HomeWorkSearch)([]model.HomeWork, int64, error){
	if info.PageSize == 0 || info.PageSize > global.GvaConfig.Mysql.LimitRecords {
		info.PageSize = global.GvaConfig.Mysql.LimitRecords
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// create db
	db := global.GvaDB.Model(&model.HomeWork{})
	var homeWorkes []model.HomeWork
	// If conditional search create search
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}

	if info.ExpireStatus{
		db = db.Where("expire_status =? " , info.ExpireStatus)
	}

	t := time.Time{}
	if info.CreatedBefore != t && info.CreatedAfter != t {
		db.Where("created_at  >= ? AND created_at  <= ?", info.CreatedBefore, info.CreatedAfter)
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return  nil, total, err
	}

	err = db.Order("id " + info.Order).Limit(limit).Offset(offset).Find(&homeWorkes).Error

	return homeWorkes, total, err
}