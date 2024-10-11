package repository

import (
    "errors"

    "interview_YangYang_20241010/models"
    "gorm.io/gorm"
)

var (
    ErrPlayerNotFound = errors.New("player not found")
)

// GetAllPlayers retrieves all players with their associated levels
func GetAllPlayers() ([]models.Player, error) {
    var players []models.Player
    result := DB.Preload("Level").Find(&players)
    return players, result.Error
}

// GetPlayerByID retrieves a player by their ID
func GetPlayerByID(id string) (*models.Player, error) {
    var player models.Player
    result := DB.Preload("Level").First(&player, "id = ?", id)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return nil, ErrPlayerNotFound
    }
    return &player, result.Error
}

// CreatePlayer adds a new player to the database
func CreatePlayer(player models.Player) (string, error) {
    result := DB.Create(&player)
    return player.ID, result.Error
}

// UpdatePlayer updates an existing player's information
func UpdatePlayer(id string, updatedPlayer models.Player) error {
    var player models.Player
    result := DB.First(&player, "id = ?", id)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return ErrPlayerNotFound
    }

    // Update fields
    player.Name = updatedPlayer.Name
    player.LevelID = updatedPlayer.LevelID

    return DB.Save(&player).Error
}

// DeletePlayer removes a player from the database
func DeletePlayer(id string) error {
    result := DB.Delete(&models.Player{}, "id = ?", id)
    if result.RowsAffected == 0 {
        return ErrPlayerNotFound
    }
    return result.Error
}