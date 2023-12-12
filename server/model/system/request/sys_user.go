package request

import (
	"github.com/ebedevelopment/next-gen-tms/server/model/common/request"
)

// Register User register structure
type Register struct {
	Email       string `json:"userName" binding:"required,max=200"`
	AuthorityId string `json:"authorityId" gorm:"comment:authorityId" binding:"required,max=20"`
}

// Login User login structure
type Login struct {
	Email     string `json:"username" binding:"required"`        // username  ,email
	Password  string `json:"password" binding:"required,min=8,max=32"` // password
	Captcha   string `json:"captcha,omitempty"`                        // code
	CaptchaId string `json:"captchaId,omitempty"`                      // codeID
}

// ChangePasswordStruct Modify password structure
// type ChangePasswordStruct struct {
// 	Email           string `json:"username" binding:"required,email"`                        // username
// 	Password        string `json:"password" binding:"required,min=8,max=32,alphanumunicode"` // password
// 	ConfirmPassword string `json:"confirmPassword" binding:"eqfield=Password,required"`      // confirm password
// }

// ChangePassword change password
type ChangePassword struct {
	OldPassword     string `json:"oldPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required,min=8,max=32"` //,alphanumunicode
	ConfirmPassword string `json:"confirmPassword"  binding:"eqfield=NewPassword,required"`
}

// ExpirePassword expire password
type ExpirePassword struct {
	Email           string `json:"email"  binding:"required,email"`
	OldPassword     string `json:"oldPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required,min=8,max=32" `  //,alphanumunicode
	ConfirmPassword string `json:"confirmPassword" binding:"eqfield=NewPassword,required"`
}

// ResetPassword reset password struct
type ResetPassword struct {
	Password        string `json:"password" binding:"required,min=8,max=32"`  //,alphanumunicode
	ConfirmPassword string `json:"confirmPassword" binding:"eqfield=Password,required"`
}

// SetUserAuth Modify  user's auth structure
type SetUserAuth struct {
	UserEmail   string `json:"userEmail" binding:"required"`
	AuthorityId string `json:"authorityId" binding:"required"` // roleID
}

// ChangeUserInfo change user info
type ChangeUserInfo struct {
	Phone     string `json:"phone"`
	HeaderImg string `json:"headerImg"`
}

// ForgetPasswordCaptcha forget password
type ForgetPasswordCaptcha struct {
	Email     string `json:"email" binding:"required,email"`
	Captcha   string `json:"captcha" binding:"required"`
	CaptchaId string `json:"captchaId" binding:"required"`
}

// UserSearch user search
type UserSearch struct {
	Email  string `json:"email" form:"email"`
	Status      string `json:"status" form:"status"`
	AuthorityId string `json:"authorityId"  form:"authorityId"`

	request.PageInfo
}


// LastPassword last password used to transfer object between fns.
type LastPassword struct {
	LastPasswords string
	Password      string
}

// Modify password structure
type PasswordRegister struct {
	Email           string `json:"-"`                        // username
	Password        string `json:"password" binding:"required,min=8,max=32"` // password  ,alphanumunicode
	ConfirmPassword string `json:"confirmPassword" binding:"eqfield=Password,required"`      // confirm password
}
