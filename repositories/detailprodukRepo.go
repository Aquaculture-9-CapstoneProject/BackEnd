package repositories

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type ProductDetailRepo interface {
	CekProdukByID(ProductID int) (*entities.Product, error)
	UpdateProductRating(ProductID int, newRating float64) error
	UpdateTotalReview(ProductID int, totalReviews int) error
}

type productDetailRepo struct {
	db *gorm.DB
}

func NewProductDetailRepo(db *gorm.DB) ProductDetailRepo {
	return &productDetailRepo{db: db}
}

func (r *productDetailRepo) CekProdukByID(ProductID int) (*entities.Product, error) {
	var product entities.Product
	err := r.db.Preload("OrderDetails.Product").Preload("Reviews.User").First(&product, ProductID).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productDetailRepo) UpdateProductRating(ProductID int, newRating float64) error {
	var product entities.Product
	if err := r.db.First(&product, ProductID).Error; err != nil {
		return err
	}
	product.Rating = newRating
	return r.db.Save(&product).Error
}

func (r *productDetailRepo) UpdateTotalReview(ProductID int, totalReviews int) error {
	return r.db.Model(&entities.Product{}).Where("id = ?", ProductID).Update("total_review", totalReviews).Error
}
