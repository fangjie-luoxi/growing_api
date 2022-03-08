package cron

import (
	"github.com/fangjie-luoxi/growing_api/global"
	"github.com/fangjie-luoxi/growing_api/models"
	"github.com/robfig/cron/v3"
	"time"
)

func Cron() {
	c := cron.New()
	_, _ = c.AddFunc("@daily", dayTask)
	c.Start()
}

func dayTask() {
	var targets []*models.GrTarget
	tn := time.Now()
	tnStr := tn.Format("2006-01-02")
	global.DBEngine.Where("begin <= ?", tnStr).Where("end >= ?", tnStr).Where("gen_task = ?", "y").Find(&targets)
	var tasks []*models.GrTask
	for _, target := range targets {
		subDay := target.End.Sub(tn) / (24 * time.Hour)
		num := (target.TtNum - target.Finish) / float64(subDay)
		tasks = append(tasks, &models.GrTask{
			UserId:     target.UserId,
			GrTargetId: target.Id,
			TkTitle:    target.TtTitle,
			TkContent:  target.TtContent,
			TtNum:      num,
			TtUnit:     target.TtUnit,
			Num:        0,
			Rm:         target.Rm,
			Status:     "r",
			Date:       &tn,
		})
	}
	if len(tasks) > 0 {
		global.DBEngine.Create(&tasks)
	}
}
