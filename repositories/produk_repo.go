package repositories

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type ProductRepo interface {
	GetAllProducts() ([]entities.Product, error)
	GetProductByID(id int) (entities.Product, error)
	CreateProduct(product entities.Product) (entities.Product, error)
	UpdateProduct(product entities.Product) (entities.Product, error)
	DeleteProduct(id int) error
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{db: db}
}

func (r *productRepo) GetAllProducts() ([]entities.Product, error) {
	var products []entities.Product
	err := r.db.Preload("Ratings").Find(&products).Error
	return products, err
}

func (r *productRepo) GetProductByID(id int) (entities.Product, error) {
	var product entities.Product
	err := r.db.Preload("Ratings").First(&product, id).Error
	return product, err
}

func (r *productRepo) CreateProduct(product entities.Product) (entities.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *productRepo) UpdateProduct(product entities.Product) (entities.Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

func (r *productRepo) DeleteProduct(id int) error {
	return r.db.Delete(&entities.Product{}, id).Error
}
