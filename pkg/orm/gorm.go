// Package orm
// orm模块, 日志默认保存路径`static/logs/orm/log.log`
package orm

import (
	"log"
	"os"
	"time"

	logIO "github.com/fangjie-luoxi/tools/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/fangjie-luoxi/growing_api/global"
)

func SetUp() *gorm.DB {
	writer := log.New(os.Stdout, "\r\n", log.LstdFlags)
	if global.Config.String("mode") == "release" {
		writer = log.New(logIO.NewLogIO("./static/logs/orm/log.log"), "\r\n", log.LstdFlags)
	}
	newLogger := logger.New(
		writer,
		logger.Config{
			SlowThreshold: 200 * time.Millisecond, // 慢SQL阈值
			LogLevel:      logger.Warn,            // 日志级别
			Colorful:      false,                  // 彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(global.Config.String("sqlConn")+"?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{
		Logger:                                   newLogger, // 日志记录
		DisableForeignKeyConstraintWhenMigrating: true,      // 自动迁移忽略外键
	})
	if err != nil {
		panic("连接数据库失败")
	}
	return db
}
