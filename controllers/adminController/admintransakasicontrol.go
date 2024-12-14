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
	// Mendapatkan parameter pagination dari query
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("per_page", "10")

	// Konversi page dan per_page menjadi integer
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page harus berupa angka"})
		return
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Per_page harus berupa angka"})
		return
	}

	// Memanggil service untuk mendapatkan detail pembayaran dengan pagination
	response, err := ctrl.adminServicesTransaksi.GetPaymentDetails(page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan response dengan status OK dan data detail pembayaran serta pagination
	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{"message": "Berhasil", "code": 200, "status": "Berhasil"},
		"data": response,
	})
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
	c.JSON(http.StatusOK, gin.H{"meta": gin.H{"message": "Berhasil Dihapus", "code": 200, "status": "Berhasil"}})
}
