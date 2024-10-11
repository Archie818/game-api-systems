package repository

import (
	"time"

	"interview_YangYang_20241010/models"
)

// CreateLog adds a new log entry to the database.
func CreateLog(logEntry models.Log) (uint, error) {
	if err := DB.Create(&logEntry).Error; err != nil {
		return 0, err
	}
	return logEntry.ID, nil
}

// QueryLogs retrieves logs based on the provided filters.
// If a filter is not provided (zero value), it is ignored.
func QueryLogs(playerID *uint, action *string, startTime *time.Time, endTime *time.Time, limit *int) ([]models.Log, error) {
	var logs []models.Log
	query := DB.Model(&models.Log{})

	if playerID != nil {
		query = query.Where("player_id = ?", *playerID)
	}

	if action != nil {
		query = query.Where("action = ?", *action)
	}

	if startTime != nil && endTime != nil {
		query = query.Where("timestamp BETWEEN ? AND ?", *startTime, *endTime)
	} else if startTime != nil {
		query = query.Where("timestamp >= ?", *startTime)
	} else if endTime != nil {
		query = query.Where("timestamp <= ?", *endTime)
	}

	if limit != nil && *limit > 0 {
		query = query.Limit(*limit)
	}

	if err := query.Order("timestamp desc").Find(&logs).Error; err != nil {
		return nil, err
	}

	return logs, nil
}