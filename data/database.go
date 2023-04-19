package data

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func New(dsn string) *gorm.DB {

	// 连接数据库
	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("数据库连接失败请检查数据库配置文件：%s", err)
	}

	// 设置数据库链接池
	// Set database connection pool
	setDatabaseConnectionPool(db)

	// 自动迁移
	// AutoMigrate
	db.AutoMigrate(&User{})

	return db

}

// 设置数据库链接池
// Set database connection pool
func setDatabaseConnectionPool(db *gorm.DB) {
	sqlDB, _ := db.DB()
	// 连接池中的最大连接数。
	// The maximum number of connections in the connection pool.
	sqlDB.SetMaxIdleConns(10)
	// 数据库的最大连接数
	// The maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// 设置连接可重用的最大时间
	// Set the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(10 * time.Second) //10秒
}
