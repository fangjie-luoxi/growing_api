package models

import (
	"time"

	"github.com/fangjie-luoxi/tools/query"
	"gorm.io/gorm"
)

// GrTask 任务表
type GrTask struct {
	Id         int            `gorm:"column:id;primaryKey"`
	CreatedAt  time.Time      `gorm:"column:created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
	UserId     int            `gorm:"column:user_id"`
	GrTargetId int            `gorm:"column:gr_target_id;default:0"`
	TkTitle    string         `gorm:"column:tk_title;size:45" binding:"max=45"`         // 任务标题
	TkContent  string         `gorm:"column:tk_content;size:255" binding:"max=255"`     // 任务内容
	TtNum      float64        `gorm:"column:tt_num"`                                    // 任务数量
	TtUnit     string         `gorm:"column:tt_unit;size:45" binding:"max=45"`          // 任务单位
	Num        float64        `gorm:"column:num"`                                       // 积分
	Rm         string         `gorm:"column:rm;size:255" binding:"max=255"`             // 备注
	Status     string         `gorm:"column:status;size:45;default:b" binding:"max=45"` // 状态 enum:s:完成,n:未完成,r:进行中,b:未开始#
	Date       *time.Time     `gorm:"column:date"`                                      // 任务日期
	User       *User          // 一对多,从表
	GrTarget   *GrTarget      // 一对多,从表
}

// pos1

func (m *GrTask) TableName() string {
	return "gr_task"
}

func (m *GrTask) GetM2MJoin(m2m string) string {
	joinStr := ""
	switch m2m {

	}
	return joinStr
}

func (m *GrTask) Get(db *gorm.DB, q *query.Query) ([]*GrTask, int64, error) {
	var res []*GrTask
	var count int64
	tx := db.Model(m)
	tx.Scopes(q.DBQuery())
	if q.Count != "F" {
		tx.Count(&count)
	}
	if q.Count == "count" {
		return nil, count, tx.Error
	}
	tx.Scopes(q.DBSelect())
	tx.Limit(q.Limit).Offset(q.Offset).Find(&res)
	return res, count, tx.Error
}

func (m *GrTask) GetOne(db *gorm.DB, q *query.Query) error {
	return db.Scopes(q.DBQuery()).First(m).Error
}

func (m *GrTask) Create(db *gorm.DB) error {
	return db.Create(m).Error
}

func (m *GrTask) Update(db *gorm.DB, v map[string]interface{}) error {
	var fieldList []string
	for k := range v {
		fieldList = append(fieldList, k)
	}
	return db.Model(m).Select(fieldList).Updates(m).Error
}

func (m *GrTask) UpdateByQuery(db *gorm.DB, q *query.Query, v map[string]interface{}) error {
	return db.Model(m).Scopes(q.DBQuery()).Updates(v).Error
}

func (m *GrTask) UpdateFull(db *gorm.DB, v map[string]interface{}) error {
	return db.Model(m).Session(&gorm.Session{FullSaveAssociations: true}).Updates(v).Error
}

func (m *GrTask) Updates(db *gorm.DB, v []*GrTask) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		tx.Model(m).Session(&gorm.Session{FullSaveAssociations: true})
		for _, item := range v {
			if err := tx.Updates(item).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// UpdateM2M 修改多对多关系
func (m *GrTask) UpdateM2M(db *gorm.DB, field string, addIds, delIds, repIds []int) error {
	var addI interface{}
	var delI interface{}
	var repI interface{}

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

func (m *GrTask) Delete(db *gorm.DB) error {
	return db.Where("id = ?", m.Id).Delete(m).Error
}

func (m *GrTask) Deletes(db *gorm.DB, ids []int) error {
	return db.Delete([]GrTask{}, ids).Error
}
