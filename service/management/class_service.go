package management

import (
	model "github.com/ebedevelopment/next-gen-tms/server/model/management"
	request "github.com/ebedevelopment/next-gen-tms/server/model/management/request"
	"github.com/ebedevelopment/next-gen-tms/server/repository/management"
)


type ClassService struct {
	classRepository management.ClassRepository
}

// CreateClass createClassRecord
func (m *ClassService) CreateClass(Class model.Class) error {
	return m.classRepository.CreateClass(Class)
}

func (m *ClassService) CheckClassName(name  string, versionName  string) bool {
	return m.classRepository.CheckClassName(name, versionName)
}

func (m *ClassService) GetClassID(id uint) (model.Class, error) {
	return m.classRepository.GetClassID(id)
}

func (m *ClassService) CheckClassID(id uint) bool {
	return m.classRepository.CheckClassID(id)
}

func (m *ClassService) GetClassList(info request.ClassSearch)([]model.Class, int64, error){
	return m.classRepository.GetClassList(info)
}