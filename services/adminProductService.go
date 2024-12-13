package services

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
)

type AdminProductUseCase interface {
	CreateAdminProduct(product *entities.Product) (*entities.Product, error)
	UpdateAdminProduct(product *entities.Product) (*entities.Product, error)
	DeleteAdminProduct(id int) error
	FindByAdminProductID(id int) (*entities.Product, error)
	GetAdminProductCount() (int64, error)
	GetAllAdminProducts(limit int) ([]entities.Product, error)
	SearchAdminProducts(nama string, kategori string, page int, limit int) ([]entities.Product, error)
}

type adminProductUseCase struct {
	repo repositories.AdminProductRepoInterface
}

func NewAdminProductService(repo repositories.AdminProductRepoInterface) *adminProductUseCase {
	return &adminProductUseCase{repo: repo}
}

func (s *adminProductUseCase) CreateAdminProduct(product *entities.Product) (*entities.Product, error) {
	return s.repo.CreateAdminProduct(product)
}

func (s *adminProductUseCase) UpdateAdminProduct(product *entities.Product) (*entities.Product, error) {
	return s.repo.UpdateAdminProduct(product)
}

func (s *adminProductUseCase) DeleteAdminProduct(id int) error {
	return s.repo.DeleteAdminProduct(id)
}

func (s *adminProductUseCase) FindByAdminProductID(id int) (*entities.Product, error) {
	return s.repo.FindByAdminProductID(id)
}

func (s *adminProductUseCase) GetAdminProductCount() (int64, error) {
	return s.repo.GetAdminProductCount()
}

func (s *adminProductUseCase) GetAllAdminProducts(limit int) ([]entities.Product, error) {
	return s.repo.GetAllAdminProducts(limit)
}

func (s *adminProductUseCase) SearchAdminProducts(nama string, kategori string, page int, limit int) ([]entities.Product, error) {
	products, err := s.repo.SearchAdminProducts(nama, kategori, page, limit)
	if err != nil {
		return nil, err
	}
	return products, nil
}
