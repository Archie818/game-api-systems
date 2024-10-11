package models

import "time"

// Challenge represents a player's participation in a challenge
type Challenge struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    PlayerID  uint      `json:"player_id" gorm:"not null"`
    Amount    float64   `json:"amount" gorm:"not null"`
    Won       bool      `json:"won"`
    CreatedAt time.Time `json:"created_at"`
}