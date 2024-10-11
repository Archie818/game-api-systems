package repository

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "interview_YangYang_20241010/models"
)

func TestPlayerCRUD(t *testing.T) {
    // Setup the test database
    db := SetupTestDB(t)
    defer TearDownTestDB(db, t)

    // Create a new player
    player := models.Player{
        Name:    "Test Player",
        LevelID: "1", // Assuming LevelID is a string
    }

    // Test CreatePlayer
    createdPlayerID, err := CreatePlayer(player)
    assert.NoError(t, err)
    assert.NotZero(t, createdPlayerID)

    // Test GetPlayerByID
    fetchedPlayer, err := GetPlayerByID(createdPlayerID)
    assert.NoError(t, err)
    assert.Equal(t, player.Name, fetchedPlayer.Name)
    assert.Equal(t, player.LevelID, fetchedPlayer.LevelID)

    // Test UpdatePlayer
    updatedPlayer := models.Player{
        ID:      createdPlayerID,
        Name:    "Updated Player",
        LevelID: "2",
    }
    err = UpdatePlayer(createdPlayerID, updatedPlayer)
    assert.NoError(t, err)

    // Verify the update
    fetchedUpdatedPlayer, err := GetPlayerByID(createdPlayerID)
    assert.NoError(t, err)
    assert.Equal(t, updatedPlayer.Name, fetchedUpdatedPlayer.Name)
    assert.Equal(t, updatedPlayer.LevelID, fetchedUpdatedPlayer.LevelID)

    // Test DeletePlayer
    err = DeletePlayer(createdPlayerID)
    assert.NoError(t, err)

    // Verify deletion
    _, err = GetPlayerByID(createdPlayerID)
    assert.Error(t, err) // Expect an error when fetching a deleted player
}