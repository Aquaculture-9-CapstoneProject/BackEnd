package services

import (
	"testing"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProductFilterRepo adalah mock untuk interface ProductFilterRepo
type MockProductFilterRepo struct {
	mock.Mock
}

func (m *MockProductFilterRepo) CariProduct(kategori string, cari string) ([]entities.Product, error) {
	args := m.Called(kategori, cari)
	return args.Get(0).([]entities.Product), args.Error(1)
}

func TestProductFilterService_CariProdukFilter(t *testing.T) {
	// Membuat mock repository
	mockRepo := new(MockProductFilterRepo)

	// Sample data untuk pengujian
	mockProducts := []entities.Product{
		{ID: 1, Nama: "Produk A", Kategori: "Ikan"},
		{ID: 2, Nama: "Produk B", Kategori: "Udang"},
	}

	// Menyiapkan mock untuk CariProduct
	mockRepo.On("CariProduct", "Ikan", "Udang").Return(mockProducts, nil)

	// Membuat ProductFilterService dengan mock repository
	productFilterService := NewProductFilterService(mockRepo)

	// Test CariProdukFilter
	t.Run("CariProdukFilter", func(t *testing.T) {
		// Memanggil metode CariProdukFilter
		result, err := productFilterService.CariProdukFilter("Ikan", "Udang")

		// Validasi hasil
		assert.Nil(t, err)       // Pastikan tidak ada error
		assert.NotNil(t, result) // Pastikan hasil tidak nil
		assert.Len(t, result, 2) // Pastikan ada 2 produk dalam hasil

		// Memastikan bahwa mock repository dipanggil dengan benar
		mockRepo.AssertExpectations(t)
	})
}
