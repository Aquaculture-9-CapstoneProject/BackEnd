package repositories

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type KeranjangRepo interface {
	GetKeranjangByUserID(userID int) ([]entities.Cart, error)
	HapusIsiKeranjang(cartID int) error
	AddToKeranjang(cart *entities.Cart) error
	RemoveFromCart(cartID string) error
}

type keranjangRepo struct {
	db *gorm.DB
}

func NewKeranjangRepo(db *gorm.DB) KeranjangRepo {
	return &keranjangRepo{db: db}
}

func (r *keranjangRepo) GetKeranjangByUserID(userID int) ([]entities.Cart, error) {
	var keranjang []entities.Cart
	if err := r.db.Preload("Product").Where("user_id = ?", userID).Find(&keranjang).Error; err != nil {
		return nil, err
	}
	return keranjang, nil
}

func (r *keranjangRepo) AddToKeranjang(cart *entities.Cart) error {
	return r.db.Create(cart).Error
}

func (r *keranjangRepo) HapusIsiKeranjang(cartID int) error {
	return r.db.Delete(&entities.Cart{}, cartID).Error
}

func (r *keranjangRepo) RemoveFromCart(cartID string) error {
	if err := r.db.Delete(&entities.Cart{}, "id = ?", cartID).Error; err != nil {
		return err
	}
	return nil
}
