package repository

import (
    "errors"

    "interview_YangYang_20241010/models"
    "gorm.io/gorm"
)

var (
    ErrLevelNotFound = errors.New("level not found")
)

// GetAllLevels retrieves all levels
func GetAllLevels() ([]models.Level, error) {
    var levels []models.Level
    result := DB.Find(&levels)
    return levels, result.Error
}

// GetLevelByID retrieves a level by its ID
func GetLevelByID(id string) (*models.Level, error) {
    var level models.Level
    result := DB.First(&level, "id = ?", id)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return nil, ErrLevelNotFound
    }
    return &level, result.Error
}

// CreateLevel adds a new level to the database
func CreateLevel(level models.Level) (string, error) {
    result := DB.Create(&level)
    return level.ID, result.Error
}