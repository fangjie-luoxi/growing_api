package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/fangjie-luoxi/growing_api/controllers"
)

// genApi 自动生成的api
func genApi(rg *gin.RouterGroup) {
	{

		// gr_integral_r路由
		rg.GET("/gr_integral_r/:id", controllers.GetGrIntegralRById)
		rg.GET("/gr_integral_r", controllers.GetGrIntegralRs)
		rg.POST("/gr_integral_r", controllers.CreateGrIntegralR)
		rg.PUT("/gr_integral_r/:id", controllers.UpdateGrIntegralR)
		rg.PUT("/gr_integral_r/full/:id", controllers.UpdateFullGrIntegralR)
		rg.PUT("/gr_integral_r", controllers.UpdateGrIntegralRs)
		rg.PUT("/gr_integral_r/query", controllers.UpdateGrIntegralRByQuery)
		rg.PATCH("/gr_integral_r/m2m/part/:id", controllers.M2MGrIntegralR)
		rg.DELETE("/gr_integral_r/:id", controllers.DelGrIntegralR)
		rg.POST("/gr_integral_r/del", controllers.DelGrIntegralRs)
		// gr_rule路由
		rg.GET("/gr_rule/:id", controllers.GetGrRuleById)
		rg.GET("/gr_rule", controllers.GetGrRules)
		rg.POST("/gr_rule", controllers.CreateGrRule)
		rg.PUT("/gr_rule/:id", controllers.UpdateGrRule)
		rg.PUT("/gr_rule/full/:id", controllers.UpdateFullGrRule)
		rg.PUT("/gr_rule", controllers.UpdateGrRules)
		rg.PUT("/gr_rule/query", controllers.UpdateGrRuleByQuery)
		rg.PATCH("/gr_rule/m2m/part/:id", controllers.M2MGrRule)
		rg.DELETE("/gr_rule/:id", controllers.DelGrRule)
		rg.POST("/gr_rule/del", controllers.DelGrRules)
		// gr_rule_record路由
		rg.GET("/gr_rule_record/:id", controllers.GetGrRuleRecordById)
		rg.GET("/gr_rule_record", controllers.GetGrRuleRecords)
		rg.POST("/gr_rule_record", controllers.CreateGrRuleRecord)
		rg.PUT("/gr_rule_record/:id", controllers.UpdateGrRuleRecord)
		rg.PUT("/gr_rule_record/full/:id", controllers.UpdateFullGrRuleRecord)
		rg.PUT("/gr_rule_record", controllers.UpdateGrRuleRecords)
		rg.PUT("/gr_rule_record/query", controllers.UpdateGrRuleRecordByQuery)
		rg.PATCH("/gr_rule_record/m2m/part/:id", controllers.M2MGrRuleRecord)
		rg.DELETE("/gr_rule_record/:id", controllers.DelGrRuleRecord)
		rg.POST("/gr_rule_record/del", controllers.DelGrRuleRecords)
		// gr_target路由
		rg.GET("/gr_target/:id", controllers.GetGrTargetById)
		rg.GET("/gr_target", controllers.GetGrTargets)
		rg.POST("/gr_target", controllers.CreateGrTarget)
		rg.PUT("/gr_target/:id", controllers.UpdateGrTarget)
		rg.PUT("/gr_target/full/:id", controllers.UpdateFullGrTarget)
		rg.PUT("/gr_target", controllers.UpdateGrTargets)
		rg.PUT("/gr_target/query", controllers.UpdateGrTargetByQuery)
		rg.PATCH("/gr_target/m2m/part/:id", controllers.M2MGrTarget)
		rg.DELETE("/gr_target/:id", controllers.DelGrTarget)
		rg.POST("/gr_target/del", controllers.DelGrTargets)
		// gr_task路由
		rg.GET("/gr_task/:id", controllers.GetGrTaskById)
		rg.GET("/gr_task", controllers.GetGrTasks)
		rg.POST("/gr_task", controllers.CreateGrTask)
		rg.PUT("/gr_task/:id", controllers.UpdateGrTask)
		rg.PUT("/gr_task/full/:id", controllers.UpdateFullGrTask)
		rg.PUT("/gr_task", controllers.UpdateGrTasks)
		rg.PUT("/gr_task/query", controllers.UpdateGrTaskByQuery)
		rg.PATCH("/gr_task/m2m/part/:id", controllers.M2MGrTask)
		rg.DELETE("/gr_task/:id", controllers.DelGrTask)
		rg.POST("/gr_task/del", controllers.DelGrTasks)
		// user路由
		rg.GET("/user/:id", controllers.GetUserById)
		rg.GET("/user", controllers.GetUsers)
		rg.POST("/user", controllers.CreateUser)
		rg.PUT("/user/:id", controllers.UpdateUser)
		rg.PUT("/user/full/:id", controllers.UpdateFullUser)
		rg.PUT("/user", controllers.UpdateUsers)
		rg.PUT("/user/query", controllers.UpdateUserByQuery)
		rg.PATCH("/user/m2m/part/:id", controllers.M2MUser)
		rg.DELETE("/user/:id", controllers.DelUser)
		rg.POST("/user/del", controllers.DelUsers)
	}
}
