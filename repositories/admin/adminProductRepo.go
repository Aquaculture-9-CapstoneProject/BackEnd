package admin

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type AdminProductRepoInterface interface {
	CreateAdminProduct(product *entities.Product) (*entities.Product, error)
	UpdateAdminProduct(product *entities.Product) (*entities.Product, error)
	DeleteAdminProduct(id int) error
	FindByAdminProductID(id int) (*entities.Product, error)
	GetAdminProductCount() (int64, error)
	GetAllAdminProducts(page int, limit int) ([]entities.Product, error)
	SearchAdminProducts(nama string, kategori string, page int, limit int) ([]entities.Product, error)
}

type adminProductRepo struct {
	db *gorm.DB
}

func NewAdminProductRepo(db *gorm.DB) *adminProductRepo {
	return &adminProductRepo{db: db}
}

func (r *adminProductRepo) CreateAdminProduct(product *entities.Product) (*entities.Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *adminProductRepo) UpdateAdminProduct(product *entities.Product) (*entities.Product, error) {
	if err := r.db.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *adminProductRepo) DeleteAdminProduct(id int) error {
	return r.db.Delete(&entities.Product{}, id).Error
}

func (r *adminProductRepo) FindByAdminProductID(id int) (*entities.Product, error) {
	var product entities.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *adminProductRepo) GetAdminProductCount() (int64, error) {
	var count int64
	if err := r.db.Model(&entities.Product{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *adminProductRepo) GetAllAdminProducts(page int, limit int) ([]entities.Product, error) {
	var products []entities.Product
	offset := (page - 1) * limit
	err := r.db.Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *adminProductRepo) SearchAdminProducts(nama string, kategori string, page int, limit int) ([]entities.Product, error) {
	var products []entities.Product
	offset := (page - 1) * limit

	db := r.db.Model(&entities.Product{})

	if nama != "" {
		db = db.Where("nama LIKE ?", "%"+nama+"%")
	}

	if kategori != "" {
		db = db.Where("kategori = ?", kategori)
	}

	err := db.Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
