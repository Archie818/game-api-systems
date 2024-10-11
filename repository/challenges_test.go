// repository/challenges_test.go
package repository

import (
	"testing"

	"interview_YangYang_20241010/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateChallenge(t *testing.T) {
	SetupTestDB(t)
	defer TearDownTestDB(TestDB, t)

	challenge := models.Challenge{
		PlayerID:  1, // Ensure a Player with ID 1 exists
		Amount:    20.01,
		Won:       false,
	}

	challengeID, err := CreateChallenge(challenge)
	assert.NoError(t, err)
	assert.NotZero(t, challengeID)
}

func TestGetChallengeByID(t *testing.T) {
	SetupTestDB(t)
	defer TearDownTestDB(TestDB, t)

	// Create a challenge first
	challenge := models.Challenge{
		PlayerID:  2,
		Amount:    20.01,
		Won:       false,
	}
	challengeID, err := CreateChallenge(challenge)
	assert.NoError(t, err)

	// Retrieve the challenge
	retrievedChallenge, err := GetChallengeByID(challengeID)
	assert.NoError(t, err)
	assert.Equal(t, uint(2), retrievedChallenge.PlayerID)
	assert.Equal(t, 20.01, retrievedChallenge.Amount)
	assert.False(t, retrievedChallenge.Won)
}

func TestUpdateChallenge(t *testing.T) {
	SetupTestDB(t)
	defer TearDownTestDB(TestDB, t)

	// Create a challenge first
	challenge := models.Challenge{
		PlayerID:  3,
		Amount:    20.01,
		Won:       false,
	}
	challengeID, err := CreateChallenge(challenge)
	assert.NoError(t, err)

	// Update the challenge's status to won
	updatedChallenge := models.Challenge{
		PlayerID:  3,
		Amount:    20.01,
		Won:       true,
	}
	err = UpdateChallenge(updatedChallenge)
	assert.NoError(t, err)

	// Retrieve and verify the update
	retrievedChallenge, err := GetChallengeByID(challengeID)
	assert.NoError(t, err)
	assert.True(t, retrievedChallenge.Won)
}
