package admincontroller

import (
	"math"
	"net/http"
	"strconv"

	adminservices "github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services/adminServices"
	"github.com/gin-gonic/gin"
)

type AdminPesananController struct {
	serviceAdmin adminservices.AdminPesananServices
}

func NewAdminPesananController(serviceAdmin adminservices.AdminPesananServices) *AdminPesananController {
	return &AdminPesananController{serviceAdmin: serviceAdmin}
}

func (pc *AdminPesananController) GetDetailedOrders(c *gin.Context) {
	// Mendapatkan query parameter untuk pagination
	page, _ := strconv.Atoi(c.Query("page"))
	perPage, _ := strconv.Atoi(c.Query("per_page"))

	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}

	// Panggil service untuk mendapatkan data pesanan
	details, totalItems, err := pc.serviceAdmin.GetDetailedOrders(page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menghitung total halaman
	totalPages := int(math.Ceil(float64(totalItems) / float64(perPage)))

	// Format respons JSON dengan pagination
	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{"message": "Berhasil", "code": 200, "status": "Berhasil"},
		"data": details,
		"pagination": gin.H{
			"page":        page,
			"per_page":    perPage,
			"total":       totalItems,
			"total_pages": totalPages,
		},
	})
}
