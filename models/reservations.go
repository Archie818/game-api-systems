package models

import "time"

// Reservation represents a reservation for a game room
type Reservation struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    RoomID    uint      `json:"room_id" gorm:"not null"`
    Date      time.Time `json:"date"`
    Time      string    `json:"time"` 
    PlayerInfo string    `json:"player_info"` 
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
