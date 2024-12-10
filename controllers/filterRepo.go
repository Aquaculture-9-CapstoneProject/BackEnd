package controllers

import (
	"net/http"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/gin-gonic/gin"
)

type ProductFilterControl struct {
	service services.ProductFilterServices
}

func NewProductFilterControl(service services.ProductFilterServices) *ProductFilterControl {
	return &ProductFilterControl{service: service}
}

func (p *ProductFilterControl) FilterProduct(c *gin.Context) {
	kategori := c.Query("kategori")
	cari := c.Query("nama")
	product, err := p.service.CariProdukFilter(kategori, cari)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data tidak ada"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": gin.H{"message": "Berhasil", "code": 200, "status": "Berhasil"}, "product": product})
}
