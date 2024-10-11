// repository/payments_test.go
package repository

import (
	"testing"

	"interview_YangYang_20241010/models"

	"github.com/stretchr/testify/assert"
)

func TestCreatePayment(t *testing.T) {
	SetupTestDB(t)
	defer TearDownTestDB(TestDB, t)

	payment := models.Payment{
		PlayerID:  1, // Ensure a Player with ID 1 exists
		Method:    "CreditCard",
		Amount:    100.00,
		Details:   `{"card_number":"4111111111111111","expiry":"12/25","cvv":"123"}`,
		Status:    "Pending",
	}

	paymentID, err := CreatePayment(payment)
	assert.NoError(t, err)
	assert.NotZero(t, paymentID)
}

func TestGetPaymentByID(t *testing.T) {
	SetupTestDB(t)
	defer TearDownTestDB(TestDB, t)

	// Create a payment first
	payment := models.Payment{
		PlayerID:      2,
		Method:        "ThirdParty",
		Amount:        50.00,
		Details:       `{"provider":"PayPal","account":"player@example.com"}`,
		Status:        "Success",
		TransactionID: "TP1234567890",
	}
	paymentID, err := CreatePayment(payment)
	assert.NoError(t, err)

	// Retrieve the payment
	retrievedPayment, err := GetPaymentByID(paymentID)
	assert.NoError(t, err)
	assert.Equal(t, uint(2), retrievedPayment.PlayerID)
	assert.Equal(t, "ThirdParty", retrievedPayment.Method)
	assert.Equal(t, 50.00, retrievedPayment.Amount)
	assert.Equal(t, "Success", retrievedPayment.Status)
	assert.Equal(t, "TP1234567890", retrievedPayment.TransactionID)
}

func TestUpdatePayment(t *testing.T) {
	SetupTestDB(t)
	defer TearDownTestDB(TestDB, t)

	// Create a payment first
	payment := models.Payment{
		PlayerID:      3,
		Method:        "BankTransfer",
		Amount:        200.00,
		Details:       `{"bank_account":"123456789","bank_code":"001"}`,
		Status:        "Pending",
	}
	paymentID, err := CreatePayment(payment)
	assert.NoError(t, err)

	// Update the payment's status to Success
	updatedPayment := models.Payment{
		PlayerID:      3,
		Method:        "BankTransfer",
		Amount:        200.00,
		Details:       `{"bank_account":"123456789","bank_code":"001"}`,
		Status:        "Success",
		TransactionID: "BT9876543210",
	}
	err = UpdatePayment(updatedPayment)
	assert.NoError(t, err)

	// Retrieve and verify the update
	retrievedPayment, err := GetPaymentByID(paymentID)
	assert.NoError(t, err)
	assert.Equal(t, "Success", retrievedPayment.Status)
	assert.Equal(t, "BT9876543210", retrievedPayment.TransactionID)
}