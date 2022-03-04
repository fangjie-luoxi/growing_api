package service

import (
	"strconv"

	"github.com/fangjie-luoxi/tools/query"
	"github.com/fangjie-luoxi/tools/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/fangjie-luoxi/growing_api/system/system/model"
)

// GetSysGroups 获取符合条件的Api
func (s Service) GetSysGroups(c *gin.Context) {
	m := model.SysGroup{}
	q, err := query.NewQuery(c, &m)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	var res []*model.SysGroup
	tx := s.engine.Model(m)
	tx.Scopes(q.DBQuery())
	var count int64
	tx.Count(&count)
	tx.Scopes(q.DBSelect())
	tx.Limit(q.Limit).Offset(q.Offset).Find(&res)
	response.PageOK(c, res, count, q.Offset, q.Limit)
}

// GetSysGroupById 获取单个
func (s Service) GetSysGroupById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, 400, "请求id错误")
		return
	}
	m := model.SysGroup{}
	if id != 0 {
		m.Id = id
	}
	q, err := query.NewQuery(c, &m)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	err = s.engine.Scopes(q.DBQuery()).First(&m).Error
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, m)
}

// CreateSysGroup 创建SysGroup
func (s Service) CreateSysGroup(c *gin.Context) {
	m := model.SysGroup{}
	err := c.ShouldBindJSON(&m)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	err = s.engine.Create(&m).Error
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.Success(c, 201, m)
}

// UpdateSysGroup 更新SysGroup
func (s Service) UpdateSysGroup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, 400, "id不能为空")
		return
	}
	m := model.SysGroup{}
	m.Id = id
	err = c.ShouldBindBodyWith(&m, binding.JSON)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	var v map[string]interface{}
	err = c.ShouldBindBodyWith(&v, binding.JSON)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	var fieldList []string
	for k := range v {
		fieldList = append(fieldList, k)
	}
	err = s.engine.Model(m).Select(fieldList).Updates(&m).Error
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, m)
}

// M2MSysGroup 修改多对多关系
func (s Service) M2MSysGroup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, 400, "请求id为空")
		return
	}

	m2mField := c.Query("m2m_field")
	if m2mField == "" {
		response.Error(c, 400, "m2m_field不能为空")
		return
	}
	ids := struct {
		Add []int
		Del []int
		Rep []int
	}{}
	err = c.ShouldBindJSON(&ids)
	if err != nil {
		response.Error(c, 400, "参数错误")
		return
	}
	m := model.SysGroup{Id: id}
	// pos61

	err = m.UpdateM2M(s.engine, m2mField, ids.Add, ids.Del, ids.Rep)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}
	// pos62

	response.Success(c, 200, "ok")
}

// DelSysGroup 删除
func (s Service) DelSysGroup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		response.Error(c, 400, "请求id错误")
		return
	}
	m := model.SysGroup{}
	m.Id = id

	err = s.engine.Where("id = ?", m.Id).Delete(&m).Error
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, "ok")
}

// DelSysGroups 批量删除
func (s Service) DelSysGroups(c *gin.Context) {
	param := struct {
		Ids []int
	}{}
	err := c.ShouldBind(&param)
	if err != nil || len(param.Ids) <= 0 {
		response.Error(c, 400, "请求ids错误")
		return
	}
	err = s.engine.Delete([]model.SysGroup{}, param.Ids).Error
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, "ok")
}
