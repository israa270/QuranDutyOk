package management

import (
	model "github.com/ebedevelopment/next-gen-tms/server/model/management"
	request "github.com/ebedevelopment/next-gen-tms/server/model/management/request"
	"github.com/ebedevelopment/next-gen-tms/server/repository/management"
)


type HomeWorkService struct {
	HomeWorkRepository management.HomeWorkRepository
}

// CreateHomeWork createHomeWorkRecord
func (m *HomeWorkService) CreateHomeWork(homeWork model.HomeWork) error {
	return m.HomeWorkRepository.CreateHomeWork(homeWork)
}

func (m *HomeWorkService) CheckHomeWorkName(name  string, versionName  string) bool {
	return m.HomeWorkRepository.CheckHomeWorkName(name, versionName)
}

func (m *HomeWorkService) GetHomeWorkID(name  string, versionName  string) (uint, error) {
	return m.HomeWorkRepository.GetHomeWorkID(name, versionName)
}

func (m *HomeWorkService) GetHomeWorkList(info request.HomeWorkSearch)([]model.HomeWork, int64, error){
	return m.HomeWorkRepository.GetHomeWorkList(info)
}