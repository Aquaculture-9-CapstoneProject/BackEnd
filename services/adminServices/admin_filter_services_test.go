package adminservices

import (
	"fmt"
	"testing"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository untuk AdminFilterRepo
type MockAdminFilterRepo struct {
	mock.Mock
}

func (m *MockAdminFilterRepo) GetPaymentsByStatus(status string) ([]entities.Payment, error) {
	args := m.Called(status)
	// Mengembalikan nil dan error jika terjadi error
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Payment), args.Error(1)
}

func (m *MockAdminFilterRepo) GetPaymentsByStatusBarang(statusBarang string) ([]entities.Payment, error) {
	args := m.Called(statusBarang)
	// Mengembalikan nil dan error jika terjadi error
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Payment), args.Error(1)
}

// Unit test untuk AdminFilterServices
func TestAdminFilterServices_GetPaymentsByStatus(t *testing.T) {
	mockRepo := new(MockAdminFilterRepo)
	service := NewAdminFilterServices(mockRepo)

	// Setup ekspektasi
	status := "paid"
	mockRepo.On("GetPaymentsByStatus", status).Return([]entities.Payment{
		{ID: 1, Status: "paid", Jumlah: 1000},
		{ID: 2, Status: "paid", Jumlah: 2000},
	}, nil)

	// Panggil metode
	result, err := service.GetPaymentsByStatus(status)

	// Validasi hasil
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "paid", result[0].Status)
	mockRepo.AssertExpectations(t)
}

func TestAdminFilterServices_GetPaymentsByStatusBarang(t *testing.T) {
	mockRepo := new(MockAdminFilterRepo)
	service := NewAdminFilterServices(mockRepo)

	// Setup ekspektasi
	statusBarang := "in_stock"
	mockRepo.On("GetPaymentsByStatusBarang", statusBarang).Return([]entities.Payment{
		{ID: 1, StatusBarang: "in_stock", Jumlah: 500},
		{ID: 2, StatusBarang: "in_stock", Jumlah: 700},
	}, nil)

	// Panggil metode
	result, err := service.GetPaymentsByStatusBarang(statusBarang)

	// Validasi hasil
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "in_stock", result[0].StatusBarang)
	mockRepo.AssertExpectations(t)
}

func TestAdminFilterServices_GetPaymentsByStatus_Error(t *testing.T) {
	mockRepo := new(MockAdminFilterRepo)
	service := NewAdminFilterServices(mockRepo)

	// Setup ekspektasi untuk error
	status := "failed"
	mockRepo.On("GetPaymentsByStatus", status).Return(nil, fmt.Errorf("database error"))

	// Panggil metode
	result, err := service.GetPaymentsByStatus(status)

	// Validasi hasil
	assert.Error(t, err)
	assert.Nil(t, result) // Memastikan hasil adalah nil
	mockRepo.AssertExpectations(t)
}

func TestAdminFilterServices_GetPaymentsByStatusBarang_Error(t *testing.T) {
	mockRepo := new(MockAdminFilterRepo)
	service := NewAdminFilterServices(mockRepo)

	// Setup ekspektasi untuk error
	statusBarang := "out_of_stock"
	mockRepo.On("GetPaymentsByStatusBarang", statusBarang).Return(nil, fmt.Errorf("database error"))

	// Panggil metode
	result, err := service.GetPaymentsByStatusBarang(statusBarang)

	// Validasi hasil
	assert.Error(t, err)
	assert.Nil(t, result) // Memastikan hasil adalah nil
	mockRepo.AssertExpectations(t)
}

func TestAdminFilterServices_GetPaymentsByStatus_EmptyResult(t *testing.T) {
	mockRepo := new(MockAdminFilterRepo)
	service := NewAdminFilterServices(mockRepo)

	// Setup ekspektasi untuk hasil kosong
	status := "unpaid"
	mockRepo.On("GetPaymentsByStatus", status).Return([]entities.Payment{}, nil)

	// Panggil metode
	result, err := service.GetPaymentsByStatus(status)

	// Validasi hasil
	assert.NoError(t, err)
	assert.Len(t, result, 0)
	mockRepo.AssertExpectations(t)
}

func TestAdminFilterServices_GetPaymentsByStatusBarang_EmptyResult(t *testing.T) {
	mockRepo := new(MockAdminFilterRepo)
	service := NewAdminFilterServices(mockRepo)

	// Setup ekspektasi untuk hasil kosong
	statusBarang := "out_of_stock"
	mockRepo.On("GetPaymentsByStatusBarang", statusBarang).Return([]entities.Payment{}, nil)

	// Panggil metode
	result, err := service.GetPaymentsByStatusBarang(statusBarang)

	// Validasi hasil
	assert.NoError(t, err)
	assert.Len(t, result, 0)
	mockRepo.AssertExpectations(t)
}

func TestAdminFilterServices_GetPaymentsByStatus_EmptyInput(t *testing.T) {
	mockRepo := new(MockAdminFilterRepo)
	service := NewAdminFilterServices(mockRepo)

	// Setup ekspektasi untuk input kosong
	status := ""
	mockRepo.On("GetPaymentsByStatus", status).Return(nil, fmt.Errorf("invalid input"))

	// Panggil metode
	result, err := service.GetPaymentsByStatus(status)

	// Validasi hasil
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestAdminFilterServices_GetPaymentsByStatusBarang_EmptyInput(t *testing.T) {
	mockRepo := new(MockAdminFilterRepo)
	service := NewAdminFilterServices(mockRepo)

	// Setup ekspektasi untuk input kosong
	statusBarang := ""
	mockRepo.On("GetPaymentsByStatusBarang", statusBarang).Return(nil, fmt.Errorf("invalid input"))

	// Panggil metode
	result, err := service.GetPaymentsByStatusBarang(statusBarang)

	// Validasi hasil
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
