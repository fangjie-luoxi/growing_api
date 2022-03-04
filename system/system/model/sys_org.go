package model

import (
	"time"

	"gorm.io/gorm"
)

// SysOrg 机构表
type SysOrg struct {
	Id         int            `gorm:"column:id;primaryKey"` // 机构id
	CreatedAt  time.Time      `gorm:"column:created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
	OrgCode    string         `gorm:"column:org_code;size:45" binding:"max=45"`    // 机构编码
	OrgName    string         `gorm:"column:org_name;size:128" binding:"max=128"`  // 机构名称
	Alias      string         `gorm:"column:alias;size:128" binding:"max=128"`     // 机构名称别名
	LeaderId   int            `gorm:"column:leader_id"`                            // 负责人id
	LeaderName string         `gorm:"column:leader_name;size:45" binding:"max=45"` // 负责人名字
	Phone      string         `gorm:"column:phone;size:45" binding:"max=45"`       // 负责人电话
	Email      string         `gorm:"column:email;size:45" binding:"max=45"`       // 邮箱
	Founder    string         `gorm:"column:founder;size:45" binding:"max=45"`     // 创建人
	Pid        int            `gorm:"column:pid;default:0"`                        // 父机构id
	Sort       int            `gorm:"column:sort"`                                 // 排序
	Children   []*SysOrg      `gorm:"-" json:"children"`                           // 子机构
	SysApis    []*SysApi      // 一对多,主表
	SysDepts   []*SysDept     // 一对多,主表
	SysGroups  []*SysGroup    // 一对多,主表
	SysOrgConf *SysOrgConf    // 一对一,主表
	SysMenus   []*SysMenu     `gorm:"many2many:sys_org_has_sys_menu;"` // 多对多
	SysPosts   []*SysPost     // 一对多,主表
	SysRoles   []*SysRole     // 一对多,主表
	Users      []*User        // 一对多,主表
}

func (m *SysOrg) TableName() string {
	return "sys_org"
}

func (m *SysOrg) GetM2MJoin(m2m string) string {
	joinStr := ""
	switch m2m {
	case "SysMenu":
		joinStr = "JOIN sys_org_has_sys_menu ON sys_org_has_sys_menu.sys_org_id = sys_org.id JOIN sys_menu ON sys_org_has_sys_menu.sys_menu_id = sys_menu.id"

	}
	return joinStr
}
