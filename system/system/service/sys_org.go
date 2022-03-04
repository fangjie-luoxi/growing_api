package service

import (
	"github.com/fangjie-luoxi/tools/response"
	"strconv"

	"github.com/fangjie-luoxi/tools/query"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/fangjie-luoxi/growing_api/system/system/model"
)

// GetSysOrgs 获取符合条件的Org
func (s Service) GetSysOrgs(c *gin.Context) {
	m := model.SysOrg{}
	q, err := query.NewQuery(c, &m)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	var res []*model.SysOrg
	tx := s.engine.Model(m)
	tx.Scopes(q.DBQuery())
	var count int64
	tx.Count(&count)
	tx.Scopes(q.DBSelect())
	tx.Limit(q.Limit).Offset(q.Offset).Find(&res)
	tree := c.Query("tree")
	if tree != "" {
		res = s.getOrgTree(res)
	}
	response.PageOK(c, res, count, q.Offset, q.Limit)
}

func (s Service) getOrgTree(orgs []*model.SysOrg) (menuChildren []*model.SysOrg) {
	treeMap := make(map[int][]*model.SysOrg)
	for _, org := range orgs {
		treeMap[org.Pid] = append(treeMap[org.Pid], org)
	}
	baseData := treeMap[0]
	for i := 0; i < len(baseData); i++ {
		s.getOrgChildrenList(baseData[i], treeMap)
	}
	return baseData
}

func (s Service) getOrgChildrenList(org *model.SysOrg, treeMap map[int][]*model.SysOrg) {
	org.Children = treeMap[org.Id]
	for i := 0; i < len(org.Children); i++ {
		s.getOrgChildrenList(org.Children[i], treeMap)
	}
}

// GetSysOrgById 获取单个
func (s Service) GetSysOrgById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, 400, "请求id错误")
		return
	}
	m := model.SysOrg{}
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

// CreateSysOrg 创建SysOrg
func (s Service) CreateSysOrg(c *gin.Context) {
	m := model.SysOrg{}
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

// UpdateSysOrg 更新SysOrg
func (s Service) UpdateSysOrg(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, 400, "id不能为空")
		return
	}
	m := model.SysOrg{}
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

// DelSysOrg 删除
func (s Service) DelSysOrg(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		response.Error(c, 400, "请求id错误")
		return
	}
	m := model.SysOrg{}
	m.Id = id

	err = s.engine.Where("id = ?", m.Id).Delete(&m).Error
	if err != nil {

		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, "ok")
}

// DelSysOrgs 批量删除
func (s Service) DelSysOrgs(c *gin.Context) {
	param := struct {
		Ids []int
	}{}
	err := c.ShouldBind(&param)
	if err != nil || len(param.Ids) <= 0 {
		response.Error(c, 400, "请求ids错误")
		return
	}
	err = s.engine.Delete([]model.SysOrg{}, param.Ids).Error
	if err != nil {

		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, "ok")
}
