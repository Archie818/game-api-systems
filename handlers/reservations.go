// handlers/reservations.go
package handlers

import (
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "interview_YangYang_20241010/models"
    "interview_YangYang_20241010/repository"
)

// ReservationInput represents the expected input for creating a reservation
type ReservationInput struct {
    RoomID     uint      `json:"room_id" binding:"required"`
    Date       string    `json:"date" binding:"required"`       // Expected format: "YYYY-MM-DD"
    Time       string    `json:"time" binding:"required"`       // Expected format: "HH:MM-HH:MM"
    PlayerInfo string    `json:"player_info" binding:"required"` // JSON or string
}

// @Summary Get reservations
// @Description Retrieve reservations with optional filters for room ID, date, and limit
// @Tags reservations
// @Accept json
// @Produce json
// @Param room_id query uint false "Room ID to filter reservations"
// @Param date query string false "Date to filter reservations (YYYY-MM-DD)"
// @Param limit query int false "Maximum number of reservations to return"
// @Success 200 {array} models.Reservation "A list of reservations"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /reservations [get]
func GetReservations(c *gin.Context) {
    roomIDStr := c.Query("room_id")
    dateStr := c.Query("date")
    limitStr := c.Query("limit")

    var roomID uint
    var err error

    if roomIDStr != "" {
        roomID64, err := strconv.ParseUint(roomIDStr, 10, 32)
        if err != nil {
            c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid room_id"})
            return
        }
        roomID = uint(roomID64)
    }

    var date time.Time
    if dateStr != "" {
        date, err = time.Parse("2006-01-02", dateStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid date format. Expected YYYY-MM-DD"})
            return
        }
    }

    var limit int
    if limitStr != "" {
        limit, err = strconv.Atoi(limitStr)
        if err != nil || limit <= 0 {
            c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid limit"})
            return
        }
    }

    reservations, err := repository.GetReservations(roomID, date, limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        return
    }
    c.JSON(http.StatusOK, reservations)
}

// @Summary Create a new reservation
// @Description Add a new reservation for a game room
// @Tags reservations
// @Accept json
// @Produce json
// @Param reservation body ReservationInput true "Reservation Information"
// @Success 201 {object} map[string]uint "Successfully created reservation ID"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /reservations [post]
func CreateReservation(c *gin.Context) {
    var input ReservationInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    // Validate date format
    date, err := time.Parse("2006-01-02", input.Date)
    if err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid date format. Expected YYYY-MM-DD"})
        return
    }

    // Optional: Validate time format (e.g., "14:00-16:00")
    // You can implement proper time validation here

    // Check if the room exists
    if _, err := repository.GetRoomByID(input.RoomID); err != nil {
        if err == repository.ErrRoomNotFound {
            c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid room ID"})
        } else {
            c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        }
        return
    }

    // Optionally, you can check for overlapping reservations here

    reservation := models.Reservation{
        RoomID:     input.RoomID,
        Date:       date,
        Time:       input.Time,
        PlayerInfo: input.PlayerInfo,
    }

    id, err := repository.CreateReservation(reservation)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        return
    }
    c.JSON(http.StatusCreated, map[string]uint{"id": id})
}