package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/gin-gonic/gin"
)

type PaymentControl struct {
	paymentServis services.PaymentServices
}

func NewPaymentController(paymentServis services.PaymentServices) *PaymentControl {
	return &PaymentControl{paymentServis: paymentServis}
}

func (ctrl *PaymentControl) TambahPayment(c *gin.Context) {
	var req struct {
		OrderID int `json:"orderID"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.OrderID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing Order ID"})
		return
	}
	orderID := req.OrderID
	_, err := ctrl.paymentServis.GetTotalAmount(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order amount"})
		return
	}
	invoiceData, err := ctrl.paymentServis.CreateInvoice(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invoice"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": gin.H{"message": "Berhasil", "code": 200, "status": "Berhasil"}, "invoiceData": invoiceData})
}

func (ctrl *PaymentControl) CheckPaymentStatus(c *gin.Context) {

	invoiceID := c.Param("invoiceID")
	if invoiceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing invoice ID"})
		return
	}

	status, err := ctrl.paymentServis.CheckPaymentStatus(invoiceID)
	if err != nil {
		log.Printf("Failed to check payment status: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check payment status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"meta": gin.H{"message": "Berhasil", "code": 200, "status": "Berhasil"}, "invoice_id": invoiceID, "status": status, "tanggal": time.Now()})
}

func (ctrl *PaymentControl) CancelPayment(c *gin.Context) {
	var requestBody struct {
		InvoiceID string `json:"invoice_id"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}
	if requestBody.InvoiceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invoice ID is required"})
		return
	}

	err := ctrl.paymentServis.CancelOrder(requestBody.InvoiceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel payment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pembayaran Berhasil Dibatalkan", "code": 200, "status": "Berhasil"})
}

func (ctrl *PaymentControl) GetPaymentByInvoiceID(c *gin.Context) {
	invoiceID := c.Param("invoiceID")

	if invoiceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invoice tidak ada"})
		return
	}

	payment, err := ctrl.paymentServis.GetPaymentByInvoiceID(invoiceID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pembayaran Tidak ada", "details": err.Error()})
		return
	}

	response := gin.H{
		"payment": gin.H{
			"id":                payment.ID,
			"invoice_id":        payment.InvoiceID,
			"status_pembayaran": payment.Status,
			"status_barang":     payment.StatusBarang,
			"jumlah":            payment.Jumlah,
			"payment_url":       payment.PaymentUrl,
			"order": gin.H{
				"id":                payment.Order.ID,
				"user_id":           payment.Order.UserID,
				"total":             payment.Order.Total,
				"biaya_layanan":     payment.Order.BiayaLayanan,
				"biaya_ongkir":      payment.Order.BiayaOngkir,
				"metode_pembayaran": payment.Order.MetodePembayaran,
				"details":           payment.Order.Details,
			},
		},
	}

	c.JSON(http.StatusOK, gin.H{"response": response, "code": 200, "status": "Berhasil"})
}

func (ctrl *PaymentControl) GetPaidOrders(c *gin.Context) {
	payments, err := ctrl.paymentServis.GetPaidOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var orders []gin.H
	for _, payment := range payments {
		orders = append(orders, gin.H{
			"id":                payment.ID,
			"invoice_id":        payment.InvoiceID,
			"jumlah":            payment.Jumlah,
			"status_pembayaran": payment.Status,
			"status_barang":     payment.StatusBarang,
			"order": gin.H{
				"id":                payment.Order.ID,
				"user_id":           payment.Order.UserID,
				"total":             payment.Order.Total,
				"biaya_layanan":     payment.Order.BiayaLayanan,
				"biaya_ongkir":      payment.Order.BiayaOngkir,
				"metode_pembayaran": payment.Order.MetodePembayaran,
				"details":           payment.Order.Details,
			},
		})
	}
	c.JSON(http.StatusOK, gin.H{"orders": orders, "code": 200, "status": "Berhasil"})
}
