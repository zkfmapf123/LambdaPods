package configs

import (
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MustInitDB(
	DB_HOST string,
	DB_PORT string,
	DB_USER string,
	DB_PASS string,
	DB_NAME string,
) *gorm.DB {
	// pgx 로 연결
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)

	config, _ := pgx.ParseConfig(dsn)

	// gorm 으로 매핑
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: stdlib.OpenDB(*config),
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	sql, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sql.SetMaxIdleConns(10)
	sql.SetMaxOpenConns(100)
	sql.SetConnMaxLifetime(time.Hour)

	if err := sql.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	return db
}
