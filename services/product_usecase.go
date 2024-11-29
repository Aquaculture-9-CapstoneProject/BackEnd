package services

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
)

type ProductUseCase interface {
	GetAllProducts() ([]entities.Product, error)
	GetProductByID(id uint) (entities.Product, error)
}

type productUseCase struct {
	repo repositories.ProductRepo
}

func NewProductUseCase(repo repositories.ProductRepo) ProductUseCase {
	return &productUseCase{repo: repo}
}

func (s *productUseCase) GetAllProducts() ([]entities.Product, error) {
	return s.repo.FindAll()
}

func (s *productUseCase) GetProductByID(id uint) (entities.Product, error) {
	return s.repo.FindByID(id)
}
