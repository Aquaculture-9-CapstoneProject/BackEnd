package services

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
)

type ProductDetailServices interface {
	LihatProductByID(ProductID int) (*entities.Product, error)
	// UpdateProductRating(ProductID int) error
}

type productDetailServices struct {
	productRepo repositories.ProductDetailRepo
	// reviewRepo  repositories.ReviewRepository
}

func NewProductDetailServices(productRepo repositories.ProductDetailRepo) ProductDetailServices {
	return &productDetailServices{productRepo: productRepo}
}

func (s *productDetailServices) LihatProductByID(ProductID int) (*entities.Product, error) {
	product, err := s.productRepo.CekProdukByID(ProductID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// func (s *productDetailServices) UpdateProductRating(productID uint) error {
// 	// Hitung jumlah ulasan unik (user_id) untuk produk
// 	totalReviews, err := s.reviewRepo.CountReviewsByProduct(productID)
// 	if err != nil {
// 		return err
// 	}

// 	// Hitung jumlah total rating untuk produk
// 	totalRating, err := s.reviewRepo.SumRatingByProduct(productID)
// 	if err != nil {
// 		return err
// 	}

// 	// Hitung rata-rata rating
// 	newRating := 0.0
// 	if totalReviews > 0 {
// 		newRating = float64(totalRating) / float64(totalReviews)
// 	}

// 	// Perbarui nilai Rating dan TotalReview di tabel Product
// 	err = s.productRepo.UpdateProductRating(productID, newRating)
// 	if err != nil {
// 		return err
// 	}
// 	err = s.productRepo.UpdateTotalReview(productID, int(totalReviews))
// 	return err
// }
