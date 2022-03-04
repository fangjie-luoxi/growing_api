package model

import (
	"time"

	"gorm.io/gorm"
)

// SysApi api接口
type SysApi struct {
	Id          int            `gorm:"column:id;primaryKey"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"`
	Path        string         `gorm:"column:path;index:idx_del;size:255" binding:"max=255"` // api路径
	Description string         `gorm:"column:description;size:255" binding:"max=255"`        // api中文描述
	ApiGroup    string         `gorm:"column:api_group;size:255" binding:"max=255"`          // api组
	Method      string         `gorm:"column:method;size:255" binding:"max=255"`             // 方法 enum:POST:创建,GET:查看,PUT:更新,DELETE:删除#
	ApiTp       string         `gorm:"column:api_tp;size:10;default:perm" binding:"max=10"`  // api类型 enum:perm:接口权限,jwt:jwt鉴权#
	SysOrgId    int            `gorm:"column:sys_org_id;default:0"`
	SysOrg      *SysOrg        // 一对多,从表
}

func (m *SysApi) TableName() string {
	return "sys_api"
}

func (m *SysApi) GetM2MJoin(s string) string {
	return ""
}
