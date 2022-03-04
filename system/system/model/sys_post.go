package model

import (
	"time"

	"gorm.io/gorm"
)

// SysPost 岗位表
type SysPost struct {
	Id        int            `gorm:"column:id;primaryKey"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
	PostCode  string         `gorm:"column:post_code;size:45" binding:"max=45"` // 岗位代码
	PostName  string         `gorm:"column:post_name;size:45" binding:"max=45"` // 岗位
	Sort      int            `gorm:"column:sort"`                               // 排序
	Status    string         `gorm:"column:status;size:10" binding:"max=1"`     // 状态 enum:r:启用,c:停用#
	Remark    string         `gorm:"column:remark;size:255" binding:"max=255"`  // 备注
	CreateBy  string         `gorm:"column:create_by;size:45" binding:"max=45"` // 创建人
	SysOrgId  int            `gorm:"column:sys_org_id;default:0"`
	SysOrg    *SysOrg        // 一对多,从表
	Users     []*User        // 一对多,主表
}

func (m *SysPost) TableName() string {
	return "sys_post"
}

func (m *SysPost) GetM2MJoin(m2m string) string {
	joinStr := ""
	switch m2m {

	}
	return joinStr
}
