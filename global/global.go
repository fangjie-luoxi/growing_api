package global

import (
	"github.com/fangjie-luoxi/tools/config"
	"github.com/fangjie-luoxi/tools/response"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/fangjie-luoxi/growing_api/pkg/cache"
	"github.com/fangjie-luoxi/growing_api/pkg/casbin"
	"github.com/fangjie-luoxi/growing_api/pkg/wechat"
	"github.com/fangjie-luoxi/growing_api/system/login"
)

var (
	DBEngine *gorm.DB           // 数据库驱动
	Config   *config.Config     // 配置
	Log      *zap.SugaredLogger // 日志
	Resp     *response.Resp     // 响应
	Cache    *cache.Cache       // 缓存
	Casbin   *casbin.Instance   // 权限认证
	Login    *login.Login       // 登录处理
	Wechat   *wechat.Wechat     // 处理
)
