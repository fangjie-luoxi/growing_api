package models

import (
	"time"

	"github.com/fangjie-luoxi/tools/query"
	"gorm.io/gorm"
)

// GrRule 规则
type GrRule struct {
	Id            int             `gorm:"column:id;primaryKey"`
	CreatedAt     time.Time       `gorm:"column:created_at"`
	UpdatedAt     time.Time       `gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt  `json:"-"`
	UserId        int             `gorm:"column:user_id"`
	ReName        string          `gorm:"column:re_name;size:45" binding:"max=45"`           // 规则名称
	Content       string          `gorm:"column:content;size:255" binding:"max=255"`         // 规则内容
	InType        string          `gorm:"column:in_type;size:10;default:i" binding:"max=10"` // 类型 enum:i:增加,o:扣除#
	Num           float64         `gorm:"column:num"`                                        // 积分
	Rm            string          `gorm:"column:rm;size:1000" binding:"max=1000"`            // 备注
	User          *User           // 一对多,从表
	GrRuleRecords []*GrRuleRecord // 一对多,主表
}

// pos1

func (m *GrRule) TableName() string {
	return "gr_rule"
}

func (m *GrRule) GetM2MJoin(m2m string) string {
	joinStr := ""
	switch m2m {

	}
	return joinStr
}

func (m *GrRule) Get(db *gorm.DB, q *query.Query) ([]*GrRule, int64, error) {
	var res []*GrRule
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

func (m *GrRule) GetOne(db *gorm.DB, q *query.Query) error {
	return db.Scopes(q.DBQuery()).First(m).Error
}

func (m *GrRule) Create(db *gorm.DB) error {
	return db.Create(m).Error
}

func (m *GrRule) Update(db *gorm.DB, v map[string]interface{}) error {
	var fieldList []string
	for k := range v {
		fieldList = append(fieldList, k)
	}
	return db.Model(m).Select(fieldList).Updates(m).Error
}

func (m *GrRule) UpdateByQuery(db *gorm.DB, q *query.Query, v map[string]interface{}) error {
	return db.Model(m).Scopes(q.DBQuery()).Updates(v).Error
}

func (m *GrRule) UpdateFull(db *gorm.DB, v map[string]interface{}) error {
	return db.Model(m).Session(&gorm.Session{FullSaveAssociations: true}).Updates(v).Error
}

func (m *GrRule) Updates(db *gorm.DB, v []*GrRule) error {
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
func (m *GrRule) UpdateM2M(db *gorm.DB, field string, addIds, delIds, repIds []int) error {
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

func (m *GrRule) Delete(db *gorm.DB) error {
	return db.Where("id = ?", m.Id).Delete(m).Error
}

func (m *GrRule) Deletes(db *gorm.DB, ids []int) error {
	return db.Delete([]GrRule{}, ids).Error
}
