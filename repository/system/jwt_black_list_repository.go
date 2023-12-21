package system

import (
	"fmt"


	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/system"
	// "github.com/ebedevelopment/next-gen-tms/server/utils"

)

// JwtRepository jwt object for fn jwt
type JwtRepository struct{}

// JsonInBlacklist json jwt
func (j *JwtRepository) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	if global.GvaDB == nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].DB)
		// response.FailWithMessage(global.Translate("sysInitDB.db"), http.StatusRepositoryUnavailable, "warning", c)
		return fmt.Errorf(global.Translate("sysInitDB.db"))
	}
	err = global.GvaDB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

// IsBlacklist JWT exist black name one within section
func (j *JwtRepository) IsBlacklist(jwt string) bool {
	if global.GvaDB == nil {
		return true
	}
	_, ok := global.BlackCache.Get(jwt)
	return ok

	// err := global.GvaDB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	// return !isNotFound

	// err := global.GvaDB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	// return !isNotFound
}

// // GetRedisJWT from redis get jwt
// func (j *JwtRepository) GetRedisJWT(userName string) (redisJWT string, err error) {
// 	redisJWT, err = global.GvaRedis.Get(context.Background(), userName).Result()
// 	return redisJWT, err
// }

// GetRedisJWT from redis get jwt
// func (j *JwtRepository) GetRedisJWT(userName string) (redisJWT string, err error) {
//    redisValue := global.GvaRedis.Get(userName)

//    log.Println("redis : ", redisValue.Err() , " ---------------", redisValue.String() )
// 	// redisJWT, err = global.GvaRedis.Get(userName).Result()
// 	return redisJWT, err
// }

// func (j *JwtRepository) GetRedisJWT(userName string) (redisJWT string, err error) {
// 	redisValue := global.GvaRedis.Get(userName)
//     if redisValue.Val() == ""{
// 		return "" , fmt.Errorf("key is not found")
// 	}

// 	 redisJWT, err = redisValue.Result()
// 	//  redisJWT, err = global.GvaRedis.Get(userName).Result()
// 	 return redisJWT, err
//  }

// // SetRedisJWT jwt set redis and setup expiration
// func (j *JwtRepository) SetRedisJWT(jwt string, userName string) (err error) {
// 	// this point expiration equal jwt expiration
// 	// timer := time.Duration(global.GvaConfig.JWT.ExpiresTime) * time.Second
// 	// err = global.GvaRedis.Set(context.Background(), userName, jwt, timer).Err()
// 	// return err

// 	dr, err := utils.ParseDuration(global.GvaConfig.JWT.ExpiresTime)
// 	if err != nil {
// 		return err
// 	}
// 	timer := dr
// 	err = global.GvaRedis.Set(context.Background(), userName, jwt, timer).Err()
// 	return err
// }

// SetRedisJWT jwt set redis and setup expiration
// func (j *JwtRepository) SetRedisJWT(jwt string, userName string) (err error) {
// 	// this point expiration equal jwt expiration
// 	// timer := time.Duration(global.GvaConfig.JWT.ExpiresTime) * time.Second
// 	// err = global.GvaRedis.Set(context.Background(), userName, jwt, timer).Err()
// 	// return err

// 	dr, err := utils.ParseDuration(global.GvaConfig.JWT.ExpiresTime)
// 	if err != nil {
// 		return err
// 	}
// 	timer := dr
// 	err = global.GvaRedis.Set(userName, jwt, timer).Err()
// 	return err
// }

// LoadAll tokens
// func (j *JwtRepository) LoadAll() {
// 	var data []string
// 	err := global.GvaDB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
// 	if err != nil {
// 		global.GvaLog.Error("load database jwt black name one failed!", zap.Error(err))
// 		return
// 	}
// 	for i := 0; i < len(data); i++ {
// 		global.BlackCache.SetDefault(data[i], struct{}{})
// 	}
// }
