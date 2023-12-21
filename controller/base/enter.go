package base

import (
	"github.com/ebedevelopment/next-gen-tms/server/service/management"
	"github.com/ebedevelopment/next-gen-tms/server/service/system"
	usercase "github.com/ebedevelopment/next-gen-tms/server/usecase/user"
)

type BaseController struct {
	adminService  system.AdminService
	loginAttemptService system.LoginAttemptService
	teacherService   management.TeacherService
    studentService   management.StudentService


	UserUseCase usercase.UserUseCase
}
