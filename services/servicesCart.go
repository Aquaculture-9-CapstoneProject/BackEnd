package services

import (
	"fmt"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
)

type KeranjangServices interface {
	TambahCart(userID int, productID int, quantity int) error
	GetCartForUser(userID int) ([]entities.Cart, float64, error)
	RemoveFromCart(cartID int) error
	Checkout(userID int) error
}

type keranjangServices struct {
	cartRepo   repositories.KeranjangRepo
	orderRepo  repositories.OrderRepo
	produkRepo repositories.ProductDetailRepo
}

func NewServicesKeranjang(cartRepo repositories.KeranjangRepo, produkRepo repositories.ProductDetailRepo, orderRepo repositories.OrderRepo) KeranjangServices {
	return &keranjangServices{
		cartRepo:   cartRepo,
		produkRepo: produkRepo,
		orderRepo:  orderRepo,
	}
}

func (s *keranjangServices) TambahCart(userID int, productID int, quantity int) error {
	// Cek apakah produk sudah ada di cart
	cartItem, err := s.cartRepo.GetKeranjangItem(userID, productID)
	if err != nil && err.Error() != "record not found" {
		return err
	}
	produk, err := s.produkRepo.CekProdukByID(productID)
	if err != nil {
		return err
	}
	subTotal := float64(quantity) * produk.Harga

	if cartItem != nil {
		newQuantity := cartItem.Kuantitas + quantity
		newSubTotal := cartItem.Subtotal + subTotal
		return s.cartRepo.UpdateKeranjangItem(cartItem.ID, newQuantity, newSubTotal)
	}
	return s.cartRepo.CreateKeranjangItem(userID, productID, quantity)
}

func (s *keranjangServices) GetCartForUser(userID int) ([]entities.Cart, float64, error) {
	carts, err := s.cartRepo.GetKeranjangByUserID(userID)
	if err != nil {
		return nil, 0, err
	}
	total := 0.0
	for _, cart := range carts {
		total += cart.Subtotal
	}
	return carts, total, nil
}

func (s *keranjangServices) RemoveFromCart(cartID int) error {
	return s.cartRepo.RemoveFromCart(cartID)
}

func (s *keranjangServices) Checkout(userID int) error {
	carts, err := s.cartRepo.GetKeranjangByUserID(userID)
	if err != nil {
		return fmt.Errorf("gagal mendapatkan keranjang: %w", err)
	}
	total := 0.0
	for _, cart := range carts {
		total += cart.Subtotal
	}

	biayaLayanan := total * 0.05
	biayaOngkir := 10000.0
	totalOrder := total + biayaLayanan + biayaOngkir
	order := &entities.Order{
		UserID:           userID,
		Total:            totalOrder,
		MetodePembayaran: "Transfer Bank",
		BiayaLayanan:     biayaLayanan,
		BiayaOngkir:      biayaOngkir,
	}
	if err := s.orderRepo.CreateOrder(order); err != nil {
		return fmt.Errorf("gagal membuat order: %w", err)
	}
	for _, cart := range carts {
		orderDetail := &entities.OrderDetail{
			OrderID:   order.ID,
			ProductID: cart.ProductID,
			UserID:    userID,
			Kuantitas: cart.Kuantitas,
			Subtotal:  cart.Subtotal,
		}
		if err := s.orderRepo.CreateOrderDetail(orderDetail); err != nil {
			return fmt.Errorf("gagal menyimpan detail order: %w", err)
		}
	}
	for _, cart := range carts {
		if err := s.cartRepo.HapusIsiKeranjang(cart.ID); err != nil {
			return fmt.Errorf("gagal menghapus produk dari keranjang: %w", err)
		}
	}

	return nil
}
