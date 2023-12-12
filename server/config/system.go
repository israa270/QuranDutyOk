package config

// System
type System struct {
	Env                     string `mapstructure:"env" json:"env" yaml:"env"` // env value

	Addr                    string    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType                  string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`                        // database type of:mysql(default)|sqlite|sql server|postgresql
	OssType                 string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`                     // Oss type of
	LimitCountIP            int    `mapstructure:"ip-limit-count" json:"ipLimitCount" yaml:"ip-limit-count"`
	LimitTimeIP             int    `mapstructure:"ip-limit-time" json:"ipLimitTime" yaml:"ip-limit-time"`
	BackendPath             string `mapstructure:"backend-path" json:"backendPath" yaml:"backend-path"`
	ServerPath              string `mapstructure:"server-path" json:"serverPath" yaml:"server-path"`
	UiServerPath            string `mapstructure:"ui-server-path" json:"uiServerPath" yaml:"ui-server-path"`
}
