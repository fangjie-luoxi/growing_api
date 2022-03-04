package model

import (
	"gorm.io/gorm"
	"time"
)

// SysDict 字典表
type SysDict struct {
	Id        int            `gorm:"column:id;primaryKey"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Category  string         `gorm:"column:category;size:45;index:idx_category" binding:"max=45"`     // 中文名
	NameCn    string         `gorm:"column:name_cn;size:45" binding:"max=45"`                         // 中文名
	NameEn    string         `gorm:"column:name_en;size:45;uniqueIndex:idx_name_en" binding:"max=45"` // 英文名
	Value     string         `gorm:"column:value"`                                                    // 值
	Status    string         `gorm:"column:status;size:10;default:r" binding:"max=10"`                // 状态 enum:r:启用,c:停用#
	Describe  string         `gorm:"column:describe;size:255" binding:"max=255"`                      // 描述
	Extend    string         `gorm:"column:extend"`                                                   // 扩展字段
}

func (m *SysDict) TableName() string {
	return "sys_dict"
}

func (m *SysDict) GetM2MJoin(m2m string) string {
	joinStr := ""
	switch m2m {

	}
	return joinStr
}
