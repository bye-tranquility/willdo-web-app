package database

import (
	"fmt"
	"log"

	"willdo/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

func InitDB(dbURL string, logger *log.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.AutoMigrate(&models.Event{}); err != nil {
		return nil, fmt.Errorf("failed to auto migrate database: %w", err)
	}

	logger.Println("[INFO] Database connection established and migrations completed")
	return db, nil
}
