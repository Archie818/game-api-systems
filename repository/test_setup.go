// repository/test_setup.go
package repository

import (
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"interview_YangYang_20241010/models" // Ensure this import is correct
)

// TestDB is the global database connection used for testing
var TestDB *gorm.DB

// SetupTestDB initializes the test database connection
func SetupTestDB(t *testing.T) *gorm.DB {
	var err error
	// Replace with your actual test database connection string
	dsn := "host=localhost user=archie password=postgres dbname=spinnerdb port=5432 sslmode=disable"
	TestDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Perform migrations for all models
	err = TestDB.AutoMigrate(
		&models.Player{},
		&models.Level{},
		&models.Room{},
		&models.Reservation{},
		&models.Challenge{},
		&models.Log{},
		&models.Payment{},
	)
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	// Assign TestDB to the repository's global DB variable
	DB = TestDB

	return TestDB
}

// TearDownTestDB closes the database connection and cleans up
func TearDownTestDB(db *gorm.DB, t *testing.T) {
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Failed to get database from gorm: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		t.Fatalf("Failed to close database: %v", err)
	}

	// Optionally, you can drop the test database or clean up test data here
}