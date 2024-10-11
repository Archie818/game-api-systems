// repository/levels_test.go
package repository

import (
	"testing"

	"interview_YangYang_20241010/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateLevel(t *testing.T) {
	SetupTestDB(t)
	defer TearDownTestDB(TestDB, t)

	level := models.Level{
		Name: "Beginner",
	}

	levelID, err := CreateLevel(level)
	assert.NoError(t, err)
	assert.NotZero(t, levelID)
}

func TestGetLevelByID(t *testing.T) {
	SetupTestDB(t)
	defer TearDownTestDB(TestDB, t)

	// Create a level first
	level := models.Level{
		Name: "Intermediate",
	}
	levelID, err := CreateLevel(level)
	assert.NoError(t, err)

	// Retrieve the level
	retrievedLevel, err := GetLevelByID(levelID)
	assert.NoError(t, err)
	assert.Equal(t, "Intermediate", retrievedLevel.Name)
}
