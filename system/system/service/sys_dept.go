package service

import (
	"strconv"

	"github.com/fangjie-luoxi/tools/query"
	"github.com/fangjie-luoxi/tools/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/fangjie-luoxi/growing_api/system/system/model"
)

// GetSysDepts 获取符合条件的Dept
func (s Service) GetSysDepts(c *gin.Context) {
	m := model.SysDept{}
	q, err := query.NewQuery(c, &m)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	var res []*model.SysDept
	tx := s.engine.Model(m)
	tx.Scopes(q.DBQuery())
	var count int64
	tx.Count(&count)
	tx.Scopes(q.DBSelect())
	tx.Limit(q.Limit).Offset(q.Offset).Find(&res)
	tree := c.Query("tree")
	if tree != "" {
		res = s.getDeptTree(res)
	}
	response.PageOK(c, res, count, q.Offset, q.Limit)
}

func (s Service) getDeptTree(dataList []*model.SysDept) (menuChildren []*model.SysDept) {
	treeMap := make(map[int][]*model.SysDept)
	for _, data := range dataList {
		treeMap[data.Pid] = append(treeMap[data.Pid], data)
	}
	baseData := treeMap[0]
	for i := 0; i < len(baseData); i++ {
		s.getDeptChildrenList(baseData[i], treeMap)
	}
	return baseData
}

func (s Service) getDeptChildrenList(data *model.SysDept, treeMap map[int][]*model.SysDept) {
	data.Children = treeMap[data.Id]
	for i := 0; i < len(data.Children); i++ {
		s.getDeptChildrenList(data.Children[i], treeMap)
	}
}

// GetSysDeptById 获取单个
func (s Service) GetSysDeptById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, 400, "请求id错误")
		return
	}
	m := model.SysDept{}
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

// CreateSysDept 创建SysDept
func (s Service) CreateSysDept(c *gin.Context) {
	m := model.SysDept{}
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

// UpdateSysDept 更新SysDept
func (s Service) UpdateSysDept(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, 400, "id不能为空")
		return
	}
	m := model.SysDept{}
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

// DelSysDept 删除
func (s Service) DelSysDept(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		response.Error(c, 400, "请求id错误")
		return
	}
	m := model.SysDept{}
	m.Id = id

	err = s.engine.Where("id = ?", m.Id).Delete(&m).Error
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, "ok")
}

// DelSysDepts 批量删除
func (s Service) DelSysDepts(c *gin.Context) {
	param := struct {
		Ids []int
	}{}
	err := c.ShouldBind(&param)
	if err != nil || len(param.Ids) <= 0 {
		response.Error(c, 400, "请求ids错误")
		return
	}
	err = s.engine.Delete([]model.SysDept{}, param.Ids).Error
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, "ok")
}
