// repository/reservations.go
package repository

import (
    "errors"
    "time"

    "interview_YangYang_20241010/models"
)

var (
    ErrReservationNotFound = errors.New("reservation not found")
)

// GetReservations retrieves reservations based on optional filters
func GetReservations(roomID uint, date time.Time, limit int) ([]models.Reservation, error) {
    var reservations []models.Reservation
    query := DB.Preload("Room").Preload("Player").Model(&models.Reservation{})

    if roomID != 0 {
        query = query.Where("room_id = ?", roomID)
    }
    if !date.IsZero() {
        formattedDate := date.Format("2006-01-02")
        query = query.Where("date = ?", formattedDate)
    }
    if limit > 0 {
        query = query.Limit(limit)
    }

    result := query.Find(&reservations)
    return reservations, result.Error
}

// CreateReservation adds a new reservation to the database
func CreateReservation(reservation models.Reservation) (uint, error) {
    result := DB.Create(&reservation)
    if result.Error != nil {
        return 0, result.Error
    }
    return reservation.ID, nil
}