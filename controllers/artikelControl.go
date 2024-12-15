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

func (ac *ArtikelController) TopArtikel(c *gin.Context) {
	limit := 3
	artikels, err := ac.service.Top3(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat mengambil artikel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    artikels,
		"message": "Artikel berhasil ditampilkan",
	})
}

func (ac *ArtikelController) GetAllForUser(c *gin.Context) {
	id := c.Param("id")
	page, err := strconv.Atoi(id)
	limit := 9

	artikels, err := ac.service.GetAll(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat mengambil artikel"})
		return
	}

	totalItems, err := ac.service.Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat menghitung total artikel"})
		return
	}

	totalPages := int((totalItems + int64(limit) - 1) / int64(limit))

	c.JSON(http.StatusOK, gin.H{
		"pagination": entities.Pagination{
			CurrentPage: page,
			TotalPages:  totalPages,
			TotalItems:  totalItems,
		},
		"data":    artikels,
		"message": "Artikel berhasil ditampilkan",
	})
}

func (ac *ArtikelController) FindAll(c *gin.Context) {
	judul := c.Query("judul")
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

	artikels, err := ac.service.FindAll(judul, kategori, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat mencari artikel"})
		return
	}

	totalItems, err := ac.service.Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat menghitung total artikel"})
		return
	}

	totalPages := int((totalItems + int64(limit) - 1) / int64(limit))

	c.JSON(http.StatusOK, gin.H{
		"pagination": entities.Pagination{
			CurrentPage: page,
			TotalPages:  totalPages,
			TotalItems:  totalItems,
		},
		"data":    artikels,
		"message": "Artikel berhasil ditampilkan",
	})
}

func (ac *ArtikelController) GetDetails(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	artikel, err := ac.service.FindByID(intID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artikel berhasil ditampilkan", "data": artikel})
}
