package adminservices

import (
	"testing"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAdminFilterRepo struct {
	mock.Mock
}

func (m *MockAdminFilterRepo) GetPaymentsByStatus(status string) ([]entities.Payment, error) {
	args := m.Called(status)
	return args.Get(0).([]entities.Payment), args.Error(1)
}

func (m *MockAdminFilterRepo) GetPaymentsByStatusBarang(statusBarang string) ([]entities.Payment, error) {
	args := m.Called(statusBarang)
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
	statusBarang := "dikirim"
	mockRepo.On("GetPaymentsByStatusBarang", statusBarang).Return([]entities.Payment{
		{ID: 1, StatusBarang: "dikirim", Jumlah: 500},
		{ID: 2, StatusBarang: "dikirim", Jumlah: 700},
	}, nil)

	// Panggil metode
	result, err := service.GetPaymentsByStatusBarang(statusBarang)

	// Validasi hasil
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "dikirim", result[0].StatusBarang)
	mockRepo.AssertExpectations(t)
}

func TestAdminFilterServices_GetPaymentsByStatusBarang_Selesai(t *testing.T) {
	mockRepo := new(MockAdminFilterRepo)
	service := NewAdminFilterServices(mockRepo)

	// Setup ekspektasi
	statusBarang := "selesai"
	mockRepo.On("GetPaymentsByStatusBarang", statusBarang).Return([]entities.Payment{
		{ID: 1, StatusBarang: "selesai", Jumlah: 1000},
		{ID: 2, StatusBarang: "selesai", Jumlah: 1500},
	}, nil)

	// Panggil metode
	result, err := service.GetPaymentsByStatusBarang(statusBarang)

	// Validasi hasil
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "selesai", result[0].StatusBarang)
	mockRepo.AssertExpectations(t)
}

func TestAdminFilterServices_GetPaymentsByStatus_Pending(t *testing.T) {
	mockRepo := new(MockAdminFilterRepo)
	service := NewAdminFilterServices(mockRepo)

	// Setup ekspektasi
	status := "pending"
	mockRepo.On("GetPaymentsByStatus", status).Return([]entities.Payment{
		{ID: 1, Status: "pending", Jumlah: 500},
		{ID: 2, Status: "pending", Jumlah: 1000},
	}, nil)

	// Panggil metode
	result, err := service.GetPaymentsByStatus(status)

	// Validasi hasil
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "pending", result[0].Status)
	mockRepo.AssertExpectations(t)
}
