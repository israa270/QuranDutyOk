package middleware

import (
	"net/http"

	"github.com/ebedevelopment/next-gen-tms/server/config"
	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/gin-gonic/gin"
)

// Cors cors row have cross domain please request and put row have OPTIONS method
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
		c.Header("Access-Control-Allow-Credentials", "true")

		//  put row have OPTIONS method
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		//  point login please request
		c.Next()
	}
}

// CorsByRules  according configure point login cross domain please request
func CorsByRules() gin.HandlerFunc {
	//  put row all
	if global.GvaConfig.Cors.Mode == "allow-all" {
		return Cors()
	}
	return func(c *gin.Context) {
		whitelist := checkCors(c.GetHeader("origin"))

		// by check, add please request top
		if whitelist != nil {
			c.Header("Access-Control-Allow-Origin", whitelist.AllowOrigin)
			c.Header("Access-Control-Allow-Headers", whitelist.AllowHeaders)
			c.Header("Access-Control-Allow-Methods", whitelist.AllowMethods)
			c.Header("Access-Control-Expose-Headers", whitelist.ExposeHeaders)
			if whitelist.AllowCredentials {
				c.Header("Access-Control-Allow-Credentials", "true")
			}
		}

		// white list name model and not by check，point login please request
		if whitelist == nil && global.GvaConfig.Cors.Mode == "strict-whitelist" && !(c.Request.Method == "GET" && c.Request.URL.Path == "/health") {
			c.AbortWithStatus(http.StatusForbidden)
		} else {
			// now hit list name one model， none yes no by check put row have OPTIONS method
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
			}
		}

		//  point login please request
		c.Next()
	}
}

func checkCors(currentOrigin string) *config.CORSWhitelist {
	for _, whitelist := range global.GvaConfig.Cors.Whitelist {
		// configure middle of cross domain top
		if currentOrigin == whitelist.AllowOrigin {
			return &whitelist
		}
	}
	return nil
}
