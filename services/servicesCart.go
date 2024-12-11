package services

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
)

type KeranjangServices interface {
	TambahCart(userID int, productID int, quantity int) error
	GetCartForUser(userID int) ([]entities.Cart, float64, error)
	RemoveFromCart(cartID string) error
	// Checkout(userID int) error
}

type keranjangServices struct {
	cartRepo repositories.KeranjangRepo
	// orderRepo *repositories.OrderRepository
	produkid repositories.ProductDetailRepo
}

func NewServicesKeranjang(cartRepo repositories.KeranjangRepo, produkid repositories.ProductDetailRepo) KeranjangServices {
	return &keranjangServices{cartRepo: cartRepo, produkid: produkid}
}

func (s *keranjangServices) TambahCart(userID int, productID int, quantity int) error {
	produk, _ := s.produkid.CekProdukByID(productID)
	subTotal := float64(quantity) * produk.Harga
	cart := &entities.Cart{
		UserID:    userID,
		ProductID: productID,
		Kuantitas: quantity,
		Subtotal:  subTotal,
	}
	return s.cartRepo.AddToKeranjang(cart)
}

func (s *keranjangServices) GetCartForUser(userID int) ([]entities.Cart, float64, error) {
	carts, _ := s.cartRepo.GetKeranjangByUserID(userID)
	total := 0.0
	for _, cart := range carts {
		total += cart.Subtotal
	}
	return carts, total, nil
}

func (s *keranjangServices) RemoveFromCart(cartID string) error {
	return s.cartRepo.RemoveFromCart(cartID)
}

// func (s *keranjangServices) Checkout(userID int) error {
// 	carts, _ := s.cartRepo.GetKeranjangByUserID(userID)
// 	total := 0.0
// 	for _, cart := range carts {
// 		total += cart.Subtotal
// 	}
// 	// biayaLayanan := total * 0.5
// 	// biayaOngkir := 10000.0
// 	// totalOrder := total + biayaLayanan + biayaOngkir
// 	// order := &entities.Order{
// 	// 	UserID:           userID,
// 	// 	Total:            totalOrder,
// 	// 	MetodePembayaran: "Transfer Bank",
// 	// 	BiayaLayanan:     biayaLayanan,
// 	// 	BiayaOngkir:      biayaOngkir,
// 	// }
// 	// if err := s.orderRepo.CreateOrder(order); err != nil {
// 	// 	return fmt.Errorf("gagal membuat order: %w", err)
// 	// }

// 	for _, cart := range carts {
// 		orderDetail := &entities.OrderDetail{
// 			OrderID: order.ID,

// 		}
// 	}

// }
