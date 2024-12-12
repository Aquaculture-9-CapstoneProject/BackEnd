package controllers

import (
	"net/http"
	"strconv"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/gin-gonic/gin"
)

type ReviewController struct {
	ReviewService services.ReviewServices
}

func NewReviewController(reviewService services.ReviewServices) *ReviewController {
	return &ReviewController{ReviewService: reviewService}
}

func (ctrl *ReviewController) AddReview(c *gin.Context) {
	var request struct {
		Rating float64 `json:"rating"`
		Ulasan string  `json:"ulasan"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	productID := c.Param("product_id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	err = ctrl.ReviewService.AddReview(userID.(int), id, request.Rating, request.Ulasan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to add review",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tambah Ulasan Berhasil", "code": 200, "status": "Berhasil"})
}

func (ctrl *ReviewController) GetReviewsByProduct(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	reviews, err := ctrl.ReviewService.GetReviewsByProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": reviews, "message": "Berhasil", "code": 200, "status": "Berhasil"})
}
