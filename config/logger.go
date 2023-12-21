package config

type MessageLogger struct {
	GeneralMsg          `json:"general"`
	SysCaptchaMsg       `json:"sysCaptcha"`
	SysSystemMsg        `json:"sysSystem"`
	SysAuthorityMsg     `json:"sysAuthority"`
	SysCasbinMsg        `json:"sysCasbin"`
	ExcelMsg            `json:"excel"`
	SysInitDBMsg        `json:"sysInitDB"`
	ApiMsg              `json:"api"`
	MerchantMsg         `json:"merchant"`
	MailMsg             `json:"mail"`
	InitMsg             `json:"init"`
	UserMsg             `json:"sysUser"`
	ReportsMsg          `json:"reports"`
	FileMsg             `json:"file"`
}

type GeneralMsg struct {
	ModifyFail         string `json:"modifyFail"`
	ModifySuccess      string `json:"modifySuccess"`
	GetDataFail        string `json:"getDataFail"`
	GetDataSuccess     string `json:"getDataSuccess"`
	DeleteFail         string `json:"deleteFail"`
	DeleteSuccess      string `json:"deleteSuccess"`
	SetupFail          string `json:"setupFail"`
	SetupSuccess       string `json:"setupSuccess"`
	CreationFail       string `json:"creationFail"`
	CreateSuccess      string `json:"createSuccess"`
	UpdateFail         string `json:"updateFail"`
	UpdateSuccess      string `json:"updateSuccess"`
	QueryFail          string `json:"queryFail"`
	QuerySuccess       string `json:"querySuccess"`
	TableDataInitFail  string `json:"tableDataInitFail"`
	ViewCreateFail     string `json:"viewCreateFail"`
	BatchDeleteFail    string `json:"batchDeleteFail"`
	BatchDeleteSuccess string `json:"batchDeleteSuccess"`

	BadRequest         string `json:"badRequest"`
	ReadFileFailed     string `json:"readFileFailed"`
	FilePath           string `json:"filePath"`
	DuplicateValueName string `json:"duplicateValueName"`
	IdNotFound         string `json:"idNotFound"`
	EnableStatus       string `json:"enableStatus"`
	InputFail          string `json:"inputFail"`
	NotFoundFile       string `json:"notFoundFile"`
	DownloadFail       string `json:"downloadFail"`
	DownloadSuccess    string `json:"downloadSuccess"`
	RemoveFile         string `json:"removeFile"`

	UploadFile          string `json:"UploadFile"`
	UsernameFail        string `json:"usernameFail"`
	UrlEmpty            string `json:"urlEmpty"`
	DownloadPermission  string `json:"downloadPermission"`
	GetFileSize         string `json:"getFileSize"`
	ParseFile           string `json:"parseFile"`
	FileFDS             string `json:"fileFDS"`
	DeleteActive        string `json:"deleteActive"`
	FileFolder          string `json:"fileFolder"`
	GenerateExcelFile   string `json:"generateExcelFile"`
	Approved            string `json:"approved"`
	UpdateStatus        string `json:"updateStatus"`
	UpdateStatusSuccess string `json:"updateStatusSuccess"`
	CalculateTotal      string `json:"calculateTotal"`
	UserPermission      string `json:"userPermission"`
	UserData            string `json:"userData"`
	DisableUser         string `json:"disableUser"`
	DisableUserFail     string `json:"disableUserFail"`
	ImportFileFail      string `json:"importFileFail"`
	TimeParsingFail     string `json:"timeParsingFail"`
	EncodeImg           string `json:"encodeImg"`
	ValidatePhone       string `json:"validatePhone"`
	SignatureFail       string `json:"signatureFail"`
	ValidateEmail       string `json:"validateEmail"`
	ImportSuccess       string `json:"importSuccess"`
	ImportQuery         string `json:"importQuery"`
	RequiredPhone       string `json:"requiredPhone"`
	LoadJson            string `json:"loadJson"`
	Lang                string `json:"lang"`
	SaveImg             string `json:"saveImg"`
	GetMenu             string `json:"getMenu"`
	DeleteKey           string `json:"deleteKey"`
	ShortLinkFail       string `json:"shortLinkFail"`
	ShortLinkSuccess    string `json:"shortLinkSuccess"`
	GetShortLink        string `json:"getShortLink"`
	GetShortLinkDB      string `json:"getShortLinkDB"`
	ContactName         string `json:"contactName"`
	StatusFail          string `json:"statusFail"`
	StatusSuccess       string `json:"statusSuccess"`
	SaveMenu            string `json:"saveMenu"`
	StatusActive        string `json:"statusActive"`
	ParentYourself      string `json:"parentYourself"`
	IconFilePath        string `json:"iconFilePath"`
	VersionNumber       string `json:"versionNumber"`
	DuplicateFileName   string `json:"duplicateFileName"`
	UpdateDownload      string `json:"updateDownload"`
	FileNotLoad         string `json:"fileNotLoad"`
	ScreenShotFileName  string `json:"screenShotFileName"`
	UnzipFile           string `json:"unzipFile"`
	UnSupportFormat     string `json:"unSupportFormat"`
	DuplicateValuePhone string `json:"duplicateValuePhone"`
	FailCount           string `json:"failCount"`
	ParseDuration       string `json:"parseDuration"`
	DeleteFile          string `json:"deleteFile"`
	CreateFile          string `json:"createFile"`
	VariableSource      string `json:"variableSource"`
	ParseJson           string `json:"parseJson"`
	ReadBodyResp        string `json:"readBodyResp"`
	UserFail            string `json:"userFail"`
}

