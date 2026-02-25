package database

import (
	"assignment/internal/infrastructure/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// log.Fatal を log.Printf に変えて、プログラムを止めないようにする
		log.Printf("⚠️ DB接続に失敗しましたが、テスト用に続行します: %v", err)
		return nil
	}
	return db
}
