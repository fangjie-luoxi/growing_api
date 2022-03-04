// Package casbin
// 权限控制
package casbin

import (
	_ "embed"
	"gorm.io/gorm"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

//go:embed rbac_model.conf
var modelStr string // 静态配置文件

// SetUp casbin的初始化程序
func SetUp(orm *gorm.DB) (*Instance, error) {
	a, err := gormadapter.NewAdapterByDB(orm)
	if err != nil {
		return nil, err
	}
	m, err := model.NewModelFromString(modelStr)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, err
	}
	// 从数据库加载策略
	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}
	// 自动保存策略
	e.EnableAutoSave(true)
	return &Instance{Enforcer: e}, nil
}

type Instance struct {
	Enforcer *casbin.Enforcer
}

func (i *Instance) AddPolicies() {
}
