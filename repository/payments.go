package repository

import (
	"errors"

	"interview_YangYang_20241010/models"

	"gorm.io/gorm"
)

// Define custom errors
var (
	ErrPaymentNotFound = errors.New("payment not found")
)

// CreatePayment adds a new payment record to the database.
func CreatePayment(payment models.Payment) (uint, error) {
	payment.Status = "Pending"
	if err := DB.Create(&payment).Error; err != nil {
		return 0, err
	}
	return payment.ID, nil
}

// GetPaymentByID retrieves a payment by its ID.
func GetPaymentByID(id uint) (*models.Payment, error) {
	var payment models.Payment
	if err := DB.First(&payment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPaymentNotFound
		}
		return nil, err
	}
	return &payment, nil
}

// UpdatePayment updates the payment record in the database.
func UpdatePayment(payment models.Payment) error {
	return DB.Save(&payment).Error
}