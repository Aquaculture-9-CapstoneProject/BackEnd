package controllers

import (
	"net/http"
	"strconv"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/gin-gonic/gin"
)

type KeranjangControl struct {
	cartService services.KeranjangServices
}

func NewCartControl(cartService services.KeranjangServices) *KeranjangControl {
	return &KeranjangControl{cartService: cartService}
}

func (ctrl *KeranjangControl) AddToCart(c *gin.Context) {
	// Get user ID from context
	userID, _ := c.Get("userID")

	// Bind request data (product_id, quantity)
	var req struct {
		ProductID int `json:"product_id" binding:"required"`
		Quantity  int `json:"quantity" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gagal input"})
		return
	}

	// Call service to add product to cart
	if err := ctrl.cartService.TambahCart(userID.(int), req.ProductID, req.Quantity); err != nil {
		// Tampilkan pesan error yang lebih jelas
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal menambahkan produk ke keranjang", "details": err.Error()})
		}
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Ditambah", "code": 200, "status": "Berhasil"})
}

func (ctrl *KeranjangControl) GetCartUser(c *gin.Context) {
	userID, exit := c.Get("userID")
	if !exit {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak terdaftar"})
		return
	}
	cart, total, err := ctrl.cartService.GetCartForUser(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan cart"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": gin.H{"message": "Berhasil", "code": 200, "status": "Berhasil"}, "Keranjang": cart, "Total": total})
}

func (ctrl *KeranjangControl) DeleteKeranjang(c *gin.Context) {
	cartIDStr := c.Param("cartID")
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
		return
	}
	if err := ctrl.cartService.RemoveFromCart(cartID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus dari cart"})
}

func (ctrl *KeranjangControl) CheckOut(c *gin.Context) {
	userID, _ := c.Get("userID")
	if err := ctrl.cartService.Checkout(userID.(int)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal melakukan checkout"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil", "code": 200, "status": "Berhasil Checkout"})
}
