package models

import (
	"time"

	"github.com/fangjie-luoxi/tools/query"
	"gorm.io/gorm"
)

// GrTarget 目标表
type GrTarget struct {
	Id        int            `gorm:"column:id;primaryKey"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
	UserId    int            `gorm:"column:user_id"`
	TtTitle   string         `gorm:"column:tt_title;size:45" binding:"max=45"`           // 目标标题
	TtContent string         `gorm:"column:tt_content;size:255" binding:"max=255"`       // 目标内容
	TtType    string         `gorm:"column:tt_type;size:45" binding:"max=45"`            // 目标类型 enum:y:年,m:月,d:日#
	TtNum     float64        `gorm:"column:tt_num"`                                      // 目标数量
	TtUnit    string         `gorm:"column:tt_unit;size:45" binding:"max=45"`            // 目标单位
	Begin     *time.Time     `gorm:"column:begin"`                                       // 开始时间
	End       *time.Time     `gorm:"column:end"`                                         // 结束时间
	Status    string         `gorm:"column:status;size:45" binding:"max=45"`             // 状态 enum:s:完成,n:未完成,r:进行中#
	Num       float64        `gorm:"column:num"`                                         // 积分
	Rm        string         `gorm:"column:rm;size:255" binding:"max=255"`               // 备注
	GenTask   string         `gorm:"column:gen_task;size:10;default:y" binding:"max=10"` // 是否生成任务 enum:y:生成,n:不生成#
	Finish    float64        `gorm:"column:finish"`                                      // 完成数量
	User      *User          // 一对多,从表
	GrTasks   []*GrTask      // 一对多,主表
}

// pos1

func (m *GrTarget) TableName() string {
	return "gr_target"
}

func (m *GrTarget) GetM2MJoin(m2m string) string {
	joinStr := ""
	switch m2m {

	}
	return joinStr
}

func (m *GrTarget) Get(db *gorm.DB, q *query.Query) ([]*GrTarget, int64, error) {
	var res []*GrTarget
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

func (m *GrTarget) GetOne(db *gorm.DB, q *query.Query) error {
	return db.Scopes(q.DBQuery()).First(m).Error
}

func (m *GrTarget) Create(db *gorm.DB) error {
	return db.Create(m).Error
}

func (m *GrTarget) Update(db *gorm.DB, v map[string]interface{}) error {
	var fieldList []string
	for k := range v {
		fieldList = append(fieldList, k)
	}
	return db.Model(m).Select(fieldList).Updates(m).Error
}

func (m *GrTarget) UpdateByQuery(db *gorm.DB, q *query.Query, v map[string]interface{}) error {
	return db.Model(m).Scopes(q.DBQuery()).Updates(v).Error
}

func (m *GrTarget) UpdateFull(db *gorm.DB, v map[string]interface{}) error {
	return db.Model(m).Session(&gorm.Session{FullSaveAssociations: true}).Updates(v).Error
}

func (m *GrTarget) Updates(db *gorm.DB, v []*GrTarget) error {
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
func (m *GrTarget) UpdateM2M(db *gorm.DB, field string, addIds, delIds, repIds []int) error {
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

func (m *GrTarget) Delete(db *gorm.DB) error {
	return db.Where("id = ?", m.Id).Delete(m).Error
}

func (m *GrTarget) Deletes(db *gorm.DB, ids []int) error {
	return db.Delete([]GrTarget{}, ids).Error
}
