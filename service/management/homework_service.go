package management

import (
	model "github.com/ebedevelopment/next-gen-tms/server/model/management"
	request "github.com/ebedevelopment/next-gen-tms/server/model/management/request"
	"github.com/ebedevelopment/next-gen-tms/server/repository/management"
)


type HomeWorkService struct {
	homeWorkRepository management.HomeWorkRepository
}

// CreateHomeWork createHomeWorkRecord
func (m *HomeWorkService) CreateHomeWork(homeWork model.HomeWork) error {
	return m.homeWorkRepository.CreateHomeWork(homeWork)
}


func (m *HomeWorkService) GetHomeWorkID(id uint) (model.HomeWork, error) {
	return m.homeWorkRepository.GetHomeWorkID(id)
}

func (m *HomeWorkService) GetHomeWorkList(info request.HomeWorkSearch)([]model.HomeWork, int64, error){
	return m.homeWorkRepository.GetHomeWorkList(info)
}


func (m *HomeWorkService) AssignHomeWorkToClass(homework model.HomeWork) error{
	return m.homeWorkRepository.AssignHomeWorkToClass(homework)
}


func (m *HomeWorkService) GetClassHomework(classId uint) ([]model.HomeworkClasses, error){
	return m.homeWorkRepository.GetClassHomework(classId)
}

