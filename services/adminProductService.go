package services

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
)

type AdminProductUseCase interface {
	CreateAdminProduct(product *entities.Product) (*entities.Product, error)
	UpdateAdminProduct(product *entities.Product) (*entities.Product, error)
	DeleteAdminProduct(id int) error
	GetAllAdminProducts(page int, limit int) ([]entities.Product, error)
	FindByAdminProductID(id int) (*entities.Product, error)
	GetAdminProductCount() (int64, error)
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

func (s *adminProductUseCase) GetAllAdminProducts(page int, limit int) ([]entities.Product, error) {
	return s.repo.FindAllAdminProducts(page, limit)
}

func (s *adminProductUseCase) FindByAdminProductID(id int) (*entities.Product, error) {
	return s.repo.FindByAdminProductID(id)
}

func (s *adminProductUseCase) GetAdminProductCount() (int64, error) {
	return s.repo.GetAdminProductCount()
}
