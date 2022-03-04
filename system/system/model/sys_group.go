package model

import (
	"gorm.io/gorm/clause"
	"time"

	"gorm.io/gorm"
)

// SysGroup 用户组表
type SysGroup struct {
	Id         int            `gorm:"column:id;primaryKey"`
	CreatedAt  time.Time      `gorm:"column:created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
	GroupCode  string         `gorm:"column:group_code;size:45" binding:"max=45"`  // 用户组编号
	GroupName  string         `gorm:"column:group_name;size:45" binding:"max=45"`  // 用户组名称
	LeaderId   int            `gorm:"column:leader_id"`                            // 负责人id
	LeaderName string         `gorm:"column:leader_name;size:45" binding:"max=45"` // 负责人
	Phone      string         `gorm:"column:phone;size:45" binding:"max=45"`       // 电话
	Sort       int            `gorm:"column:sort"`                                 // 排序
	Remark     string         `gorm:"column:remark;size:255" binding:"max=255"`    // 描述/备注
	SysOrgId   int            `gorm:"column:sys_org_id;default:0"`
	SysOrg     *SysOrg        // 一对多,从表
	Users      []*User        `gorm:"many2many:sys_group_has_user;"` // 多对多
}

func (m *SysGroup) TableName() string {
	return "sys_group"
}

func (m *SysGroup) GetM2MJoin(m2m string) string {
	joinStr := ""
	switch m2m {
	case "User":
		joinStr = "JOIN sys_group_has_user ON sys_group_has_user.sys_group_id = sys_group.id JOIN user ON sys_group_has_user.user_id = user.id"

	}
	return joinStr
}

// UpdateM2M 修改多对多关系
func (m *SysGroup) UpdateM2M(db *gorm.DB, field string, addIds, delIds, repIds []int) error {
	if field == "Users" {
		var addList []*User
		var delList []*User
		var repList []*User
		for _, id := range addIds {
			addList = append(addList, &User{Id: id})
		}
		for _, id := range delIds {
			delList = append(delList, &User{Id: id})
		}
		for _, id := range repIds {
			repList = append(repList, &User{Id: id})
		}
		if len(repIds) > 0 {
			return db.Model(m).Omit(field).Association(field).Replace(repList)
		}
		if len(addIds) > 0 {
			return db.Model(m).Debug().Omit(field + ".*").Association(field).Append(addList)
		}
		if len(delIds) > 0 {
			return db.Model(m).Omit(clause.Associations).Association(field).Delete(delList)
		}
	}
	return nil
}
