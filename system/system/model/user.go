package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户表
type User struct {
	Id            int            `gorm:"column:id;primaryKey"` // 用户id
	CreatedAt     time.Time      `gorm:"column:created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-"`
	LoginCode     string         `gorm:"column:login_code;uniqueIndex:idx_login_code;size:45" binding:"max=45"` // 登录账户
	Password      string         `gorm:"column:password;size:255" binding:"max=255"`                            // 密码
	UserCode      string         `gorm:"column:user_code;uniqueIndex:idx_user_code;size:45" binding:"max=45"`   // 用户编码
	UserName      string         `gorm:"column:user_name;index:idx_user_name;size:45" binding:"max=45"`         // 用户名称
	Birth         *time.Time     `gorm:"column:birth"`                                                          // 生日
	Sex           string         `gorm:"column:sex;size:45"`                                                    // 性别 enum:b:男,g:女,s:保密#
	CardNo        string         `gorm:"column:card_no;size:45" binding:"max=45"`                               // 证件号码
	Avatar        string         `gorm:"column:avatar;size:255" binding:"max=255"`                              // 头像
	Phone         string         `gorm:"column:phone;size:45" binding:"max=45"`                                 // 手机号码
	Email         string         `gorm:"column:email;size:255" binding:"max=255"`                               // 邮箱
	UserType      string         `gorm:"column:user_type;;size:45" binding:"max=10"`                            // 用户类型 enum:super:超级管理员,admin:机构管理员,user:普通用户#
	Status        string         `gorm:"column:status;default:r;size:45" binding:"max=45"`                      // 状态 enum:r:正常,c:禁用#
	LastLoginTime *time.Time     `gorm:"column:last_login_time"`                                                // 最后登录时间
	Address       string         `gorm:"column:address;size:255" binding:"max=255"`                             // 地址
	Profile       string         `gorm:"column:profile"`                                                        // 个人简介
	SysDeptId     int            `gorm:"column:sys_dept_id;default:0"`
	SysOrgId      int            `gorm:"column:sys_org_id;default:0"`
	SysPostId     int            `gorm:"column:sys_post_id;default:0"`
	SysGroups     []*SysGroup    `gorm:"many2many:sys_group_has_user;"` // 多对多
	SysRoles      []*SysRole     `gorm:"many2many:sys_role_has_user;"`  // 多对多
	SysDept       *SysDept       // 一对多,从表
	SysOrg        *SysOrg        // 一对多,从表
	SysPost       *SysPost       // 一对多,从表
}

func (m *User) TableName() string {
	return "user"
}

func (m *User) GetM2MJoin(m2m string) string {
	joinStr := ""
	switch m2m {
	case "SysGroup":
		joinStr = "JOIN sys_group_has_user ON sys_group_has_user.user_id = user.id JOIN sys_group ON sys_group_has_user.sys_group_id = sys_group.id"
	case "SysRole":
		joinStr = "JOIN sys_role_has_user ON sys_role_has_user.user_id = user.id JOIN sys_role ON sys_role_has_user.sys_role_id = sys_role.id"

	}
	return joinStr
}

// UpdateM2M 修改多对多关系
func (m *User) UpdateM2M(db *gorm.DB, field string, addIds, delIds, repIds []int) error {
	var addI interface{}
	var delI interface{}
	var repI interface{}

	if field == "SysGroups" {
		var addList []SysGroup
		var delList []SysGroup
		var repList []SysGroup
		for _, id := range addIds {
			addList = append(addList, SysGroup{Id: id})
		}
		for _, id := range delIds {
			delList = append(delList, SysGroup{Id: id})
		}
		for _, id := range repIds {
			repList = append(repList, SysGroup{Id: id})
		}
		addI = addList
		delI = delList
		repI = repList
	}

	if field == "SysRoles" {
		var addList []SysRole
		var delList []SysRole
		var repList []SysRole
		for _, id := range addIds {
			addList = append(addList, SysRole{Id: id})
		}
		for _, id := range delIds {
			delList = append(delList, SysRole{Id: id})
		}
		for _, id := range repIds {
			repList = append(repList, SysRole{Id: id})
		}
		addI = addList
		delI = delList
		repI = repList
	}

	if len(repIds) > 0 {
		return db.Model(m).Omit(field + ".*").Association(field).Replace(repI)
	}
	if len(addIds) > 0 {
		return db.Model(m).Omit(field + ".*").Association(field).Append(addI)
	}
	if len(delIds) > 0 {
		return db.Model(m).Omit(field + ".*").Association(field).Delete(delI)
	}
	return nil
}
