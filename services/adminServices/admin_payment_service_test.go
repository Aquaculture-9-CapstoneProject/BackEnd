package adminservices

import (
	"testing"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository untuk adminPayment
type MockAdminPaymentRepository struct {
	mock.Mock
}

func (m *MockAdminPaymentRepository) GetAdminTotalPendapatanBulanIni() (float64, error) {
	args := m.Called()
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockAdminPaymentRepository) GetAdminJumlahPesananByStatus(status []string) (int64, error) {
	args := m.Called(status)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAdminPaymentRepository) GetTotalProdukInStock() (int, error) {
	args := m.Called()
	return args.Get(0).(int), args.Error(1)
}

func (m *MockAdminPaymentRepository) GetJumlahPesananByStatus(status string) (int64, error) {
	args := m.Called(status)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAdminPaymentRepository) GetJumlahPesananDikirim() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAdminPaymentRepository) GetJumlahPesananDiterima() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAdminPaymentRepository) GetTotalPendapatan() ([]entities.TotalPendapatan, error) {
	args := m.Called()
	return args.Get(0).([]entities.TotalPendapatan), args.Error(1)
}

func (m *MockAdminPaymentRepository) GetJumlahArtikel() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAdminPaymentRepository) GetProdukDenganKategoriStokTerbanyak() ([]entities.Product, error) {
	args := m.Called()
	return args.Get(0).([]entities.Product), args.Error(1)
}

// Unit test untuk AdminPaymentService
func TestAdminPaymentService_HitungAdminPendapatanBulanIni(t *testing.T) {
	mockRepo := new(MockAdminPaymentRepository)
	service := NewAdminPaymentService(mockRepo)

	// Setup ekspektasi
	mockRepo.On("GetAdminTotalPendapatanBulanIni").Return(5000.0, nil)

	// Panggil metode
	result, err := service.HitungAdminPendapatanBulanIni()

	// Validasi hasil
	assert.NoError(t, err)
	assert.Equal(t, 5000.0, result)
	mockRepo.AssertExpectations(t)
}

func TestAdminPaymentService_GetAdminJumlahPesananBulanIni(t *testing.T) {
	mockRepo := new(MockAdminPaymentRepository)
	service := NewAdminPaymentService(mockRepo)

	// Setup ekspektasi
	status := []string{"pending", "completed"}
	mockRepo.On("GetAdminJumlahPesananByStatus", status).Return(int64(100), nil)

	// Panggil metode
	result, err := service.GetAdminJumlahPesananBulanIni(status)

	// Validasi hasil
	assert.NoError(t, err)
	assert.Equal(t, int64(100), result)
	mockRepo.AssertExpectations(t)
}

func TestAdminPaymentService_GetTotalProduk(t *testing.T) {
	mockRepo := new(MockAdminPaymentRepository)
	service := NewAdminPaymentService(mockRepo)

	// Setup ekspektasi
	mockRepo.On("GetTotalProdukInStock").Return(50, nil)

	// Panggil metode
	result, err := service.GetTotalProduk()

	// Validasi hasil
	assert.NoError(t, err)
	assert.Equal(t, 50, result)
	mockRepo.AssertExpectations(t)
}

func TestAdminPaymentService_GetJumlahPesananDikirim(t *testing.T) {
	mockRepo := new(MockAdminPaymentRepository)
	service := NewAdminPaymentService(mockRepo)

	// Setup ekspektasi
	mockRepo.On("GetJumlahPesananDikirim").Return(int64(25), nil)

	// Panggil metode
	result, err := service.GetJumlahPesananDikirim()

	// Validasi hasil
	assert.NoError(t, err)
	assert.Equal(t, int64(25), result)
	mockRepo.AssertExpectations(t)
}

func TestAdminPaymentService_GetProdukDenganKategoriStokTerbanyak(t *testing.T) {
	mockRepo := new(MockAdminPaymentRepository)
	service := NewAdminPaymentService(mockRepo)

	// Setup ekspektasi
	products := []entities.Product{
		{ID: 1, Nama: "Produk A", Kategori: "Category 1", Stok: 100},
		{ID: 2, Nama: "Produk B", Kategori: "Category 1", Stok: 200},
	}
	mockRepo.On("GetProdukDenganKategoriStokTerbanyak").Return(products, nil)

	// Panggil metode
	result, err := service.GetProdukDenganKategoriStokTerbanyak()

	// Validasi hasil
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Produk B", result[1].Nama)
	mockRepo.AssertExpectations(t)
}
