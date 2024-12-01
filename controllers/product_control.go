package controllers

import (
	"net/http"
	"strconv"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service services.ProductUseCase
}

func NewProductController(service services.ProductUseCase) *ProductController {
	return &ProductController{service: service}
}

func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	products, err := c.service.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve products",
			"data":    nil,
		})
		return
	}

	var formattedProducts []gin.H

	for _, product := range products {
		var totalRating int
		ratingCount := len(product.Ratings)

		if ratingCount > 0 {
			for _, rating := range product.Ratings {
				totalRating += rating.Rating
			}
		}

		averageRating := 0.0
		if ratingCount > 0 {
			averageRating = float64(totalRating) / float64(ratingCount)
		}

		formattedProducts = append(formattedProducts, gin.H{
			"nama":     product.Nama,
			"harga":    product.Harga,
			"gambar":   product.Gambar,
			"rating":   averageRating,
			"kategori": product.Kategori,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully",
		"data":    formattedProducts,
	})
}

func (c *ProductController) GetProductByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	product, averageRating, err := c.service.GetProductByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"product":        product,
		"average_rating": averageRating,
	})
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var product entities.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdProduct, err := c.service.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, createdProduct)
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var product entities.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ID = id
	updatedProduct, err := c.service.UpdateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedProduct)
}

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.service.DeleteProduct(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
