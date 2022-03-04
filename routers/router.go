package routers

import (
	"context"

	logIO "github.com/fangjie-luoxi/tools/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/fangjie-luoxi/growing_api/global"
	"github.com/fangjie-luoxi/growing_api/middleware"
	"github.com/fangjie-luoxi/growing_api/system/system/service"
	"github.com/fangjie-luoxi/growing_api/system/upload"
)

func NewRouter() *gin.Engine {
	gin.SetMode(global.Config.DefaultString("mode", "debug"))
	r := gin.New()
	if global.Config.String("mode") == "release" {
		loggerIO := logIO.NewLogIO("./static/logs/gin/log.log") // 路由日志
		r.Use(gin.LoggerWithWriter(loggerIO))
	} else {
		r.Use(gin.Logger())
	}
	r.Use(gin.Recovery())
	// 跨域中间件
	r.Use(cors.Default())
	_ = r.SetTrustedProxies(nil)
	// api组
	apiGroup := r.Group(global.Config.DefaultString("appName", "api"))
	apiGroup.Static("/static", "./static") // 静态文件夹
	// 授权中间件
	jwt, err := middleware.JWT(global.Login)
	if err != nil {
		global.Log.Fatal("创建jwt失败", err)
	}
	global.Login.OpenApi(apiGroup, jwt) // 登录开放接口
	if global.Config.DefaultBool("jwt", false) {
		apiGroup.Use(jwt.MiddlewareFunc()) // 开启jwt鉴权
		if global.Config.DefaultBool("perm", false) {
			apiGroup.Use(middleware.NewAuthorizer(global.Casbin.Enforcer, global.DBEngine, global.Cache))
		}
	}
	apiGroup.POST("/upload/files", upload.FilesUpload) // 多文件上传
	global.Login.Api(apiGroup, jwt)                    // 登录路由
	if global.Config.DefaultBool("system", false) {
		systemService := service.NewService(context.Background(), global.DBEngine, global.Casbin.Enforcer, global.Cache.C)
		systemService.SysApi(apiGroup) // 系统管理路由
	}
	genApi(apiGroup)  // 自动生成的路由
	ruleApi(apiGroup) // 自定义路由
	return r
}
