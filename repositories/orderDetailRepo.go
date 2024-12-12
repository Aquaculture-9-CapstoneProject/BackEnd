package repositories

import (
	"errors"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type OrderRepo interface {
	CheckIfUserExists(userID int) (bool, error)
	CreateOrder(order *entities.Order) error
	CreateOrderDetail(orderDetail *entities.OrderDetail) error
	ReduceStock(productID int, quantity int) error
	GetOrderForCheckout(userID int) ([]entities.Order, error)
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{db: db}
}

func (r *orderRepo) CheckIfUserExists(userID int) (bool, error) {
	var user entities.User
	if err := r.db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *orderRepo) CreateOrderDetail(orderDetail *entities.OrderDetail) error {
	return r.db.Create(orderDetail).Error
}

func (r *orderRepo) CreateOrder(order *entities.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepo) ReduceStock(productID int, quantity int) error {
	var product entities.Product
	if err := r.db.First(&product, productID).Error; err != nil {
		return err
	}
	if product.Stok < quantity {
		return errors.New("stok Tidak Cukup")
	}
	product.Stok -= quantity
	return r.db.Save(&product).Error
}

func (r *orderRepo) GetOrderForCheckout(userID int) ([]entities.Order, error) {
	var orders []entities.Order
	if err := r.db.Preload("Details.Product").Preload("Details.User").Preload("User").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
