package service

import (
	"github.com/gin-gonic/gin"
)

// SysApi 系统管理api
func (s *Service) SysApi(r *gin.RouterGroup) {
	sys := r.Group("system")
	// api管理
	sys.GET("/sys_api", s.GetSysApis)
	sys.GET("/sys_api/:id", s.GetSysApiById)
	sys.POST("/sys_api", s.CreateSysApi)
	sys.PUT("/sys_api/:id", s.UpdateSysApi)
	sys.DELETE("/sys_api/:id", s.DelSysApi)
	sys.POST("/sys_api/del", s.DelSysApis)
	// 角色管理
	sys.GET("/sys_role", s.GetSysRoles)
	sys.GET("/sys_role/:id", s.GetSysRoleById)
	sys.POST("/sys_role", s.CreateSysRole)
	sys.PUT("/sys_role/:id", s.UpdateSysRole)
	sys.DELETE("/sys_role/:id", s.DelSysRole)
	sys.POST("/sys_role/del", s.DelSysRoles)
	sys.POST("/sys_role/menus", s.RoleMenus)
	sys.POST("/sys_role/apis", s.RoleApis)
	sys.GET("/sys_role/apis/:id", s.GetRoleApis)
	// 用户管理
	sys.GET("/user", s.GetUsers)
	sys.GET("/user/:id", s.GetUserById)
	sys.POST("/user", s.CreateUser)
	sys.PUT("/user/:id", s.UpdateUser)
	sys.DELETE("/user/:id", s.DelUser)
	sys.POST("/user/del", s.DelUsers)
	sys.PATCH("/user/m2m/part/:id", s.M2MUser)
	// 菜单管理
	sys.GET("/sys_menu/user", s.GetUserMenus)
	sys.GET("/sys_menu/:id", s.GetSysMenuById)
	sys.GET("/sys_menu", s.GetSysMenus)
	sys.POST("/sys_menu", s.CreateSysMenu)
	sys.PUT("/sys_menu/:id", s.UpdateSysMenu)
	sys.DELETE("/sys_menu/:id", s.DelSysMenu)
	sys.POST("/sys_menu/del", s.DelSysMenus)
	// 机构管理
	sys.GET("/sys_org", s.GetSysOrgs)
	sys.GET("/sys_org/:id", s.GetSysOrgById)
	sys.POST("/sys_org", s.CreateSysOrg)
	sys.PUT("/sys_org/:id", s.UpdateSysOrg)
	sys.DELETE("/sys_org/:id", s.DelSysOrg)
	sys.POST("/sys_org/del", s.DelSysOrgs)
	// 机构设置管理
	sys.GET("/sys_org_conf", s.GetSysOrgConfs)
	sys.GET("/sys_org_conf/:id", s.GetSysOrgConfById)
	sys.POST("/sys_org_conf", s.CreateSysOrgConf)
	sys.PUT("/sys_org_conf/:id", s.UpdateSysOrgConf)
	sys.DELETE("/sys_org_conf/:id", s.DelSysOrgConf)
	sys.POST("/sys_org_conf/del", s.DelSysOrgConfs)
	// 部门管理
	sys.GET("/sys_dept", s.GetSysDepts)
	sys.GET("/sys_dept/:id", s.GetSysDeptById)
	sys.POST("/sys_dept", s.CreateSysDept)
	sys.PUT("/sys_dept/:id", s.UpdateSysDept)
	sys.DELETE("/sys_dept/:id", s.DelSysDept)
	sys.POST("/sys_dept/del", s.DelSysDepts)
	// 用户组管理
	sys.GET("/sys_group", s.GetSysGroups)
	sys.GET("/sys_group/:id", s.GetSysGroupById)
	sys.POST("/sys_group", s.CreateSysGroup)
	sys.PUT("/sys_group/:id", s.UpdateSysGroup)
	sys.DELETE("/sys_group/:id", s.DelSysGroup)
	sys.POST("/sys_group/del", s.DelSysGroups)
	sys.PATCH("/sys_group/m2m/part/:id", s.M2MSysGroup)
	// 岗位管理
	sys.GET("/sys_post", s.GetSysPosts)
	sys.GET("/sys_post/:id", s.GetSysPostById)
	sys.POST("/sys_post", s.CreateSysPost)
	sys.PUT("/sys_post/:id", s.UpdateSysPost)
	sys.DELETE("/sys_post/:id", s.DelSysPost)
	sys.POST("/sys_post/del", s.DelSysPosts)
	// 系统设置
	sys.GET("/sys_conf", s.GetSysConfs)
	sys.GET("/sys_conf/current", s.GetCurrentSysConf) // 获取当前系统配置
	sys.GET("/sys_conf/:id", s.GetSysConfById)
	sys.POST("/sys_conf", s.CreateSysConf)
	sys.PUT("/sys_conf/:id", s.UpdateSysConf)
	sys.DELETE("/sys_conf/:id", s.DelSysConf)
	sys.POST("/sys_conf/del", s.DelSysConfs)
	// 系统字典管理
	sys.GET("/sys_dict", s.GetSysDicts)
	sys.GET("/sys_dict/:id", s.GetSysDictById)
	sys.POST("/sys_dict", s.CreateSysDict)
	sys.PUT("/sys_dict/:id", s.UpdateSysDict)
	sys.DELETE("/sys_dict/:id", s.DelSysDict)
	sys.POST("/sys_dict/del", s.DelSysDicts)
}
