package repository

import (
    "log"
    "interview_YangYang_20241010/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection and performs migrations
func InitDB() {
    dsn := "host=db user=archie password=postgres dbname=spinnerdb port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Perform migrations
    err = DB.AutoMigrate(&models.Player{}, &models.Level{}, &models.Room{}, &models.Reservation{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }
}