package services

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
)

type ProductUseCase interface {
	GetAllProducts() ([]entities.Product, error)
	GetProductByID(id int) (entities.Product, float64, error)
	CreateProduct(product entities.Product) (entities.Product, error)
	UpdateProduct(product entities.Product) (entities.Product, error)
	DeleteProduct(id int) error
}

type productUseCase struct {
	repo repositories.ProductRepo
}

func NewProductUseCase(repo repositories.ProductRepo) ProductUseCase {
	return &productUseCase{repo: repo}
}

func (s *productUseCase) GetAllProducts() ([]entities.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *productUseCase) GetProductByID(id int) (entities.Product, float64, error) {
	product, err := s.repo.GetProductByID(id)
	if err != nil {
		return product, 0, err
	}

	var totalRating int
	ratingCount := len(product.Ratings)

	if ratingCount > 0 {
		for _, rating := range product.Ratings {
			totalRating += rating.Rating
		}
		averageRating := float64(totalRating) / float64(ratingCount)
		return product, averageRating, nil
	}

	return product, 0, nil
}

func (s *productUseCase) CreateProduct(product entities.Product) (entities.Product, error) {
	return s.repo.CreateProduct(product)
}

func (s *productUseCase) UpdateProduct(product entities.Product) (entities.Product, error) {
	return s.repo.UpdateProduct(product)
}

func (s *productUseCase) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}
