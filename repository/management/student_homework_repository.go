package management

import (
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	model "github.com/ebedevelopment/next-gen-tms/server/model/management"
	"github.com/ebedevelopment/next-gen-tms/server/model/management/request"
)




type StudentHomeWorkRepository struct {
}

// CreateStudentHomework createStudentHomeWorkRecord
func (m *StudentHomeWorkRepository) CreateStudentHomeworks(stWork []model.StudentHomeWorks) (err error) {
	 err = global.GvaDB.Create(&stWork).Error
	 return err
}


func (m *StudentHomeWorkRepository) UpdateStudentHomeworks(stWork request.UpdateStudentHomeWork, updatedBy string) (err error) {
	var stHomeWork model.StudentHomeWorks
	err = global.GvaDB.Where("student_id = ? AND home_work_id =?",stWork.StudentId, stWork.HomeWorkId).First(&stHomeWork).Updates(map[string]interface{}{"status_changed": "done", "status_change_date": time.Now().Format("2006-01-02 15:04:05"), "updated_by":updatedBy}).Error
	return err
}


func (m *StudentHomeWorkRepository) GetStudentHomeWorks(info request.StudentHomeWorkSearch) ([]model.StudentHomeWorks, int64, error) {
	if info.PageSize == 0 || info.PageSize > global.GvaConfig.Mysql.LimitRecords {
		info.PageSize = global.GvaConfig.Mysql.LimitRecords
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// create db
	db := global.GvaDB.Model(&model.StudentHomeWorks{})
	var studentsHomeWork []model.StudentHomeWorks
	// If conditional search create search

	if info.StudentId != 0 {
		db = db.Where("student_id = ?", info.StudentId)
	}

	if info.Status != ""{
		db = db.Where("status = ?", info.Status)
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

	err = db.Order("id " + info.Order).Limit(limit).Offset(offset).Find(&studentsHomeWork).Error

	return studentsHomeWork, total, err
}