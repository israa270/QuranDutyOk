package config

// Server config struct
type Server struct {
	JWT     JWT        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap        `mapstructure:"zap" json:"zap" yaml:"zap"`

	System  System     `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha    `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Login   UserConfig `mapstructure:"user-config" json:"userConfig" yaml:"userConfig"`

	// gorm
	MysqlInfo  MysqlInfo `mapstructure:"mysql-info" json:"mysqlInfo" yaml:"mysql-info"`
	Mysql  Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

	Local Local `mapstructure:"local" json:"local" yaml:"local"`

	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	Timer Timer `mapstructure:"timer" json:"timer" yaml:"timer"`

	// cross domain configure
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`

	// added  to support multi-language
	Language Language `mapstructure:"language" json:"language" yaml:"language"`


	MenuTranslator MenuTranslator `mapstructure:"menu-translator" json:"menuTranslator" yaml:"menu-translator"`

	LoggerPath     LoggerPath      `mapstructure:"logger-path" json:"loggerPath" yaml:"logger-path"`
}
