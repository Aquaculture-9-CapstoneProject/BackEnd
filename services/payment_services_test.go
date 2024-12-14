package services

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository for PaymentsRepo
type MockPaymentsRepo struct {
	mock.Mock
}

func (m *MockPaymentsRepo) GetOrderTotalAmount(orderID int) (float64, error) {
	args := m.Called(orderID)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockPaymentsRepo) SavePaymentStatus(invoiceID, status string, orderID int, amount float64, invoiceURL string) (*entities.Payment, error) {
	args := m.Called(invoiceID, status, orderID, amount, invoiceURL)
	return args.Get(0).(*entities.Payment), args.Error(1)
}

func (m *MockPaymentsRepo) GetPaymentStatus(invoiceID string) (string, error) {
	args := m.Called(invoiceID)
	return args.String(0), args.Error(1)
}

func (m *MockPaymentsRepo) UpdatePaymentStatus(invoiceID, status string) error {
	args := m.Called(invoiceID, status)
	return args.Error(0)
}

func (m *MockPaymentsRepo) UpdateBarangStatusAsync(invoiceID string) error {
	args := m.Called(invoiceID)
	return args.Error(0)
}

func (m *MockPaymentsRepo) CancelOrder(invoiceID string) error {
	args := m.Called(invoiceID)
	return args.Error(0)
}

func (m *MockPaymentsRepo) GetPaymentByInvoiceID(invoiceID string) (*entities.Payment, error) {
	args := m.Called(invoiceID)
	return args.Get(0).(*entities.Payment), args.Error(1)
}

func (m *MockPaymentsRepo) GetPaymentsByUserID(userID int) ([]entities.Payment, error) {
	args := m.Called(userID)
	return args.Get(0).([]entities.Payment), args.Error(1)
}

func TestPaymentServices_CreateInvoice(t *testing.T) {
	// Set the XENDIT_SECRET_KEY for testing
	os.Setenv("XENDIT_SECRET_KEY", "xnd_development_eYRvZS9CD63DN6gfokhdUcDomCY66UUHMBn8sPnkTrARJTqbUIDYaMnuAIRkd")

	// Create the mock repository and payment service
	mockRepo := new(MockPaymentsRepo)
	service := NewPaymentServices(mockRepo)

	// Test setup
	orderID := 1
	expectedAmount := 1000.0
	expectedInvoicePrefix := "675dab0943bd94e64350a548" // Ensure the prefix is correct
	expectedInvoiceID := expectedInvoicePrefix          // Expected invoice ID

	// Return the payment object with the expected InvoiceID
	expectedPayment := &entities.Payment{
		InvoiceID:  expectedInvoiceID,
		Status:     "PENDING",
		OrderID:    orderID,
		Jumlah:     expectedAmount,
		PaymentUrl: "https://checkout-staging.xendit.co/",
	}

	// Setup mock expectations
	mockRepo.On("GetOrderTotalAmount", orderID).Return(expectedAmount, nil)
	mockRepo.On("SavePaymentStatus", mock.AnythingOfType("string"), "PENDING", orderID, expectedAmount, mock.MatchedBy(func(url string) bool {
		return strings.HasPrefix(url, "https://checkout-staging.xendit.co/")
	})).Return(expectedPayment, nil)

	// Call the service method
	result, err := service.CreateInvoice(orderID)

	// Assert expectations
	assert.Nil(t, err)
	assert.NotNil(t, result)

	// Type assert the invoice_id to string and check the prefix
	invoiceID, ok := result["invoice_id"].(string)
	if !ok {
		t.Fatal("invoice_id is not a string")
	}
	fmt.Println("Generated invoiceID:", invoiceID) // Debugging line

	// Check if the invoiceID starts with the expected prefix
	assert.True(t, strings.HasPrefix(invoiceID, "675dad"))
	assert.Equal(t, "order_1", result["external_id"])

	// Ensure the mock expectations were met
	mockRepo.AssertExpectations(t)
}
