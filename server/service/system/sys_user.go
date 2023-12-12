package system

import (
	"fmt"
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/system"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
	repository "github.com/ebedevelopment/next-gen-tms/server/repository/system"
	commoncase "github.com/ebedevelopment/next-gen-tms/server/usecase/common"
	useCase "github.com/ebedevelopment/next-gen-tms/server/usecase/excel"
	pdfCase "github.com/ebedevelopment/next-gen-tms/server/usecase/pdf"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
)

// UserService user
type UserService struct {
	userRepository repository.UserRepository
	excelUsecase   useCase.ExcelUseCase
	pdfUsecase      pdfCase.PDFUseCase
}

// Register User register
func (s *UserService) Register(r sysReq.Register,code string , createdBy string) (system.SysUser, error) {
	return s.userRepository.Register(r, code ,createdBy)
}

func (s *UserService) ChangeUserImageProfile(email string,imgProfile string) (err error) {
	return s.userRepository.ChangeUserImageProfile(email, imgProfile)
}

// Login User login
func (s *UserService) Login(u *system.SysUser, ipAddress string) (*system.SysUser, error) {
	return s.userRepository.Login(u, ipAddress)
}

// ChangePassword update User password
func (s *UserService) ChangePassword(u *system.SysUser) error {
	return s.userRepository.ChangePassword(u)
}

// ConfirmPassword update User password
func (s *UserService) ConfirmPassword(u *system.SysUser) error {
	return s.userRepository.ConfirmPassword(u)
}

// SetUserInfo setupUser info
func (s *UserService) SetUserInfo(req system.SysUser) error {
	return s.userRepository.SetUserInfo(req)
}

// GetUserInfo get user info
func (s *UserService) GetUserInfo(id uint) (*system.SysUser, error) {
	return s.userRepository.GetUserInfo(id)
}

// FindUserById  get user by id
func (s *UserService) FindUserById(id int) (*system.SysUser, error) {
	return s.userRepository.FindUserById(id)
}

// func (s *UserService) CheckUserById(id uint) bool {
// 	return s.userRepository.CheckUserById(id)
// }

// GetUserEmail
func (s *UserService) GetUserEmail(id uint) (string, bool) {
	return s.userRepository.GetUserEmail(id)
}

// GetUserByEmail
func (s *UserService) GetUserByEmail(email string) (*system.SysUser, error) {
	return s.userRepository.GetUserByEmail(email)
}

// CheckUserByEmail
func (s *UserService) CheckUserByEmail(email string) bool {
	return s.userRepository.CheckUserByEmail(email)
}

//----user admin



// DeleteUser deleteUser
func (s *UserService) DeleteUser(id int) error {
	return s.userRepository.DeleteUser(id)
}

// DisableUser
func (s *UserService) UpdateUserStatus(u *system.SysUser) error {
	return s.userRepository.UpdateUserStatus(u)
}

// GetUserInfoList paging data
func (s *UserService) GetUserInfoList(info sysReq.UserSearch) ([]system.SysUser, int64, error) {
	return s.userRepository.GetUserInfoList(info)
}


// ParseUsersInfoList2Excel Parse Users Info List to Excel
func (s *UserService) ParseUsersInfoList2Excel(users []system.SysUser, format string) (string, error) {

	header := []string{"Email", "Authority Name", "Phone", "Register Time", "User Status", "Last login Time"}
	data := [][]string{}

	if format == utils.Excel {
		data = append(data, header)
	}

	for _, p := range users {
		var s []string
		s = append(s, p.Email, p.Phone, p.RegisterTime.Format("2006-01-02 15:04:05"), commoncase.UserStatus(p.UserStatus), p.LastLoginTime)

		data = append(data, s)
	}

	var filePath string
	if format == utils.PDF {
		filePath = global.GvaConfig.Excel.Dir + "userExport_" + fmt.Sprintf("%v", time.Now().Unix()) + ".pdf"
		return s.pdfUsecase.CreatePDF(filePath, header, data, "Users Data", pdfCase.User)
	}

	filePath = global.GvaConfig.Excel.Dir + "userExport_" + fmt.Sprintf("%v", time.Now().Unix()) + ".xlsx"

	return s.excelUsecase.GenerateExcelSheet(data, filePath)
}
