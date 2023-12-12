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

func (m *ClassService) GetClassID(name  string, versionName  string) (uint, error) {
	return m.classRepository.GetClassID(name, versionName)
}

func (m *ClassService) CheckClassID(classId uint) bool {
	return m.classRepository.CheckClassID(classId)
}

func (m *ClassService) GetClassList(info request.ClassSearch)([]model.Class, int64, error){
	return m.classRepository.GetClassList(info)
}