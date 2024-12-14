package controllers

import (
	"fmt"
	"mime/multipart"
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
	var bindFile struct {
		File *multipart.FileHeader `form:"gambar" binding:"required"`
	}

	if err := c.ShouldBind(&bindFile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Gagal mengikat file: %s", err.Error())})
		return
	}

	file := bindFile.File
	filePath := "./uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan gambar"})
		return
	}

	linkFile := "https://www.bluebay.my.id/uploads/" + file.Filename
	artikel.Gambar = linkFile
	artikel.Judul = c.PostForm("judul")
	artikel.Deskripsi = c.PostForm("deskripsi")
	artikel.Kategori = c.PostForm("kategori")

	createdArtikel, err := ac.service.Create(&artikel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat menambah artikel"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Artikel berhasil ditambahkan",
		"data":    createdArtikel,
	})
}

func (ac *ArtikelController) Update(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var bindFile struct {
		File *multipart.FileHeader `form:"gambar"`
	}

	var artikel entities.Artikel
	if err := c.ShouldBind(&bindFile); err != nil && bindFile.File != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal mengikat file: " + err.Error()})
		return
	}

	existingArtikel, err := ac.service.FindByID(intID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
		return
	}

	if bindFile.File != nil {
		file := bindFile.File
		filePath := "./uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan gambar"})
			return
		}
		artikel.Gambar = "https://www.bluebay.my.id/uploads/" + file.Filename
	} else {
		artikel.Gambar = existingArtikel.Gambar
	}

	artikel.ID = existingArtikel.ID
	artikel.Judul = c.PostForm("judul")
	artikel.Deskripsi = c.PostForm("deskripsi")
	artikel.Kategori = c.PostForm("kategori")
	artikel.CreatedAt = existingArtikel.CreatedAt

	updatedArtikel, err := ac.service.Update(&artikel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat memperbarui artikel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Artikel berhasil diperbarui",
		"data":    updatedArtikel,
	})
}

func (ac *ArtikelController) Delete(c *gin.Context) {
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

	err = ac.service.Delete(artikel.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat menghapus artikel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artikel berhasil dihapus"})
}

func (ac *ArtikelController) GetAllForAdmin(c *gin.Context) {
	id := c.Param("id")
	page, err := strconv.Atoi(id)
	limit := 10

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
