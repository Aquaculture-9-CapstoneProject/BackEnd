package admincontroller

import (
	"net/http"
	"strconv"

	adminservices "github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services/adminServices"
	"github.com/gin-gonic/gin"
)

type AdminTransaksiControl struct {
	adminServicesTransaksi adminservices.AdminTransaksiService
}

func NewAdminTransaksiController(adminServicesTransaksi adminservices.AdminTransaksiService) *AdminTransaksiControl {
	return &AdminTransaksiControl{adminServicesTransaksi: adminServicesTransaksi}
}

func (ctrl *AdminTransaksiControl) GetPaymentDetails(c *gin.Context) {
	// Memanggil service untuk mendapatkan detail pembayaran
	details, err := ctrl.adminServicesTransaksi.GetPaymentDetails()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Mengembalikan response dengan status OK dan data detail pembayaran
	c.JSON(http.StatusOK, gin.H{"data": details})
}

func (ctrl *AdminTransaksiControl) DeletePaymentByID(c *gin.Context) {
	// Mendapatkan ID dari parameter URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	// Memanggil service untuk menghapus payment
	err = ctrl.adminServicesTransaksi.DeletePaymentByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, gin.H{"message": "Payment berhasil dihapus"})
}
