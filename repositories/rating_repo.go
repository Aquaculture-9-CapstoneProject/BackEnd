package repositories

import (
	"gorm.io/gorm"
)

type RatingRepo interface {
}

type ratingRepo struct {
	db *gorm.DB
}

func NewRatingRepo(db *gorm.DB) RatingRepo {
	return &ratingRepo{db: db}
}
