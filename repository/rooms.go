// repository/rooms.go
package repository

import (
    "errors"

    "interview_YangYang_20241010/models"
    "gorm.io/gorm"
)

var (
    ErrRoomNotFound = errors.New("room not found")
)

// GetAllRooms retrieves all game rooms with their reservations
func GetAllRooms() ([]models.Room, error) {
    var rooms []models.Room
    result := DB.Preload("Reservations").Find(&rooms)
    return rooms, result.Error
}

// GetRoomByID retrieves a room by its ID
func GetRoomByID(id uint) (*models.Room, error) {
    var room models.Room
    result := DB.Preload("Reservations").First(&room, "id = ?", id)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return nil, ErrRoomNotFound
    }
    return &room, result.Error
}

// CreateRoom adds a new room to the database
func CreateRoom(room models.Room) (uint, error) {
    result := DB.Create(&room)
    if result.Error != nil {
        return 0, result.Error
    }
    return room.ID, nil
}

// UpdateRoom updates an existing room's information
func UpdateRoom(id uint, updatedRoom models.Room) error {
    var room models.Room
    result := DB.First(&room, "id = ?", id)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return ErrRoomNotFound
    }

    // Update fields
    room.Name = updatedRoom.Name
    room.Description = updatedRoom.Description
    room.Status = updatedRoom.Status

    return DB.Save(&room).Error
}

// DeleteRoom removes a room from the database
func DeleteRoom(id uint) error {
    result := DB.Delete(&models.Room{}, "id = ?", id)
    if result.RowsAffected == 0 {
        return ErrRoomNotFound
    }
    return result.Error
}