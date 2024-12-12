package services

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
)

type ReviewServices interface {
	AddReview(userID, productID int, rating float64, comment string) error
	GetUserReview(userID, productID int) (*entities.Review, error)
	GetReviewsByProduct(productID int) ([]entities.Review, error)
}

type reviewServices struct {
	repoReview   repositories.RatingRepo
	servieDetail ProductDetailServices
}

func NewServiceRating(repoReview repositories.RatingRepo, servieDetail ProductDetailServices) ReviewServices {
	return &reviewServices{repoReview: repoReview, servieDetail: servieDetail}
}

func (s *reviewServices) AddReview(userID, productID int, rating float64, comment string) error {
	review := entities.Review{
		UserID:    userID,
		ProductID: productID,
		Rating:    rating,
		Ulasan:    comment,
	}

	if err := s.repoReview.AddReview(&review); err != nil {
		return err
	}
	return s.servieDetail.UpdateProductRating(productID)
}

func (s *reviewServices) GetUserReview(userID, productID int) (*entities.Review, error) {
	return s.repoReview.GetReviewByUserAndProduct(userID, productID)
}

func (s *reviewServices) GetReviewsByProduct(productID int) ([]entities.Review, error) {
	return s.repoReview.GetReviewsByProductID(productID)
}
