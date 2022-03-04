package service

import (
	"strconv"

	"github.com/fangjie-luoxi/tools/query"
	"github.com/fangjie-luoxi/tools/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/json-iterator/go"

	"github.com/fangjie-luoxi/growing_api/system/system/model"
)

// GetCurrentSysConf 获取当前系统配置
func (s Service) GetCurrentSysConf(c *gin.Context) {
	var res []*model.SysConf
	err := s.engine.Limit(10).Find(&res).Error
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}
	if len(res) > 0 {
		tp := c.Query("layout")
		if tp != "" {
			bytes, err := jsoniter.Marshal(res[0])
			if err != nil {
				response.Error(c, 500, err.Error())
				return
			}
			var conf model.SysConfLayout
			err = jsoniter.Unmarshal(bytes, &conf)
			if err != nil {
				response.Error(c, 500, err.Error())
				return
			}
			response.Success(c, 200, conf)
			return
		}
		response.Success(c, 200, res[0])
		return
	} else {
		newConf := model.SysConf{
			Version:     "v1.0.0",
			NavTheme:    "light",
			HeaderTheme: "dark",
			Layout:      "mix",
		}
		s.engine.Create(&newConf)
		response.Success(c, 200, newConf)
	}
}

// GetSysConfs 获取符合条件的Api
func (s Service) GetSysConfs(c *gin.Context) {
	m := model.SysConf{}
	q, err := query.NewQuery(c, &m)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	var res []*model.SysConf
	tx := s.engine.Model(m)
	tx.Scopes(q.DBQuery())
	var count int64
	tx.Count(&count)
	tx.Scopes(q.DBSelect())
	tx.Limit(q.Limit).Offset(q.Offset).Find(&res)
	response.PageOK(c, res, count, q.Offset, q.Limit)
}

// GetSysConfById 获取单个
func (s Service) GetSysConfById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, 400, "请求id错误")
		return
	}
	m := model.SysConf{}
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

// CreateSysConf 创建SysConf
func (s Service) CreateSysConf(c *gin.Context) {
	m := model.SysConf{}
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

// UpdateSysConf 更新SysConf
func (s Service) UpdateSysConf(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, 400, "id不能为空")
		return
	}
	m := model.SysConf{}
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

// DelSysConf 删除
func (s Service) DelSysConf(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		response.Error(c, 400, "请求id错误")
		return
	}
	m := model.SysConf{}
	m.Id = id

	err = s.engine.Where("id = ?", m.Id).Delete(&m).Error
	if err != nil {

		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, "ok")
}

// DelSysConfs 批量删除
func (s Service) DelSysConfs(c *gin.Context) {
	param := struct {
		Ids []int
	}{}
	err := c.ShouldBind(&param)
	if err != nil || len(param.Ids) <= 0 {
		response.Error(c, 400, "请求ids错误")
		return
	}
	err = s.engine.Delete([]model.SysConf{}, param.Ids).Error
	if err != nil {

		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, "ok")
}
