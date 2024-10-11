// repository/logs_test.go
package repository

import (
	"testing"
	"time"

	"interview_YangYang_20241010/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateLog(t *testing.T) {
	SetupTestDB(t)
	defer TearDownTestDB(TestDB, t)

	logEntry := models.Log{
		PlayerID:  1, // Ensure a Player with ID 1 exists
		Action:    "Login",
		Details:   "Player logged in successfully.",
		Timestamp: time.Now(),
	}

	logID, err := CreateLog(logEntry)
	assert.NoError(t, err)
	assert.NotZero(t, logID)
}
func TestGetLogByID(t *testing.T) {
	SetupTestDB(t)
	defer TearDownTestDB(TestDB, t)

	// Create a log first
	logEntry := models.Log{
		PlayerID:  2,
		Action:    "Register",
		Details:   "Player registered successfully.",
		Timestamp: time.Now(),
	}
	logID, err := CreateLog(logEntry)
	assert.NoError(t, err)

	// Retrieve the log
	retrievedLogs, err := QueryLogs(&logID, nil, nil, nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(retrievedLogs)) // Assuming there's only one log with the given ID
	retrievedLog := retrievedLogs[0]
	assert.Equal(t, uint(2), retrievedLog.PlayerID)
	assert.Equal(t, "Register", retrievedLog.Action)
	assert.Equal(t, "Player registered successfully.", retrievedLog.Details)
}