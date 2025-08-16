package configs

import (
	"fmt"
	"log"
	"time"

	"github.com/zkfmapf123/lambda-pods/domains"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MustInitDB(
	DB_HOST string,
	DB_PORT string,
	DB_USER string,
	DB_PASS string,
	DB_NAME string,
) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)

	// gorm 으로 매핑
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	sql, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// migrate
	if err := MigrateDB(db); err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}

	sql.SetMaxIdleConns(10)
	sql.SetMaxOpenConns(100)
	sql.SetConnMaxLifetime(time.Hour)

	if err := sql.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	return db
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		&domains.User{},
	)
}

func CloseDB(db *gorm.DB) error {
	sql, err := db.DB()
	if err != nil {
		return err
	}

	sql.Close()
	return nil
}
