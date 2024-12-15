package repositories

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type ArtikelRepoInterface interface {
	GetAll(page int, limit int) ([]entities.Artikel, error)
	Top3(limit int) ([]entities.Artikel, error)
	FindAll(judul string, kategori string, page int, limit int) ([]entities.Artikel, error)
	FindByID(id int) (*entities.Artikel, error)
	Count() (int64, error)
}

type artikelRepo struct {
	db *gorm.DB
}

func NewArtikelRepo(db *gorm.DB) *artikelRepo {
	return &artikelRepo{db: db}
}

func (r *artikelRepo) GetAll(page int, limit int) ([]entities.Artikel, error) {
	var artikels []entities.Artikel
	offset := (page - 1) * limit
	err := r.db.Limit(limit).Offset(offset).Find(&artikels).Error
	if err != nil {
		return nil, err
	}
	return artikels, nil
}

func (r *artikelRepo) Top3(limit int) ([]entities.Artikel, error) {
	var artikels []entities.Artikel
	db := r.db.Model(&entities.Artikel{})
	db = db.Where("kategori LIKE ?", "%resep%")

	err := db.Limit(limit).Find(&artikels).Error
	if err != nil {
		return nil, err
	}

	return artikels, nil
}

func (r *artikelRepo) FindAll(judul string, kategori string, page int, limit int) ([]entities.Artikel, error) {
	var artikels []entities.Artikel
	offset := (page - 1) * limit

	db := r.db.Model(&entities.Artikel{})

	if judul != "" {
		db = db.Where("judul LIKE ?", "%"+judul+"%")
	}

	if kategori != "" {
		db = db.Where("kategori = ?", kategori)
	}

	err := db.Limit(limit).Offset(offset).Find(&artikels).Error
	if err != nil {
		return nil, err
	}

	return artikels, nil
}

func (r *artikelRepo) FindByID(id int) (*entities.Artikel, error) {
	var artikel entities.Artikel
	if err := r.db.First(&artikel, id).Error; err != nil {
		return nil, err
	}
	return &artikel, nil
}

func (r *artikelRepo) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&entities.Artikel{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
