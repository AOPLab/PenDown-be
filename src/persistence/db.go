package persistence

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(10)           // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxOpenConns(100)          // SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetConnMaxLifetime(time.Hour) // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.

	return DB, err
}

func InitTestDB(mock_db *gorm.DB) {
	DB = mock_db
}
