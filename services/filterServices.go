package services

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
)

type ProductFilterServices interface {
	CariProdukFilter(kategori string, cari string) ([]entities.Product, error)
}

type productFilterService struct {
	repo repositories.ProductFilterRepo
}

func NewProductFilterService(repo repositories.ProductFilterRepo) ProductFilterServices {
	return &productFilterService{repo: repo}
}

func (s *productFilterService) CariProdukFilter(kategori string, cari string) ([]entities.Product, error) {
	return s.repo.CariProduct(kategori, cari)
}
