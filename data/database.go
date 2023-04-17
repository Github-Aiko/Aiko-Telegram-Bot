package data

import (
	"fmt"
	"log"
	"time"

	"github.com/Github-Aiko/Aiko-Telegram-Bot/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(cfg *config.Config) *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.APPs.Database.User,
		cfg.APPs.Database.Pass,
		cfg.APPs.Database.IP,
		cfg.APPs.Database.Port,
		cfg.APPs.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("The database connection failed. Please check the database configuration file: %s", err)
		panic(err)
	}

	setDatabaseConnectionPool(db)

	return db

}

// Set up a database connection pool
func setDatabaseConnectionPool(db *gorm.DB) {
	sqlDB, _ := db.DB()
	// The maximum number of connections in the connection pool.
	sqlDB.SetMaxIdleConns(10)
	// The maximum number of connections to the database
	sqlDB.SetMaxOpenConns(100)
	// Set the maximum time a connection can be reused
	sqlDB.SetConnMaxLifetime(10 * time.Second) //10 seconds
}
