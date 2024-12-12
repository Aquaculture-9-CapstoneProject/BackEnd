package repositories

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type KeranjangRepo interface {
	GetKeranjangByUserID(userID int) ([]entities.Cart, error)
	HapusIsiKeranjang(cartID int) error
	AddToKeranjang(cart *entities.Cart) error
	RemoveFromCart(cartID int) error // Ganti parameter ke int untuk konsistensi
	GetKeranjangItem(userID, productID int) (*entities.Cart, error)
	UpdateKeranjangItem(cartID, newQuantity int, newSubTotal float64) error
	CreateKeranjangItem(userID, productID, quantity int) error
}

type keranjangRepo struct {
	db *gorm.DB
}

func NewKeranjangRepo(db *gorm.DB) KeranjangRepo {
	return &keranjangRepo{db: db}
}

// AddToKeranjang untuk menambah produk ke dalam keranjang
func (r *keranjangRepo) AddToKeranjang(cart *entities.Cart) error {
	return r.db.Create(cart).Error
}

// GetKeranjangByUserID untuk mendapatkan produk yang ada dalam cart berdasarkan UserID
func (r *keranjangRepo) GetKeranjangByUserID(userID int) ([]entities.Cart, error) {
	var carts []entities.Cart
	if err := r.db.Preload("Product").Where("user_id = ?", userID).Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *keranjangRepo) HapusIsiKeranjang(cartID int) error {
	return r.db.Delete(&entities.Cart{}, cartID).Error
}

func (r *keranjangRepo) RemoveFromCart(cartID int) error {
	if err := r.db.Delete(&entities.Cart{}, "id = ?", cartID).Error; err != nil {
		return err
	}
	return nil
}

func (r *keranjangRepo) GetProductByID(productID int) (*entities.Product, error) {
	var product entities.Product
	if err := r.db.First(&product, productID).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// GetKeranjangItem untuk mendapatkan item keranjang berdasarkan userID dan productID
func (r *keranjangRepo) GetKeranjangItem(userID, productID int) (*entities.Cart, error) {
	var cart entities.Cart
	err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *keranjangRepo) UpdateKeranjangItem(cartID, newQuantity int, newSubTotal float64) error {
	return r.db.Model(&entities.Cart{}).Where("id = ?", cartID).Updates(map[string]interface{}{
		"kuantitas": newQuantity,
		"sub_total": newSubTotal,
	}).Error
}

func (r *keranjangRepo) CreateKeranjangItem(userID, productID, quantity int) error {
	product := entities.Product{}
	err := r.db.First(&product, productID).Error
	if err != nil {
		return err
	}
	subTotal := product.Harga * float64(quantity)

	cart := entities.Cart{
		UserID:    userID,
		ProductID: productID,
		Kuantitas: quantity,
		Subtotal:  subTotal,
	}

	return r.db.Create(&cart).Error
}
