// handlers/rooms.go
package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "interview_YangYang_20241010/models"
    "interview_YangYang_20241010/repository"
)

// @Summary Get all game rooms
// @Description Retrieve a list of all game rooms with their details
// @Tags rooms
// @Accept json
// @Produce json
// @Success 200 {array} models.Room "A list of game rooms"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /rooms [get]
func GetRooms(c *gin.Context) {
    rooms, err := repository.GetAllRooms()
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        return
    }
    c.JSON(http.StatusOK, rooms)
}

// @Summary Create a new game room
// @Description Add a new game room with specified name and description
// @Tags rooms
// @Accept json
// @Produce json
// @Param room body models.Room true "Room Information"
// @Success 201 {object} map[string]uint "Successfully created room ID"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /rooms [post]
func CreateRoom(c *gin.Context) {
    var room models.Room
    if err := c.ShouldBindJSON(&room); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    // Validate that the room name is provided
    if room.Name == "" {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Room name is required"})
        return
    }

    id, err := repository.CreateRoom(room)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        return
    }
    c.JSON(http.StatusCreated, map[string]uint{"id": id})
}

// @Summary Get room details by ID
// @Description Retrieve detailed information of a game room by its ID
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path uint true "Room ID"
// @Success 200 {object} models.Room "Room details"
// @Failure 404 {object} models.ErrorResponse "Room not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /rooms/{id} [get]
func GetRoomByID(c *gin.Context) {
    var roomID uint
    if err := c.ShouldBindUri(&roomID); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid room ID"})
        return
    }

    room, err := repository.GetRoomByID(roomID)
    if err != nil {
        if err == repository.ErrRoomNotFound {
            c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Room not found"})
        } else {
            c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        }
        return
    }
    c.JSON(http.StatusOK, room)
}

// @Summary Update room information
// @Description Update the information of an existing game room
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path uint true "Room ID"
// @Param room body models.Room true "Updated Room Information"
// @Success 200 {object} models.SuccessResponse "Update status"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "Room not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /rooms/{id} [put]
func UpdateRoom(c *gin.Context) {
    var roomID uint
    if err := c.ShouldBindUri(&roomID); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid room ID"})
        return
    }

    var room models.Room
    if err := c.ShouldBindJSON(&room); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    // Validate that the room name is provided
    if room.Name == "" {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Room name is required"})
        return
    }

    err := repository.UpdateRoom(roomID, room)
    if err != nil {
        if err == repository.ErrRoomNotFound {
            c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Room not found"})
        } else {
            c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        }
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse{Status: "updated"})
}

// @Summary Delete a room
// @Description Remove a game room from the system by its ID
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path uint true "Room ID"
// @Success 200 {object} models.SuccessResponse "Deletion status"
// @Failure 404 {object} models.ErrorResponse "Room not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /rooms/{id} [delete]
func DeleteRoom(c *gin.Context) {
    var roomID uint
    if err := c.ShouldBindUri(&roomID); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid room ID"})
        return
    }

    err := repository.DeleteRoom(roomID)
    if err != nil {
        if err == repository.ErrRoomNotFound {
            c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Room not found"})
        } else {
            c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        }
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse{Status: "deleted"})
}