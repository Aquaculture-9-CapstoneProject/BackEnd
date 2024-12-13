// services/order_service.go

package services

import (
	"fmt"
	"time"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
)

type OrderDetailService interface {
	PlaceOrder(userID int, productID int, quantity int) error
	GetOrderForCheckout(userID int) ([]entities.Order, error)
}

type orderService struct {
	orderRepo  repositories.OrderRepo
	repoProduk repositories.ProductDetailRepo
}

func NeworderService(orderRepo repositories.OrderRepo, repoProduk repositories.ProductDetailRepo) OrderDetailService {
	return &orderService{orderRepo: orderRepo, repoProduk: repoProduk}
}

func (s *orderService) PlaceOrder(userID int, productID int, quantity int) error {

	userExists, err := s.orderRepo.CheckIfUserExists(userID)
	if err != nil {
		return fmt.Errorf("gagal memeriksa user: %w", err)
	}
	if !userExists {
		return fmt.Errorf("user tidak ditemukan")
	}

	product, err := s.repoProduk.CekProdukByID(productID)
	if err != nil {
		return fmt.Errorf("produk tidak ditemukan: %w", err)
	}

	subTotal := float64(quantity) * product.Harga
	biayaLayanan := 2000.0
	biayaOngkir := 10000.0
	total := subTotal + biayaLayanan + biayaOngkir

	order := &entities.Order{
		UserID:           userID,
		Total:            total,
		MetodePembayaran: "E-Wallet",
		BiayaLayanan:     biayaLayanan,
		BiayaOngkir:      biayaOngkir,
		CreatedAt:        time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := s.orderRepo.CreateOrder(order); err != nil {
		return fmt.Errorf("gagal menyimpan order: %w", err)
	}

	orderDetail := &entities.OrderDetail{
		OrderID:   order.ID,
		ProductID: productID,
		UserID:    userID,
		Kuantitas: quantity,
		Subtotal:  subTotal,
	}

	if err := s.orderRepo.CreateOrderDetail(orderDetail); err != nil {
		return fmt.Errorf("gagal menyimpan order detail: %w", err)
	}

	if err := s.orderRepo.ReduceStock(productID, quantity); err != nil {
		return fmt.Errorf("gagal mengurangi stok: %w", err)
	}

	return nil
}

func (s *orderService) GetOrderForCheckout(userID int) ([]entities.Order, error) {
	return s.orderRepo.GetOrderForCheckout(userID)
}
