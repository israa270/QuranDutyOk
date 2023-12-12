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

func (m *HomeWorkRepository) CheckHomeWorkName(name string, versionName string) bool {
	var homeWork model.HomeWork
	err := global.GvaDB.Where("name = ? AND version_name =?", name, versionName).First(&homeWork).Error
	return err == nil
}

func (m *HomeWorkRepository) GetHomeWorkID(name string, versionName string) (uint, error) {
	var homeWork model.HomeWork
	err := global.GvaDB.Where("name = ? AND version_name =?", name, versionName).First(&homeWork).Error
	return homeWork.ID, err
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