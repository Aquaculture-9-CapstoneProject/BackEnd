package repositories

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type ProductRepo interface {
	FindAll() ([]entities.Product, error)
	FindByID(id int) (entities.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{db: db}
}

func (r *productRepo) FindAll() ([]entities.Product, error) {
	var products []entities.Product
	err := r.db.Preload("Kategori").Preload("Rating").Find(&products).Error
	return products, err
}

func (r *productRepo) FindByID(id int) (entities.Product, error) {
	var product entities.Product
	err := r.db.Preload("Kategori").Preload("Rating").First(&product, id).Error
	return product, err
}
