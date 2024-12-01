package repositories

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type RatingRepo interface {
	GetAllRatings() ([]entities.Rating, error)
	GetRatingByID(id int) (entities.Rating, error)
	CreateRating(rating entities.Rating) error
	UpdateRating(rating entities.Rating) error
	DeleteRating(id int) error
}

type ratingRepo struct {
	db *gorm.DB
}

func NewRatingRepo(db *gorm.DB) RatingRepo {
	return &ratingRepo{db: db}
}

func (r *ratingRepo) GetAllRatings() ([]entities.Rating, error) {
	var ratings []entities.Rating
	err := r.db.Find(&ratings).Error
	return ratings, err
}

func (r *ratingRepo) GetRatingByID(id int) (entities.Rating, error) {
	var rating entities.Rating
	err := r.db.First(&rating, id).Error
	return rating, err
}

func (r *ratingRepo) CreateRating(rating entities.Rating) error {
	return r.db.Create(&rating).Error
}

func (r *ratingRepo) UpdateRating(rating entities.Rating) error {
	return r.db.Save(&rating).Error
}

func (r *ratingRepo) DeleteRating(id int) error {
	return r.db.Delete(&entities.Rating{}, id).Error
}
