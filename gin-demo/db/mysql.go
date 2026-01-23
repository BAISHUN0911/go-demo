package db

import (
	"fmt"
	"log"
	"time"

	"gin-demo/configx"
	"gin-demo/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMySQL() {
	cfg := configx.MySQL

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)

	// 添加自动迁移
	err = DB.AutoMigrate(model.AllModels()...)
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}

	log.Println("MySQL连接成功")
}
