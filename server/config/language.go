package config

// Language struct
type Language struct {
	Language string `mapstructure:"language" json:"language" yaml:"language"`
	Dir      string `mapstructure:"dir" json:"dir" yaml:"dir"`
}
