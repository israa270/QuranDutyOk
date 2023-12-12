package system

import (
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
)

// SysUser  system user
type SysUser struct {
	global.GvaModel
	Password    string       `json:"-"  gorm:"comment:User login password"`    // User login password                               // User nickname
	AuthorityId string       `json:"authorityId" gorm:"comment:User roleID"`   // User role ID
	// Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:User role"`

	Phone string `json:"phone,omitempty"  gorm:"comment:User phone number"` // User phone
	Email string `json:"username"  gorm:"comment:User Mail;"`                  // User Mail


	UserType

	IsVerified    bool      `json:"isVerified" gorm:"column:is_verified"`
	RegisterTime  time.Time `json:"registerTime" gorm:"column:register_time"`
	Status        string    `json:"status" gorm:"column:status"`  //for pending, verify , active
	UserStatus    bool      `json:"userStatus" gorm:"userStatus"` //for only [disable | enable]
	LastLoginTime string    `json:"lastLoginTime" gorm:"column:last_login_time"`
	ActivateDate  string    `json:"ActivateDate" gorm:"column:activate_date"` // when login
	LastPasswords string    `json:"-" gorm:"column:last_passwords;type:text;"`

	CreatedBy string `json:"createdBy,omitempty" gorm:"column:created_by"`
	UpdatedBy string `json:"updatedBy,omitempty" gorm:"column:updated_by"`
}

// TableName SysUser table Name
func (SysUser) TableName() string {
	return "sys_users"
}

// status [active - inactive - disable]

//User Types
// "superAdmin": false,
// "resellerAdmin": false,
// "merchantAdmin": false,

type UserType struct {
	SuperAdmin    bool `json:"superAdmin,omitempty" gorm:"column:super_admin"`
	OrgAdmin      bool `json:"orgAdmin,omitempty"  gorm:"column:org_admin"`
	MerchantAdmin bool `json:"merchantAdmin,omitempty" gorm:"column:merchant_admin"`

	MID uint `json:"mid,omitempty" gorm:"column:mid"`
}
