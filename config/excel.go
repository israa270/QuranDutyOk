package config

// Excel struct
type Excel struct {
	Dir string `mapstructure:"dir" json:"dir" yaml:"dir"`
	DirTemplate string `mapstructure:"dir-template" json:"dirTemplate" yaml:"dir-template"`
}
