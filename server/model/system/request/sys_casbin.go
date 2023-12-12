package request

// CasbinInfo structure
type CasbinInfo struct {
	Path   string `json:"path"`   // path
	Method string `json:"method"` // method
}

//CasbinInReceive Casbin structure for input parameters
type CasbinInReceive struct {
	AuthorityId string       `json:"authorityId"  binding:"required"` // permission id
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

// DefaultCasbin default apis
func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/menu/getMenuList", Method: "GET"},
		{Path: "/menu/getMenuAuthority/:authorityId", Method: "GET"},
	
		{Path: "/base/login", Method: "POST"},
		{Path: "/base/captcha", Method: "GET"},
		{Path: "/user/logout", Method: "GET"},

		{Path: "/base/forgetPassword", Method: "GET"},
		{Path: "/base/verifyAccount", Method: "GET"},
		{Path: "/base/resetPasswordLink",Method: "GET"},
		{Path: "/base/passwordConfirm/:verificationCode", Method: "PUT"},
		{Path: "/base/resetPassword/:token", Method: "PUT"},
		{Path: "/base/expirePassword", Method: "PUT"},
	
		{Path: "/user/changePassword", Method: "PUT"},
        {Path: "/user/uploadProfile", Method: "POST"},

		{Path: "/user/setUserInfo", Method: "PUT"},
		{Path: "/user/getUserInfo", Method: "GET"},

		{Path: "/base/CheckPassword",Method: "GET"},
		{Path:"/base/checkPasswordReset", Method: "GET"},
		// download excel download
		{Path: "/excel/downloadTemplate", Method: "GET"},
	}
}
