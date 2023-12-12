package config

// Timer
type Timer struct {
	Start  bool     `mapstructure:"start" json:"start" yaml:"start"` // yesnostart
	Spec   string   `mapstructure:"spec" json:"spec" yaml:"spec"`    // CRON table expression
	Detail []Detail `mapstructure:"detail" json:"detail" yaml:"detail"`
}

// Detail
type Detail struct {
	TableName    string `mapstructure:"tableName" json:"tableName" yaml:"tableName"`          // need cleaningoftable Name
	CompareField string `mapstructure:"compareField" json:"compareField" yaml:"compareField"` // when a comparison is requiredbetweenoffield
	Interval     string `mapstructure:"interval" json:"interval" yaml:"interval"`             // time between interval
}
