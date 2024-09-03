package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// Config holds the application configuration
type Config struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

// LoadConfig initializes the configuration for the application
func LoadConfig() (*Config, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	// Initialize Logger
	logger, err := initLogger()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	// Initialize Database
	db, err := initDB(logger)
	if err != nil {
		logger.Error("failed to initialize database", zap.Error(err))
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	return &Config{
		DB:     db,
		Logger: logger,
	}, nil
}

// initLogger initializes the zap logger
func initLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()

	// Adjust log level if needed (info, warn, error, etc.)
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	// Adjust encoding if you want a different format (console, json, etc.)
	config.Encoding = "json"

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	zap.ReplaceGlobals(logger)
	return logger, nil
}

// initDB initializes the GORM database connection
func initDB(logger *zap.Logger) (*gorm.DB, error) {
	// Example DSN for a PostgreSQL connection
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "yourusername"),
		getEnv("DB_PASSWORD", "yourpassword"),
		getEnv("DB_NAME", "yourdbname"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_TIMEZONE", "UTC"),
	)

	// Set up GORM logger
	gormLogger := gormlogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		gormlogger.Config{
			SlowThreshold:             time.Second,     // Slow SQL query threshold
			LogLevel:                  gormlogger.Info, // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,            // Enable color
		},
	)

	// Connect to the database using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, err
	}

	// Ping the database to check the connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// getEnv retrieves environment variables or returns a default value if not set
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
