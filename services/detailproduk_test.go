package services

import (
	"testing"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/stretchr/testify/mock"
)

// MockRatingRepo for RatingRepo
type MockRatingRepo struct {
	mock.Mock
}

func (m *MockRatingRepo) CountReviewsByProduct(ProductID int) (int, error) {
	args := m.Called(ProductID)
	return args.Int(0), args.Error(1)
}

func (m *MockRatingRepo) SumRatingByProduct(ProductID int) (float64, error) {
	args := m.Called(ProductID)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockRatingRepo) AddReview(review *entities.Review) error {
	args := m.Called(review)
	return args.Error(0)
}

func (m *MockRatingRepo) GetReviewByUserAndProduct(userID int, productID int) (*entities.Review, error) {
	args := m.Called(userID, productID)
	return args.Get(0).(*entities.Review), args.Error(1)
}

// Modify GetReviewsByProductID to return []entities.Review instead of []*entities.Review
func (m *MockRatingRepo) GetReviewsByProductID(ProductID int) ([]entities.Review, error) {
	args := m.Called(ProductID)
	return args.Get(0).([]entities.Review), args.Error(1)
}

// MockProductDetailRepo for ProductDetailRepo
type MockProductDetailRepo struct {
	mock.Mock
}

func (m *MockProductDetailRepo) CekProdukByID(ProductID int) (*entities.Product, error) {
	args := m.Called(ProductID)
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (m *MockProductDetailRepo) UpdateProductRating(productID int, rating float64) error {
	args := m.Called(productID, rating)
	return args.Error(0)
}

func (m *MockProductDetailRepo) UpdateTotalReview(productID int, totalReviews int) error {
	args := m.Called(productID, totalReviews)
	return args.Error(0)
}

// TestProductDetailServices
func TestProductDetailServices(t *testing.T) {
	// Create mocks
	mockReviewRepo := new(MockRatingRepo)
	mockProductRepo := new(MockProductDetailRepo)

	// Setup expectations
	mockReviewRepo.On("CountReviewsByProduct", 1).Return(10, nil)
	mockReviewRepo.On("SumRatingByProduct", 1).Return(50.0, nil)

	mockProductRepo.On("CekProdukByID", 1).Return(&entities.Product{ID: 1, Nama: "Product 1"}, nil)
	mockProductRepo.On("UpdateProductRating", 1, 5.0).Return(nil)
	mockProductRepo.On("UpdateTotalReview", 1, 10).Return(nil)

	// Create service with mocks
	service := NewProductDetailServices(mockProductRepo, mockReviewRepo)

	// Test LihatProductByID
	product, err := service.LihatProductByID(1)
	if err != nil {
		t.Fatal(err)
	}
	if product.ID != 1 {
		t.Errorf("Expected product ID 1, got %d", product.ID)
	}

	// Test UpdateProductRating
	err = service.UpdateProductRating(1)
	if err != nil {
		t.Fatal(err)
	}

	// Assert expectations
	mockReviewRepo.AssertExpectations(t)
	mockProductRepo.AssertExpectations(t)
}
