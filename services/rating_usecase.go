package services

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
)

type RatingUseCase interface {
	GetAllRatings() ([]entities.Rating, error)
	GetRatingByID(id int) (entities.Rating, error)
	CreateRating(rating entities.Rating) error
	UpdateRating(rating entities.Rating) error
	DeleteRating(id int) error
}

type ratingUseCase struct {
	repo repositories.RatingRepo
}

func NewRatingUseCase(repo repositories.RatingRepo) RatingUseCase {
	return &ratingUseCase{repo: repo}
}

func (s *ratingUseCase) GetAllRatings() ([]entities.Rating, error) {
	return s.repo.GetAllRatings()
}

func (s *ratingUseCase) GetRatingByID(id int) (entities.Rating, error) {
	return s.repo.GetRatingByID(id)
}

func (s *ratingUseCase) CreateRating(rating entities.Rating) error {
	return s.repo.CreateRating(rating)
}

func (s *ratingUseCase) UpdateRating(rating entities.Rating) error {
	return s.repo.UpdateRating(rating)
}

func (s *ratingUseCase) DeleteRating(id int) error {
	return s.repo.DeleteRating(id)
}
