package controllers

import (
	"net/http"
	"strconv"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUseCase services.ProductUseCase
}

func NewProductController(productUseCase services.ProductUseCase) *ProductController {
	return &ProductController{productUseCase: productUseCase}
}

func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	products, err := c.productUseCase.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (c *ProductController) GetProductByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	product, err := c.productUseCase.GetProductByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	ctx.JSON(http.StatusOK, product)
}