type SysCaptchaMsg struct {
	VCodeFail    string `json:"vCodeFail"`
	VCodeSuccess string `json:"vCodeSuccess"`
}

type SysSystemMsg struct {
	RebootFail    string `json:"rebootFail"`
	RebootSuccess string `json:"rebootSuccess"`
	CpuFail       string `json:"cpuFail"`
	RamFail       string `json:"ramFail"`
	DiskFail      string `json:"diskFail"`
}

type SysAuthorityMsg struct {
	RoleExist           string `json:"roleExist"`
	RoleDelete          string `json:"roleDelete"`
	RoleUsers           string `json:"roleUsers"`
	SuperAdminRole      string `json:"superAdminRole"`
	AuthorityName       string `json:"authorityName"`
	AuthorityCodeLen    string `json:"authorityCodeLen"`
	PrivilegesAuthority string `json:"privilegesAuthority"`
}

type SysCasbinMsg struct {
	UpdateCasbin     string `json:"updateCasbin"`
	CasbinPermission string `json:"casbinPermission"`
}

type ExcelMsg struct {
	ExcelFail            string `json:"excelFail"`
	DownloadTemplateFail string `json:"downloadTemplateFail"`
}

type SysInitDBMsg struct {
	AutoCreateDBFail    string `json:"autoCreateDBFail"`
	AutoCreateDBFailErr string `json:"autoCreateDBFailErr"`
	AutoCreateDBSuccess string `json:"autoCreateDBSuccess"`
	DB                  string `json:"db"`
}

