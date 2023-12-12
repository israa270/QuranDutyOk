package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ebedevelopment/next-gen-tms/server/utils"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/common/response"
	claimCase "github.com/ebedevelopment/next-gen-tms/server/usecase/user/claim"
	"github.com/gin-gonic/gin"
)


func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// jwt info x-token
		// token := c.Request.Header.Get("x-token")
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "Not logged in or no legal access", http.StatusUnauthorized, "error", c)
			c.Abort()
			return
		}
		if jwtService.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "Your of account is logged in offsite or the token is invalid", http.StatusUnauthorized, "error", c)
			c.Abort()
			return
		}
		j := claimCase.NewJWT()
		// parseToken
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == ErrTokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "Authorization already expired", http.StatusUnauthorized, "error", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), http.StatusUnauthorized, "error", c)
			c.Abort()
			return
		}
		// User cover delete of logic  require  want optimize this point compare be if require  want  please  since row open
		//if err, _ = userService.FindUserByUuid(claims.UUID.String()); err != nil {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			// claims.ExpiresAt = time.Now().Unix() + global.GvaConfig.JWT.ExpiresTime
			dr, _ := utils.ParseDuration(global.GvaConfig.JWT.ExpiresTime)
			claims.ExpiresAt = time.Now().Add(dr).Unix()
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		c.Set("claims", claims)
		c.Next()
	}
}
