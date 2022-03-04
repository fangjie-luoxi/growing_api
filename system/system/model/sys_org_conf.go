package model

import (
	"time"

	"gorm.io/gorm"
)

// SysOrgConf 机构配置
type SysOrgConf struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index:idx_del" json:"-"`
	SysOrgId  int            `gorm:"type:int;comment:o2o"` // o2o
	SysOrg    *SysOrg        // 一对一,从表
}

func (m *SysOrgConf) TableName() string {
	return "sys_org_conf"
}

func (m *SysOrgConf) GetM2MJoin(m2mStr string) string {
	joinStr := ""
	switch m2mStr {

	}
	return joinStr
}
