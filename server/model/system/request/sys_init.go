package request

import (
	"fmt"

	"github.com/ebedevelopment/next-gen-tms/server/config"
)

// InitDB struct
type InitDB struct {
	DBType   string `json:"dbType"`                      // database type
	Host     string `json:"host"`                        // server address
	WPort     string `json:"wPort"`                        // database connect port
	UserName string `json:"userName" binding:"required"` // database username
	Password string `json:"password"`                    // database password
	DBName   string `json:"dbName" binding:"required"`   // database name
	Writer     bool   `json:"-"`
	ENV     string   `json:"-"`
}

// MysqlEmptyDsn mysql null database establish library link
func (i *InitDB) MysqlEmptyDsn(writer bool) string {
	if i.Host == "" {
		i.Host = "127.0.0.1"
	}
	if i.WPort == "" {
		i.WPort = "6446"
	}


		return fmt.Sprintf("%s:%s@tcp(%s:%s)/", i.UserName, i.Password, i.Host, i.WPort)

}


// PgsqlEmptyDsn pgsql null database establish library link
func (i *InitDB) PgsqlEmptyDsn(writer bool) string {
	if i.Host == "" {
		i.Host = "127.0.0.1"
	}
	if i.WPort == "" {
		i.WPort = "3306"
	}



	return "host=" + i.Host + " user=" + i.UserName + " password=" + i.Password + " port=" + i.WPort + " " + "sslmode=disable TimeZone=Asia/Shanghai"

}

// ToMysqlConfig convert config.Mysql
func (i *InitDB) ToMysqlConfig() config.Mysql {
	return config.Mysql{
		Path:         i.Host,
		WPort:         i.WPort,
		DbName:       i.DBName,
		Username:     i.UserName,
		Password:     i.Password,
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		LogMode:      "error",
		Config:       "charset=utf8mb4&parseTime=True&loc=Local",
		LogZap:        true,
		LimitRecords:  1000,
	}
}

// // ToPgsqlConfig convert config.Pgsql
// func (i *InitDB) ToPgsqlConfig() config.Pgsql {
// 	return config.Pgsql{
// 		Path:         i.Host,
// 		WPort:        i.WPort,
// 		RPort:        i.RPort,
// 		DbName:       i.DBName,
// 		Username:     i.UserName,
// 		Password:     i.Password,
// 		MaxIdleConns: 10,
// 		MaxOpenConns: 100,
// 		LogMode:      "error",
// 		Config:       "sslmode=disable TimeZone=Asia/Shanghai",
// 	}
// }
