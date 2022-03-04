package model

import (
	"time"

	"gorm.io/gorm"
)

// SysMenu 菜单表
type SysMenu struct {
	Id                 int            `gorm:"column:id;primaryKey"`
	CreatedAt          time.Time      `gorm:"column:created_at"`
	UpdatedAt          time.Time      `gorm:"column:updated_at"`
	DeletedAt          gorm.DeletedAt `json:"-"`
	Name               string         `gorm:"column:name;size:45" binding:"max=45"`           // 菜单名称
	Path               string         `gorm:"column:path;size:255" binding:"max=255"`         // 路由path
	HideChildrenInMenu string         `gorm:"column:hide_children_in_menu;default:0;size:10"` // 是否展示子菜单
	HideInMenu         string         `gorm:"column:hide_in_menu;default:0;size:10"`          // 是否展示菜单
	Authority          string         `gorm:"column:authority;size:255" binding:"max=255"`    // 权限
	Pid                int            `gorm:"column:pid;default:0"`                           // 父菜单id
	Sort               int            `gorm:"column:sort"`                                    // 排序
	Icon               string         `gorm:"column:icon;size:255" binding:"max=255"`         // 图标
	Children           []*SysMenu     `gorm:"-" json:"children"`                              // 子菜单
	SysOrgs            []*SysOrg      `gorm:"many2many:sys_org_has_sys_menu;"`                // 多对多
	SysRoles           []*SysRole     `gorm:"many2many:sys_role_has_sys_menu;"`               // 多对多
}

func (m *SysMenu) TableName() string {
	return "sys_menu"
}

func (m *SysMenu) GetM2MJoin(m2m string) string {
	joinStr := ""
	switch m2m {
	case "SysOrg":
		joinStr = "JOIN sys_org_has_sys_menu ON sys_org_has_sys_menu.sys_menu_id = sys_menu.id JOIN sys_org ON sys_org_has_sys_menu.sys_org_id = sys_org.id"
	case "SysRole":
		joinStr = "JOIN sys_role_has_sys_menu ON sys_role_has_sys_menu.sys_menu_id = sys_menu.id JOIN sys_role ON sys_role_has_sys_menu.sys_role_id = sys_role.id"
	}
	return joinStr
}
