package handlers

import (
	"net/http"
	"strconv"
	"time"

	"interview_YangYang_20241010/models"
	"interview_YangYang_20241010/repository"

	"github.com/gin-gonic/gin"
)

// LogRequest represents the request body for creating a new log.
type LogRequest struct {
	PlayerID uint   `json:"player_id" binding:"required"`
	Action   string `json:"action" binding:"required"`   // e.g., Register, Login, Logout, Enter Room, Exit Room, Participate in Challenge, Challenge Result
	Details  string `json:"details" binding:"required"`  // Additional information about the action
}

// LogResponse represents the response after creating a new log.
type LogResponse struct {
	ID uint `json:"id"`
}

// @Summary Retrieve Game Logs
// @Description Retrieve a list of game logs with optional filters.
// @Tags Logs
// @Accept json
// @Produce json
// @Param player_id query uint false "Filter by Player ID"
// @Param action query string false "Filter by Action Type (e.g., Register, Login)"
// @Param start_time query string false "Filter logs from this time (RFC3339 format)"
// @Param end_time query string false "Filter logs up to this time (RFC3339 format)"
// @Param limit query int false "Maximum number of logs to return"
// @Success 200 {array} models.Log "A list of game logs"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /logs [get]
func GetLogs(c *gin.Context) {
	var playerID *uint
	var action *string
	var startTime *time.Time
	var endTime *time.Time
	var limit *int

	// Parse query parameters
	if pid := c.Query("player_id"); pid != "" {
		parsedPID, err := strconv.ParseUint(pid, 10, 64)
		if err == nil {
			pidUint := uint(parsedPID)
			playerID = &pidUint
		}
	}

	if a := c.Query("action"); a != "" {
		action = &a
	}

	if st := c.Query("start_time"); st != "" {
		parsedST, err := time.Parse(time.RFC3339, st)
		if err == nil {
			startTime = &parsedST
		}
	}

	if et := c.Query("end_time"); et != "" {
		parsedET, err := time.Parse(time.RFC3339, et)
		if err == nil {
			endTime = &parsedET
		}
	}

	if l := c.Query("limit"); l != "" {
		parsedL, err := strconv.Atoi(l)
		if err == nil && parsedL > 0 {
			limit = &parsedL
		}
	}

	// Query logs from the repository
	logs, err := repository.QueryLogs(playerID, action, startTime, endTime, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve logs"})
		return
	}

	// Respond with the logs
	c.JSON(http.StatusOK, logs)
}

// @Summary Create a Game Log
// @Description Create a new game operation log.
// @Tags Logs
// @Accept json
// @Produce json
// @Param log body LogRequest true "Game Log Information"
// @Success 200 {object} LogResponse "New Log ID"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /logs [post]
func CreateLog(c *gin.Context) {
	var req LogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new log entry
	logEntry := models.Log{
		PlayerID:  req.PlayerID,
		Action:    req.Action,
		Details:   req.Details,
		Timestamp: time.Now(),
	}

	logID, err := repository.CreateLog(logEntry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create log"})
		return
	}

	// Respond with the new log ID
	c.JSON(http.StatusOK, LogResponse{ID: logID})
}