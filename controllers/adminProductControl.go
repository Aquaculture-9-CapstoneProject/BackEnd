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
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := ac.service.CreateAdminProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product added successfully",
		"data":    createdProduct,
	})
}

func (ac *AdminProductController) UpdateAdminProduct(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var product entities.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingProduct, err := ac.service.FindByAdminProductID(intID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	product.ID = existingProduct.ID

	updatedProduct, err := ac.service.UpdateAdminProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product update successfully",
		"data":    updatedProduct,
	})
}

func (ac *AdminProductController) DeleteAdminProduct(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	product, err := ac.service.FindByAdminProductID(intID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	err = ac.service.DeleteAdminProduct(product.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func (ac *AdminProductController) GetAdminProductDetails(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	product, err := ac.service.FindByAdminProductID(intID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalItems, err := ac.service.GetAdminProductCount()
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
		"data": products,
	})
}
