package models

import (
	"time"
)

// Log represents a player's game operation
type Log struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	PlayerID    uint      `json:"player_id" gorm:"not null"`
	Action      string    `json:"action" gorm:"not null"` // e.g., Register, Login, Logout, Enter Room, Exit Room, Participate in Challenge, Challenge Result
	Timestamp   time.Time `json:"timestamp" gorm:"autoCreateTime"` // Automatically set to current time on creation
	Details     string    `json:"details"`                            // Additional information about the action
}