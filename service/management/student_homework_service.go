package management

import (
	model "github.com/ebedevelopment/next-gen-tms/server/model/management"
	request "github.com/ebedevelopment/next-gen-tms/server/model/management/request"
	"github.com/ebedevelopment/next-gen-tms/server/repository/management"
)

type StudentHomeWorkService struct {
	studentHomeWorkRepository management.StudentHomeWorkRepository
}

// CreateHomeWork createHomeWorkRecord
func (m *StudentHomeWorkService) CreateStudentHomeworks(homeWork []model.StudentHomeWorks) error {
	return m.studentHomeWorkRepository.CreateStudentHomeworks(homeWork)
}

func (m *StudentHomeWorkService) UpdateStudentHomeworks(homeWork request.UpdateStudentHomeWork, updatedBy string) error {
	return m.studentHomeWorkRepository.UpdateStudentHomeworks(homeWork, updatedBy)
}

func (m *StudentHomeWorkService) GetStudentHomeWorks(info request.StudentHomeWorkSearch) ([]model.StudentHomeWorks, int64, error){
	return m.studentHomeWorkRepository.GetStudentHomeWorks(info)
}
