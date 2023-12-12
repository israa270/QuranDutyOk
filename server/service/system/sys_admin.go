package system

import (
	repository "github.com/ebedevelopment/next-gen-tms/server/repository/system"
	model "github.com/ebedevelopment/next-gen-tms/server/model/system"
)

type AdminService struct{
	adminRepository repository.AdminRepository
}


// func (m *AdminService) GetAdminByUserName(name  string) (model.Admin, error ) {
// 	return m.adminRepository.GetAdminByUserName(name)
// }

func (m *AdminService) CreateAdmin(admins  []model.Admin) error {
	return m.adminRepository.CreateAdmin(admins)
}


func (m *AdminService) CheckAdminByUserName(userName  string) bool {
	return m.adminRepository.CheckAdminByUserName(userName)
}


func (s *AdminService) Login(u *model.Admin) (*model.Admin, error) {
	return s.adminRepository.Login(u)
}