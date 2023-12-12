package middleware

import (
	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/gin-gonic/gin"
)

// LanguageHandler config language handler
func LanguageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		lang := c.Request.FormValue("lang")
		accept := c.Request.Header.Get("Accept-Language")

		if lang == "" && accept == ""{
			lang ="en"
		}

		global.GvaTranslator.SetTranslatorLanguage(lang, accept)

		c.Next()
	}
}
