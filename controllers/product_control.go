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
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Meta": gin.H{"Message": "Berhasil", "Code": 200, "Status": "Berhasil"}, "Data": produk})
}

func (ctrl *ProductIkanController) GetPopulerProduk(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "4"))
	produk, err := ctrl.service.GetAllProductPopuler(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Meta": gin.H{"Message": err.Error(), "Code": 401, "Status": "eror"}, "Data": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Meta": gin.H{"Message": "Berhasil", "Code": 200, "Status": "Berhasil"}, "Data": produk})
}
