package controllers

import (
	"net/http"

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
	userID, err := c.Get("userID")
	if !err {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "id user error"})
		return
	}
	var req struct {
		ProductID int `json:"product_id" binding:"required"`
		Quantity  int `json:"quantity" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gagal input"})
		return
	}
	if err := ctrl.cartService.TambahCart(userID.(int), req.ProductID, req.Quantity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak ada"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil", "code": 200, "status": "Data Berhasil Ditambah"})
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
	cartID := c.Param("cart_id")
	err := ctrl.cartService.RemoveFromCart(cartID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus dari cart"})
}

// func (ctrl *KeranjangControl) CheckOut(c *gin.Context) {
// 	userID, _ := c.Get("userID")
// 	if err := ctrl.cartService.Checkout(userID.(int)); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal melakukan checkout"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Berhasil", "code": 200, "status": "Berhasil Checkout"})
// }
