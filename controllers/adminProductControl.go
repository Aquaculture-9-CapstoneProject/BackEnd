package controllers

import (
	"net/http"
	"strconv"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/gin-gonic/gin"
)

type AdminProductController struct {
	service services.AdminProductUseCase
}

func NewAdminProductController(service services.AdminProductUseCase) *AdminProductController {
	return &AdminProductController{service: service}
}

func (ac *AdminProductController) CreateAdminProduct(c *gin.Context) {
	var product entities.Product

	file, err := c.FormFile("gambar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gambar tidak ditemukan"})
		return
	}

	filePath := "./uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan gambar"})
		return
	}

	product.Gambar = filePath
	product.Nama = c.PostForm("nama")
	product.Deskripsi = c.PostForm("deskripsi")
	product.Keunggulan = c.PostForm("keunggulan")
	product.Harga, _ = strconv.ParseFloat(c.PostForm("harga"), 64)
	product.Variasi = c.PostForm("variasi")
	product.Kategori = c.PostForm("kategori")
	product.KotaAsal = c.PostForm("kota_asal")
	product.Stok, _ = strconv.Atoi(c.PostForm("stok"))
	product.Status = c.PostForm("status")
	product.TipsPenyimpanan = c.PostForm("tips_penyimpanan")

	createdProduct, err := ac.service.CreateAdminProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat menambah produk"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Produk Berhasil Ditambahkan",
		"data":    createdProduct,
	})
}

func (ac *AdminProductController) UpdateAdminProduct(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var product entities.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal mengikat data JSON"})
		return
	}

	existingProduct, err := ac.service.FindByAdminProductID(intID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	product.ID = existingProduct.ID
	product.Gambar = existingProduct.Gambar
	product.Rating = existingProduct.Rating
	product.TotalReview = existingProduct.TotalReview
	product.Terjual = existingProduct.Terjual

	updatedProduct, err := ac.service.UpdateAdminProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat memperbarui produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Produk berhasil diperbarui",
		"data":    updatedProduct,
	})
}

func (ac *AdminProductController) DeleteAdminProduct(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	product, err := ac.service.FindByAdminProductID(intID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	err = ac.service.DeleteAdminProduct(product.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat menghapus produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
}

func (ac *AdminProductController) GetAdminProductDetails(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	product, err := ac.service.FindByAdminProductID(intID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil ditampilkan", "data": product})
}

func (ac *AdminProductController) GetAllAdminProducts(c *gin.Context) {
	id := c.Param("id")
	page, err := strconv.Atoi(id)
	limit := 15

	products, err := ac.service.GetAllAdminProducts(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat mengambil produk"})
		return
	}

	totalItems, err := ac.service.GetAdminProductCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat menghitung total produk"})
		return
	}

	totalPages := int((totalItems + int64(limit) - 1) / int64(limit))

	c.JSON(http.StatusOK, gin.H{
		"pagination": entities.Pagination{
			CurrentPage: page,
			TotalPages:  totalPages,
			TotalItems:  totalItems,
		},
		"data":    products,
		"message": "Produk berhasil ditampilkan",
	})
}

func (ac *AdminProductController) SearchAdminProducts(c *gin.Context) {
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
		limit = 15
	}

	products, err := ac.service.SearchAdminProducts(nama, kategori, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat mencari produk"})
		return
	}

	totalItems, err := ac.service.GetAdminProductCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat menghitung total produk"})
		return
	}

	totalPages := int((totalItems + int64(limit) - 1) / int64(limit))

	c.JSON(http.StatusOK, gin.H{
		"pagination": entities.Pagination{
			CurrentPage: page,
			TotalPages:  totalPages,
			TotalItems:  totalItems,
		},
		"data":    products,
		"message": "Produk berhasil ditampilkan",
	})
}
