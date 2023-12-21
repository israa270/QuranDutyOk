package config

import (
	// "fmt"

	// "github.com/ebedevelopment/next-gen-tms/server/global"
	// "go.uber.org/zap"
)

// Mysql struct
type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                             // server address
	WPort         string `mapstructure:"w-port" json:"wPort" yaml:"w-port"`                             // port
	RPort         string `mapstructure:"r-port" json:"rPort" yaml:"r-port"` 
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                       // advanced configure
	DbName       string `mapstructure:"db-name" json:"dbName" yaml:"db-name"`                     // database name
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                 // database username
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                 // database password
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"` // null idle of max number of connections
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"` // open arrive database of max number of connections
	LogMode      string `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`                  // yes no turn on Gorm global log
	LogZap       bool   `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`                     // yes no by zap write log file
	LimitRecords int     `mapstructure:"limit-records" json:"limitRecords" yaml:"limit-records"`
    // ENV          string   `mapstructure:"env" json:"env" yaml:"env"`
}



// Dsn db connection
func (m *Mysql) Dsn(writer bool) string {
	if writer{
	    return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.WPort + ")/" + m.DbName + "?" + m.Config

	}
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.RPort + ")/" + m.DbName + "?" + m.Config
}

// GetLogMode get log db
func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
