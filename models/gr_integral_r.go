package models

import (
	"time"

	"github.com/fangjie-luoxi/tools/query"
	"gorm.io/gorm"
)

// GrIntegralR 积分记录
type GrIntegralR struct {
	Id        int            `gorm:"column:id;primaryKey"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
	UserId    int            `gorm:"column:user_id"`
	InType    string         `gorm:"column:in_type;size:10;default:i" binding:"max=10"` // 类型 enum:i:增加,o:扣除#
	Num       int            `gorm:"column:num"`                                        // 增加或扣除的积分
	ReTpye    string         `gorm:"column:re_tpye;size:10" binding:"max=10"`           // 关联(积分变动所关联的项) enum:tt:目标,tk:任务,re:规则#
	ReId      int            `gorm:"column:re_id"`                                      // 关联所对应的id
	Desc      string         `gorm:"column:desc;size:255" binding:"max=255"`            // 描述
	User      *User          // 一对多,从表
}

// pos1

func (m *GrIntegralR) TableName() string {
	return "gr_integral_r"
}

func (m *GrIntegralR) GetM2MJoin(m2m string) string {
	joinStr := ""
	switch m2m {

	}
	return joinStr
}

func (m *GrIntegralR) Get(db *gorm.DB, q *query.Query) ([]*GrIntegralR, int64, error) {
	var res []*GrIntegralR
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

func (m *GrIntegralR) GetOne(db *gorm.DB, q *query.Query) error {
	return db.Scopes(q.DBQuery()).First(m).Error
}

func (m *GrIntegralR) Create(db *gorm.DB) error {
	return db.Create(m).Error
}

func (m *GrIntegralR) Update(db *gorm.DB, v map[string]interface{}) error {
	var fieldList []string
	for k := range v {
		fieldList = append(fieldList, k)
	}
	return db.Model(m).Select(fieldList).Updates(m).Error
}

func (m *GrIntegralR) UpdateByQuery(db *gorm.DB, q *query.Query, v map[string]interface{}) error {
	return db.Model(m).Scopes(q.DBQuery()).Updates(v).Error
}

func (m *GrIntegralR) UpdateFull(db *gorm.DB, v map[string]interface{}) error {
	return db.Model(m).Session(&gorm.Session{FullSaveAssociations: true}).Updates(v).Error
}

func (m *GrIntegralR) Updates(db *gorm.DB, v []*GrIntegralR) error {
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
func (m *GrIntegralR) UpdateM2M(db *gorm.DB, field string, addIds, delIds, repIds []int) error {
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

func (m *GrIntegralR) Delete(db *gorm.DB) error {
	return db.Where("id = ?", m.Id).Delete(m).Error
}

func (m *GrIntegralR) Deletes(db *gorm.DB, ids []int) error {
	return db.Delete([]GrIntegralR{}, ids).Error
}
