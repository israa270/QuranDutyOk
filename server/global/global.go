package global

import (
	// "sync"

	// pool "github.com/bitleak/go-redis-pool"
	"github.com/ebedevelopment/next-gen-tms/server/utils/timer"
	"github.com/ebedevelopment/next-gen-tms/server/utils/translate"

	// "github.com/shaj13/go-guardian/auth/strategies/ldap"
	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/ebedevelopment/next-gen-tms/server/config"

	// "github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GvaDB     *gorm.DB
	GvaDBRead  *gorm.DB
	GvaDBList map[string]*gorm.DB
	// GvaRedis  *redis.Client

	// GvaRedis  *pool.Pool
	// GvaRedisCluster  *pool.Pool
	GvaConfig config.Server
	GvaVP     *viper.Viper

	GvaLog                *zap.Logger
	GvaTimer              timer.Timer = timer.NewTimerTask()
	GvaConcurrencyControl             = &singleflight.Group{}

	BlackCache local_cache.Cache
	// lock       sync.RWMutex

	GvaTranslator translate.Translator
	GvaMenu       map[string]translate.Menu //support menu translate
	GvaLoggerMessage  = make(map[string]*config.MessageLogger)

	// GvaLdap *ldap.Config
)

// GetGlobalDBByDBName by name get db list middle of db
// func GetGlobalDBByDBName(dbName string) *gorm.DB {
// 	lock.RLock()
// 	defer lock.RUnlock()
// 	return GvaDBList[dbName]
// }

// MustGetGlobalDBByDBName by name get db if not exist but panic
// func MustGetGlobalDBByDBName(dbName string) *gorm.DB {
// 	lock.RLock()
// 	defer lock.RUnlock()
// 	db, ok := GvaDBList[dbName]
// 	if !ok || db == nil {
// 		panic("db no init")
// 	}
// 	return db
// }

func Translate(msg string) string {
	if GvaTranslator.IsInit {
		message := GvaTranslator.TranslateMessage(msg)
		return message
	}

	return msg
}