type ApiMsg struct {
	Role                    string `json:"role"`
	SystemUser              string `json:"systemUser"`
	UserLoginRequired       string `json:"userLoginRequired"`
	AddMenu                 string `json:"addMenu"`
	AddMenuRole             string `json:"addMenuRole"`
	ChangePassword          string `json:"changePassword"`
	ConfirmPassword         string `json:"confirmPassword"`
	ResetPasswordConfirm    string `json:"resetPasswordConfirm"`
	ChangeRoleAPIPermission string `json:"changeRoleAPIPermission"`
	CreateAPI               string `json:"createAPI"`
	CreateRole              string `json:"createRole"`
	DeleteAPI               string `json:"deleteAPI"`
	DeleteAPIByID           string `json:"deleteAPIByID"`
	DeleteMenu              string `json:"deleteMenu"`
	DeleteRole              string `json:"deleteRole"`
	DeleteUsers             string `json:"deleteUsers"`
	FileUploadDownload      string `json:"fileUploadDownload"`
	GetAPIByID              string `json:"getAPIByID"`
	GetAPIList              string `json:"getAPIList"`
	GetAllAPI               string `json:"getAllAPI"`
	GetDynamicRoute         string `json:"getDynamicRoute"`
	GetMenuByID             string `json:"getMenuByID"`
	GetMenuList             string `json:"getMenuList"`
	GetMenuRole             string `json:"getMenuRole"`
	GetMenuTree             string `json:"getMenuTree"`
	GetPermissionList       string `json:"getPermissionList"`
	GetPermissionNames      string `json:"getPermissionNames"`
	GetPermissionMap        string `json:"getPermissionMap"`
	GetPermByRole           string `json:"getPermByRole"`
	GetRoleList             string `json:"getRoleList"`
	GetRolesMenu            string `json:"getRolesMenu"`
	GetAuthorityCode        string `json:"getAuthorityCode"`
	GetSelfInfo             string `json:"getSelfInfo"`
	GetUsersList            string `json:"getUsersList"`
	JwtAddedToBlackList     string `json:"jwtAddedToBlackList"`
	Menu                    string `json:"menu"`
	ModifyUserRole          string `json:"modifyUserRole"`
	OptRecord               string `json:"optRecord"`
	ResetUserPassword       string `json:"resetUserPassword"`
	ForgetPassword          string `json:"forgetPassword"`
	SetUserInfo             string `json:"setUserInfo"`
	SystemService           string `json:"systemService"`
	UpdateAPI               string `json:"updateAPI"`
	UpdateMenu              string `json:"updateMenu"`
	UpdateRole              string `json:"updateRole"`
	UserRegistration        string `json:"userRegistration"`
	DuplicateApi            string `json:"duplicateApi"`
	DuplicateApiPath        string `json:"duplicateApiPath"`
	ActiveAccount           string `json:"activeAccount"`
	DisableUser             string `json:"disableUser"`
	UserLogout              string `json:"userLogout"`
}

type MerchantMsg struct {
	MerchantNotExist       string `json:"merchantNotExist"`
	MerchantNotExistStatus string `json:"merchantNotExistStatus"`
	MerchantIdExist        string `json:"merchantIdExist"`
	MerchantName           string `json:"merchantName"`
	MerchantEmail          string `json:"merchantEmail"`
	MerchantPhone          string `json:"merchantPhone"`
	MerchantTerminals      string `json:"merchantTerminals"`
	MerchantPermission     string `json:"merchantPermission"`
	DisableMerchant        string `json:"disableMerchant"`
	ShopMerchant           string `json:"shopMerchant"`
	MidRequired            string `json:"midRequired"`
	MerchantExistOrg       string `json:"merchantExistOrg"`
}

type MailMsg struct {
	SendEmailFail      string `json:"sendEmailFail"`
	SendEmailSuccess   string `json:"sendEmailSuccess"`
	EmailTokenNotValid string `json:"emailTokenNotValid"`
	EmailVerifyBefore  string `json:"emailVerifyBefore"`
	ActiveLink         string `json:"activeLink"`
}

type InitMsg struct {
	DbAlreadyExist string `json:"dbAlreadyExist"`
	DbNotExist     string `json:"dbNotExist"`
	InitDB         string `json:"initDB"`
	DbAlreadyInit  string `json:"dbAlreadyInit"`
	TableFail      string `json:"tableFail"`
	TableSuccess   string `json:"tableSuccess"`
	RedisFail      string `json:"redisFail"`
	RedisSuccess   string `json:"redisSuccess"`
	RouterSuccess  string `json:"routerSuccess"`
	AuthExpire     string `json:"authExpire"`
	NotLog         string `json:"notLog"`
	TokenNotValid  string `json:"tokenNotValid"`
	TokenExpire    string `json:"tokenExpire"`
	TokenNotActive string `json:"tokenNotActive"`
	TokenEven      string `json:"tokenEven"`
	NotToken       string `json:"notToken"`
}

