package initialize

import (
	"net/http"
	"time"

	_ "github.com/ebedevelopment/next-gen-tms/server/docs"
	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/middleware"
	"github.com/ebedevelopment/next-gen-tms/server/router"
	"github.com/gin-contrib/gzip"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	// ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Routers initialization total routing
func Routers() *gin.Engine {
	Router := gin.Default()

	Router.Use(ginzap.Ginzap(global.GvaLog, time.RFC3339, true))
	Router.Use(ginzap.RecoveryWithZap(global.GvaLog, true))
	// if you want not use nginx agent before side page ，can with update web/.env.production
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	//row pack order npm run build exist open under 4 row notes
	// Router.LoadHTMLGlob("./dist/*.html")
	// Router.Static("/favicon.ico", "./dist/favicon.ico")

	Router.Static("/resource/excel", "./resource/excel")
	Router.Static("/resource/template", "./resource/template")

	Router.Static("/assets/icons", "./assets/icons")

	Router.Static("/resource/xmlFiles/push", "./resource/xmlFiles/push")
	Router.Static("/downloads", "./downloads")
	// Router.StaticFS("/resource/xmlFiles/push",http.Dir("/resource/xmlFiles/push"))

	Router.StaticFS(global.GvaConfig.Local.Path, http.Dir(global.GvaConfig.Local.Path)) // forUser top image and file address
	// Router.Use(middleware.LoadTls())  // if require  want use https please open this middle member before to  core/server.go use start up model change for Router.RunTLS("port"," ofcre/pemfile","ofKeyFile")
	global.GvaLog.Debug("use middleware logger")
	// cross domain， as require cross domain can open under of notes
	Router.Use(middleware.Cors()) // cors row all cross domain please request

	Router.Use(gzip.Gzip(gzip.DefaultCompression)) //default compression

	//Router.Use(middleware.CorsByRules()) //  Play cross -domain requests in accordance with the rules of configuration
	global.GvaLog.Debug("use middleware cors")
	//allow multi language handling
	global.GvaLog.Debug("use middleware translator")
	Router.Use(middleware.LanguageHandler()) // add global language handler middleware
	// end of adding
	// Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.GET("/", func(c *gin.Context) {
		// return no data
		c.JSON(http.StatusOK, gin.H{})
	})

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	global.GvaLog.Debug("register swagger handler")

	// get routing group instance
	systemRouter := router.GroupRouterApp.System
	tmsRouter := router.GroupRouterApp.Tms
	managementRouter := router.GroupRouterApp.Management

	PublicGroup := Router.Group("")
	{
		// health
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // register base be routing not do authentication
		// systemRouter.InitInitRouter(PublicGroup) // automatic initialization related
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{

		systemRouter.InitUserRouter(PrivateGroup) // registerUser routing

		systemRouter.InitSysOperationRecordRouter(PrivateGroup) // operations record
		systemRouter.InitMenuRouter(PrivateGroup)
		tmsRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // file upload load be routing

		tmsRouter.InitManufacturerRouter(PrivateGroup)
		tmsRouter.InitDataExcelRouter(PrivateGroup)

		managementRouter.InitTeacherRouter(PrivateGroup)
		managementRouter.InitClassRouter(PrivateGroup)
		managementRouter.InitStudentRouter(PrivateGroup)
		managementRouter.InitHomeWorkRouter(PrivateGroup)
	}

	global.GvaLog.Debug("router register success")
	return Router
}
