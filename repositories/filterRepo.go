package repositories

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type ProductFilterRepo interface {
	CariProduct(kategori string, cari string) ([]entities.Product, error)
}

type productFilterRepo struct {
	db *gorm.DB
}

func NewProductFilterRepo(db *gorm.DB) ProductFilterRepo {
	return &productFilterRepo{db: db}
}

func (r *productFilterRepo) CariProduct(kategori string, cari string) ([]entities.Product, error) {
	var product []entities.Product
	query := r.db.Model(&entities.Product{})
	if kategori != "" {
		query = query.Where("Kategori = ?", kategori)
	}
	if cari != "" {
		query = query.Where("Nama LIKE ?", "%"+cari+"%")
	}
	err := query.Find(&product).Error
	return product, err
}
