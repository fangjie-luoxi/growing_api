package service

import (
	"strconv"

	"github.com/fangjie-luoxi/tools/query"
	"github.com/fangjie-luoxi/tools/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/fangjie-luoxi/growing_api/system/system/model"
)

// RoleMenus 更新角色菜单关系
func (s Service) RoleMenus(c *gin.Context) {
	param := struct {
		RoleId  int `binding:"required"`
		MenuIds []int
	}{}
	err := c.ShouldBindJSON(&param)
	if err != nil || param.RoleId == 0 {
		response.Error(c, 400, err.Error())
		return
	}
	var menus []model.SysMenu
	for _, id := range param.MenuIds {
		menus = append(menus, model.SysMenu{Id: id})
	}

	m := model.SysRole{Id: param.RoleId}
	err = s.engine.Model(&m).Omit("SysMenus.*").Association("SysMenus").Replace(menus)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}
	response.Success(c, 200, "OK")
}

// RoleApis 更新角色api关系
func (s Service) RoleApis(c *gin.Context) {
	param := struct {
		RoleId int `binding:"required"`
		ApiIds []int
	}{}
	err := c.ShouldBindJSON(&param)
	if err != nil || param.RoleId == 0 {
		response.Error(c, 400, err.Error())
		return
	}
	roleId := "role::" + strconv.Itoa(param.RoleId)
	_, err = s.casbinEnforcer.RemovePolicy(roleId)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}
	_ = s.casbinEnforcer.LoadPolicy()
	if len(param.ApiIds) <= 0 {
		response.Success(c, 200, "OK")
		return
	}
	var apis []model.SysApi
	err = s.engine.Find(&apis, param.ApiIds).Error
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}
	var rules [][]string
	for _, api := range apis {
		rules = append(rules, []string{roleId, api.Path, api.Method, "api", strconv.Itoa(api.Id)})
	}
	_, err = s.casbinEnforcer.AddPolicies(rules)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}
	response.Success(c, 200, "OK")
}

// GetRoleApis 获取角色对应的api
func (s Service) GetRoleApis(c *gin.Context) {
	idStr := c.Param("id")
	rules := s.casbinEnforcer.GetPermissionsForUser("role::" + idStr)
	var apiIds []int
	for _, rule := range rules {
		if len(rule) < 5 {
			continue
		}
		apiId, err := strconv.Atoi(rule[4])
		if err != nil {
			continue
		}
		apiIds = append(apiIds, apiId)
	}
	response.Success(c, 200, apiIds)
}

// GetSysRoles 获取符合条件的Role
func (s Service) GetSysRoles(c *gin.Context) {
	m := model.SysRole{}
	q, err := query.NewQuery(c, &m)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	var res []*model.SysRole
	tx := s.engine.Model(m)
	tx.Scopes(q.DBQuery())
	var count int64
	tx.Count(&count)
	tx.Scopes(q.DBSelect())
	tx.Limit(q.Limit).Offset(q.Offset).Find(&res)
	response.PageOK(c, res, count, q.Offset, q.Limit)
}

// GetSysRoleById 获取单个
func (s Service) GetSysRoleById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, 400, "请求id错误")
		return
	}
	m := model.SysRole{}
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

// CreateSysRole 创建SysRole
func (s Service) CreateSysRole(c *gin.Context) {
	m := model.SysRole{}
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

// UpdateSysRole 更新SysRole
func (s Service) UpdateSysRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, 400, "id不能为空")
		return
	}
	m := model.SysRole{}
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

// DelSysRole 删除
func (s Service) DelSysRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		response.Error(c, 400, "请求id错误")
		return
	}
	m := model.SysRole{}
	m.Id = id

	err = s.engine.Where("id = ?", m.Id).Delete(&m).Error
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	_, _ = s.casbinEnforcer.DeleteRole("role::" + idStr)

	response.Success(c, 200, "ok")
}

// DelSysRoles 批量删除
func (s Service) DelSysRoles(c *gin.Context) {
	param := struct {
		Ids []int
	}{}
	err := c.ShouldBind(&param)
	if err != nil || len(param.Ids) <= 0 {
		response.Error(c, 400, "请求ids错误")
		return
	}
	err = s.engine.Delete([]model.SysRole{}, param.Ids).Error
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	for _, id := range param.Ids {
		_, _ = s.casbinEnforcer.DeleteRole("role::" + strconv.Itoa(id))
	}

	response.Success(c, 200, "ok")
}
