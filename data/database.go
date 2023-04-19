package data

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(dsn string) (*gorm.DB, error) {

	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %s", err)
	}

	// Set database connection pool
	setDatabaseConnectionPool(db)

	// AutoMigrate
	if err := db.AutoMigrate(&User{}); err != nil {
		return nil, fmt.Errorf("failed to auto migrate database: %s", err)
	}

	log.Printf("successful database connection established with DSN: %s", dsn)
	return db, nil

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
