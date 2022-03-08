package routers

import (
	"github.com/fangjie-luoxi/growing_api/rules/finish"
	"github.com/gin-gonic/gin"
)

func ruleApi(rg *gin.RouterGroup) {
	rule := rg.Group("rule")
	rule.POST("/task/finish/:id", finish.Task)     // 完成任务
	rule.POST("/target/finish/:id", finish.Target) // 完成目标
	rule.POST("/rule/finish/:id", finish.Rule)     // 触发规则
}