type UserMsg struct {
	LoginFail               string `json:"loginFail"`
	UserNameOrPasswordError string `json:"userNameOrPasswordError"`
	GetTokenFail            string `json:"getTokenFail"`
	LoginSuccess            string `json:"loginSuccess"`
	LoginStatusFail         string `json:"loginStatusFail"`
	JwtInvalidationFailed   string `json:"jwtInvalidationFailed"`
	RegistrationFail        string `json:"registrationFail"`
	DuplicatedUserName      string `json:"duplicatedUserName"`
	DuplicatedEmail         string `json:"duplicatedEmail"`
	RegistrationSuccess     string `json:"registrationSuccess"`
	ChangePW                string `json:"changePW"`
	DeleteUserFail          string `json:"deleteUserFail"`
	ResetPWFail             string `json:"resetPWFail"`
	ResetPWSuccess          string `json:"resetPWSuccess"`
	IpAddressBlock          string `json:"ipAddressBlock"`
	UserBlock               string `json:"userBlock"`
	EmailNotFound           string `json:"emailNotFound"`
	EmailSendFail           string `json:"emailSendFail"`
	EmailSendSuccess        string `json:"emailSendSuccess"`
	UserLoginAttempt        string `json:"userLoginAttempt"`
	IpAddressAttempt        string `json:"ipAddressAttempt"`
	UpdateLastLoginFail     string `json:"updateLastLoginFail"`
	LogLoginFail            string `json:"logLoginFail"`
	UpdateBlockTime         string `json:"updateBlockTime"`
	ClearIpCount            string `json:"clearIpCount"`
	ActiveSuccess           string `json:"activeSuccess"`
	UpdateUserStatusFail    string `json:"updateUserStatusFail"`
	PasswordFail            string `json:"passwordFail"`
	UserRegister            string `json:"userRegister"`
	UserRegisterFail        string `json:"userRegisterFail"`
	UserStatus              string `json:"userStatus"`
	UserExist               string `json:"userExist"`
	DisableYourSelf         string `json:"disableYourSelf"`
	UserDisable             string `json:"userDisable"`
	UserLogoutFail          string `json:"UserLogoutFail"`
	UserLogoutSuccess       string `json:"UserLogoutSuccess"`
	UserEmail               string `json:"userEmail"`
	WrongPassword           string `json:"wrongPassword"`
	PasswordExpire          string `json:"passwordExpire"`
	OldNotSamePassword      string `json:"oldNotSamePassword"`
	PasswordSameLast        string `json:"passwordSameLast"`
	EmailFail               string `json:"emailFail"`
	ImageProfile            string `json:"imageProfile"`
	UserRole                string `json:"userRole"`
	ResetPassword           string `json:"resetPassword"`
	RoleNotExist            string `json:"roleNotExist"`
	AdminPermission         string `json:"adminPermission"`
	LoginFailWithCaptcha    string `json:"loginFailWithCaptcha"`
	UpdateIpAddress         string `json:"updateIpAddress"`
	UserTypeFail            string `json:"userTypeFail"`
	UpdateExpireFlag        string `json:"updateExpireFlag"`
	ExpirePassword          string `json:"expirePassword"`
	WrongPsd                string `json:"wrongPsd"`
	UpdateLoginAttempt      string `json:"updateLoginAttempt"`
	ResetPasswordFlag       string `json:"resetPasswordFlag"`
	ResetPasswordLink       string `json:"resetPasswordLink"`
	UserPassword            string `json:"userPassword"`
	ResendRegister          string `json:"resendRegister"`
	UserNotRegister         string   `json:"userNotRegister"`
	LdapUser                string  `json:"ldapUser"`
	UserVerified            string   `json:"userVerified"`
	ValidatePassword        string   `json:"validatePassword"`
}




type ReportsMsg struct {
	ParameterReportsFail string `json:"parameterReportsFail"`
	InfoFail             string `json:"infoFail"`
	UpdateReportTask     string `json:"updateReportTask"`
	SaveReportFile       string `json:"saveReportFile"`
}


type FileMsg struct {
	ReceiveFile   string `json:"receiveFile"`
	LimitFileSize string `json:"limitFileSize"`
	LinkFilePath  string `json:"linkFilePath"`
	UploadSuccess string `json:"uploadSuccess"`
}
