package model

import (
	"gorm.io/gorm"
	"time"
)

// SysConf 系统设置表
type SysConf struct {
	Id        int            `gorm:"column:id;primaryKey"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
	// 图标、favicon是展示在浏览器标签页上的内容，严格来说它是属于浏览器 meta 的一部分，
	// 浏览器认为 favicon 不会经常改动做了非常强的缓存。所以我们并没有做动态修改 favicon 的方案
	// Favicon string `gorm:"column:favicon;size:255" binding:"max=255"`
	Logo    string `gorm:"column:logo;size:255" binding:"max=255"`  // logo
	Title   string `gorm:"column:title;size:45" binding:"max=45"`   // 系统标题
	Version string `gorm:"column:version;size:45" binding:"max=45"` // 版本

	NavTheme     string `gorm:"column:nav_theme;size:45" binding:"max=45"`    // 菜单的颜色 'realDark'|"light" | "dark"
	HeaderTheme  string `gorm:"column:header_theme;size:45" binding:"max=45"` // 头部颜色 "light" | "dark"
	HeaderHeight int    `gorm:"column:header_height"`                         // 头部菜单的高度
	Layout       string `gorm:"column:layout;size:45" binding:"max=45"`       // layout格式 'side' | 'top' | 'mix'
	ContentWidth int    `gorm:"column:content_width"`                         // 内容宽度
	// PrimaryColor string `gorm:"column:primary_color;size:45" binding:"max=45"` // 主题颜色、暂不支持
}

func (m *SysConf) TableName() string {
	return "sys_conf"
}

func (m *SysConf) GetM2MJoin(m2m string) string {
	joinStr := ""
	switch m2m {

	}
	return joinStr
}

// SysConfLayout 符合Layout格式的SysConf
type SysConfLayout struct {
	Id           int    `json:"-"`
	Logo         string `json:"logo"`         // logo
	Title        string `json:"title"`        // 系统标题
	NavTheme     string `json:"navTheme"`     // 菜单颜色
	HeaderTheme  string `json:"headerTheme"`  // 头部菜单的颜色 "light" | "dark"
	HeaderHeight int    `json:"headerHeight"` // 头部菜单的高度
	Layout       string `json:"layout"`       // layout布局 'side' | 'top' | 'mix'
	ContentWidth int    `json:"contentWidth"` // 内容宽度
}
