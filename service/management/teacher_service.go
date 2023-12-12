package management


import (

"github.com/ebedevelopment/next-gen-tms/server/repository/management"
model "github.com/ebedevelopment/next-gen-tms/server/model/management"
request "github.com/ebedevelopment/next-gen-tms/server/model/management/request"
)


type TeacherService struct {
	teacherRepository management.TeacherRepository
}





// CreateTeacher createTeacherRecord
func (m *TeacherService) CreateTeacher(teacher model.Teacher) error {
	return m.teacherRepository.CreateTeacher(teacher)
}

func (m *TeacherService) CheckTeacherName(name  string) bool {
	return m.teacherRepository.CheckTeacherName(name)
}


func (m *TeacherService) CheckTeacherExist(id uint) bool{
	return m.teacherRepository.CheckTeacherExist(id)
}
 

func (m *TeacherService) GetTeacherList(info request.ListSearch)([]model.Teacher, int64, error){
   return m.teacherRepository.GetTeacherList(info)
}

func (m *TeacherService) CheckTeacherByUserName(userName  string) bool {
	return m.teacherRepository.CheckTeacherByUserName(userName)
}


func (s *TeacherService) Login(u *model.Teacher) (*model.Teacher, error) {
	return s.teacherRepository.Login(u)
}