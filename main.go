package main

import (
	_ "embed"

	"github.com/fangjie-luoxi/tools/config"
	"github.com/fangjie-luoxi/tools/log"
	"github.com/fangjie-luoxi/tools/response"

	"github.com/fangjie-luoxi/growing_api/global"
	"github.com/fangjie-luoxi/growing_api/pkg/cache"
	"github.com/fangjie-luoxi/growing_api/pkg/casbin"
	"github.com/fangjie-luoxi/growing_api/pkg/orm"
	"github.com/fangjie-luoxi/growing_api/pkg/wechat"
	"github.com/fangjie-luoxi/growing_api/routers"
	"github.com/fangjie-luoxi/growing_api/rules/cron"
	"github.com/fangjie-luoxi/growing_api/system/login"
)

//go:embed conf/app.yaml
var confStr string // 静态配置文件

func main() {
	setUp()
	cron.Cron()

	r := routers.NewRouter()
	_ = r.Run(":" + global.Config.DefaultString("httpPort", "8080"))
}

// 初始化
func setUp() {
	var err error
	global.Config = config.SetUp(confStr)     // 初始化配置文件
	global.DBEngine = orm.SetUp()             // 初始化数据库
	global.Log = log.Setup("")                // 初始化日志
	global.Resp = response.NewResp("antd", 4) // 初始化响应
	global.Cache = cache.SetUp()              // 初始化缓存

	global.Casbin, err = casbin.SetUp(global.DBEngine) // 初始化casbin
	if err != nil {
		panic("初始化casbin失败: " + err.Error())
	}
	// 微信相关初始化
	if global.Config.DefaultBool("wechat.open", false) {
		global.Wechat = wechat.NewWechat(global.Config)
	}
	global.Login = login.NewLogin(global.DBEngine, global.Config, global.Wechat) // 登录处理
}
