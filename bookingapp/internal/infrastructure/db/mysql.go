package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	User string
	Pass string
	Host string // e.g. 127.0.0.1
	Port int    // e.g. 3306
	Name string // database name
}

func Open(c Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		c.User, c.Pass, c.Host, c.Port, c.Name,
	)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		// ここで必要ならLoggerやNamingStrategyを設定
	})
}

// 便利: Ping（接続確認）
func Ping(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	return sqlDB.Ping()
}
