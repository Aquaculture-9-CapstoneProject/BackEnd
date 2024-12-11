package repositories

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type ProdukIkanRepo interface {
	GetTermurah(limit int) ([]entities.Product, error)
	GetPopuler(limit int) ([]entities.Product, error)
}

type produkIkanRepo struct {
	db *gorm.DB
}

func NewProdukIkanRepo(db *gorm.DB) ProdukIkanRepo {
	return &produkIkanRepo{db: db}
}

func (r *produkIkanRepo) GetTermurah(limit int) ([]entities.Product, error) {
	var produk []entities.Product
	err := r.db.Select("id", "gambar", "nama", "kategori", "harga", "rating").Order("harga asc").Limit(limit).Find(&produk).Error
	return produk, err
}

func (r *produkIkanRepo) GetPopuler(limit int) ([]entities.Product, error) {
	var produk []entities.Product
	err := r.db.Select("id", "gambar", "nama", "kategori", "harga", "rating").Order("rating desc").Limit(limit).Find(&produk).Error
	return produk, err
}
