package controllers

import (
	"net/http"
	"time"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/gin-gonic/gin"
)

type OrderControl struct {
	orderService services.OrderDetailService
}

func NewOrderControl(orderService services.OrderDetailService) *OrderControl {
	return &OrderControl{orderService: orderService}
}

func (ctrl *OrderControl) PlaceOrder(c *gin.Context) {
	userID, err := c.Get("userID")
	if !err {
		c.JSON(http.StatusUnauthorized, gin.H{"Message": "User Tidak Ditemukan"})
		return
	}
	var req struct {
		ProductID int `json:"product_id" binding:"required"`
		Quantity  int `json:"quantity" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Data Tidak Valid"})
		return
	}
	if req.Quantity == 0 {
		req.Quantity = 1
	}
	if err := ctrl.orderService.PlaceOrder(userID.(int), req.ProductID, req.Quantity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order Berhasil Dibuat", "code": 200, "status": "Berhasil"})
}

func (ctrl *OrderControl) GetOrderForCheckout(c *gin.Context) {
	// Ambil userID dari konteks
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak terdaftar"})
		return
	}

	// Panggil service untuk mendapatkan data order
	orders, err := ctrl.orderService.GetOrderForCheckout(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data order"})
		return
	}

	// Pastikan `User` di setiap elemen menjadi nil
	for i := range orders {
		orders[i].User = entities.User{}
	}

	// Kirimkan response ke client
	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"message": "Berhasil",
			"code":    200,
			"status":  "Berhasil",
		},
		"orders":  orders,
		"tanggal": time.Now().Format(time.RFC3339),
	})
}
