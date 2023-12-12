package management

import (
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	model "github.com/ebedevelopment/next-gen-tms/server/model/management"
	"github.com/ebedevelopment/next-gen-tms/server/model/management/request"
)

type ClassRepository struct{}

// CreateClass createClassRecord
func (m *ClassRepository) CreateClass(Class model.Class) (err error) {
	err = global.GvaDB.Create(&Class).Error

	return err
}

func (m *ClassRepository) CheckClassName(name string, versionName string) bool {
	var class model.Class
	err := global.GvaDB.Where("name = ? AND version_name =?", name, versionName).First(&class).Error
	return err == nil
}

func (m *ClassRepository) GetClassID(name string, versionName string) (uint, error) {
	var class model.Class
	err := global.GvaDB.Where("name = ? AND version_name =?", name, versionName).First(&class).Error
	return class.ID, err
}




func (m *ClassRepository) GetClassList(info request.ClassSearch)([]model.Class, int64, error){
	if info.PageSize == 0 || info.PageSize > global.GvaConfig.Mysql.LimitRecords {
		info.PageSize = global.GvaConfig.Mysql.LimitRecords
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// create db
	db := global.GvaDB.Model(&model.Class{})
	var classes []model.Class
	// If conditional search create search
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.VersionName != "" {
		db = db.Where("version_name = ?", "%"+info.VersionName+"%")
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

	err = db.Order("id " + info.Order).Limit(limit).Offset(offset).Find(&classes).Error

	return classes, total, err
}