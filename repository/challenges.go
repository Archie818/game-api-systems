package repository

import (
    "errors"
    "time"

    "interview_YangYang_20241010/models"
    "gorm.io/gorm"
)

// Define custom errors
var (
    ErrPlayerNotAllowed   = errors.New("player can only participate once per minute")
    ErrChallengeNotFound  = errors.New("challenge not found")
)

// CreateChallenge adds a new challenge to the database after validating participation rules.
func CreateChallenge(challenge models.Challenge) (uint, error) {
    // Check if the player has participated in the last minute
    oneMinuteAgo := time.Now().Add(-1 * time.Minute)
    var count int64
    if err := DB.Model(&models.Challenge{}).
        Where("player_id = ? AND created_at > ?", challenge.PlayerID, oneMinuteAgo).
        Count(&count).Error; err != nil {
        return 0, err
    }
    if count > 0 {
        return 0, ErrPlayerNotAllowed
    }

    // Create the challenge
    if err := DB.Create(&challenge).Error; err != nil {
        return 0, err
    }

    return challenge.ID, nil
}

// GetRecentChallengeResults retrieves the most recent challenges up to the specified limit.
func GetRecentChallengeResults(limit int) ([]models.Challenge, error) {
    var challenges []models.Challenge
    if err := DB.Order("created_at desc").Limit(limit).Find(&challenges).Error; err != nil {
        return nil, err
    }
    return challenges, nil
}

// GetChallengeByID retrieves a challenge by its ID.
func GetChallengeByID(id uint) (*models.Challenge, error) {
    var challenge models.Challenge
    if err := DB.First(&challenge, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrChallengeNotFound
        }
        return nil, err
    }
    return &challenge, nil
}

// UpdateChallenge updates the outcome of a challenge.
func UpdateChallenge(challenge models.Challenge) error {
    return DB.Save(&challenge).Error
}

// GetPlayerParticipationCount retrieves the total number of participations by a player.
func GetPlayerParticipationCount(playerID uint) (int, error) {
    var count int64
    if err := DB.Model(&models.Challenge{}).
        Where("player_id = ?", playerID).
        Count(&count).Error; err != nil {
        return 0, err
    }
    return int(count), nil
}