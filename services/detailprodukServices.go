package services

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
)

type ProductDetailServices interface {
	LihatProductByID(ProductID int) (*entities.Product, error)
	UpdateProductRating(ProductID int) error
}

type productDetailServices struct {
	productRepo repositories.ProductDetailRepo
	reviewRepo  repositories.RatingRepo
}

func NewProductDetailServices(productRepo repositories.ProductDetailRepo, reviewRepo repositories.RatingRepo) ProductDetailServices {
	return &productDetailServices{productRepo: productRepo, reviewRepo: reviewRepo}
}

func (s *productDetailServices) LihatProductByID(ProductID int) (*entities.Product, error) {
	product, err := s.productRepo.CekProdukByID(ProductID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productDetailServices) UpdateProductRating(productID int) error {
	// Mendapatkan jumlah ulasan untuk produk
	totalReviews, err := s.reviewRepo.CountReviewsByProduct(productID)
	if err != nil {
		return err
	}

	// Mendapatkan jumlah rating total untuk produk
	totalRating, err := s.reviewRepo.SumRatingByProduct(productID)
	if err != nil {
		return err
	}

	// Menghitung rating baru dengan pembagian total rating / total ulasan
	newRating := 0.0
	if totalReviews > 0 {
		newRating = totalRating / float64(totalReviews) // totalRating sekarang bertipe float64
	}

	err = s.productRepo.UpdateProductRating(productID, newRating)
	if err != nil {
		return err
	}

	// Memperbarui jumlah ulasan produk di repository
	err = s.productRepo.UpdateTotalReview(productID, int(totalReviews))
	return err
}
