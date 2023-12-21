package config

// import "time"

// UserConfig
type UserConfig struct {
    UserBlockNum       int           `mapstructure:"user-block-num" json:"userBlockNum" yaml:"user-block-num"`
	UserBlockTime      string `mapstructure:"user-block-time" json:"userBlockTime" yaml:"user-block-time"`
	PasswordExpired    int           `mapstructure:"password-expired" json:"passwordExpired" yaml:"password-expired"`
	LoginAttemptCountTime string `mapstructure:"login-attempt-count-time" json:"loginAttemptCountTime" yaml:"login-attempt-count-time"`
	AuthorityCode       int    `mapstructure:"authority-code" json:"authorityCode" yaml:"authority-code"`
	LastPasswordChangeNum int           `mapstructure:"last-password-change-num" json:"lastPasswordChangeNum" yaml:"last-password-change-num"`
	ResetPasswordLength int    `mapstructure:"reset-password-length" json:"resetPasswordLength" yaml:"reset-password-length"`
    
	StatusTrue             string `mapstructure:"status-true" json:"statusTrue" yaml:"status-true"`
	StatusFalse            string `mapstructure:"status-false" json:"statusFalse" yaml:"status-false"`
}
