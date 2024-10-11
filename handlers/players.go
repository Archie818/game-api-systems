package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "interview_YangYang_20241010/models"
    "interview_YangYang_20241010/repository"
)

// @Summary Get all players
// @Description Retrieve a list of all players with their level information
// @Tags players
// @Accept json
// @Produce json
// @Success 200 {array} models.Player "A list of players"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /players [get]
func GetPlayers(c *gin.Context) {
    players, err := repository.GetAllPlayers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        return
    }
    c.JSON(http.StatusOK, players)
}

// @Summary Register a new player
// @Description Create a new player with a specified level
// @Tags players
// @Accept json
// @Produce json
// @Param player body models.Player true "Player Information"
// @Success 201 {object} map[string]string "Successfully created player ID"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /players [post]
func CreatePlayer(c *gin.Context) {
    var player models.Player
    if err := c.ShouldBindJSON(&player); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    // Validate that the level exists
    if _, err := repository.GetLevelByID(player.LevelID); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid level ID"})
        return
    }

    id, err := repository.CreatePlayer(player)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        return
    }
    c.JSON(http.StatusCreated, map[string]string{"id": id})
}

// @Summary Get player by ID
// @Description Retrieve detailed information of a player by their ID
// @Tags players
// @Accept json
// @Produce json
// @Param id path string true "Player ID"
// @Success 200 {object} models.Player "Player details"
// @Failure 404 {object} models.ErrorResponse "Player not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /players/{id} [get]
func GetPlayerByID(c *gin.Context) {
    id := c.Param("id")
    player, err := repository.GetPlayerByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Player not found"})
        return
    }
    c.JSON(http.StatusOK, player)
}

// @Summary Update player information
// @Description Update the information of an existing player
// @Tags players
// @Accept json
// @Produce json
// @Param id path string true "Player ID"
// @Param player body models.Player true "Updated Player Information"
// @Success 200 {object} models.SuccessResponse "Update status"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /players/{id} [put]
func UpdatePlayer(c *gin.Context) {
    id := c.Param("id")
    var player models.Player
    if err := c.ShouldBindJSON(&player); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    // Validate that the level exists
    if _, err := repository.GetLevelByID(player.LevelID); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid level ID"})
        return
    }

    err := repository.UpdatePlayer(id, player)
    if err != nil {
        if err == repository.ErrPlayerNotFound {
            c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Player not found"})
        } else {
            c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        }
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse{Status: "updated"})
}

// @Summary Delete a player
// @Description Remove a player from the system by their ID
// @Tags players
// @Accept json
// @Produce json
// @Param id path string true "Player ID"
// @Success 200 {object} models.SuccessResponse "Deletion status"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /players/{id} [delete]
func DeletePlayer(c *gin.Context) {
    id := c.Param("id")
    err := repository.DeletePlayer(id)
    if err != nil {
        if err == repository.ErrPlayerNotFound {
            c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Player not found"})
        } else {
            c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        }
        return
    }
    c.JSON(http.StatusOK, models.SuccessResponse{Status: "deleted"})
}