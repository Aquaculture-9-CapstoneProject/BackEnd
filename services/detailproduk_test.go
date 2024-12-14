package services_test

import (
	"errors"
	"testing"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductDetailRepo struct {
	mock.Mock
}

func (m *MockProductDetailRepo) CekProdukByID(ProductID int) (*entities.Product, error) {
	args := m.Called(ProductID)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockProductDetailRepo) UpdateProductRating(productID int, newRating float64) error {
	args := m.Called(productID, newRating)
	return args.Error(0)
}

func (m *MockProductDetailRepo) UpdateTotalReview(productID int, totalReviews int) error {
	args := m.Called(productID, totalReviews)
	return args.Error(0)
}

func (m *MockRatingRepo) CountReviewsByProduct(productID int) (int64, error) {
	args := m.Called(productID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockRatingRepo) SumRatingByProduct(productID int) (float64, error) {
	args := m.Called(productID)
	return args.Get(0).(float64), args.Error(1)
}

func TestLihatProductByID(t *testing.T) {
	mockProductRepo := new(MockProductDetailRepo)
	mockRatingRepo := new(MockRatingRepo)
	service := services.NewProductDetailServices(mockProductRepo, mockRatingRepo)

	t.Run("success case", func(t *testing.T) {
		mockProduct := &entities.Product{ID: 1, Nama: "Test Product"}
		mockProductRepo.On("CekProdukByID", 1).Return(mockProduct, nil)

		product, err := service.LihatProductByID(1)

		assert.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, 1, product.ID)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("product not found", func(t *testing.T) {
		mockProductRepo.On("CekProdukByID", 1).Return(nil, errors.New("product not found"))

		product, err := service.LihatProductByID(1)

		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Equal(t, "product not found", err.Error())
		mockProductRepo.AssertExpectations(t)
	})
}

func TestUpdateProductRating(t *testing.T) {
	mockProductRepo := new(MockProductDetailRepo)
	mockRatingRepo := new(MockRatingRepo)
	service := services.NewProductDetailServices(mockProductRepo, mockRatingRepo)

	t.Run("success case", func(t *testing.T) {
		mockRatingRepo.On("CountReviewsByProduct", 1).Return(int64(5), nil)
		mockRatingRepo.On("SumRatingByProduct", 1).Return(20.0, nil)
		mockProductRepo.On("UpdateProductRating", 1, 4.0).Return(nil)
		mockProductRepo.On("UpdateTotalReview", 1, 5).Return(nil)

		err := service.UpdateProductRating(1)

		assert.NoError(t, err)
		mockRatingRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("error in CountReviewsByProduct", func(t *testing.T) {
		mockRatingRepo.On("CountReviewsByProduct", 1).Return(int64(0), errors.New("db error"))

		err := service.UpdateProductRating(1)

		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
		mockRatingRepo.AssertExpectations(t)
	})

	t.Run("error in SumRatingByProduct", func(t *testing.T) {
		mockRatingRepo.On("CountReviewsByProduct", 1).Return(int64(5), nil)
		mockRatingRepo.On("SumRatingByProduct", 1).Return(0.0, errors.New("db error"))

		err := service.UpdateProductRating(1)

		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
		mockRatingRepo.AssertExpectations(t)
	})

	t.Run("error in UpdateProductRating", func(t *testing.T) {
		mockRatingRepo.On("CountReviewsByProduct", 1).Return(int64(5), nil)
		mockRatingRepo.On("SumRatingByProduct", 1).Return(20.0, nil)
		mockProductRepo.On("UpdateProductRating", 1, 4.0).Return(errors.New("update error"))

		err := service.UpdateProductRating(1)

		assert.Error(t, err)
		assert.Equal(t, "update error", err.Error())
		mockRatingRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
	})
}
