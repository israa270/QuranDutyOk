package config

type MenuTranslator struct{
	Dir      string `mapstructure:"dir" json:"dir" yaml:"dir"`
}


type LoggerPath struct{
	Dir    string  `mapstructure:"dir" json:"dir" yaml:"dir"`
}