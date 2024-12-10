package controllers

import (
	"net/http"
	"strconv"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/gin-gonic/gin"
)

type ProductIkanController struct {
	service services.ProductUseCase
}

func NewProductIkanController(service services.ProductUseCase) *ProductIkanController {
	return &ProductIkanController{service: service}
}

func (ctrl *ProductIkanController) GetTermurahProduk(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	produk, err := ctrl.service.GetProdukTermurah(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": gin.H{"message": err.Error(), "code": 401, "status": "eror"}, "data": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": gin.H{"message": "Berhasil", "code": 200, "status": "Berhasil"}, "data": produk})
}

func (ctrl *ProductIkanController) GetPopulerProduk(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "4"))
	produk, err := ctrl.service.GetAllProductPopuler(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"meta": gin.H{"message": err.Error(), "code": 401, "status": "eror"}, "data": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": gin.H{"message": "Berhasil", "code": 200, "status": "Berhasil"}, "data": produk})
}
