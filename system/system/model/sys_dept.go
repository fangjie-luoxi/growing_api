package model

import (
	"time"

	"gorm.io/gorm"
)

// SysDept 部门表
type SysDept struct {
	Id         int            `gorm:"column:id;primaryKey"`
	CreatedAt  time.Time      `gorm:"column:created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
	DeptCode   string         `gorm:"column:dept_code;size:255" binding:"max=255"` // 部门编码
	DeptName   string         `gorm:"column:dept_name;size:255" binding:"max=255"` // 部门名称
	Pid        int            `gorm:"column:pid;default:0"`                        // 上级部门
	LeaderId   int            `gorm:"column:leader_id"`                            // 领导id
	LeaderName string         `gorm:"column:leader_name;size:45" binding:"max=45"` // 领导
	Email      string         `gorm:"column:email;size:45" binding:"max=45"`       // 邮箱
	Phone      string         `gorm:"column:phone;size:45" binding:"max=45"`       // 手机号码
	Status     string         `gorm:"column:status;default:r;size:1"`              // 状态 enum:r:启用,c:停用#
	SysOrgId   int            `gorm:"column:sys_org_id;default:0"`
	Children   []*SysDept     `gorm:"-" json:"children"` // 子部门
	SysOrg     *SysOrg        // 一对多,从表
	Users      []*User        // 一对多,主表
}

func (m *SysDept) TableName() string {
	return "sys_dept"
}

func (m *SysDept) GetM2MJoin(m2m string) string {
	joinStr := ""
	switch m2m {

	}
	return joinStr
}
