package controllers

import (
	"net/http"
	"strconv"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/gin-gonic/gin"
)

type ArtikelController struct {
	service services.ArtikelUseCase
}

func NewArtikelController(service services.ArtikelUseCase) *ArtikelController {
	return &ArtikelController{service: service}
}

func (ac *ArtikelController) Create(c *gin.Context) {
	var artikel entities.Artikel
	if err := c.ShouldBindJSON(&artikel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin, err := ac.service.GetAdminByID(artikel.AdminID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Admin ID"})
		return
	}

	createdArtikel, err := ac.service.Create(&artikel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createdArtikel.Admin = *admin

	c.JSON(http.StatusCreated, gin.H{
		"message": "Artikel added successfully",
		"data":    createdArtikel,
	})
}

func (ac *ArtikelController) Update(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var artikel entities.Artikel
	if err := c.ShouldBindJSON(&artikel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingArtikel, err := ac.service.FindByID(intID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel not found"})
		return
	}

	artikel.ID = existingArtikel.ID
	artikel.CreatedAt = existingArtikel.CreatedAt

	admin, err := ac.service.GetAdminByID(artikel.AdminID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Admin ID"})
		return
	}

	updatedArtikel, err := ac.service.Update(&artikel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedArtikel.Admin = *admin

	c.JSON(http.StatusOK, gin.H{
		"message": "Artikel update successfully",
		"data":    updatedArtikel,
	})
}

func (ac *ArtikelController) Delete(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	artikel, err := ac.service.FindByID(intID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel not found"})
		return
	}

	err = ac.service.Delete(artikel.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artikel deleted successfully"})
}

func (ac *ArtikelController) GetAll(c *gin.Context) {
	nama := c.Query("nama")
	kategori := c.Query("kategori")
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 9
	}

	artikels, err := ac.service.GetAll(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalItems, err := ac.service.Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := int((totalItems + int64(limit) - 1) / int64(limit))

	c.JSON(http.StatusOK, gin.H{
		"pagination": entities.Pagination{
			CurrentPage: page,
			TotalPages:  totalPages,
			TotalItems:  totalItems,
		},
		"data": artikels,
	})
}

func (ac *ArtikelController) GetDetails(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	artikel, err := ac.service.FindByID(intID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": artikel})
}
