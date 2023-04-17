package data

import (
	"fmt"
	"log"

	"github.com/Github-Aiko/Aiko-Telegram-Bot/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(cfg config.Config) *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.APPs.Database.User,
		cfg.APPs.Database.Pass,
		cfg.APPs.Database.IP,
		cfg.APPs.Database.Port,
		cfg.APPs.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("数据库连接失败请检查数据库配置文件：%s", err)
	}

	return db

}
