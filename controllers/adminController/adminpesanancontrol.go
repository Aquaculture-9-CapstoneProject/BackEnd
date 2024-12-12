package admincontroller

import (
	"net/http"

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
	details, err := pc.serviceAdmin.GetDetailedOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"meta": gin.H{"message": "Berhasil", "code": 200, "status": "Berhasil"}, "data": details})
}
