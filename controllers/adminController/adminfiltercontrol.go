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

	// Membuat slice baru untuk menyimpan data yang dibutuhkan saja
	var responsePayments []gin.H
	for _, payment := range payments {
		// Membuat objek baru hanya dengan id dan nama
		responsePayments = append(responsePayments, gin.H{
			"id":                payment.ID,
			"id_pesanan":        payment.InvoiceID,
			"status":            payment.Status,
			"statusbarang":      payment.StatusBarang,
			"jumlah":            payment.Jumlah,
			"orderid":           payment.OrderID,
			"created_at":        payment.Order.CreatedAt,
			"metode_pembayaran": payment.Order.MetodePembayaran,
		})
	}

	// Mengirimkan response dengan hanya data yang dibutuhkan
	c.JSON(http.StatusOK, gin.H{"data": responsePayments})
}

func (ctrl *AdminFilterController) GetPaymentsByStatusBarang(c *gin.Context) {
	statusBarang := c.DefaultQuery("status_barang", "DIKIRIM")

	payments, err := ctrl.service.GetPaymentsByStatusBarang(statusBarang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data pembayaran"})
		return
	}

	var responsePayments []gin.H
	for _, payment := range payments {
		// Default nama pengguna jika tidak ada
		namapengguna := "Unknown"
		namaproduk := "unknown"

		// Mengakses nama pengguna dari OrderDetail pertama, jika ada
		if len(payment.Order.Details) > 0 {
			namapengguna = payment.Order.Details[0].User.NamaLengkap
		}
		if len(payment.Order.Details) > 0 {
			namaproduk = payment.Order.Details[0].Product.Nama
		}

		// Membuat response JSON
		responsePayments = append(responsePayments, gin.H{
			"order_id":        payment.ID,
			"namapengguna":    namapengguna,
			"statusbarang":    payment.StatusBarang,
			"nominal":         payment.Jumlah,
			"produk":          namaproduk,
			"status":          payment.StatusBarang,
			"tanggaldanwaktu": payment.Order.CreatedAt,
		})
	}
	c.JSON(http.StatusOK, gin.H{"data": responsePayments})
}
