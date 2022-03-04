package service

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/fangjie-luoxi/growing_api/system/system/model"
)

type Service struct {
	ctx            context.Context  // 上下文
	engine         *gorm.DB         // 数据库驱动
	casbinEnforcer *casbin.Enforcer // 权限控制
	cache          *cache.Cache     // 缓存
}

func NewService(ctx context.Context, engine *gorm.DB, enforcer *casbin.Enforcer, cache *cache.Cache) *Service {
	model.SetUp(engine)
	return &Service{
		ctx:            ctx,
		engine:         engine,
		casbinEnforcer: enforcer,
		cache:          cache,
	}
}
