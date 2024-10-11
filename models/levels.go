package models

type Level struct {
    ID   string `json:"id" gorm:"primaryKey"`
    Name string `json:"name" gorm:"unique"`
}