// handlers/payments.go
package handlers

import (
	"fmt"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"interview_YangYang_20241010/models"
	"interview_YangYang_20241010/repository"

	"github.com/gin-gonic/gin"
)

// PaymentRequest represents the request body for creating a new payment.
type PaymentRequest struct {
	PlayerID uint          `json:"player_id" binding:"required"`
	Method   string        `json:"method" binding:"required"` // e.g., CreditCard, BankTransfer, ThirdParty, Blockchain
	Amount   float64       `json:"amount" binding:"required,gt=0"`
	Details  json.RawMessage `json:"details" binding:"required"` // Specific details based on payment method
}

// PaymentResponse represents the response after processing a payment.
type PaymentResponse struct {
	Status        string `json:"status"`
	TransactionID string `json:"transaction_id,omitempty"`
	ErrorMessage  string `json:"error_message,omitempty"`
}

// @Summary Process a Payment
// @Description Process a payment using various payment methods.
// @Tags Payments
// @Accept json
// @Produce json
// @Param payment body PaymentRequest true "Payment Information"
// @Success 200 {object} models.Payment "Payment Status"
// @Failure 400 {object} models.Payment "Bad Request"
// @Failure 500 {object} models.Payment "Internal Server Error"
// @Router /payments [post]
func ProcessPayment(c *gin.Context) {
	var req PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, PaymentResponse{ErrorMessage: err.Error()})
		return
	}

	// Validate payment method
	switch req.Method {
	case "CreditCard", "BankTransfer", "ThirdParty", "Blockchain":
		// Valid methods
	default:
		c.JSON(http.StatusBadRequest, PaymentResponse{ErrorMessage: "Invalid payment method"})
		return
	}

	// Create a new payment record
	payment := models.Payment{
		PlayerID: req.PlayerID,
		Method:   req.Method,
		Amount:   req.Amount,
		Details:  string(req.Details),
		Status:   "Pending",
	}

	paymentID, err := repository.CreatePayment(payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, PaymentResponse{ErrorMessage: "Failed to create payment record"})
		return
	}

	// Simulate payment processing asynchronously
	go handlePaymentProcessing(paymentID)

	// Respond with payment status
	c.JSON(http.StatusOK, PaymentResponse{
		Status: "Payment processing initiated",
	})
}

// @Summary Get Payment Details
// @Description Retrieve detailed information about a specific payment.
// @Tags Payments
// @Accept json
// @Produce json
// @Param id path uint true "Payment ID"
// @Success 200 {object} models.Payment "Payment Details"
// @Failure 404 {object} models.ErrorResponse "Payment Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /payments/{id} [get]
func GetPaymentDetails(c *gin.Context) {
	idParam := c.Param("id")
	id, err := parseUint(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	payment, err := repository.GetPaymentByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrPaymentNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve payment"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

// parseUint converts a string to uint, handling errors.
func parseUint(s string) (uint, error) {
	var i uint64
	_, err := fmt.Sscanf(s, "%d", &i)
	return uint(i), err
}

// handlePaymentProcessing simulates payment processing based on the payment method.
func handlePaymentProcessing(paymentID uint) {
	// Retrieve the payment record
	payment, err := repository.GetPaymentByID(paymentID)
	if err != nil {
		// Handle error (logging can be added here)
		return
	}

	// Simulate different payment methods
	var transactionID string
	var status string
	var errorMessage string

	switch payment.Method {
	case "CreditCard":
		transactionID, status, errorMessage = processCreditCardPayment(payment)
	case "BankTransfer":
		transactionID, status, errorMessage = processBankTransfer(payment)
	case "ThirdParty":
		transactionID, status, errorMessage = processThirdPartyPayment(payment)
	case "Blockchain":
		transactionID, status, errorMessage = processBlockchainPayment(payment)
	default:
		status = "Failed"
		errorMessage = "Unsupported payment method"
	}

	// Update the payment record based on processing result
	if status == "Success" {
		payment.Status = "Success"
		payment.TransactionID = transactionID
	} else {
		payment.Status = "Failed"
		payment.ErrorMessage = errorMessage
	}

	repository.UpdatePayment(*payment)
}

// Simulated payment processing functions

func processCreditCardPayment(payment *models.Payment) (string, string, string) {
	// Simulate processing delay
	time.Sleep(2 * time.Second)

	// Simulate success
	return "CC1234567890", "Success", ""
}

func processBankTransfer(payment *models.Payment) (string, string, string) {
	// Simulate processing delay
	time.Sleep(3 * time.Second)

	// Simulate failure
	return "", "Failed", "Insufficient funds"
}

func processThirdPartyPayment(payment *models.Payment) (string, string, string) {
	// Simulate processing delay
	time.Sleep(1 * time.Second)

	// Simulate success
	return "TP0987654321", "Success", ""
}

func processBlockchainPayment(payment *models.Payment) (string, string, string) {
	// Simulate processing delay
	time.Sleep(4 * time.Second)

	// Simulate success
	return "BC5678901234", "Success", ""
}
