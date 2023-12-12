package system

import (
	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/system"
	repository "github.com/ebedevelopment/next-gen-tms/server/repository/system"
	"go.uber.org/zap"
)

// JwtService jwt object for fn jwt
type JwtService struct {
	jwtRepository repository.JwtRepository
}

// JsonInBlacklist json jwt
func (j *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist)  error {
	return j.jwtRepository.JsonInBlacklist(jwtList)
}

// IsBlacklist JWT exist black name one within section
func (j *JwtService) IsBlacklist(jwt string) bool {
	return j.jwtRepository.IsBlacklist(jwt)
}

// // GetRedisJWT from redis get jwt
// func (j *JwtService) GetRedisJWT(userName string) ( string,  error) {
// 	return j.jwtRepository.GetRedisJWT(userName)
// }

// // SetRedisJWT jwt set redis and setup expiration
// func (j *JwtService) SetRedisJWT(jwt string, userName string)  error {
// 	return j.jwtRepository.SetRedisJWT(jwt, userName)
// }

func LoadAll(){
	var data []string
	err := global.GvaDB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.GvaLog.Error("load database jwt black name one failed!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	}
}
// // LoadAll tokens
// func (j *JwtService) LoadAll() {
// 	 j.jwtRepository.LoadAll()
// }