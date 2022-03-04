package service

import (
	"strconv"

	"github.com/fangjie-luoxi/tools/query"
	"github.com/fangjie-luoxi/tools/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/fangjie-luoxi/growing_api/system/system/model"
)

// GetUsers 获取符合条件的Role
func (s Service) GetUsers(c *gin.Context) {
	m := model.User{}
	q, err := query.NewQuery(c, &m)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	var res []*model.User
	tx := s.engine.Model(m)
	tx.Scopes(q.DBQuery())
	var count int64
	tx.Count(&count)
	tx.Scopes(q.DBSelect())
	tx.Debug().Limit(q.Limit).Offset(q.Offset).Find(&res)
	response.PageOK(c, res, count, q.Offset, q.Limit)
}

// GetUserById 获取单个
func (s Service) GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, 400, "请求id错误")
		return
	}
	m := model.User{}
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

// CreateUser 创建User
func (s Service) CreateUser(c *gin.Context) {
	m := model.User{}
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

// UpdateUser 更新User
func (s Service) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, 400, "id不能为空")
		return
	}
	m := model.User{}
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

// M2MUser 修改多对多关系
func (s Service) M2MUser(c *gin.Context) {
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
	m := model.User{Id: id}
	// pos61

	err = m.UpdateM2M(s.engine, m2mField, ids.Add, ids.Del, ids.Rep)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}
	// pos62
	s.upRole(idStr, m2mField, ids.Add, ids.Del, ids.Rep)

	response.Success(c, 200, "ok")
}

func (s Service) upRole(idStr, field string, addIds, delIds, repIds []int) {
	if field != "SysRoles" {
		return
	}
	if len(addIds) > 0 {
		_, _ = s.casbinEnforcer.AddRolesForUser("user::"+idStr, idToRole(addIds))
	}
	if len(delIds) > 0 {
		_, _ = s.casbinEnforcer.DeleteRolesForUser("user::" + idStr)
	}
	if len(repIds) > 0 {
		_, _ = s.casbinEnforcer.DeleteRolesForUser("user::" + idStr)
		_, _ = s.casbinEnforcer.AddRolesForUser("user::"+idStr, idToRole(repIds))
	}
}

func idToRole(ids []int) []string {
	var roles []string
	for _, id := range ids {
		roles = append(roles, "role::"+strconv.Itoa(id))
	}
	return roles
}

// DelUser 删除
func (s Service) DelUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		response.Error(c, 400, "请求id错误")
		return
	}
	m := model.User{}
	m.Id = id

	err = s.engine.Where("id = ?", m.Id).Delete(&m).Error
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, "ok")
}

// DelUsers 批量删除
func (s Service) DelUsers(c *gin.Context) {
	param := struct {
		Ids []int
	}{}
	err := c.ShouldBind(&param)
	if err != nil || len(param.Ids) <= 0 {
		response.Error(c, 400, "请求ids错误")
		return
	}
	err = s.engine.Delete([]model.User{}, param.Ids).Error
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, "ok")
}
