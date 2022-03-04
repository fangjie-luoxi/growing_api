package service

import (
	"strconv"

	"github.com/fangjie-luoxi/tools/jwt"
	"github.com/fangjie-luoxi/tools/query"
	"github.com/fangjie-luoxi/tools/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/fangjie-luoxi/growing_api/system/system/model"
)

// GetUserMenus 获取用户菜单
func (s Service) GetUserMenus(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uid, ok := claims["UserId"].(float64)
	if !ok || uid == 0.0 {
		response.Error(c, 400, "获取用户信息失败")
		return
	}
	user := model.User{Id: int(uid)}
	err := s.engine.Preload("SysRoles").First(&user).Error
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	var menus []*model.SysMenu
	if len(user.SysRoles) <= 0 {
		c.JSON(200, []model.SysMenu{})
		return
	}
	var roleIds []int
	for _, role := range user.SysRoles {
		roleIds = append(roleIds, role.Id)
	}
	err = s.engine.Model(&user.SysRoles).Order("sort").Association("SysMenus").Find(&menus)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}
	res := s.getMenuTree(menus)
	c.JSON(200, res)
	return
}

// GetSysMenus 获取符合条件的菜单
func (s Service) GetSysMenus(c *gin.Context) {
	m := model.SysMenu{}
	q, err := query.NewQuery(c, &m)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	var res []*model.SysMenu
	tx := s.engine.Model(m)
	tx.Scopes(q.DBQuery())
	tx.Scopes(q.DBSelect())
	tx.Limit(q.Limit).Offset(q.Offset).Find(&res)

	tree := c.Query("tree")
	if tree != "" {
		res = s.getMenuTree(res)
	}
	response.PageOK(c, res, int64(len(res)), q.Offset, q.Limit)
}

// GetSysMenuById 获取单个
func (s Service) GetSysMenuById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, 400, "请求id错误")
		return
	}
	m := model.SysMenu{}
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

func (s Service) getMenuTree(menus []*model.SysMenu) (menuChildren []*model.SysMenu) {
	treeMap := make(map[int][]*model.SysMenu)
	for _, menu := range menus {
		treeMap[menu.Pid] = append(treeMap[menu.Pid], menu)
	}
	baseMenus := treeMap[0]
	for i := 0; i < len(baseMenus); i++ {
		s.getMenuChildrenList(baseMenus[i], treeMap)
	}
	return baseMenus
}

func (s Service) getMenuChildrenList(menu *model.SysMenu, treeMap map[int][]*model.SysMenu) {
	menu.Children = treeMap[menu.Id]
	for i := 0; i < len(menu.Children); i++ {
		s.getMenuChildrenList(menu.Children[i], treeMap)
	}
}

// CreateSysMenu 创建SysMenu
func (s Service) CreateSysMenu(c *gin.Context) {
	m := model.SysMenu{}
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

// UpdateSysMenu 更新SysMenu
func (s Service) UpdateSysMenu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, 400, "id不能为空")
		return
	}
	m := model.SysMenu{}
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

// DelSysMenu 删除
func (s Service) DelSysMenu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		response.Error(c, 400, "请求id错误")
		return
	}
	m := model.SysMenu{}
	m.Id = id

	err = s.engine.Where("id = ?", m.Id).Delete(&m).Error
	if err != nil {

		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, "ok")
}

// DelSysMenus 批量删除
func (s Service) DelSysMenus(c *gin.Context) {
	param := struct {
		Ids []int
	}{}
	err := c.ShouldBind(&param)
	if err != nil || len(param.Ids) <= 0 {
		response.Error(c, 400, "请求ids错误")
		return
	}
	err = s.engine.Delete([]model.SysMenu{}, param.Ids).Error
	if err != nil {

		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, "ok")
}
