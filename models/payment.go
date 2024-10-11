package models

import (
	"time"
)

// Payment represents a payment transaction made by a player
type Payment struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	PlayerID      uint      `json:"player_id" gorm:"not null"`
	Method        string    `json:"method" gorm:"not null"` // e.g., CreditCard, BankTransfer, ThirdParty, Blockchain
	Amount        float64   `json:"amount" gorm:"not null"`
	Details       string    `json:"details" gorm:"type:text"` // JSON string containing payment method details
	Status        string    `json:"status" gorm:"not null"`  // e.g., Pending, Success, Failed
	TransactionID string    `json:"transaction_id"`            // Populated on success
	ErrorMessage  string    `json:"error_message"`             // Populated on failure
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}