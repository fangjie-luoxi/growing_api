package models

import (
	"time"

	"github.com/fangjie-luoxi/tools/query"
	"gorm.io/gorm"
)

// User 用户表
type User struct {
	Id           int            `gorm:"column:id;primaryKey"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-"`
	Integral     float64        `gorm:"column:integral"`
	GrIntegralRs []*GrIntegralR // 一对多,主表
	GrRules      []*GrRule      // 一对多,主表
	GrTargets    []*GrTarget    // 一对多,主表
	GrTasks      []*GrTask      // 一对多,主表
}

// pos1

func (m *User) TableName() string {
	return "user"
}

func (m *User) GetM2MJoin(m2m string) string {
	joinStr := ""
	switch m2m {

	}
	return joinStr
}

func (m *User) Get(db *gorm.DB, q *query.Query) ([]*User, int64, error) {
	var res []*User
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

func (m *User) GetOne(db *gorm.DB, q *query.Query) error {
	return db.Scopes(q.DBQuery()).First(m).Error
}

func (m *User) Create(db *gorm.DB) error {
	return db.Create(m).Error
}

func (m *User) Update(db *gorm.DB, v map[string]interface{}) error {
	var fieldList []string
	for k := range v {
		fieldList = append(fieldList, k)
	}
	return db.Model(m).Select(fieldList).Updates(m).Error
}

func (m *User) UpdateByQuery(db *gorm.DB, q *query.Query, v map[string]interface{}) error {
	return db.Model(m).Scopes(q.DBQuery()).Updates(v).Error
}

func (m *User) UpdateFull(db *gorm.DB, v map[string]interface{}) error {
	return db.Model(m).Session(&gorm.Session{FullSaveAssociations: true}).Updates(v).Error
}

func (m *User) Updates(db *gorm.DB, v []*User) error {
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
func (m *User) UpdateM2M(db *gorm.DB, field string, addIds, delIds, repIds []int) error {
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

func (m *User) Delete(db *gorm.DB) error {
	return db.Where("id = ?", m.Id).Delete(m).Error
}

func (m *User) Deletes(db *gorm.DB, ids []int) error {
	return db.Delete([]User{}, ids).Error
}
