package admincontroller

import (
	"net/http"

	adminservices "github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services/adminServices"
	"github.com/gin-gonic/gin"
)

type ProductExportController struct {
	Service adminservices.ProductExportService
}

func NewProductExportController(service adminservices.ProductExportService) *ProductExportController {
	return &ProductExportController{Service: service}
}

func (ctrl *ProductExportController) ExportToCSV(c *gin.Context) {
	err := ctrl.Service.ExportToCSV()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengekspor produk ke CSV"})
		return
	}
	c.Header("Content-Type", "application/csv")
	c.Header("Content-Disposition", "attachment; filename=products.csv")
	c.File("products.csv")
}
