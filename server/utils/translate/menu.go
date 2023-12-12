package translate

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Menu struct {
	Monitoring       string `json:"monitoring"`
	Dashboard        string `json:"dashboard"`
	TerminalLocation string `json:"terminalLocation"`

	Management string `json:"management"`

	Organization        string `json:"organization"`
	OrganizationRead    string `json:"organizationRead"`
	OrganizationDelete  string `json:"organizationDelete"`
	OrganizationDisable string `json:"organizationDisable"`
	OrganizationCreate  string `json:"organizationCreate"`
	OrganizationEdit    string `json:"organizationEdit"`
	OrganizationImport  string `json:"organizationImport"`
	OrganizationExport  string `json:"organizationExport"`

	Merchant        string `json:"merchant"`
	MerchantRead    string `json:"merchantRead"`
	MerchantDelete  string `json:"merchantDelete"`
	MerchantDisable string `json:"merchantDisable"`
	MerchantCreate  string `json:"merchantCreate"`
	MerchantEdit    string `json:"merchantEdit"`
	MerchantImport  string `json:"merchantImport"`
	MerchantExport  string `json:"merchantExport"`

	Terminal        string `json:"terminal"`
	TerminalRead    string `json:"terminalRead"`
	TerminalDelete  string `json:"terminalDelete"`
	TerminalStatus string `json:"terminalStatus"`
	TerminalCreate  string `json:"terminalCreate"`
	TerminalEdit    string `json:"terminalEdit"`
	TerminalImport  string `json:"terminalImport"`
	TerminalDownloadTemplate string `json:"terminalDownloadTemplate"`
	TerminalExport  string `json:"terminalExport"`
	TerminalMove    string `json:"terminalMove"`
	TerminalGroupTerminal  string `json:"terminalGroupTerminal"`
  
	TerminalGroup     string `json:"terminalGroup"`
	TerminalGroupRead string `json:"terminalGroupRead"`
	TerminalGroupEdit string `json:"terminalGroupEdit"`
	TerminalGroupDelete string `json:"terminalGroupDelete"`
	TerminalGroupCreate string `json:"terminalGroupCreate"`
	TerminalGroupDeleteTerminal string `json:"terminalGroupDeleteTerminal"`
    TerminalGroupUpdateStatus  string `json:"terminalGroupUpdateStatus"`
    TerminalGroupAddTerminal   string  `json:"terminalGroupAddTerminal"`
	TerminalGroupImportTerminal string `json:"terminalGroupImportTerminal"`


	Application string `json:"application"`
	AppCreate   string `json:"appCreate"`
	AppDetails  string `json:"appDetails"`
	AppUpdate    string `json:"appUpdate"`
	AppDelete    string  `json:"appDelete"`
	AppStatus    string   `json:"appStatus"`
    AppVersionDelete string  `json:"appVersionDelete"`
	AppVersionStatus string  `json:"appVersionStatus"`
	DownloadApp      string `json:"downloadApp"`
    
	DownloadParamFile string  `json:"downloadParamFile"`
	DeleteParamFile   string  `json:"deleteParamFile"`
    AddParamFile     string   `json:"addParamFile"`

	AppEditModel  string   `json:"appEditModel"`
	AppDistributeOrg string `json:"appDistributeOrg"`
    
	ParamVariable  string  `json:"paramVariable"`
	ParamVariableRead string `json:"paramVariableRead"`
	ParamVariableCreate string `json:"paramVariableCreate"`
	ParamVariableDelete string `json:"paramVariableDelete"`
    ParamVariableExport string `json:"paramVariableExport"`

	

	PushTask          string `json:"pushTask"`
	PushAppTask       string `json:"pushAppTask"`
	PushParameterFile string `json:"pushParameterFile"`
	PushFirmwareTask  string `json:"pushFirmwareTask"`
	PushMessageTask   string `json:"pushMessageTask"`

	Firmware string `json:"firmware"`
	FirmwareCreate string `json:"firmwareCreate"`
	FirmwareEdit string `json:"firmwareEdit"`
	FirmwareDelete string `json:"firmwareDelete"`
	FirmwareDownload string `json:"firmwareDownload"`
	FirmwareRead  string `json:"firmwareRead"`
	FirmwareStatus string `json:"firmwareStatus"`

    PushTemplate   string `json:"pushTemplate"`
	PushTemp   string `json:"pushTemp"`

	GeneralSetting string `json:"generalSetting"`

	Manufacturer        string `json:"manufacturer"`
	ManufacturerRead    string `json:"manufacturerRead"`
	ManufacturerDelete  string `json:"manufacturerDelete"`
	ManufacturerDisable string `json:"manufacturerDisable"`
	ManufacturerCreate  string `json:"manufacturerCreate"`
	ManufacturerEdit    string `json:"manufacturerEdit"`
	ManufacturerImport  string `json:"manufacturerImport"`
	ManufacturerExport  string `json:"manufacturerExport"`

	Platform        string `json:"platform"`
	PlatformRead    string `json:"platformRead"`
	PlatformDelete  string `json:"platformDelete"`
	PlatformDisable string `json:"platformDisable"`
	PlatformCreate  string `json:"platformCreate"`
	// PlatformEdit    string `json:"platformEdit"`
	PlatformImport  string `json:"platformImport"`
	PlatformExport  string `json:"platformExport"`

	Model        string `json:"model"`
	ModelRead    string `json:"modelRead"`
	ModelDelete  string `json:"modelDelete"`
	ModelDisable string `json:"modelDisable"`
	ModelCreate  string `json:"modelCreate"`
	ModelEdit    string `json:"modelEdit"`
	ModelImport  string `json:"modelImport"`
	ModelExport  string `json:"modelExport"`

	MerchantType        string `json:"merchantType"`
	MerchantTypeRead    string `json:"merchantTypeRead"`
	MerchantTypeDelete  string `json:"merchantTypeDelete"`
	MerchantTypeDisable string `json:"merchantTypeDisable"`
	MerchantTypeCreate  string `json:"merchantTypeCreate"`
	MerchantTypeEdit    string `json:"merchantTypeEdit"`
	MerchantTypeImport  string `json:"merchantTypeImport"`
	MerchantTypeExport  string `json:"merchantTypeExport"`

	AppCategory        string `json:"appCategory"`
	AppCategoryRead    string `json:"appCategoryRead"`
	AppCategoryDelete  string `json:"appCategoryDelete"`
	AppCategoryDisable string `json:"appCategoryDisable"`
	AppCategoryCreate  string `json:"appCategoryCreate"`
	AppCategoryEdit    string `json:"appCategoryEdit"`
	AppCategoryImport  string `json:"appCategoryImport"`
	AppCategoryExport  string `json:"appCategoryExport"`

	Puk string `json:"puk"`

	System            string `json:"system"`
	User              string `json:"user"`
	UserFull          string `json:"userFull"`
	UserReadOnly      string `json:"userReadOnly"`
	UserCreate        string `json:"userCreate"`
	UserDisable       string `json:"userDisable"`
	UserDelete        string `json:"userDelete"`
	UserResetPassword string `json:"userResetPassword"`

	Role         string `json:"role"`
	RoleFull     string `json:"roleFull"`
	RoleReadOnly string `json:"roleReadOnly"`
	RoleCreate   string `json:"roleCreate"`
	RoleEdit     string `json:"roleEdit"`
	RoleDelete   string `json:"roleDelete"`

	AuditLog string `json:"auditLog"`
	
	ParameterVariableLog string  `json:"parameterVariableLog"`

	Config       string `json:"config"`
	ConfigView   string `json:"configView"`
	ConfigUpdate string `json:"configUpdate"`

	Report string `json:"report"`
}

func ParseJsonFile(filePath string) (*Menu, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Println(err)
		}
	}(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)

	var menu Menu
	err = json.Unmarshal(byteValue, &menu)
	if err != nil {
		log.Println(err)
		//return nil, err
	}

	return &menu, nil
}
