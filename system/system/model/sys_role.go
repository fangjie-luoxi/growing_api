package model

import (
	"time"

	"gorm.io/gorm"
)

// SysRole 角色表
type SysRole struct {
	Id        int            `gorm:"column:id;primaryKey"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
	RoleCode  string         `gorm:"column:role_code;size:45" binding:"max=45"`       // 角色编码
	RoleName  string         `gorm:"column:role_name;size:45" binding:"max=45"`       // 角色名称
	Status    string         `gorm:"column:status;default:r;size:10" binding:"max=1"` // 状态 enum:r:正常,c:禁用#
	Remark    string         `gorm:"column:remark;size:255" binding:"max=255"`        // 备注
	SysOrgId  int            `gorm:"column:sys_org_id;default:0"`
	SysOrg    *SysOrg        // 一对多,从表
	SysMenus  []*SysMenu     `gorm:"many2many:sys_role_has_sys_menu;"` // 多对多
	Users     []*User        `gorm:"many2many:sys_role_has_user;"`     // 多对多
}

func (m *SysRole) TableName() string {
	return "sys_role"
}

func (m *SysRole) GetM2MJoin(m2m string) string {
	joinStr := ""
	switch m2m {
	case "SysMenu":
		joinStr = "JOIN sys_role_has_sys_menu ON sys_role_has_sys_menu.sys_role_id = sys_role.id JOIN sys_menu ON sys_role_has_sys_menu.sys_menu_id = sys_menu.id"
	case "User":
		joinStr = "JOIN sys_role_has_user ON sys_role_has_user.sys_role_id = sys_role.id JOIN user ON sys_role_has_user.user_id = user.id"

	}
	return joinStr
}
