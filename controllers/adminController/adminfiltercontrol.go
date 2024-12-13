package admincontroller

import (
	"net/http"

	adminservices "github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services/adminServices"
	"github.com/gin-gonic/gin"
)

type AdminFilterController struct {
	service adminservices.AdminFilterServices
}

func NewAdminFilterController(service adminservices.AdminFilterServices) *AdminFilterController {
	return &AdminFilterController{service: service}
}

func (ctrl *AdminFilterController) GetPaymentsByStatus(c *gin.Context) {
	status := c.DefaultQuery("status", "PAID")
	payments, err := ctrl.service.GetPaymentsByStatus(status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payments})
}

func (ctrl *AdminFilterController) GetPaymentsByStatusBarang(c *gin.Context) {
	statusBarang := c.DefaultQuery("status_barang", "DIKIRIM")

	payments, err := ctrl.service.GetPaymentsByStatusBarang(statusBarang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data pembayaran"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payments})
}
