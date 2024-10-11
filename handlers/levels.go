// handlers/levels.go
package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "interview_YangYang_20241010/models"
    "interview_YangYang_20241010/repository"
)

// @Summary Get all levels
// @Description Retrieve a list of all levels
// @Tags levels
// @Accept json
// @Produce json
// @Success 200 {array} models.Level "A list of levels"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /levels [get]
func GetLevels(c *gin.Context) {
    levels, err := repository.GetAllLevels()
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        return
    }
    c.JSON(http.StatusOK, levels)
}

// @Summary Add a new level
// @Description Create a new level in the system
// @Tags levels
// @Accept json
// @Produce json
// @Param level body models.Level true "Level Information"
// @Success 201 {object} map[string]string "Successfully created level ID"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /levels [post]
func CreateLevel(c *gin.Context) {
    var level models.Level
    if err := c.ShouldBindJSON(&level); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    // Validate that the level name is provided
    if level.Name == "" {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Level name is required"})
        return
    }

    id, err := repository.CreateLevel(level)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        return
    }
    c.JSON(http.StatusCreated, map[string]string{"id": id})
}