// repository/reservations_test.go
package repository

import (
	"testing"
	"time"

	"interview_YangYang_20241010/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateReservation(t *testing.T) {
	SetupTestDB(t)
	defer TearDownTestDB(TestDB, t)

	reservation := models.Reservation{
		RoomID:     1, // Ensure a Room with ID 1 exists
		Date:       time.Now(), // Set the reservation date
		Time:       "14:00", // Set the reservation time
		PlayerInfo: "Player 1", // Set player information
	}

	reservationID, err := CreateReservation(reservation)
	assert.NoError(t, err)
	assert.NotZero(t, reservationID)
}

func TestGetReservationByID(t *testing.T) {
	SetupTestDB(t)
	defer TearDownTestDB(TestDB, t)

	// Create a reservation first
	reservation := models.Reservation{
		RoomID:     1,
		Date:       time.Now(),
		Time:       "15:00",
		PlayerInfo: "Player 2",
	}
	reservationID, err := CreateReservation(reservation)
	assert.NoError(t, err)

	// Retrieve the reservation
	retrievedReservations, err := GetReservations(reservationID, time.Now(), 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, retrievedReservations)
	assert.Equal(t, "Player 2", retrievedReservations[0].PlayerInfo)
}
