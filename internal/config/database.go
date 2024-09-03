package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	_ "modernc.org/sqlite" // Importing the modernc.org/sqlite driver
)

// DB is the global database connection
var DB *gorm.DB

// InitDatabase initializes the database connection with detailed logging
func InitDatabase() {
	var err error
	dbPath := os.Getenv("SQLITE_DB_PATH")
	if dbPath == "" {
		dbPath = "./data/test.db" // Default path if the environment variable is not set
	}

	// Use "sqlite" as the driver name with modernc.org/sqlite
	sqlDB, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("Failed to open SQLite database: %v", err)
	}

	// Configure GORM Logger for detailed output
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Output logs to standard output
		logger.Config{
			SlowThreshold:             time.Second, // Log SQL queries that take more than 1 second
			LogLevel:                  logger.Info, // Log level (Silent, Error, Warn, Info)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Enable colorful logs
		},
	)

	// Open the database with the configured logger
	DB, err = gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{
		Logger: newLogger, // Attach the custom logger
	})
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
