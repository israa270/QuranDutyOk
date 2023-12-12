package claim

import (
	"fmt"
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/system/request"
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
	"github.com/ebedevelopment/next-gen-tms/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	SigningKey []byte
}


// NewJWT object
func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GvaConfig.JWT.SigningKey),
	}
}

// CreateClaims claims jwt
func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	bf, _ := utils.ParseDuration(global.GvaConfig.JWT.BufferTime)
	ep, _ := utils.ParseDuration(global.GvaConfig.JWT.ExpiresTime)
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // buffer time1  day  buffer time within can get new of token new token  this time
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,    // sign name effect time between
			ExpiresAt: time.Now().Add(ep).Unix(),    // expiration 7  day configure file
			Issuer:    global.GvaConfig.JWT.Issuer, // sign name of send row send
		},
	}
	
	// claims := request.CustomClaims{
	// 	BaseClaims: baseClaims,
	// 	BufferTime: global.GvaConfig.JWT.BufferTime, // buffer time1  day  buffer time within can get new of token new token  this time
	// 	StandardClaims: jwt.StandardClaims{
	// 		NotBefore: time.Now().Unix() - 1000,                             // sign name effect time between
	// 		ExpiresAt: time.Now().Unix() + global.GvaConfig.JWT.ExpiresTime, // expiration 7  day configure file
	// 		Issuer:    global.GvaConfig.JWT.Issuer,                          // sign name of send row send
	// 	},
	// 	// RegisteredClaims: jwt.RegisteredClaims{
	// 	// 	NotBefore: jwt.NewNumericDate(time.Now() - 1000),                              // sign name effect time between
	// 	// 	ExpiresAt: jwt.NewNumericDate(time.Now().Unix() + global.GvaConfig.JWT.ExpiresTime), // expiration 7  day configure file
	// 	// 	Issuer: global.GvaConfig.JWT.Issuer,

	// 	// },
	// 	// jwt.RegisteredClaims{
	// 	// 	// A usual scenario is to set the expiration time relative to the current time
	// 	// 	ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	// 	// 	IssuedAt:  jwt.NewNumericDate(time.Now()),
	// 	// 	NotBefore: jwt.NewNumericDate(time.Now()),
	// 	// 	Issuer:    "test",
	// 	// 	Subject:   "somebody",
	// 	// 	ID:        "1",
	// 	// 	Audience:  []string{"somebody_else"},
	// 	// },
	// }
	return claims
}

// CreateToken create token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken old token change new token
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := global.GvaConcurrencyControl.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken parse token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf(global.Translate("init.tokenEven"))
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, fmt.Errorf(global.Translate("init.tokenExpire"))
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, fmt.Errorf(global.Translate("init.tokenNotActive"))
			} else {
				return nil, fmt.Errorf(global.Translate("init.notToken"))
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, fmt.Errorf(global.Translate("init.notToken"))

	} else {
		return nil, fmt.Errorf(global.Translate("init.notToken"))
	}
}


// GetClaims get claims for jwt
func  GetClaims(c *gin.Context) (*sysReq.CustomClaims, error) {
	// token := c.Request.Header.Get("x-token")
	token := c.Request.Header.Get("Authorization")

	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.GvaLog.Error("from Gin of Context middle get from jwt parse info failed, please check please request top yes no exist x-token and claims yes no for structure")
	}
	return claims, err
}