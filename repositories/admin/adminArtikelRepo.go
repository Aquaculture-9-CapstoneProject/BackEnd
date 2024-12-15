package admin

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type AdminArtikelRepoInterface interface {
	Create(artikel *entities.Artikel) (*entities.Artikel, error)
	Update(artikel *entities.Artikel) (*entities.Artikel, error)
	Delete(id int) error
	GetAll(page int, limit int) ([]entities.Artikel, error)
	FindAll(judul string, kategori string, page int, limit int) ([]entities.Artikel, error)
	FindByID(id int) (*entities.Artikel, error)
	Count() (int64, error)
}

type adminArtikelRepo struct {
	db *gorm.DB
}

func NewAdminArtikelRepo(db *gorm.DB) *adminArtikelRepo {
	return &adminArtikelRepo{db: db}
}

func (r *adminArtikelRepo) Create(artikel *entities.Artikel) (*entities.Artikel, error) {
	if err := r.db.Create(artikel).Error; err != nil {
		return nil, err
	}
	return artikel, nil
}

func (r *adminArtikelRepo) Update(artikel *entities.Artikel) (*entities.Artikel, error) {
	if err := r.db.Save(artikel).Error; err != nil {
		return nil, err
	}
	return artikel, nil
}

func (r *adminArtikelRepo) Delete(id int) error {
	return r.db.Delete(&entities.Artikel{}, id).Error
}

func (r *adminArtikelRepo) GetAll(page int, limit int) ([]entities.Artikel, error) {
	var artikels []entities.Artikel
	offset := (page - 1) * limit
	err := r.db.Limit(limit).Offset(offset).Find(&artikels).Error
	if err != nil {
		return nil, err
	}
	return artikels, nil
}

func (r *adminArtikelRepo) FindAll(judul string, kategori string, page int, limit int) ([]entities.Artikel, error) {
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

func (r *adminArtikelRepo) FindByID(id int) (*entities.Artikel, error) {
	var artikel entities.Artikel
	if err := r.db.First(&artikel, id).Error; err != nil {
		return nil, err
	}
	return &artikel, nil
}

func (r *adminArtikelRepo) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&entities.Artikel{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
