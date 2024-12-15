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

// Mock GetDetailedOrders dengan pagination
func (m *MockAdminPesananRepo) GetDetailedOrders(page, perPage int) ([]map[string]interface{}, int64, error) {
	args := m.Called(page, perPage)
	return args.Get(0).([]map[string]interface{}), args.Get(1).(int64), args.Error(2)
}

func TestAdminPesananServices_GetDetailedOrders_EmptyData(t *testing.T) {
	mockRepo := new(MockAdminPesananRepo)
	service := NewPesananServices(mockRepo)
	page := 1
	perPage := 10

	// Setup ekspektasi jika data kosong
	mockRepo.On("GetDetailedOrders", page, perPage).Return([]map[string]interface{}{}, int64(0), nil)

	// Panggil metode
	result, totalItems, err := service.GetDetailedOrders(page, perPage)

	// Validasi hasil
	assert.NoError(t, err)
	assert.Empty(t, result)
	assert.Equal(t, int64(0), totalItems)
	mockRepo.AssertExpectations(t)
}

func TestAdminPesananServices_GetDetailedOrders_StructValidation(t *testing.T) {
	mockRepo := new(MockAdminPesananRepo)
	service := NewPesananServices(mockRepo)
	page := 1
	perPage := 10

	// Setup ekspektasi jika data berhasil diambil
	mockRepo.On("GetDetailedOrders", page, perPage).Return([]map[string]interface{}{
		{"order_id": 1, "product_name": "Ikan nila", "status": "dikirim"},
		{"order_id": 2, "product_name": "ikan salmon", "status": "diterima"},
	}, int64(2), nil)

	// Panggil metode
	result, totalItems, err := service.GetDetailedOrders(page, perPage)

	// Validasi hasil
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, int64(2), totalItems)

	// Memeriksa apakah struktur data yang dikembalikan sesuai
	assert.Contains(t, result[0], "order_id")
	assert.Contains(t, result[0], "product_name")
	assert.Contains(t, result[0], "status")

	assert.Equal(t, "Ikan nila", result[0]["product_name"])
	assert.Equal(t, 1, result[0]["order_id"])
	assert.Equal(t, "dikirim", result[0]["status"])

	mockRepo.AssertExpectations(t)
}

func TestAdminPesananServices_GetDetailedOrders_Error(t *testing.T) {
	mockRepo := new(MockAdminPesananRepo)
	service := NewPesananServices(mockRepo)
	page := 1
	perPage := 10

	// Setup ekspektasi jika terjadi error di repository
	mockRepo.On("GetDetailedOrders", page, perPage).Return([]map[string]interface{}{}, int64(0), errors.New("repository error"))

	// Panggil metode
	result, totalItems, err := service.GetDetailedOrders(page, perPage)

	// Validasi hasil error
	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Equal(t, int64(0), totalItems)
	assert.Equal(t, "gagal mendapatkan detail pesanan: repository error", err.Error())
	mockRepo.AssertExpectations(t)
}
