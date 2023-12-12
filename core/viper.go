package core

import (
	"flag"
	"fmt"
	"os"

	// "time"

	// "github.com/songzhibin97/gkit/cache/local_cache"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/utils"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Viper
func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // firstLevel:order row > env variable > default value
			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
				config = utils.ConfigFile
				fmt.Printf("you is use config of default value,config of path for %v\n", utils.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("you is use GvaConfig env variable,config of path for %v\n", config)
			}
		} else {
			fmt.Printf("you is use order row parameter transfer of value,config of path for %v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("you is use func Viper() transfer of value,config of path for %v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GvaConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GvaConfig); err != nil {
		fmt.Println(err)
	}
	// root
	// know root location arrive correspond migration location ,ensure root path have effect
	// global.GvaConfig.AutoCode.Root, _ = filepath.Abs("..")


	// global.BlackCache = local_cache.NewCache(
	// 	local_cache.SetDefaultExpire(time.Second * time.Duration(global.GvaConfig.JWT.ExpiresTime)),
	// )
	return v
}
