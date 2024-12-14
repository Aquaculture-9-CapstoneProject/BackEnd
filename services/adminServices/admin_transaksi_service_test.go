package adminservices

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository untuk AdminTransaksiRepo
type MockAdminTransaksiRepo struct {
	mock.Mock
}

func (m *MockAdminTransaksiRepo) GetPaymentDetails(page, perPage int) ([]map[string]interface{}, int64, error) {
	args := m.Called(page, perPage)
	// Jika error, kembalikan slice kosong dan error yang sesuai
	if args.Get(2) != nil {
		return nil, 0, args.Error(2)
	}
	return args.Get(0).([]map[string]interface{}), args.Get(1).(int64), nil
}

func (m *MockAdminTransaksiRepo) DeletePaymentByID(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Unit test untuk AdminTransaksiService.GetPaymentDetails error
func TestAdminTransaksiService_GetPaymentDetails_Error(t *testing.T) {
	mockRepo := new(MockAdminTransaksiRepo)
	service := NewAdminTransaksiServices(mockRepo)

	// Setup ekspektasi ketika terjadi error
	mockRepo.On("GetPaymentDetails", 1, 10).Return(nil, int64(0), errors.New("repository error"))

	// Panggil metode
	result, err := service.GetPaymentDetails(1, 10)

	// Validasi hasil error
	assert.Error(t, err)
	assert.Nil(t, result)

	// Pastikan error yang diterima sesuai dengan yang diharapkan
	assert.Equal(t, "gagal mendapatkan detail pembayaran: repository error", err.Error())
	mockRepo.AssertExpectations(t)
}
