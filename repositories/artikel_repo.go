package repositories

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type ArtikelRepoInterface interface {
	Create(artikel *entities.Artikel) (*entities.Artikel, error)
	Update(artikel *entities.Artikel) (*entities.Artikel, error)
	Delete(id int) error
	FindAll(nama string, kategori string, page int, limit int) ([]entities.Artikel, error)
	FindByID(id int) (*entities.Artikel, error)
	GetAdminByID(id int) (*entities.Admin, error)
	Count() (int64, error)
}

type artikelRepo struct {
	db *gorm.DB
}

func NewArtikelRepo(db *gorm.DB) *artikelRepo {
	return &artikelRepo{db: db}
}

func (r *artikelRepo) Create(artikel *entities.Artikel) (*entities.Artikel, error) {
	if err := r.db.Create(artikel).Error; err != nil {
		return nil, err
	}
	return artikel, nil
}

func (r *artikelRepo) Update(artikel *entities.Artikel) (*entities.Artikel, error) {
	if err := r.db.Save(artikel).Error; err != nil {
		return nil, err
	}
	return artikel, nil
}

func (r *artikelRepo) Delete(id int) error {
	return r.db.Delete(&entities.Artikel{}, id).Error
}

func (r *artikelRepo) FindAll(nama string, kategori string, page int, limit int) ([]entities.Artikel, error) {
	var artikels []entities.Artikel
	offset := (page - 1) * limit

	db := r.db.Model(&entities.Artikel{})

	if nama != "" {
		db = db.Where("nama LIKE ?", "%"+nama+"%")
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

func (r *artikelRepo) GetAdminByID(id int) (*entities.Admin, error) {
	var admin entities.Admin
	if err := r.db.First(&admin, id).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *artikelRepo) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&entities.Artikel{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
