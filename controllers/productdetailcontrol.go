package controllers

import (
	"net/http"
	"strconv"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/gin-gonic/gin"
)

type ProductDetailControl struct {
	Productservices services.ProductDetailServices
}

func NewProductDetailControl(Productservices services.ProductDetailServices) *ProductDetailControl {
	return &ProductDetailControl{Productservices: Productservices}
}

func (ctrl *ProductDetailControl) CekDetailProdukByID(c *gin.Context) {
	productIdParam := c.Param("id")
	productID, err := strconv.Atoi(productIdParam)
	if err != nil || productID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input ID Gagal"})
		return
	}
	product, err := ctrl.Productservices.LihatProductByID(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"eror": "Product tidak ada"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": gin.H{"message": "Berhasil", "code": 200, "status": "Berhasil"}, "Produk": product})
}
