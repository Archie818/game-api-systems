package models

type Player struct {
    ID      string `json:"id" gorm:"primaryKey"`
    Name    string `json:"name"`
    LevelID string `json:"level_id"`
    levle Level `json:"level" gorm:"foreignKey:LevelID"`
}