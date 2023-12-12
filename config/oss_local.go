package config

// Local struct
type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"` // local file path
	Download string  `mapstructure:"download" json:"download" yaml:"download"`
}
