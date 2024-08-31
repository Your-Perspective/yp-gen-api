package config

import (
	"database/sql"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

// DB is the global database connection
var DB *gorm.DB

// InitDatabase initializes the database connection
func InitDatabase() {
	var err error
	// Use "sqlite" as the driver name with modernc.org/sqlite
	sqlDB, err := sql.Open("sqlite", "test.db")
	if err != nil {
		log.Fatalf("Failed to open SQLite database: %v", err)
	}

	// Pass the sqlDB to GORM
	DB, err = gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
}

// CloseDatabase closes the database connection
func CloseDatabase() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Failed to close the database: %v", err)
	}
}
