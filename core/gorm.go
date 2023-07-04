package core

import (
	"go_blog/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() (db *gorm.DB) {
	if global.Config.Mysql.Host == "" {
		global.Log.Warnln(" Mysql Host 配置为空 取消连接数据库")
	}
	dsn := global.Config.Mysql.GetDsn()
	if global.Config.System.Env == "dev" {
		global.MysqlLog = logger.Default.LogMode(logger.Info)
	} else {
		global.MysqlLog = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: global.MysqlLog,
	})

	if err != nil {
		global.Log.Fatalf("gorm connect error:%v\n dsn: %v\n", err, dsn)
		return
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(10)            // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //连接最大复用时间 不能超过mysql的waite_timeout
	return db
}
