// repositories/product_repository.go
package repositories

import (
	"fmt"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts() ([]entities.Product, error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) GetAllProducts() ([]entities.Product, error) {
	var products []entities.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, fmt.Errorf("failed to get products: %v", err)
	}
	return products, nil
}
