package management

import (
	"github.com/ebedevelopment/next-gen-tms/server/repository/management"

	model "github.com/ebedevelopment/next-gen-tms/server/model/management"
	request "github.com/ebedevelopment/next-gen-tms/server/model/management/request"
)

type StudentService struct{
	studentRepository management.StudentRepository
}


// CreateStudent createStudentRecord
func (m *StudentService) CreateStudent(Student model.Student) error {
	return m.studentRepository.CreateStudent(Student)
}

func (m *StudentService) MoveStudent(Student model.Student) error {
	return m.studentRepository.MoveStudent(Student)
}



func (m *StudentService) CheckStudentName(name  string) bool {
	return m.studentRepository.CheckStudentName(name)
}


func (m *StudentService) CheckStudentExist(id uint) bool{
	return m.studentRepository.CheckStudentExist(id)
}
 

func (m *StudentService) GetStudentList(info request.StudentSearch)([]model.Student, int64, error){
   return m.studentRepository.GetStudentList(info)
}

func (m *StudentService) CheckStudentByUserName(userName  string) bool {
	return m.studentRepository.CheckStudentByUserName(userName)
}


func (s *StudentService) Login(u *model.Student) (*model.Student, error) {
	return s.studentRepository.Login(u)
}