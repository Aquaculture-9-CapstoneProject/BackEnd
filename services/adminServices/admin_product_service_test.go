package adminservices

import (
	"testing"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock untuk AdminProductRepoInterface
type MockAdminProductRepo struct {
	mock.Mock
}

func (m *MockAdminProductRepo) CreateAdminProduct(product *entities.Product) (*entities.Product, error) {
	args := m.Called(product)
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (m *MockAdminProductRepo) UpdateAdminProduct(product *entities.Product) (*entities.Product, error) {
	args := m.Called(product)
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (m *MockAdminProductRepo) DeleteAdminProduct(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockAdminProductRepo) FindByAdminProductID(id int) (*entities.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (m *MockAdminProductRepo) GetAdminProductCount() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAdminProductRepo) GetAllAdminProducts(page int, limit int) ([]entities.Product, error) {
	args := m.Called(page, limit)
	return args.Get(0).([]entities.Product), args.Error(1)
}

func (m *MockAdminProductRepo) SearchAdminProducts(nama string, kategori string, page int, limit int) ([]entities.Product, error) {
	args := m.Called(nama, kategori, page, limit)
	return args.Get(0).([]entities.Product), args.Error(1)
}

func TestAdminProductService(t *testing.T) {
	mockRepo := new(MockAdminProductRepo)
	service := NewAdminProductService(mockRepo)

	// Test CreateAdminProduct
	product := &entities.Product{ID: 1, Nama: "Produk A"}
	mockRepo.On("CreateAdminProduct", product).Return(product, nil)

	createdProduct, err := service.CreateAdminProduct(product)
	assert.NoError(t, err)
	assert.Equal(t, product, createdProduct)

	// Test UpdateAdminProduct
	product.Nama = "Produk A Updated"
	mockRepo.On("UpdateAdminProduct", product).Return(product, nil)

	updatedProduct, err := service.UpdateAdminProduct(product)
	assert.NoError(t, err)
	assert.Equal(t, product, updatedProduct)

	// Test DeleteAdminProduct
	mockRepo.On("DeleteAdminProduct", product.ID).Return(nil)

	err = service.DeleteAdminProduct(product.ID)
	assert.NoError(t, err)

	// Test FindByAdminProductID
	mockRepo.On("FindByAdminProductID", product.ID).Return(product, nil)

	foundProduct, err := service.FindByAdminProductID(product.ID)
	assert.NoError(t, err)
	assert.Equal(t, product, foundProduct)

	// Test GetAdminProductCount
	mockRepo.On("GetAdminProductCount").Return(int64(1), nil)

	count, err := service.GetAdminProductCount()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)

	// Test GetAllAdminProducts
	mockRepo.On("GetAllAdminProducts", 1, 10).Return([]entities.Product{*product}, nil)

	products, err := service.GetAllAdminProducts(1, 10)
	assert.NoError(t, err)
	assert.Len(t, products, 1)

	// Test SearchAdminProducts
	mockRepo.On("SearchAdminProducts", "Produk A", "", 1, 10).Return([]entities.Product{*product}, nil)

	searchResults, err := service.SearchAdminProducts("Produk A", "", 1, 10)
	assert.NoError(t, err)
	assert.Len(t, searchResults, 1)

	// Verifikasi bahwa semua metode mock dipanggil
	mockRepo.AssertExpectations(t)
}
