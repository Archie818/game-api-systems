// repository/rooms_test.go
package repository

import (
	"testing"

	"interview_YangYang_20241010/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateRoom(t *testing.T) {
	SetupTestDB(t)
	defer TearDownTestDB(TestDB, t)

	room := models.Room{
		Name:        "Room A",
		Description: "First game room",
		Status:      "available",
	}

	roomID, err := CreateRoom(room)
	assert.NoError(t, err)
	assert.NotZero(t, roomID)
}

func TestGetRoomByID(t *testing.T) {
	SetupTestDB(t)
	defer TearDownTestDB(TestDB, t)

	// Create a room first
	room := models.Room{
		Name:        "Room B",
		Description: "Second game room",
		Status:      "occupied",
	}
	roomID, err := CreateRoom(room)
	assert.NoError(t, err)

	// Retrieve the room
	retrievedRoom, err := GetRoomByID(roomID)
	assert.NoError(t, err)
	assert.Equal(t, "Room B", retrievedRoom.Name)
	assert.Equal(t, "Second game room", retrievedRoom.Description)
	assert.Equal(t, "occupied", retrievedRoom.Status)
}

func TestUpdateRoom(t *testing.T) {
	SetupTestDB(t)
	defer TearDownTestDB(TestDB, t)

	// Create a room first
	room := models.Room{
		Name:        "Room C",
		Description: "Third game room",
		Status:      "available",
	}
	roomID, err := CreateRoom(room)
	assert.NoError(t, err)

	// Update the room's status
	updatedRoom := models.Room{
		Name:        "Room C",
		Description: "Third game room updated",
		Status:      "occupied",
	}
	err = UpdateRoom(roomID, updatedRoom)
	assert.NoError(t, err)

	// Retrieve and verify the update
	retrievedRoom, err := GetRoomByID(roomID)
	assert.NoError(t, err)
	assert.Equal(t, "Room C", retrievedRoom.Name)
	assert.Equal(t, "Third game room updated", retrievedRoom.Description)
	assert.Equal(t, "occupied", retrievedRoom.Status)
}

func TestDeleteRoom(t *testing.T) {
	SetupTestDB(t)
	defer TearDownTestDB(TestDB, t)

	// Create a room first
	room := models.Room{
		Name:        "Room D",
		Description: "Fourth game room",
		Status:      "available",
	}
	roomID, err := CreateRoom(room)
	assert.NoError(t, err)

	// Delete the room
	err = DeleteRoom(roomID)
	assert.NoError(t, err)

	// Attempt to retrieve the deleted room
	_, err = GetRoomByID(roomID)
	assert.Error(t, err)
}