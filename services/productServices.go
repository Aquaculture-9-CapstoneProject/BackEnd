package services

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
)

type ProductUseCase interface {
	GetProdukTermurah(limit int) ([]entities.Product, error)
	GetAllProductPopuler(limit int) ([]entities.Product, error)
}

type productUseCase struct {
	repo repositories.ProdukIkanRepo
}

func NewProductIkanServices(repo repositories.ProdukIkanRepo) ProductUseCase {
	return &productUseCase{repo: repo}
}

func (s *productUseCase) GetProdukTermurah(limit int) ([]entities.Product, error) {
	return s.repo.GetTermurah(limit)
}

func (s *productUseCase) GetAllProductPopuler(limit int) ([]entities.Product, error) {
	return s.repo.GetPopuler(limit)
}
