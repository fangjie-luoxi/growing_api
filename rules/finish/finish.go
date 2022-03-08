package finish

import (
	"github.com/fangjie-luoxi/growing_api/global"
	"github.com/fangjie-luoxi/growing_api/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Task(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		global.Resp.Error(c, 400, err, "请求id错误")
		return
	}
	tx := global.DBEngine.Begin()
	m := models.GrTask{Id: id}
	tx.Preload("GrTarget").First(&m)
	m.Status = "s"

	if m.GrTargetId != 0 {
		// todo 修改目标
		m.GrTarget.Finish += m.TtNum
	}
	err = tx.Save(&m).Error
	if err != nil {
		global.Resp.Error(c, 500, err, "")
		tx.Rollback()
		return
	}
	if m.Num != 0 && m.GrTargetId == 0 {
		integral := models.GrIntegralR{
			UserId: m.UserId,
			InType: "i",
			Num:    m.Num,
			ReTpye: "tt",
			ReId:   m.Id,
			Desc:   m.TkTitle,
		}
		tx.Create(&integral)
		user := models.User{Id: m.UserId}
		tx.First(&user)
		user.Integral = user.Integral + m.Num
		tx.Save(&user)
	}
	tx.Commit()
	global.Resp.OK(c, 200, m)
}

func Target(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		global.Resp.Error(c, 400, err, "请求id错误")
		return
	}
	m := models.GrTarget{Id: id}
	tx := global.DBEngine.Begin()

	tx.First(&m)
	m.Status = "s"
	err = tx.Save(&m).Error
	if err != nil {
		global.Resp.Error(c, 500, err, "")
		return
	}
	if m.Num != 0 {
		integral := models.GrIntegralR{
			UserId: m.UserId,
			InType: "i",
			Num:    m.Num,
			ReTpye: "tk",
			ReId:   m.Id,
			Desc:   m.TtTitle,
		}
		tx.Create(&integral)
		user := models.User{Id: m.UserId}
		tx.First(&user)
		user.Integral = user.Integral + m.Num
		tx.Save(&user)
	}
	tx.Commit()
	global.Resp.OK(c, 200, m)
}

func Rule(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		global.Resp.Error(c, 400, err, "请求id错误")
		return
	}
	tx := global.DBEngine.Begin()
	m := models.GrRule{Id: id}
	err = tx.First(&m).Error
	if err != nil {
		global.Resp.Error(c, 500, err, "")
		tx.Rollback()
		return
	}
	tn := time.Now()
	ruleR := models.GrRuleRecord{
		GrRuleId: m.Id,
		Date:     &tn,
	}
	tx.Create(&ruleR)
	if m.Num != 0 {
		integral := models.GrIntegralR{
			UserId: m.UserId,
			InType: m.InType,
			Num:    m.Num,
			ReTpye: "re",
			ReId:   m.Id,
			Desc:   m.ReName,
		}
		tx.Create(&integral)
		user := models.User{Id: m.UserId}
		tx.First(&user)
		num := m.Num
		if m.InType == "o" {
			num = num * -1
		}
		user.Integral = user.Integral + num
		tx.Save(&user)
	}
	tx.Commit()
	global.Resp.OK(c, 200, m)
}
