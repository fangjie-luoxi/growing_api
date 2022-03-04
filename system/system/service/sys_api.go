package service

import (
	"strconv"

	"github.com/fangjie-luoxi/tools/query"
	"github.com/fangjie-luoxi/tools/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/fangjie-luoxi/growing_api/system/system/model"
)

// GetSysApis 获取符合条件的Api
func (s Service) GetSysApis(c *gin.Context) {
	m := model.SysApi{}
	q, err := query.NewQuery(c, &m)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	var res []*model.SysApi
	tx := s.engine.Model(m)
	tx.Scopes(q.DBQuery())
	var count int64
	tx.Count(&count)
	tx.Scopes(q.DBSelect())
	tx.Limit(q.Limit).Offset(q.Offset).Find(&res)
	response.PageOK(c, res, count, q.Offset, q.Limit)
}

// GetSysApiById 获取单个User
func (s Service) GetSysApiById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	m := model.SysApi{}
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

// CreateSysApi 创建SysApi
func (s Service) CreateSysApi(c *gin.Context) {
	m := model.SysApi{}
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

	go s.cache.Delete("ApiCacheJwt")
	go s.cache.Delete("ApiCachePerm:" + strconv.Itoa(m.SysOrgId))
	response.Success(c, 201, m)
}

// UpdateSysApi 更新SysApi
func (s Service) UpdateSysApi(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	m := model.SysApi{}
	m.Id = id
	err = c.ShouldBindBodyWith(&m, binding.JSON)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	var oldApi model.SysApi
	oldErr := s.engine.First(&oldApi, id).Error

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

	if oldErr == nil {
		oldPolicys := s.casbinEnforcer.GetFilteredPolicy(1, oldApi.Path, oldApi.Method, "", "", "")
		var newPolicys [][]string
		for _, policy := range oldPolicys {
			var newPolicy []string
			newPolicy = append(newPolicy, policy...)
			newPolicy[1] = m.Path
			newPolicy[2] = m.Method
			newPolicys = append(newPolicys, newPolicy)
		}
		_, _ = s.casbinEnforcer.UpdatePolicies(oldPolicys, newPolicys)
	}

	go s.cache.Delete("ApiCacheJwt")
	go s.cache.Delete("ApiCachePerm:" + strconv.Itoa(m.SysOrgId))
	response.Success(c, 200, m)
}

// DelSysApi 删除
func (s Service) DelSysApi(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		response.Error(c, 400, err.Error())
		return
	}
	m := model.SysApi{}
	m.Id = id
	var oldApi model.SysApi
	oldErr := s.engine.First(&oldApi, id).Error

	err = s.engine.Where("id = ?", m.Id).Delete(&m).Error
	if err != nil {

		response.Error(c, 500, err.Error())
		return
	}

	if oldErr == nil {
		_, _ = s.casbinEnforcer.RemoveFilteredPolicy(1, oldApi.Path, oldApi.Method, "", "", "")
	}

	go s.cache.Delete("ApiCacheJwt")
	go s.cache.Delete("ApiCachePerm:" + strconv.Itoa(m.SysOrgId))
	response.Success(c, 200, "ok")
}

// DelSysApis 批量删除
func (s Service) DelSysApis(c *gin.Context) {
	param := struct {
		Ids []int
	}{}
	err := c.ShouldBind(&param)
	if err != nil || len(param.Ids) <= 0 {
		response.Error(c, 400, "请求ids错误")
		return
	}

	var oldApis []*model.SysApi
	oldErr := s.engine.Find(&oldApis, param.Ids).Error

	err = s.engine.Delete([]model.SysApi{}, param.Ids).Error
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}
	if oldErr == nil {
		for _, api := range oldApis {
			_, _ = s.casbinEnforcer.RemoveFilteredPolicy(1, api.Path, api.Method, "", "", "")
		}
		if len(oldApis) > 0 {
			go s.cache.Delete("ApiCacheJwt")
			go s.cache.Delete("ApiCachePerm:" + strconv.Itoa(oldApis[0].SysOrgId))
		}
	}

	response.Success(c, 200, "ok")
}
