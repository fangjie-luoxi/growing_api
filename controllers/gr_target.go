package controllers

import (
	"strconv"

	"github.com/fangjie-luoxi/tools/query"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/fangjie-luoxi/growing_api/global"
	"github.com/fangjie-luoxi/growing_api/models"
)

// GetGrTargets 获取符合条件的GrTarget
func GetGrTargets(c *gin.Context) {
	m := models.GrTarget{}
	q, err := query.NewQuery(c, &m)
	if err != nil {
		global.Resp.Error(c, 400, err, "")
		return
	}
	// pos11

	v, count, err := m.Get(global.DBEngine, q)
	if err != nil {
		global.Log.Error(err)
		global.Resp.Error(c, 400, err, "")
		return
	}
	if q.Count == "count" {
		c.JSON(200, count)
		return
	}
	var res interface{}
	res = v
	if len(q.Select) > 0 && len(v) > 0 {
		res = q.GetSelectMap(v, c.Query("fields"))
	}
	// pos12

	if q.Count == "F" {
		global.Resp.OK(c, 200, res)
		return
	}
	global.Resp.PageOK(c, res, count, q.Offset, q.Limit)
}

// GetGrTargetById 获取单个GrTarget
func GetGrTargetById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		global.Resp.Error(c, 400, err, "请求id错误")
		return
	}
	m := models.GrTarget{}
	if id != 0 {
		m.Id = id
	}
	q, err := query.NewQuery(c, &m)
	if err != nil {
		global.Resp.Error(c, 400, err, "")
		return
	}
	// pos21

	err = m.GetOne(global.DBEngine, q)
	if err != nil {
		global.Log.Error(err)
		global.Resp.Error(c, 500, err, "")
		return
	}
	// pos22

	global.Resp.OK(c, 200, m)
}

// CreateGrTarget 创建GrTarget
func CreateGrTarget(c *gin.Context) {
	m := models.GrTarget{}
	err := c.ShouldBindJSON(&m)
	if err != nil {
		global.Resp.Error(c, 400, err, "")
		return
	}
	// pos31

	err = m.Create(global.DBEngine)
	if err != nil {
		global.Resp.Error(c, 400, err, "")
		return
	}
	// pos32

	global.Resp.OK(c, 201, m)
}

// UpdateGrTarget 更新GrTarget
func UpdateGrTarget(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		global.Resp.Error(c, 400, err, "id不能为空")
		return
	}
	m := models.GrTarget{}
	m.Id = id
	err = c.ShouldBindBodyWith(&m, binding.JSON)
	if err != nil {
		global.Resp.Error(c, 400, err, "")
		return
	}

	var v map[string]interface{}
	err = c.ShouldBindBodyWith(&v, binding.JSON)
	if err != nil {
		global.Resp.Error(c, 400, err, "")
		return
	}
	// pos41

	err = m.Update(global.DBEngine, v)
	if err != nil {
		global.Log.Error(err)
		global.Resp.Error(c, 500, err, "")
		return
	}
	// pos42

	global.Resp.OK(c, 200, m)
}

// UpdateFullGrTarget 更新关联字段
func UpdateFullGrTarget(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		global.Resp.Error(c, 400, err, "id不能为空")
		return
	}
	m := models.GrTarget{}
	m.Id = id
	err = c.ShouldBindBodyWith(&m, binding.JSON)
	if err != nil {
		global.Resp.Error(c, 400, err, "")
		return
	}
	var v map[string]interface{}
	err = c.ShouldBindBodyWith(&v, binding.JSON)
	if err != nil {
		global.Resp.Error(c, 400, err, "")
		return
	}
	// pos43

	err = m.UpdateFull(global.DBEngine, v)
	if err != nil {
		global.Log.Error(err)
		global.Resp.Error(c, 500, err, "")
		return
	}
	// pos44

	global.Resp.OK(c, 200, m)
}

// UpdateGrTargets 更新多个
func UpdateGrTargets(c *gin.Context) {
	m := models.GrTarget{}
	var v []*models.GrTarget
	err := c.ShouldBindJSON(&v)
	if err != nil {
		global.Resp.Error(c, 400, err, "请求参数错误")
		return
	}
	// pos51

	err = m.Updates(global.DBEngine, v)
	if err != nil {
		global.Log.Error(err)
		global.Resp.Error(c, 500, err, "")
		return
	}
	// pos52

	global.Resp.OK(c, 200, "ok")
}

// UpdateGrTargetByQuery 通过查询条件更新多个
func UpdateGrTargetByQuery(c *gin.Context) {
	var v map[string]interface{}
	err := c.ShouldBindBodyWith(&v, binding.JSON)
	if err != nil {
		global.Resp.Error(c, 400, err, "")
		return
	}
	m := models.GrTarget{}
	q, err := query.NewQuery(c, &m)
	if err != nil {
		global.Resp.Error(c, 400, err, "")
		return
	}
	err = m.UpdateByQuery(global.DBEngine, q, v)
	if err != nil {
		global.Log.Error(err)
		global.Resp.Error(c, 500, err, "")
		return
	}
	global.Resp.OK(c, 200, "ok")
}

// M2MGrTarget 修改多对多关系
func M2MGrTarget(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		global.Resp.Error(c, 400, err, "请求id为空")
		return
	}

	m2mField := c.Query("m2m_field")
	if m2mField == "" {
		global.Resp.Error(c, 400, err, "m2m_field不能为空")
		return
	}
	ids := struct {
		Add []int
		Del []int
		Rep []int
	}{}
	err = c.ShouldBindJSON(&ids)
	if err != nil {
		global.Resp.Error(c, 400, err, "参数错误")
		return
	}
	m := models.GrTarget{Id: id}
	// pos61

	err = m.UpdateM2M(global.DBEngine, m2mField, ids.Add, ids.Del, ids.Rep)
	if err != nil {
		global.Log.Error(err)
		global.Resp.Error(c, 500, err, "")
		return
	}
	// pos62

	global.Resp.OK(c, 200, "ok")
}

// DelGrTarget 删除
func DelGrTarget(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		global.Resp.Error(c, 400, err, "请求id错误")
		return
	}
	m := models.GrTarget{}
	m.Id = id
	// pos71

	err = m.Delete(global.DBEngine)
	if err != nil {
		global.Log.Error(err)
		global.Resp.Error(c, 500, err, "")
		return
	}
	// pos72

	global.Resp.OK(c, 200, "ok")
}

// DelGrTargets 批量删除
func DelGrTargets(c *gin.Context) {
	param := struct {
		Ids []int
	}{}
	err := c.ShouldBind(&param)
	if err != nil || len(param.Ids) <= 0 {
		global.Resp.Error(c, 400, err, "请求ids错误")
		return
	}

	m := models.GrTarget{}
	// pos81

	err = m.Deletes(global.DBEngine, param.Ids)
	if err != nil {
		global.Log.Error(err)
		global.Resp.Error(c, 500, err, "")
		return
	}
	// pos82

	global.Resp.OK(c, 200, "ok")
}
