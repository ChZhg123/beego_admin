package inits

import (
	"blog/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)


func init()  {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.CONFIG.Mysql.Username,
		global.CONFIG.Mysql.Password,
		global.CONFIG.Mysql.Host,
		global.CONFIG.Mysql.Port,
		global.CONFIG.Mysql.DataBase,
	)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{
		Logger: newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil{
		panic("failed to connect database")
	}
	sqlDB,_ := db.DB()
	global.DB = db
	global.SqlDB = sqlDB
}

