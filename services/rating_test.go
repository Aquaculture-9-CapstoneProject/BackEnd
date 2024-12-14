package services_test

import (
	"errors"
	"testing"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRatingRepo struct {
	mock.Mock
}

type MockProductDetailServices struct {
	mock.Mock
}

func (m *MockRatingRepo) AddReview(review *entities.Review) error {
	args := m.Called(review)
	return args.Error(0)
}

func (m *MockRatingRepo) GetReviewByUserAndProduct(userID, productID int) (*entities.Review, error) {
	args := m.Called(userID, productID)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Review), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRatingRepo) GetReviewsByProductID(productID int) ([]entities.Review, error) {
	args := m.Called(productID)
	if args.Get(0) != nil {
		return args.Get(0).([]entities.Review), args.Error(1)
	}
	return nil, args.Error(1)
}

type MockProductDetailService struct {
	mock.Mock
}

func (m *MockProductDetailService) UpdateProductRating(productID int) error {
	args := m.Called(productID)
	return args.Error(0)
}

func TestReviewServices_AddReview(t *testing.T) {
	mockRepo := new(MockRatingRepo)
	mockDetailService := new(MockProductDetailService)
	service := services.NewServiceRating(mockRepo, mockDetailService)

	t.Run("success case", func(t *testing.T) {
		review := &entities.Review{
			UserID:    1,
			ProductID: 1,
			Rating:    4.5,
			Ulasan:    "Great product!",
		}
		mockRepo.On("AddReview", review).Return(nil)
		mockDetailService.On("UpdateProductRating", 1).Return(nil)

		err := service.AddReview(1, 1, 4.5, "Great product!")

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
		mockDetailService.AssertExpectations(t)
	})

	t.Run("error adding review", func(t *testing.T) {
		review := &entities.Review{
			UserID:    1,
			ProductID: 1,
			Rating:    4.5,
			Ulasan:    "Great product!",
		}
		mockRepo.On("AddReview", review).Return(errors.New("db error"))

		err := service.AddReview(1, 1, 4.5, "Great product!")

		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("error updating product rating", func(t *testing.T) {
		review := &entities.Review{
			UserID:    1,
			ProductID: 1,
			Rating:    4.5,
			Ulasan:    "Great product!",
		}
		mockRepo.On("AddReview", review).Return(nil)
		mockDetailService.On("UpdateProductRating", 1).Return(errors.New("update error"))

		err := service.AddReview(1, 1, 4.5, "Great product!")

		assert.Error(t, err)
		assert.Equal(t, "update error", err.Error())
		mockRepo.AssertExpectations(t)
		mockDetailService.AssertExpectations(t)
	})
}

func TestReviewServices_GetUserReview(t *testing.T) {
	mockRepo := new(MockRatingRepo)
	service := services.NewServiceRating(mockRepo, nil)

	t.Run("success case", func(t *testing.T) {
		expectedReview := &entities.Review{
			UserID:    1,
			ProductID: 1,
			Rating:    4.5,
			Ulasan:    "Great product!",
		}
		mockRepo.On("GetReviewByUserAndProduct", 1, 1).Return(expectedReview, nil)

		review, err := service.GetUserReview(1, 1)

		assert.NoError(t, err)
		assert.NotNil(t, review)
		assert.Equal(t, expectedReview, review)
		mockRepo.AssertExpectations(t)
	})

	t.Run("review not found", func(t *testing.T) {
		mockRepo.On("GetReviewByUserAndProduct", 1, 1).Return(nil, errors.New("review not found"))

		review, err := service.GetUserReview(1, 1)

		assert.Error(t, err)
		assert.Nil(t, review)
		assert.Equal(t, "review not found", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestReviewServices_GetReviewsByProduct(t *testing.T) {
	mockRepo := new(MockRatingRepo)
	service := services.NewServiceRating(mockRepo, nil)

	t.Run("success case", func(t *testing.T) {
		expectedReviews := []entities.Review{
			{UserID: 1, ProductID: 1, Rating: 4.5, Ulasan: "Great product!"},
			{UserID: 2, ProductID: 1, Rating: 5.0, Ulasan: "Excellent!"},
		}
		mockRepo.On("GetReviewsByProductID", 1).Return(expectedReviews, nil)

		reviews, err := service.GetReviewsByProduct(1)

		assert.NoError(t, err)
		assert.NotNil(t, reviews)
		assert.Equal(t, expectedReviews, reviews)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error fetching reviews", func(t *testing.T) {
		mockRepo.On("GetReviewsByProductID", 1).Return(nil, errors.New("db error"))

		reviews, err := service.GetReviewsByProduct(1)

		assert.Error(t, err)
		assert.Nil(t, reviews)
		assert.Equal(t, "db error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
