package repositories

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type RatingRepo interface {
	GetReviewsByProductID(productID int) ([]entities.Review, error)
	AddReview(review *entities.Review) error
	GetReviewByUserAndProduct(userID, productID int) (*entities.Review, error)
	CountReviewsByProduct(productID int) (int, error)
	SumRatingByProduct(productID int) (int, error)
}

type ratingRepo struct {
	db *gorm.DB
}

func NewRatingRepo(db *gorm.DB) RatingRepo {
	return &ratingRepo{db: db}
}

func (r *ratingRepo) GetReviewsByProductID(productID int) ([]entities.Review, error) {
	var reviews []entities.Review
	err := r.db.Preload("User").Where("product_id = ?", productID).Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *ratingRepo) AddReview(review *entities.Review) error {
	return r.db.Create(review).Error
}

func (r *ratingRepo) GetReviewByUserAndProduct(userID, productID int) (*entities.Review, error) {
	var review entities.Review
	err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&review).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ratingRepo) CountReviewsByProduct(productID int) (int, error) {
	var count int64
	err := r.db.Model(&entities.Review{}).Where("product_id = ?", productID).Distinct("user_id").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *ratingRepo) SumRatingByProduct(productID int) (int, error) {
	var totalRating int
	err := r.db.Model(&entities.Review{}).Where("product_id = ?", productID).Select("SUM(rating)").Scan(&totalRating).Error
	if err != nil {
		return 0, err
	}
	return totalRating, nil
}
