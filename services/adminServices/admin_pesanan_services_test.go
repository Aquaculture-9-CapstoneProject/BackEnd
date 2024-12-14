package adminservices

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository untuk AdminPesananRepo
type MockAdminPesananRepo struct {
	mock.Mock
}

func (m *MockAdminPesananRepo) GetDetailedOrders() ([]map[string]interface{}, error) {
	args := m.Called()
	// Pastikan nilai yang dikembalikan sesuai dengan ekspektasi.
	return args.Get(0).([]map[string]interface{}), args.Error(1)
}

// Unit test untuk AdminPesananServices.GetDetailedOrders ketika data kosong
func TestAdminPesananServices_GetDetailedOrders_EmptyData(t *testing.T) {
	mockRepo := new(MockAdminPesananRepo)
	service := NewPesananServices(mockRepo)

	// Setup ekspektasi jika data kosong
	mockRepo.On("GetDetailedOrders").Return([]map[string]interface{}{}, nil)

	// Panggil metode
	result, err := service.GetDetailedOrders()

	// Validasi hasil
	assert.NoError(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

// Unit test untuk AdminPesananServices.GetDetailedOrders ketika data dengan struktur yang benar
func TestAdminPesananServices_GetDetailedOrders_StructValidation(t *testing.T) {
	mockRepo := new(MockAdminPesananRepo)
	service := NewPesananServices(mockRepo)

	// Setup ekspektasi jika berhasil
	mockRepo.On("GetDetailedOrders").Return([]map[string]interface{}{
		{"order_id": 1, "product_name": "Produk A", "status": "dikirim"},
		{"order_id": 2, "product_name": "Produk B", "status": "diterima"},
	}, nil)

	// Panggil metode
	result, err := service.GetDetailedOrders()

	// Validasi hasil
	assert.NoError(t, err)
	assert.Len(t, result, 2)

	// Memeriksa apakah struktur data yang dikembalikan sesuai
	assert.Contains(t, result[0], "order_id")
	assert.Contains(t, result[0], "product_name")
	assert.Contains(t, result[0], "status")

	assert.Equal(t, "Produk A", result[0]["product_name"])
	assert.Equal(t, 1, result[0]["order_id"])
	assert.Equal(t, "dikirim", result[0]["status"])

	mockRepo.AssertExpectations(t)
}

// Unit test untuk AdminPesananServices.GetDetailedOrders ketika ada error
func TestAdminPesananServices_GetDetailedOrders_Error(t *testing.T) {
	mockRepo := new(MockAdminPesananRepo)
	service := NewPesananServices(mockRepo)

	// Setup ekspektasi jika gagal
	mockRepo.On("GetDetailedOrders").Return([]map[string]interface{}{}, errors.New("repository error"))

	// Panggil metode
	result, err := service.GetDetailedOrders()

	// Validasi hasil error
	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Equal(t, "gagal mendapatkan detail pesanan: repository error", err.Error())
	mockRepo.AssertExpectations(t)
}
