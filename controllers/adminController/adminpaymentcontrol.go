package admincontroller

import (
	"net/http"

	adminservices "github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services/adminServices"
	"github.com/gin-gonic/gin"
)

type AdminPaymentController struct {
	adminPaymentService adminservices.AdminPaymentService
}

func NewAdminPaymentController(adminPaymentService adminservices.AdminPaymentService) *AdminPaymentController {
	return &AdminPaymentController{adminPaymentService: adminPaymentService}
}

func (ctrl *AdminPaymentController) GetAdminTotalPendapatanBulanIni(c *gin.Context) {
	totalPendapatan, err := ctrl.adminPaymentService.HitungAdminPendapatanBulanIni()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"totalPendapatan": totalPendapatan, "message": "Berhasil", "code": 200, "status": "Berhasil"})
}

func (ctrl *AdminPaymentController) GetAdminJumlahPesananBulanIni(c *gin.Context) {
	status := []string{"Paid", "Expired"}
	jumlahPesanan, err := ctrl.adminPaymentService.GetAdminJumlahPesananBulanIni(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"totalPesanan": jumlahPesanan, "message": "Berhasil", "code": 200, "status": "Berhasil"})
}

func (ctrl *AdminPaymentController) GetTotalProduk(c *gin.Context) {
	totalProduk, err := ctrl.adminPaymentService.GetTotalProduk()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"totalProduk": totalProduk, "message": "Berhasil", "code": 200, "status": "Berhasil"})
}

func (ctrl *AdminPaymentController) GetJumlahStatus(c *gin.Context) {
	paidCount, err := ctrl.adminPaymentService.GetJumlahPesananByStatus("PAID")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expiredCount, err := ctrl.adminPaymentService.GetJumlahPesananByStatus("EXPIRED")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalCount := paidCount + expiredCount

	c.JSON(http.StatusOK, gin.H{"jumlahBerhasil": paidCount, "jumlahGagal": expiredCount, "totalPesananBulanini": totalCount, "message": "Berhasil", "code": 200, "status": "Berhasil"})
}

func (ctrl *AdminPaymentController) GetJumlahPesananDikirim(c *gin.Context) {
	count, err := ctrl.adminPaymentService.GetJumlahPesananDikirim()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"totalPesananDikrim": count, "message": "Berhasil", "code": 200, "status": "Berhasil"})
}

func (ctrl *AdminPaymentController) GetJumlahPesananDiterima(c *gin.Context) {
	count, err := ctrl.adminPaymentService.GetJumlahPesananDiterima()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"totalPesananDiterima": count, "message": "Berhasil", "code": 200, "status": "Berhasil"})
}

func (ctrl *AdminPaymentController) TampilkanTotalPendapatan(c *gin.Context) {
	data, err := ctrl.adminPaymentService.GetTotalPendapatBulan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
	}
	c.JSON(http.StatusOK, gin.H{"message": data, "code": 200, "status": "Berhasil"})
}

func (ctrl *AdminPaymentController) GetJumlahArtikel(c *gin.Context) {
	count, err := ctrl.adminPaymentService.GetJumlahArtikel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"totalArtikel": count, "message": "Berhasil", "code": 200, "status": "Berhasil"})
}
