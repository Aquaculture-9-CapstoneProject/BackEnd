package services_test

import (
	"testing"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProdukIkanRepo struct {
	mock.Mock
}

func (m *MockProdukIkanRepo) GetTermurah(limit int) ([]entities.Product, error) {
	args := m.Called(limit)
	return args.Get(0).([]entities.Product), args.Error(1)
}

func (m *MockProdukIkanRepo) GetPopuler(limit int) ([]entities.Product, error) {
	args := m.Called(limit)
	return args.Get(0).([]entities.Product), args.Error(1)
}

func (m *MockProdukIkanRepo) GetSemuaProduk() ([]entities.Product, error) {
	args := m.Called()
	return args.Get(0).([]entities.Product), args.Error(1)
}

func TestProductUseCaseProdukTermurah(t *testing.T) {
	mockRepo := new(MockProdukIkanRepo)
	service := services.NewProductIkanServices(mockRepo)
	t.Run("succes case", func(t *testing.T) {
		mockProducts := []entities.Product{
			{ID: 1, Nama: "Ikan Tuna", Harga: 10000},
			{ID: 2, Nama: "Ikan Tongkol", Harga: 12000},
		}
		mockRepo.On("GetTermurah", 2).Return(mockProducts, nil)
		result, err := service.GetProdukTermurah(2)
		assert.NoError(t, err)
		assert.Equal(t, mockProducts, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetAllProdukPopuler(t *testing.T) {
	mockRepo := new(MockProdukIkanRepo)
	service := services.NewProductIkanServices(mockRepo)

	t.Run("succes case", func(t *testing.T) {
		mockProducts := []entities.Product{
			{ID: 3, Nama: "ikan salmon", Rating: 3},
			{ID: 4, Nama: "ikan kakap", Rating: 4},
		}
		mockRepo.On("GetPopuler", 2).Return(mockProducts, nil)
		result, err := service.GetAllProductPopuler(2)
		assert.NoError(t, err)
		assert.Equal(t, mockProducts, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestTampilkanSemua(t *testing.T) {
	mockRepo := new(MockProdukIkanRepo)
	service := services.NewProductIkanServices(mockRepo)

	t.Run("succes case", func(t *testing.T) {
		mockProduk := []entities.Product{
			{ID: 5, Nama: "ikan salmon", Harga: 10000, Variasi: "spesial"},
			{ID: 6, Nama: "ikan kakap", Harga: 20000, Variasi: "spesial"},
		}
		mockRepo.On("GetSemuaProduk").Return(mockProduk, nil)
		result, err := service.GetAllProduct()
		assert.NoError(t, err)
		assert.Equal(t, mockProduk, result)
		mockRepo.AssertExpectations(t)
	})

}
