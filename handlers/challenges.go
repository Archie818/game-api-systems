package handlers

import (
    "errors"
    "math/rand"
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "interview_YangYang_20241010/models"
    "interview_YangYang_20241010/repository"
)

// ChallengeRequest represents the request body for creating a challenge.
type ChallengeRequest struct {
    PlayerID uint `json:"player_id" binding:"required"`
}

// ChallengeResponse represents the response after creating a challenge.
type ChallengeResponse struct {
    Status string `json:"status"`
    ID     uint   `json:"id,omitempty"`
    Error  string `json:"error,omitempty"`
}

// SuccessResponse represents a generic success response.
type SuccessResponse struct {
    Status string `json:"status"`
}

// @Summary Participate in a Challenge
// @Description Players can participate in an endless challenge by paying 20.01.
// @Tags Challenges
// @Accept json
// @Produce json
// @Param challenge body ChallengeRequest true "Challenge Participation"
// @Success 200 {object} ChallengeResponse "Challenge started"
// @Failure 400 {object} ChallengeResponse "Bad Request"
// @Failure 500 {object} ChallengeResponse "Internal Server Error"
// @Router /challenges [post]
func ParticipateChallenge(c *gin.Context) {
    var req ChallengeRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, ChallengeResponse{Error: err.Error()})
        return
    }

    // Initialize a new challenge with a fixed amount of 20.01
    challenge := models.Challenge{
        PlayerID: req.PlayerID,
        Amount:   20.01,
        Won:      false,
    }

    // Attempt to create the challenge
    challengeID, err := repository.CreateChallenge(challenge)
    if err != nil {
        if errors.Is(err, repository.ErrPlayerNotAllowed) {
            c.JSON(http.StatusBadRequest, ChallengeResponse{Error: "Player can only participate once per minute"})
            return
        }
        c.JSON(http.StatusInternalServerError, ChallengeResponse{Error: "Failed to create challenge"})
        return
    }

    // Respond immediately with the challenge status
    c.JSON(http.StatusOK, ChallengeResponse{
        Status: "challenge started",
        ID:     challengeID,
    })

    // Process the challenge outcome after 30 seconds in the background
    go processChallengeOutcome(challengeID)
}

// processChallengeOutcome determines the outcome of a challenge after a delay.
func processChallengeOutcome(challengeID uint) {
    // Wait for 30 seconds before determining the outcome
    time.Sleep(30 * time.Second)

    // Retrieve the challenge from the database
    challenge, err := repository.GetChallengeByID(challengeID)
    if err != nil {
        // Log the error if necessary (not implemented here)
        return
    }

    // Retrieve the player's total number of participations to adjust win probability
    participationCount, err := repository.GetPlayerParticipationCount(challenge.PlayerID)
    if err != nil {
        // Log the error if necessary (not implemented here)
        return
    }

    // Base win probability is 1%, increases by 1% per participation
    winProbability := float64(participationCount) * 1.0
    if winProbability > 100.0 {
        winProbability = 100.0
    }

    // Generate a random float between 0 and 100
    rand.Seed(time.Now().UnixNano())
    randomNumber := rand.Float64() * 100.0

    // Determine if the player wins based on the probability
    if randomNumber < winProbability {
        challenge.Won = true
    } else {
        challenge.Won = false
    }

    // Update the challenge outcome in the database
    repository.UpdateChallenge(*challenge)
}

// @Summary Get Recent Challenge Results
// @Description Retrieve a list of recent challenge results.
// @Tags Challenges
// @Accept json
// @Produce json
// @Param limit query int false "Maximum number of results to return"
// @Success 200 {array} models.Challenge "A list of recent challenges"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /challenges/results [get]
func GetChallengeResults(c *gin.Context) {
    // Retrieve 'limit' from query parameters; default to 10 if not provided
    limitParam := c.Query("limit")
    limit := 10
    if limitParam != "" {
        parsedLimit, err := strconv.Atoi(limitParam)
        if err == nil && parsedLimit > 0 {
            limit = parsedLimit
        }
    }

    // Fetch recent challenge results from the repository
    challenges, err := repository.GetRecentChallengeResults(limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve challenge results"})
        return
    }

    // Respond with the list of challenges
    c.JSON(http.StatusOK, challenges)
}